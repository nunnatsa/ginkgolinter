package linter

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	gotypes "go/types"
	"time"

	"github.com/go-toolsmith/astcopy"
	"golang.org/x/tools/go/analysis"

	"github.com/nunnatsa/ginkgolinter/internal/expression"
	"github.com/nunnatsa/ginkgolinter/internal/expression/actual"
	"github.com/nunnatsa/ginkgolinter/internal/expression/matcher"
	"github.com/nunnatsa/ginkgolinter/internal/expression/value"
	"github.com/nunnatsa/ginkgolinter/internal/ginkgohandler"
	"github.com/nunnatsa/ginkgolinter/internal/gomegahandler"
	"github.com/nunnatsa/ginkgolinter/internal/intervals"
	"github.com/nunnatsa/ginkgolinter/internal/reports"
	"github.com/nunnatsa/ginkgolinter/types"
)

// The ginkgolinter enforces standards of using ginkgo and gomega.
//
// For more details, look at the README.md file

const (
	linterName                     = "ginkgo-linter"
	wrongLengthWarningTemplate     = "wrong length assertion"
	wrongCapWarningTemplate        = "wrong cap assertion"
	wrongNilWarningTemplate        = "wrong nil assertion"
	wrongBoolWarningTemplate       = "wrong boolean assertion"
	wrongErrWarningTemplate        = "wrong error assertion"
	wrongCompareWarningTemplate    = "wrong comparison assertion"
	doubleNegativeWarningTemplate  = "avoid double negative assertion"
	valueInEventually              = "use a function call in %[1]s. This actually checks nothing, because %[1]s receives the function returned value, instead of function itself, and this value is never changed"
	comparePointerToValue          = "comparing a pointer to a value will always fail"
	missingAssertionMessage        = linterName + `: %q: missing assertion method. Expected %s`
	focusContainerFound            = linterName + ": Focus container found. This is used only for local debug and should not be part of the actual source code. Consider to replace with %q"
	focusSpecFound                 = linterName + ": Focus spec found. This is used only for local debug and should not be part of the actual source code. Consider to remove it"
	compareDifferentTypes          = "use %[1]s with different types: Comparing %[2]s with %[3]s; either change the expected value type if possible, or use the BeEquivalentTo() matcher, instead of %[1]s()"
	matchErrorArgWrongType         = "the MatchError matcher used to assert a non error type (%s)"
	matchErrorWrongTypeAssertion   = "MatchError first parameter (%s) must be error, string, GomegaMatcher or func(error)bool are allowed"
	matchErrorMissingDescription   = "missing function description as second parameter of MatchError"
	matchErrorRedundantArg         = "redundant MatchError arguments; consider removing them"
	matchErrorNoFuncDescription    = "The second parameter of MatchError must be the function description (string)"
	forceExpectToTemplate          = "must not use %s with %s"
	useBeforeEachTemplate          = "use BeforeEach() to assign variable %s"
	onlyUseTimeDurationForInterval = "only use time.Duration for timeout and polling in Eventually() or Consistently()"
)

const ( // gomega actuals
	omega                  = "Ω"
	expect                 = "Expect"
	expectWithOffset       = "ExpectWithOffset"
	eventually             = "Eventually"
	eventuallyWithOffset   = "EventuallyWithOffset"
	consistently           = "Consistently"
	consistentlyWithOffset = "ConsistentlyWithOffset"
)

type GinkgoLinter struct {
	config *types.Config
}

// NewGinkgoLinter return new ginkgolinter object
func NewGinkgoLinter(config *types.Config) *GinkgoLinter {
	return &GinkgoLinter{
		config: config,
	}
}

// Run is the main assertion function
func (l *GinkgoLinter) Run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		fileConfig := l.config.Clone()

		cm := ast.NewCommentMap(pass.Fset, file, file.Comments)

		fileConfig.UpdateFromFile(cm)

		gomegaHndlr := gomegahandler.GetGomegaHandler(file)
		ginkgoHndlr := ginkgohandler.GetGinkgoHandler(file)

		if gomegaHndlr == nil && ginkgoHndlr == nil { // no gomega or ginkgo imports => no use in gomega in this file; nothing to do here
			continue
		}

		ast.Inspect(file, func(n ast.Node) bool {
			if ginkgoHndlr != nil {
				goDeeper := false
				spec, ok := n.(*ast.ValueSpec)
				if ok {
					for _, val := range spec.Values {
						if exp, ok := val.(*ast.CallExpr); ok {
							if bool(fileConfig.ForbidFocus) && checkFocusContainer(pass, ginkgoHndlr, exp) {
								goDeeper = true
							}

							if bool(fileConfig.ForbidSpecPollution) && checkAssignmentsInContainer(pass, ginkgoHndlr, exp) {
								goDeeper = true
							}
						}
					}
				}
				if goDeeper {
					return true
				}
			}

			stmt, ok := n.(*ast.ExprStmt)
			if !ok {
				return true
			}

			config := fileConfig.Clone()

			if comments, ok := cm[stmt]; ok {
				config.UpdateFromComment(comments)
			}

			// search for function calls
			assertionExp, ok := stmt.X.(*ast.CallExpr)
			if !ok {
				return true
			}

			if ginkgoHndlr != nil {
				goDeeper := false
				if bool(config.ForbidFocus) && checkFocusContainer(pass, ginkgoHndlr, assertionExp) {
					goDeeper = true
				}
				if bool(config.ForbidSpecPollution) && checkAssignmentsInContainer(pass, ginkgoHndlr, assertionExp) {
					goDeeper = true
				}
				if goDeeper {
					return true
				}
			}

			// no more ginkgo checks. From here it's only gomega. So if there is no gomega handler, exit here. This is
			// mostly to prevent nil pointer error.
			if gomegaHndlr == nil {
				return true
			}

			assertionFunc, ok := assertionExp.Fun.(*ast.SelectorExpr)
			if !ok {
				checkNoAssertion(pass, assertionExp, gomegaHndlr)
				return true
			}

			if !isAssertionFunc(assertionFunc.Sel.Name) {
				checkNoAssertion(pass, assertionExp, gomegaHndlr)
				return true
			}

			actualExpr := gomegaHndlr.GetActualExpr(assertionFunc)
			if actualExpr == nil {
				return true
			}

			return checkExpression(pass, config, assertionExp, gomegaHndlr, getTimePkg(file))
		})
	}
	return nil, nil
}

func checkAssignmentsInContainer(pass *analysis.Pass, ginkgoHndlr ginkgohandler.Handler, exp *ast.CallExpr) bool {
	foundSomething := false
	if ginkgoHndlr.IsWrapContainer(exp) {
		for _, arg := range exp.Args {
			if fn, ok := arg.(*ast.FuncLit); ok {
				if fn.Body != nil {
					if checkAssignments(pass, fn.Body.List) {
						foundSomething = true
					}
					break
				}
			}
		}
	}

	return foundSomething
}

func checkAssignments(pass *analysis.Pass, list []ast.Stmt) bool {
	foundSomething := false
	for _, stmt := range list {
		switch st := stmt.(type) {
		case *ast.DeclStmt:
			if gen, ok := st.Decl.(*ast.GenDecl); ok {
				if gen.Tok != token.VAR {
					continue
				}
				for _, spec := range gen.Specs {
					if valSpec, ok := spec.(*ast.ValueSpec); ok {
						if checkAssignmentsValues(pass, valSpec.Names, valSpec.Values) {
							foundSomething = true
						}
					}
				}
			}

		case *ast.AssignStmt:
			for i, val := range st.Rhs {
				if !is[*ast.FuncLit](val) {
					if id, isIdent := st.Lhs[i].(*ast.Ident); isIdent && id.Name != "_" {
						reportNoFix(pass, id.Pos(), useBeforeEachTemplate, id.Name)
						foundSomething = true
					}
				}
			}

		case *ast.IfStmt:
			if st.Body != nil {
				if checkAssignments(pass, st.Body.List) {
					foundSomething = true
				}
			}
			if st.Else != nil {
				if block, isBlock := st.Else.(*ast.BlockStmt); isBlock {
					if checkAssignments(pass, block.List) {
						foundSomething = true
					}
				}
			}
		}
	}

	return foundSomething
}

func checkAssignmentsValues(pass *analysis.Pass, names []*ast.Ident, values []ast.Expr) bool {
	foundSomething := false
	for i, val := range values {
		if !is[*ast.FuncLit](val) {
			reportNoFix(pass, names[i].Pos(), useBeforeEachTemplate, names[i].Name)
			foundSomething = true
		}
	}

	return foundSomething
}

func checkFocusContainer(pass *analysis.Pass, ginkgoHndlr ginkgohandler.Handler, exp *ast.CallExpr) bool {
	foundFocus := false
	isFocus, id := ginkgoHndlr.GetFocusContainerName(exp)
	if isFocus {
		reportNewName(pass, id, id.Name[1:], focusContainerFound, id.Name)
		foundFocus = true
	}

	if id != nil && ginkgohandler.IsContainer(id.Name) {
		for _, arg := range exp.Args {
			if ginkgoHndlr.IsFocusSpec(arg) {
				reportNoFix(pass, arg.Pos(), focusSpecFound)
				foundFocus = true
			} else if callExp, ok := arg.(*ast.CallExpr); ok {
				if checkFocusContainer(pass, ginkgoHndlr, callExp) { // handle table entries
					foundFocus = true
				}
			}
		}
	}

	return foundFocus
}

func checkExpression(pass *analysis.Pass, config types.Config, assertionExp *ast.CallExpr, handler gomegahandler.Handler, timePkg string) bool {
	expr := astcopy.CallExpr(assertionExp)

	reportBuilder := reports.NewBuilder(pass.Fset, expr)
	gexp, ok := expression.New(assertionExp, pass, handler, timePkg)
	if !ok || gexp == nil {
		return false
	}

	goNested := false
	if gexp.IsAsync() {
		goNested = checkAsyncAssertion(gexp, pass, config, reportBuilder)
	} else {
		if config.ForceExpectTo {
			var newName, assertFuncName string
			newName, assertFuncName, goNested = gexp.ForceExpectTo()
			if newName != "" {
				reportBuilder.AddIssue(true, forceExpectToTemplate, gexp.GetActualFuncName(), assertFuncName)
			}
		}

		goNested = doCheckExpression(pass, gexp, config, reportBuilder) || goNested
	}

	if reportBuilder.HasReport() {
		reportBuilder.SetFixOffer(pass.Fset, gexp.Clone)
		pass.Report(reportBuilder.Build())
	}

	return goNested
}

func doCheckExpression(pass *analysis.Pass, gexp *expression.GomegaExpression, config types.Config, reportBuilder *reports.Builder) bool {
	actualArgType := gexp.Actual.Arg.ArgType()

	if !bool(config.SuppressLen) && actualArgType.Is(actual.LenFuncActualArgType) {
		return checkLengthMatcher(gexp, reportBuilder)

	} else if !bool(config.SuppressLen) && actualArgType.Is(actual.CapFuncActualArgType) {
		return checkCapMatcher(gexp, reportBuilder)
	} else if actualArgType.Is(actual.NilComparisonActualArgType) {
		return checkNilMatcher(gexp, config, reportBuilder)
	} else if actualArgType.Is(actual.ComparisonActualArgType) {
		if !startCheckComparison(gexp) {
			return false
		}

		actl, ok := gexp.Actual.Arg.(*actual.FuncComparisonPayload)
		if ok && !bool(config.SuppressLen) {
			if actualArgType.Is(actual.LenComparisonActualArgType) {
				return handleLenComparison(gexp, actl, reportBuilder)
			} else if actualArgType.Is(actual.CapComparisonActualArgType) {
				return handleCapComparison(gexp, actl, reportBuilder)
			}
		} else {
			return bool(config.SuppressCompare) || checkComparison(gexp, reportBuilder)
		}
	} else if checkMatchError(pass, gexp, reportBuilder) {
		return false
	} else if actualArgType.Is(actual.ErrActualArgType) {
		return bool(config.SuppressErr) || checkNilError(gexp, reportBuilder)

	} else if checkPointerComparison(gexp, config, reportBuilder) {
		return false
	} else if !handleAssertionOnly(gexp, config, reportBuilder) {
		return false
	} else if !config.SuppressTypeCompare {
		return !checkEqualWrongType(gexp, reportBuilder)
	}

	return true
}

func checkMatchError(pass *analysis.Pass, gexp *expression.GomegaExpression, reportBuilder *reports.Builder) bool {
	mtchr := gexp.Matcher.GetMatcherInfo()
	switch m := mtchr.(type) {
	case matcher.MatchErrorMatcher:
		return checkMatchErrorMatcher(gexp, gexp.Matcher, m, pass, reportBuilder)
	case *matcher.MultipleMatchersMatcher:
		res := false
		for i := range m.Len() {
			nested := m.At(i)
			if specific, ok := nested.GetMatcherInfo().(matcher.MatchErrorMatcher); ok {
				if valid := checkMatchErrorMatcher(gexp, gexp.Matcher, specific, pass, reportBuilder); valid {
					res = true
				}
			}
		}
		return res
	default:
		return false
	}
}

func checkMatchErrorMatcher(gexp *expression.GomegaExpression, mtchr *matcher.Matcher, mtchrInfo matcher.MatchErrorMatcher, pass *analysis.Pass, reportBuilder *reports.Builder) bool {
	if !gexp.Actual.Arg.ArgType().Is(actual.ErrorTypeArgType) {
		reportBuilder.AddIssue(false, matchErrorArgWrongType, goFmt(pass.Fset, gexp.Actual.GetActualArg()))
	}

	switch m := mtchrInfo.(type) {
	case *matcher.InvalidMatchErrorMatcher:
		reportBuilder.AddIssue(false, matchErrorWrongTypeAssertion, goFmt(pass.Fset, mtchr.Clone.Args[0]))

	case *matcher.MatchErrorMatcherWithErrFunc:
		if m.NumArgs() == m.AllowedNumArgs() {
			if !m.IsSecondArgString() {
				reportBuilder.AddIssue(false, matchErrorNoFuncDescription)
			}
			return true
		}

		if m.NumArgs() == 1 {
			reportBuilder.AddIssue(false, matchErrorMissingDescription)
			return true
		}

	case *matcher.MatchErrorMatcherWithErr,
		*matcher.MatchErrorMatcherWithMatcher,
		*matcher.MatchErrorMatcherWithString:
		// continue
	default:
		return false
	}

	if mtchrInfo.NumArgs() == mtchrInfo.AllowedNumArgs() {
		return true
	}

	if mtchrInfo.NumArgs() > mtchrInfo.AllowedNumArgs() {
		var newArgsSuggestion []ast.Expr
		for i := 0; i < mtchrInfo.AllowedNumArgs(); i++ {
			newArgsSuggestion = append(newArgsSuggestion, mtchr.Clone.Args[i])
		}
		mtchr.Clone.Args = newArgsSuggestion
		reportBuilder.AddIssue(false, matchErrorRedundantArg)
		return true
	}
	return false
}

func checkEqualWrongType(gexp *expression.GomegaExpression, reportBuilder *reports.Builder) bool {
	return checkEqualDifferentTypes(gexp, gexp.Matcher, false, reportBuilder)
}

func checkEqualDifferentTypes(gexp *expression.GomegaExpression, mtchr *matcher.Matcher, parentPointer bool, reportBuilder *reports.Builder) bool {
	actualType := gexp.Actual.ArgGOType()

	if parentPointer {
		if t, ok := actualType.(*gotypes.Pointer); ok {
			actualType = t.Elem()
		}
	}

	var (
		matcherType gotypes.Type
		matcherName string
	)

	switch specificMatcher := mtchr.GetMatcherInfo().(type) {
	case *matcher.EqualMatcher:
		matcherType = specificMatcher.GetType()
		matcherName = specificMatcher.MatcherName()

	case *matcher.BeIdenticalToMatcher:
		matcherType = specificMatcher.GetType()
		matcherName = specificMatcher.MatcherName()

	case *matcher.HaveValueMatcher:
		return checkEqualDifferentTypes(gexp, specificMatcher.GetNested(), true, reportBuilder)

	case *matcher.MultipleMatchersMatcher:
		foundIssue := false
		for i := range specificMatcher.Len() {
			if checkEqualDifferentTypes(gexp, specificMatcher.At(i), parentPointer, reportBuilder) {
				foundIssue = true
			}

		}
		return foundIssue

	case *matcher.EqualNilMatcher:
		matcherType = specificMatcher.GetType()
		matcherName = specificMatcher.MatcherName()

	case *matcher.WithTransformMatcher:
		nested := specificMatcher.GetNested()
		switch specificNested := nested.GetMatcherInfo().(type) {
		case *matcher.EqualMatcher:
			matcherType = specificNested.GetType()
			matcherName = specificNested.MatcherName()

		case *matcher.BeIdenticalToMatcher:
			matcherType = specificNested.GetType()
			matcherName = specificNested.MatcherName()

		default:
			return false
		}

		actualType = specificMatcher.GetFuncType()
	default:
		return false
	}

	if !gotypes.Identical(matcherType, actualType) {
		if isImplementing(matcherType, actualType) || isImplementing(actualType, matcherType) {
			return false
		}

		reportBuilder.AddIssue(false, compareDifferentTypes, matcherName, actualType, matcherType)
		return true
	}

	return false
}

func isImplementing(ifs, impl gotypes.Type) bool {
	if gotypes.IsInterface(ifs) {

		var (
			theIfs *gotypes.Interface
			ok     bool
		)

		for {
			theIfs, ok = ifs.(*gotypes.Interface)
			if ok {
				break
			}
			ifs = ifs.Underlying()
		}

		return gotypes.Implements(impl, theIfs)
	}
	return false
}

func checkPointerComparison(gexp *expression.GomegaExpression, config types.Config, reportBuilder *reports.Builder) bool {
	actl, ok := gexp.Actual.Arg.(*actual.RegularArgPayload)
	if !ok {
		return false
	}

	if !actl.IsPointer() {
		return false
	}

	switch mtchr := gexp.Matcher.GetMatcherInfo().(type) {
	case *matcher.EqualMatcher:
		if mtchr.IsPointer() || mtchr.IsInterface() {
			return false
		}

	case *matcher.BeEquivalentToMatcher:
		if mtchr.IsPointer() || mtchr.IsInterface() || mtchr.IsNil() {
			return false
		}

	case *matcher.BeIdenticalToMatcher:
		if mtchr.IsPointer() || mtchr.IsInterface() || mtchr.IsNil() {
			return false
		}

	case *matcher.EqualNilMatcher:
		return false

	case *matcher.BeTrueMatcher,
		*matcher.BeFalseMatcher,
		*matcher.BeNumericallyMatcher,
		*matcher.EqualTrueMatcher,
		*matcher.EqualFalseMatcher:

	default:
		return false
	}

	handleAssertionOnly(gexp, config, reportBuilder)
	gexp.SetMatcherHaveValue()
	reportBuilder.AddIssue(true, comparePointerToValue)
	return true
}

// check async assertion does not assert function call. This is a real bug in the test. In this case, the assertion is
// done on the returned value, instead of polling the result of a function, for instance.
func checkAsyncAssertion(gexp *expression.GomegaExpression, pass *analysis.Pass, config types.Config, reportBuilder *reports.Builder) bool {
	asyncArg := gexp.Actual.GetAsyncArg()

	if !config.SuppressAsync {
		if !asyncArg.IsValid() {
			gexp.Actual.AppendWithArgsMethod()

			funcName := gexp.Actual.FuncName()
			reportBuilder.AddIssue(true, valueInEventually, funcName)
		}

		if config.ValidateAsyncIntervals {
			if asyncArg.TooManyTimeouts() {
				reportBuilder.AddIssue(false, "timeout defined more than once")
			}

			if asyncArg.TooManyPolling() {
				reportBuilder.AddIssue(false, "polling defined more than once")
			}

			timeoutDuration := checkInterval(gexp, asyncArg.Timeout(), reportBuilder)
			pollingDuration := checkInterval(gexp, asyncArg.Polling(), reportBuilder)

			if timeoutDuration > 0 && pollingDuration > 0 && pollingDuration > timeoutDuration {
				reportBuilder.AddIssue(false, "timeout must not be shorter than the polling interval")
			}
		}
	}

	if !config.SuppressErr {
		if gexp.Actual.Arg.ArgType().Is(actual.ErrorTypeArgType) {
			if !checkNilError(gexp, reportBuilder) {
				return true
			}
		}
	}

	checkMatchError(pass, gexp, reportBuilder)

	handleAssertionOnly(gexp, config, reportBuilder)

	return true
}

func checkInterval(gexp *expression.GomegaExpression, durVal intervals.DurationValue, reportBuilder *reports.Builder) time.Duration {
	if durVal != nil {
		switch to := durVal.(type) {
		case *intervals.RealDurationValue, *intervals.UnknownDurationTypeValue:

		case *intervals.NumericDurationValue:
			if checkNumericInterval(gexp.Actual.Clone, to) {
				reportBuilder.AddIssue(true, onlyUseTimeDurationForInterval)
			}

		case *intervals.UnknownDurationValue:
			reportBuilder.AddIssue(true, onlyUseTimeDurationForInterval)
		}

		return durVal.Duration()
	}

	return 0
}

func checkNumericInterval(intervalMethod *ast.CallExpr, interval intervals.DurationValue) bool {
	if interval != nil {
		if numVal, ok := interval.(intervals.NumericValue); ok {
			if offset := numVal.GetOffset(); offset > 0 {
				intervalMethod.Args[offset] = numVal.GetDurationExpr()
				return true
			}
		}
	}

	return false
}

func startCheckComparison(gexp *expression.GomegaExpression) bool {
	switch gexp.Matcher.GetMatcherInfo().(type) {
	case *matcher.BeTrueMatcher, *matcher.EqualTrueMatcher:
	case *matcher.BeFalseMatcher, *matcher.EqualFalseMatcher:
		gexp.ReverseAssertionFuncLogic()

	default:
		return false
	}

	return true
}

func checkComparison(gexp *expression.GomegaExpression, reportBuilder *reports.Builder) bool {
	actl, ok := gexp.Actual.Arg.(actual.ComparisonActualPayload)
	if !ok {
		return true
	}

	switch actl.GetOp() {
	case token.EQL:
		handleEqualComparison(gexp, actl)

	case token.NEQ:
		gexp.ReverseAssertionFuncLogic()
		handleEqualComparison(gexp, actl)
	case token.GTR, token.GEQ, token.LSS, token.LEQ:
		if !actl.GetRight().IsValueNumeric() {
			return true
		}

		gexp.SetMatcherBeNumerically(actl.GetOp(), actl.GetRight().GetValueExpr())

	default:
		return true
	}

	gexp.ReplaceActual(actl.GetLeft().GetValueExpr())

	reportBuilder.AddIssue(true, wrongCompareWarningTemplate)
	return false
}

func handleEqualComparison(gexp *expression.GomegaExpression, actual actual.ComparisonActualPayload) {
	if actual.GetRight().IsValueZero() {
		gexp.SetMatcherBeZero()
	} else {
		left := actual.GetLeft()
		arg := actual.GetRight().GetValueExpr()
		if left.IsInterface() || left.IsPointer() {
			gexp.SetMatcherBeIdenticalTo(arg)
		} else {
			gexp.SetMatcherEqual(arg)
		}
	}
}

func handleLenComparison(gexp *expression.GomegaExpression, actual *actual.FuncComparisonPayload, reportBuilder *reports.Builder) bool {
	if op := actual.GetOp(); op == token.NEQ {
		gexp.ReverseAssertionFuncLogic()
	} else if op != token.EQL {
		return false
	}

	if actual.IsValueZero() {
		gexp.SetMatcherBeEmpty()
	} else {
		gexp.SetMatcherLen(actual.GetValueExpr())
	}

	gexp.ReplaceActual(actual.GetFuncArg())

	reportBuilder.AddIssue(true, wrongLengthWarningTemplate)
	return true
}

func handleCapComparison(gexp *expression.GomegaExpression, actual *actual.FuncComparisonPayload, reportBuilder *reports.Builder) bool {
	if op := actual.GetOp(); op == token.NEQ {
		gexp.ReverseAssertionFuncLogic()
	} else if op != token.EQL {
		return false
	}

	gexp.SetMatcherCap(actual.GetValueExpr())
	gexp.ReplaceActual(actual.GetFuncArg())

	reportBuilder.AddIssue(true, wrongCapWarningTemplate)
	return true
}

// Check if matcher function is in one of the patterns we want to avoid
func checkLengthMatcher(gexp *expression.GomegaExpression, reportBuilder *reports.Builder) bool {
	switch mtchr := gexp.Matcher.GetMatcherInfo().(type) {
	case *matcher.EqualMatcher:
		gexp.SetLenNumericMatcher()
	case *matcher.BeZeroMatcher:
		gexp.SetMatcherBeEmpty()
	case *matcher.BeNumericallyMatcher:
		if handleLenBeNumerically(gexp, mtchr) {
			return true
		}
	default:
		return true
	}

	reportLengthAssertion(gexp, reportBuilder)
	return false
}

// Check if matcher function is in one of the patterns we want to avoid
func checkCapMatcher(gexp *expression.GomegaExpression, reportBuilder *reports.Builder) bool {
	//witch matcherFuncName {
	switch mtchr := gexp.Matcher.GetMatcherInfo().(type) {
	case *matcher.EqualMatcher:
		gexp.SetMatcherCap(mtchr.GetValueExpr())

	case *matcher.BeZeroMatcher:
		gexp.SetMatcherCapZero()

	case *matcher.BeNumericallyMatcher:
		if handleCapBeNumerically(gexp, mtchr) {
			return true
		}

	default:
		return true
	}

	reportCapAssertion(gexp, reportBuilder)
	return false
}

// Check if matcher function is in one of the patterns we want to avoid
func checkNilMatcher(gexp *expression.GomegaExpression, config types.Config, reportBuilder *reports.Builder) bool {
	actl, ok := gexp.Actual.Arg.(*actual.NilComparisonPayload)
	if !ok {
		return true
	}

	isErr := false
	if actl.IsError() {
		if !config.SuppressErr {
			isErr = true
		}
	} else {
		if config.SuppressNil {
			return true
		}
	}

	switch gexp.Matcher.GetMatcherInfo().(type) {
	case *matcher.EqualTrueMatcher, *matcher.BeTrueMatcher:
	case *matcher.EqualFalseMatcher, *matcher.BeFalseMatcher:
		gexp.ReverseAssertionFuncLogic()
	default:
		return true
	}
	handleNilBeBoolMatcher(gexp, actl, reportBuilder, isErr)
	return false
}

func handleNilBeBoolMatcher(gexp *expression.GomegaExpression, actual *actual.NilComparisonPayload, reportBuilder *reports.Builder, isErr bool) {
	template := wrongNilWarningTemplate
	if isErr {
		template = wrongErrWarningTemplate
		if actual.IsFunc() {
			gexp.SetMatcherSucceed()
		} else {
			gexp.ReverseAssertionFuncLogic()
			gexp.SetMatcherHaveOccurred()
		}
	} else {
		gexp.SetMatcherBeNil()
	}

	gexp.ReplaceActual(actual.GetValueExpr())

	if actual.GetOp() == token.NEQ {
		gexp.ReverseAssertionFuncLogic()
	}

	reportBuilder.AddIssue(true, template)
}

func checkNilError(gexp *expression.GomegaExpression, reportBuilder *reports.Builder) bool {

	switch gexp.Matcher.GetMatcherInfo().(type) {
	case *matcher.EqualNilMatcher, *matcher.BeNilMatcher:
	default:
		return true
	}

	if v, ok := gexp.Actual.Arg.(value.Valuer); ok && v.IsFunc() || gexp.Actual.Arg.ArgType().Is(actual.ErrFuncActualArgType) {
		gexp.SetMatcherSucceed()
	} else {
		gexp.ReverseAssertionFuncLogic()
		gexp.SetMatcherHaveOccurred()
	}

	reportBuilder.AddIssue(true, wrongErrWarningTemplate)
	return false
}

// handleAssertionOnly checks use-cases when the actual value is valid, but only the assertion should be fixed
// it handles:
//
//	Equal(nil) => BeNil()
//	Equal(true) => BeTrue()
//	Equal(false) => BeFalse()
//	HaveLen(0) => BeEmpty()
func handleAssertionOnly(gexp *expression.GomegaExpression, config types.Config, reportBuilder *reports.Builder) bool {
	var template string
	switch gexp.Matcher.GetMatcherInfo().(type) {
	case *matcher.EqualTrueMatcher:
		gexp.SetMatcherBeTrue()
		template = wrongBoolWarningTemplate

	case *matcher.EqualFalseMatcher:
		if gexp.IsNegativeAssertion() {
			gexp.ReverseAssertionFuncLogic()
			gexp.SetMatcherBeTrue()
		} else {
			gexp.SetMatcherBeFalse()
		}
		template = wrongBoolWarningTemplate

	case *matcher.EqualNilMatcher:
		if config.SuppressNil {
			return true
		}

		template = wrongNilWarningTemplate
		gexp.SetMatcherBeNil()

	case *matcher.BeFalseMatcher:
		if gexp.IsNegativeAssertion() {
			gexp.ReverseAssertionFuncLogic()
			gexp.SetMatcherBeTrue()
			template = doubleNegativeWarningTemplate
		} else {
			return true
		}

	case *matcher.HaveLenZeroMatcher:
		if config.AllowHaveLen0 {
			return true
		}

		gexp.SetMatcherBeEmpty()

		template = wrongLengthWarningTemplate

	default:
		return true
	}

	reportBuilder.AddIssue(true, template)
	return false
}

// For the BeNumerically matcher, we want to avoid the assertion of length to be > 0 or >= 1, or just == number
func handleLenBeNumerically(gexp *expression.GomegaExpression, matcher *matcher.BeNumericallyMatcher) bool {
	op := matcher.GetOp()
	val := matcher.GetValue()
	isValZero := val.String() == "0"
	isValOne := val.String() == "1"

	if (op == token.GTR && isValZero) || (op == token.GEQ && isValOne) {
		gexp.ReverseAssertionFuncLogic()
		gexp.SetMatcherBeEmpty()
	} else if op == token.EQL {
		gexp.SetLenNumericMatcher()
	} else if op == token.NEQ {
		gexp.ReverseAssertionFuncLogic()
		gexp.SetLenNumericMatcher()
	} else {
		return true
	}

	return false
}

// For the BeNumerically matcher, we want to avoid the assertion of length to be > 0 or >= 1, or just == number
func handleCapBeNumerically(gexp *expression.GomegaExpression, matcher *matcher.BeNumericallyMatcher) bool {
	op := matcher.GetOp()
	val := matcher.GetValue()
	isValZero := val.String() == "0"
	isValOne := val.String() == "1"

	if (op == token.GTR && isValZero) || (op == token.GEQ && isValOne) {
		gexp.ReverseAssertionFuncLogic()
		gexp.SetMatcherCapZero()
	} else if op == token.EQL {
		gexp.SetMatcherCap(matcher.GetValueExpr())
	} else if op == token.NEQ {
		gexp.ReverseAssertionFuncLogic()
		gexp.SetMatcherCap(matcher.GetValueExpr())
	} else {
		return true
	}

	return false
}

func isAssertionFunc(name string) bool {
	switch name {
	case "To", "ToNot", "NotTo", "Should", "ShouldNot":
		return true
	}
	return false
}

func reportLengthAssertion(gexp *expression.GomegaExpression, reportBuilder *reports.Builder) {
	gexp.ReplaceActualWithItsFirstArg()

	reportBuilder.AddIssue(true, wrongLengthWarningTemplate)
}

func reportCapAssertion(gexp *expression.GomegaExpression, reportBuilder *reports.Builder) {
	gexp.ReplaceActualWithItsFirstArg()

	reportBuilder.AddIssue(true, wrongCapWarningTemplate)
}

func reportNewName(pass *analysis.Pass, id *ast.Ident, newName string, messageTemplate, oldExpr string) {
	pass.Report(analysis.Diagnostic{
		Pos:     id.Pos(),
		Message: fmt.Sprintf(messageTemplate, newName),
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: fmt.Sprintf("should replace %s with %s", oldExpr, newName),
				TextEdits: []analysis.TextEdit{
					{
						Pos:     id.Pos(),
						End:     id.End(),
						NewText: []byte(newName),
					},
				},
			},
		},
	})
}

func reportNoFix(pass *analysis.Pass, pos token.Pos, message string, args ...any) {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	pass.Report(analysis.Diagnostic{
		Pos:     pos,
		Message: message,
	})
}

func goFmt(fset *token.FileSet, x ast.Expr) string {
	var b bytes.Buffer
	_ = printer.Fprint(&b, fset, x)
	return b.String()
}

func checkNoAssertion(pass *analysis.Pass, expr *ast.CallExpr, handler gomegahandler.Handler) {
	funcName, ok := handler.GetActualFuncName(expr)
	if ok {
		var allowedFunction string
		switch funcName {
		case expect, expectWithOffset:
			allowedFunction = `"To()", "ToNot()" or "NotTo()"`
		case eventually, eventuallyWithOffset, consistently, consistentlyWithOffset:
			allowedFunction = `"Should()" or "ShouldNot()"`
		case omega:
			allowedFunction = `"Should()", "To()", "ShouldNot()", "ToNot()" or "NotTo()"`
		default:
			return
		}
		reportNoFix(pass, expr.Pos(), missingAssertionMessage, funcName, allowedFunction)
	}
}

func is[T any](x any) bool {
	_, matchType := x.(T)
	return matchType
}

func getTimePkg(file *ast.File) string {
	timePkg := "time"
	for _, imp := range file.Imports {
		if imp.Path.Value == `"time"` && imp.Name != nil {
			timePkg = imp.Name.Name
		}
	}

	return timePkg
}
