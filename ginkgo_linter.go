package ginkgolinter

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/nunnatsa/ginkgolinter/reverseassertion"
)

// The ginkgolinter enforces standards of using ginkgo and gomega.
//
// The current checks are:
// * enforce right length assertion - warn for assertion of len(something):
//
//   This check finds the following patterns and suggests an alternative
//   * Expect(len(something)).To(Equal(number)) ===> Expect(x).To(HaveLen(number))
//   * ExpectWithOffset(1, len(something)).ShouldNot(Equal(0)) ===> ExpectWithOffset(1, something).ShouldNot(BeEmpty())
//   * Ω(len(something)).NotTo(BeZero()) ===> Ω(something).NotTo(BeEmpty())
//   * Expect(len(something)).To(BeNumerically(">", 0)) ===> Expect(something).ToNot(BeEmpty())
//   * Expect(len(something)).To(BeNumerically(">=", 1)) ===> Expect(something).ToNot(BeEmpty())
//   * Expect(len(something)).To(BeNumerically("==", number)) ===> Expect(something).To(HaveLen(number))
//
// * enforce right nil assertion - warn for assertion of x == nil:
//   This check finds the following patterns and suggests an alternative
//   * Expect(x == nil).Should(Equal(true))
//   * Expect(nil == x).Should(BeTrue())
//   * Expect(x != nil).Should(Equal(false))
//   * Expect(nil == x).Should(BeFalse())

const (
	linterName                 = "ginkgo-linter"
	wrongLengthWarningTemplate = linterName + ": wrong length assertion; consider using `%s` instead"
	wrongNilWarningTemplate    = linterName + ": wrong nil assertion; consider using `%s` instead"
	beEmpty                    = "BeEmpty"
	beNil                      = "BeNil"
	haveLen                    = "HaveLen"
	expect                     = "Expect"
	omega                      = "Ω"
	expectWithOffset           = "ExpectWithOffset"

	suppressPrefix                 = "ginkgo-linter:"
	suppressLengthAssertionWarning = suppressPrefix + "ignore-len-assert-warning"
	suppressNilAssertionWarning    = suppressPrefix + "ignore-nil-assert-warning"
)

type boolean bool

func (b *boolean) Set(value string) error {
	if b == nil {
		return errors.New("trying to set nil parameter")
	}
	switch strings.ToLower(value) {
	case "true":
		*b = true
	case "false":
		*b = false
	default:
		return errors.New(value + " is not a boolean value")

	}
	return nil
}

func (b boolean) String() string {
	if b {
		return "true"
	}
	return "false"
}

type ginkgoLinter struct {
	suppress *suppress
}

// NewAnalyzer returns an Analyzer - the package interface with nogo
func NewAnalyzer() *analysis.Analyzer {
	linter := ginkgoLinter{
		suppress: &suppress{
			Len: false,
			Nil: false,
		},
	}

	a := &analysis.Analyzer{
		Name: "ginkgolinter",
		Doc: `enforces standards of using ginkgo and gomega
currently, the linter searches for following:
* wrong length assertions. We want to assert the item rather than its length.
For example:
	Expect(len(x)).Should(Equal(1))
This should be replaced with:
	Expect(x)).Should(HavelLen(1))
	
* wrong nil assertions. We want to assert the item rather than a comparison result.
For example:
	Expect(x == nil).Should(BeTrue())
This should be replaced with:
	Expect(x).Should(BeNil())
	`,
		Run:              linter.run,
		RunDespiteErrors: true,
	}

	a.Flags.Init("ginkgolinter", flag.ExitOnError)
	a.Flags.Var(&linter.suppress.Len, "suppress-len-assertion", "suppress warning for wrong length assertions")
	a.Flags.Var(&linter.suppress.Nil, "suppress-nil-assertion", "suppress warning for wrong nil assertions")

	return a
}

//var Analyzer = NewAnalyzer()

// main assertion function
func (l *ginkgoLinter) run(pass *analysis.Pass) (interface{}, error) {
	if l.suppress.allTrue() {
		return nil, nil
	}

	for _, file := range pass.Files {
		fileSuppress := l.suppress.clone()

		cm := ast.NewCommentMap(pass.Fset, file, file.Comments)

		fileSuppress.updateFromFile(cm)
		if fileSuppress.allTrue() {
			continue
		}

		handler := getGomegaHandler(file)
		if handler == nil { // no gomega import => no use in gomega in this file; nothing to do here
			continue
		}

		ast.Inspect(file, func(n ast.Node) bool {

			stmt, ok := n.(*ast.ExprStmt)
			if !ok {
				return true
			}

			exprSuppress := fileSuppress.clone()

			if comments, ok := cm[stmt]; ok {
				exprSuppress.updateFromComment(comments)
			}

			// search for function calls
			assertionExp, ok := stmt.X.(*ast.CallExpr)
			if !ok {
				return true
			}

			assertionFunc, ok := assertionExp.Fun.(*ast.SelectorExpr)
			if !ok || !isAssertionFunc(assertionFunc.Sel.Name) {
				return true
			}

			actualExpr, ok := assertionFunc.X.(*ast.CallExpr)
			if !ok {
				return true
			}

			actualArg := getActualArg(actualExpr, handler)
			if actualArg == nil {
				return true
			}

			if !bool(exprSuppress.Len) && isActualIsLenFunc(actualArg) {
				oldExpr := goFmt(pass.Fset, assertionExp)

				return checkLengthMatcher(assertionExp, pass, handler, oldExpr)
			} else if !exprSuppress.Nil {
				if nilable, compOp := getNilableFromComparison(actualArg); nilable != nil {
					oldExpr := goFmt(pass.Fset, assertionExp)
					return checkNilMatcher(assertionExp, pass, nilable, handler, compOp == token.NEQ, oldExpr)
				}
			}
			return true

		})
	}
	return nil, nil
}

func getGomegaHandler(file *ast.File) gomegaHandler {
	for _, imp := range file.Imports {
		if imp.Path.Value != `"github.com/onsi/gomega"` {
			continue
		}

		switch name := imp.Name.String(); {
		case len(name) == 0:
			return nil // no gomega import; this file does not use gomega
		case name == ".":
			return gomegaDotHandler{}
		case name == "<nil>": // import with no local name
			return gomegaNameHandler("gomega")
		default:
			return gomegaNameHandler(name)
		}
	}

	return nil
}

// gomegaHandler provide different handling, depend on the way gomega was imported, whether
// in imported with "." name, custome name or without any name.
type gomegaHandler interface {
	getActualFuncName(*ast.CallExpr) (string, bool)
	replaceFunction(*ast.CallExpr, *ast.Ident)
}

// this handler is used when importing gomega with dot; i.e.
// import . "github.com/onsi/gomega"
//
type gomegaDotHandler struct{}

func (gomegaDotHandler) getActualFuncName(expr *ast.CallExpr) (string, bool) {
	actualFunc, ok := expr.Fun.(*ast.Ident)
	if !ok {
		return "", false
	}

	return actualFunc.Name, true
}

func (gomegaDotHandler) replaceFunction(caller *ast.CallExpr, newExpr *ast.Ident) {
	caller.Fun = newExpr
}

// this handler is used when importing gomega without name; i.e.
// import "github.com/onsi/gomega"
//
// or with a custom name; e.g.
// import customname "github.com/onsi/gomega"
//
type gomegaNameHandler string

func (g gomegaNameHandler) getActualFuncName(expr *ast.CallExpr) (string, bool) {
	selector, ok := expr.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", false
	}

	x, ok := selector.X.(*ast.Ident)
	if !ok {
		return "", false
	}

	if x.Name != string(g) {
		return "", false
	}

	return selector.Sel.Name, true
}

func (gomegaNameHandler) replaceFunction(caller *ast.CallExpr, newExpr *ast.Ident) {
	caller.Fun.(*ast.SelectorExpr).Sel = newExpr
}

// Check if the "actual" argument is a call to the golang built-in len() function
func isActualIsLenFunc(actualArg ast.Expr) bool {
	lenArgExp, ok := actualArg.(*ast.CallExpr)
	if !ok {
		return false
	}

	lenFunc, ok := lenArgExp.Fun.(*ast.Ident)
	return ok && lenFunc.Name == "len"
}

// Check if matcher function is in one of the patterns we want to avoid
func checkLengthMatcher(exp *ast.CallExpr, pass *analysis.Pass, handler gomegaHandler, oldExp string) bool {
	matcher, ok := exp.Args[0].(*ast.CallExpr)
	if !ok {
		return true
	}

	matcherFuncName, ok := handler.getActualFuncName(matcher)
	if !ok {
		return true
	}

	switch matcherFuncName {
	case "Equal":
		handleEqualMatcher(matcher, pass, exp, handler, oldExp)
		return false

	case "BeZero":
		handleBeZero(pass, exp, handler, oldExp)
		return false

	case "BeNumerically":
		return handleBeNumerically(matcher, pass, exp, handler, oldExp)

	case "Not":
		reverseAssertionFuncLogic(exp)
		exp.Args[0] = exp.Args[0].(*ast.CallExpr).Args[0]
		return checkLengthMatcher(exp, pass, handler, oldExp)

	default:
		return true
	}
}

// Check if matcher function is in one of the patterns we want to avoid
func checkNilMatcher(exp *ast.CallExpr, pass *analysis.Pass, nilable ast.Expr, handler gomegaHandler, notEqual bool, oldExp string) bool {
	matcher, ok := exp.Args[0].(*ast.CallExpr)
	if !ok {
		return true
	}

	matcherFuncName, ok := handler.getActualFuncName(matcher)
	if !ok {
		return true
	}

	switch matcherFuncName {
	case "Equal":
		handleEqualNilMatcher(matcher, pass, exp, handler, nilable, notEqual, oldExp)

	case "BeTrue":
		handleNilBeBoolMatcher(pass, exp, handler, nilable, notEqual, oldExp)

	case "BeFalse":
		reverseAssertionFuncLogic(exp)
		handleNilBeBoolMatcher(pass, exp, handler, nilable, notEqual, oldExp)

	case "Not":
		reverseAssertionFuncLogic(exp)
		exp.Args[0] = exp.Args[0].(*ast.CallExpr).Args[0]
		return checkNilMatcher(exp, pass, nilable, handler, notEqual, oldExp)

	default:
		return true
	}
	return false
}

// checks that the function is an assertion's actual function and return the "actual" parameter. If the function
// is not assertion's actual function, return nil.
func getActualArg(actualExpr *ast.CallExpr, handler gomegaHandler) ast.Expr {
	funcName, ok := handler.getActualFuncName(actualExpr)
	if !ok {
		return nil
	}

	switch funcName {
	case expect, omega:
		return actualExpr.Args[0]
	case expectWithOffset:
		return actualExpr.Args[1]
	default:
		return nil
	}
}

// Replace the len function call by its parameter, to create a fix suggestion
func replaceLenActualArg(actualExpr *ast.CallExpr, handler gomegaHandler) {
	name, ok := handler.getActualFuncName(actualExpr)
	if !ok {
		return
	}

	switch name {
	case expect, omega:
		arg := actualExpr.Args[0]
		if isActualIsLenFunc(arg) {
			// replace the len function call by its parameter, to create a fix suggestion
			actualExpr.Args[0] = arg.(*ast.CallExpr).Args[0]
		}
	case expectWithOffset:
		arg := actualExpr.Args[1]
		if isActualIsLenFunc(arg) {
			// replace the len function call by its parameter, to create a fix suggestion
			actualExpr.Args[1] = arg.(*ast.CallExpr).Args[0]
		}
	}
}

// Replace the nil comparison with the compared object, to create a fix suggestion
func replaceNilActualArg(actualExpr *ast.CallExpr, handler gomegaHandler, nilable ast.Expr) bool {
	actualFuncName, ok := handler.getActualFuncName(actualExpr)
	if !ok {
		return false
	}

	switch actualFuncName {
	case expect, omega:
		actualExpr.Args[0] = nilable
		return true

	case expectWithOffset:
		actualExpr.Args[1] = nilable
		return true

	default:
		return false
	}
}

// For the BeNumerically matcher, we want to avoid the assertion of length to be > 0 or >= 1, or just == number
func handleBeNumerically(matcher *ast.CallExpr, pass *analysis.Pass, exp *ast.CallExpr, handler gomegaHandler, oldExp string) bool {
	opExp, ok1 := matcher.Args[0].(*ast.BasicLit)
	valExp, ok2 := matcher.Args[1].(*ast.BasicLit)

	if ok1 && ok2 {
		op := opExp.Value
		val := valExp.Value

		caller := exp.Args[0].(*ast.CallExpr)
		if (op == `">"` && val == "0") || (op == `">="` && val == "1") {
			reverseAssertionFuncLogic(exp)
			handler.replaceFunction(caller, ast.NewIdent(beEmpty))
			exp.Args[0].(*ast.CallExpr).Args = nil
			reportLengthAssertion(pass, exp, handler, oldExp)
			return false
		} else if op == `"=="` {
			if val == "0" {
				handler.replaceFunction(caller, ast.NewIdent(beEmpty))
				exp.Args[0].(*ast.CallExpr).Args = nil
			} else {
				handler.replaceFunction(caller, ast.NewIdent(haveLen))
				exp.Args[0].(*ast.CallExpr).Args = []ast.Expr{valExp}
			}

			reportLengthAssertion(pass, exp, handler, oldExp)
			return false
		} else if op == `"!="` {
			reverseAssertionFuncLogic(exp)

			if val == "0" {
				handler.replaceFunction(caller, ast.NewIdent(beEmpty))
				exp.Args[0].(*ast.CallExpr).Args = nil
			} else {
				handler.replaceFunction(caller, ast.NewIdent(haveLen))
				exp.Args[0].(*ast.CallExpr).Args = []ast.Expr{valExp}
			}

			reportLengthAssertion(pass, exp, handler, oldExp)
			return false
		}
	}
	return true
}

func reverseAssertionFuncLogic(exp *ast.CallExpr) {
	assertionFunc := exp.Fun.(*ast.SelectorExpr).Sel
	assertionFunc.Name = reverseassertion.ChangeAssertionLogic(assertionFunc.Name)
}

func handleEqualMatcher(matcher *ast.CallExpr, pass *analysis.Pass, exp *ast.CallExpr, handler gomegaHandler, oldExp string) {
	equalTo, ok := matcher.Args[0].(*ast.BasicLit)
	caller := exp.Args[0].(*ast.CallExpr)
	if ok {
		if equalTo.Value == "0" {
			handler.replaceFunction(caller, ast.NewIdent(beEmpty))
			exp.Args[0].(*ast.CallExpr).Args = nil
		} else {
			handler.replaceFunction(caller, ast.NewIdent(haveLen))
			exp.Args[0].(*ast.CallExpr).Args = []ast.Expr{equalTo}
		}
	} else {
		handler.replaceFunction(caller, ast.NewIdent(haveLen))
		exp.Args[0].(*ast.CallExpr).Args = []ast.Expr{matcher.Args[0]}
	}
	reportLengthAssertion(pass, exp, handler, oldExp)
}

func handleBeZero(pass *analysis.Pass, exp *ast.CallExpr, handler gomegaHandler, oldExp string) {
	exp.Args[0].(*ast.CallExpr).Args = nil
	handler.replaceFunction(exp.Args[0].(*ast.CallExpr), ast.NewIdent(beEmpty))
	reportLengthAssertion(pass, exp, handler, oldExp)
}

func handleEqualNilMatcher(matcher *ast.CallExpr, pass *analysis.Pass, exp *ast.CallExpr, handler gomegaHandler, nilable ast.Expr, notEqual bool, oldExp string) {
	equalTo, ok := matcher.Args[0].(*ast.Ident)
	if !ok {
		return
	}

	if equalTo.Name == "false" {
		reverseAssertionFuncLogic(exp)
	} else if equalTo.Name != "true" {
		return
	}

	handler.replaceFunction(exp.Args[0].(*ast.CallExpr), ast.NewIdent(beNil))
	exp.Args[0].(*ast.CallExpr).Args = nil

	reportNilAssertion(pass, exp, handler, nilable, notEqual, oldExp)
}

func handleNilBeBoolMatcher(pass *analysis.Pass, exp *ast.CallExpr, handler gomegaHandler, nilable ast.Expr, notEqual bool, oldExp string) {
	handler.replaceFunction(exp.Args[0].(*ast.CallExpr), ast.NewIdent(beNil))
	exp.Args[0].(*ast.CallExpr).Args = nil

	reportNilAssertion(pass, exp, handler, nilable, notEqual, oldExp)
}

func isAssertionFunc(name string) bool {
	switch name {
	case "To", "ToNot", "NotTo", "Should", "ShouldNot":
		return true
	}
	return false
}

func reportLengthAssertion(pass *analysis.Pass, expr *ast.CallExpr, handler gomegaHandler, oldExpr string) {
	replaceLenActualArg(expr.Fun.(*ast.SelectorExpr).X.(*ast.CallExpr), handler)

	report(pass, expr, wrongLengthWarningTemplate, oldExpr)
}

func reportNilAssertion(pass *analysis.Pass, expr *ast.CallExpr, handler gomegaHandler, nilable ast.Expr, notEqual bool, oldExpr string) {
	changed := replaceNilActualArg(expr.Fun.(*ast.SelectorExpr).X.(*ast.CallExpr), handler, nilable)
	if !changed {
		return
	}

	if notEqual {
		reverseAssertionFuncLogic(expr)
	}

	report(pass, expr, wrongNilWarningTemplate, oldExpr)
}

func report(pass *analysis.Pass, expr *ast.CallExpr, messageTemplate, oldExpr string) {
	newExp := goFmt(pass.Fset, expr)
	pass.Report(analysis.Diagnostic{
		Pos:     expr.Pos(),
		Message: fmt.Sprintf(messageTemplate, newExp),
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: fmt.Sprintf("should replace %s with %s", oldExpr, newExp),
				TextEdits: []analysis.TextEdit{
					{
						Pos:     expr.Pos(),
						End:     expr.End(),
						NewText: []byte(newExp),
					},
				},
			},
		},
	})
}

func getNilableFromComparison(actualArg ast.Expr) (ast.Expr, token.Token) {
	bin, ok := actualArg.(*ast.BinaryExpr)
	if !ok {
		return nil, token.ILLEGAL
	}

	if bin.Op == token.EQL || bin.Op == token.NEQ {
		if isNil(bin.Y) {
			return bin.X, bin.Op
		} else if isNil(bin.X) {
			return bin.Y, bin.Op
		}
	}

	return nil, token.ILLEGAL
}

func isNil(expr ast.Expr) bool {
	nilObject, ok := expr.(*ast.Ident)
	return ok && nilObject.Name == "nil" && nilObject.Obj == nil
}

func goFmt(fset *token.FileSet, x ast.Expr) string {
	var b bytes.Buffer
	_ = printer.Fprint(&b, fset, x)
	return b.String()
}

type suppress struct {
	Len boolean
	Nil boolean
}

func (s suppress) allTrue() bool {
	return bool(s.Len) && bool(s.Nil)
}

func (s suppress) clone() suppress {
	return suppress{
		Len: s.Len,
		Nil: s.Nil,
	}
}

func (s *suppress) updateFromComment(commentGroup []*ast.CommentGroup) {
	for _, cmntList := range commentGroup {
		if s.allTrue() {
			break
		}

		for _, cmnt := range cmntList.List {
			commentLines := strings.Split(cmnt.Text, "\n")
			for _, comment := range commentLines {
				comment = strings.TrimPrefix(comment, "//")
				comment = strings.TrimPrefix(comment, "/*")
				comment = strings.TrimSuffix(comment, "*/")
				comment = strings.TrimSpace(comment)

				s.Len = s.Len || (comment == suppressLengthAssertionWarning)
				s.Nil = s.Nil || (comment == suppressNilAssertionWarning)
			}
		}
	}
}

func (s *suppress) updateFromFile(cm ast.CommentMap) {

	for key, commentGroup := range cm {
		if s.allTrue() {
			break
		}

		if _, ok := key.(*ast.GenDecl); ok {
			s.updateFromComment(commentGroup)
		}
	}
}
