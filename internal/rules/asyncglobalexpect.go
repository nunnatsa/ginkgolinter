package rules

import (
	"github.com/nunnatsa/ginkgolinter/gomegaanalyzer"
	"github.com/nunnatsa/ginkgolinter/internal/expression"
	"github.com/nunnatsa/ginkgolinter/internal/reports"
	"github.com/nunnatsa/ginkgolinter/types"
)

type AsyncGlobalExpect struct {
	res *gomegaanalyzer.Result
}

func NewAsyncGlobalExpect(res *gomegaanalyzer.Result) *AsyncGlobalExpect {
	return &AsyncGlobalExpect{res: res}
}

func (r AsyncGlobalExpect) Apply(gexp *expression.GomegaExpression, config types.Config, reportBuilder *reports.Builder) bool {
	if !bool(config.ForbidAsyncGlobalAssertions) || gexp.IsAsync() || gexp.IsUsingGomegaVar() {
		return false
	}

	if pos, inScope := r.res.InAsyncScope(gexp.Pos()); inScope {
		reportBuilder.AddIssue(false, "using a global %s in an async assertion at %v; either use a Gomega variable, or StopTrying(...).Now()", gexp.GetActualFuncName(), pos)
	}

	return false
}
