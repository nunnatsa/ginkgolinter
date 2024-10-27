package rules

import (
	"github.com/nunnatsa/ginkgolinter/internal/expression"
	"github.com/nunnatsa/ginkgolinter/internal/expression/actual"
	"github.com/nunnatsa/ginkgolinter/internal/expression/matcher"
	"github.com/nunnatsa/ginkgolinter/internal/expression/value"
	"github.com/nunnatsa/ginkgolinter/internal/reports"
	"github.com/nunnatsa/ginkgolinter/types"
)

type ErrorEqualNilRule struct{}

func (ErrorEqualNilRule) isApplied(gexp *expression.GomegaExpression, config types.Config) bool {
	return !bool(config.SuppressErr) &&
		gexp.Actual.Arg.ArgType().Is(actual.ErrorTypeArgType) &&
		gexp.Matcher.GetMatcherInfo().Type().Is(matcher.BeNilMatcherType|matcher.EqualNilMatcherType)
}

func (r ErrorEqualNilRule) Apply(gexp *expression.GomegaExpression, config types.Config, reportBuilder *reports.Builder) bool {
	if !r.isApplied(gexp, config) {
		return false
	}

	if v, ok := gexp.Actual.Arg.(value.Valuer); ok && v.IsFunc() || gexp.Actual.Arg.ArgType().Is(actual.ErrFuncActualArgType) {
		gexp.SetMatcherSucceed()
	} else {
		gexp.ReverseAssertionFuncLogic()
		gexp.SetMatcherHaveOccurred()
	}

	reportBuilder.AddIssue(true, wrongErrWarningTemplate)

	return true
}
