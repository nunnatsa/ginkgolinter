package rules

import (
	"go/token"

	"github.com/nunnatsa/ginkgolinter/internal/expression"
	"github.com/nunnatsa/ginkgolinter/internal/expression/actual"
	"github.com/nunnatsa/ginkgolinter/internal/expression/matcher"
	"github.com/nunnatsa/ginkgolinter/internal/reports"
	"github.com/nunnatsa/ginkgolinter/types"
)

const wrongCapWarningTemplate = "wrong cap assertion"

// CapRule does not allow using the cap() function in actual with numeric comparison.
// it suggests to use the HaveLen matcher, instead.
type CapRule struct{}

func (r *CapRule) Apply(gexp *expression.GomegaExpression, config types.Config, reportBuilder *reports.Builder) bool {

	if !r.isApplied(gexp, config) {
		return false
	}

	if r.fixExpression(gexp) {
		reportBuilder.AddIssue(true, wrongCapWarningTemplate)
		return true
	}
	return false
}

func (r *CapRule) isApplied(gexp *expression.GomegaExpression, config types.Config) bool {
	if config.SuppressLen {
		return false
	}

	matcherType := gexp.Matcher.GetMatcherInfo().Type()
	if gexp.Actual.Arg.ArgType().Is(actual.CapFuncActualArgType) {
		if matcherType.Is(matcher.EqualMatcherType | matcher.BeZeroMatcherType) {
			return true
		}

		if matcherType.Is(matcher.BeNumericallyMatcherType) {
			mtchr := gexp.Matcher.GetMatcherInfo().(*matcher.BeNumericallyMatcher)
			return mtchr.GetOp() == token.EQL || mtchr.GetOp() == token.NEQ || matcherType.Is(matcher.EqualZero|matcher.GreaterThanZero)
		}
	}

	if gexp.Actual.Arg.ArgType().Is(actual.CapComparisonActualArgType) && matcherType.Is(matcher.BeTrueMatcherType|matcher.BeFalseMatcherType|matcher.EqualBoolValueMatcherType) {
		return true
	}

	return false
}

func (r *CapRule) fixExpression(gexp *expression.GomegaExpression) bool {
	if gexp.Actual.Arg.ArgType().Is(actual.CapFuncActualArgType) {
		return r.fixEqual(gexp)
	}

	if gexp.Actual.Arg.ArgType().Is(actual.CapComparisonActualArgType) {
		return r.fixComparison(gexp)
	}

	return false
}

func (r *CapRule) fixEqual(gexp *expression.GomegaExpression) bool {
	matcherInfo := gexp.Matcher.GetMatcherInfo()
	switch mtchr := matcherInfo.(type) {
	case *matcher.EqualMatcher:
		gexp.SetMatcherCap(mtchr.GetValueExpr())

	case *matcher.BeZeroMatcher:
		gexp.SetMatcherCapZero()

	case *matcher.BeNumericallyMatcher:
		if !r.handleBeNumerically(gexp, mtchr) {
			return false
		}

	default:
		return false
	}

	gexp.ReplaceActualWithItsFirstArg()

	return true
}

func (r *CapRule) fixComparison(gexp *expression.GomegaExpression) bool {
	actl := gexp.Actual.Arg.(*actual.FuncComparisonPayload)
	if op := actl.GetOp(); op == token.NEQ {
		gexp.ReverseAssertionFuncLogic()
	} else if op != token.EQL {
		return false
	}

	gexp.SetMatcherCap(actl.GetValueExpr())
	gexp.ReplaceActual(actl.GetFuncArg())

	if gexp.Matcher.GetMatcherInfo().Type().Is(matcher.BoolValueFalse) {
		gexp.ReverseAssertionFuncLogic()
	}

	return true
}

func (r *CapRule) handleBeNumerically(gexp *expression.GomegaExpression, matcher *matcher.BeNumericallyMatcher) bool {
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
		return false
	}

	return true
}
