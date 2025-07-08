package rules

import (
	"github.com/nunnatsa/ginkgolinter/internal/expression"
	"github.com/nunnatsa/ginkgolinter/internal/reports"
	"github.com/nunnatsa/ginkgolinter/types"
)

const assertionDescriptionWarningTemplate = "missing assertion description"

type AssertionDescriptionRule struct{}

func (r *AssertionDescriptionRule) Apply(gexp *expression.GomegaExpression, config types.Config, reportBuilder *reports.Builder) bool {
	if !r.isApplied(gexp, config) {
		return false
	}

	reportBuilder.AddIssue(false, assertionDescriptionWarningTemplate)
	return true
}

func (r *AssertionDescriptionRule) isApplied(gexp *expression.GomegaExpression, config types.Config) bool {
	if !config.ForceAssertionDescription {
		return false
	}

	return !r.hasDescription(gexp)
}

func (r *AssertionDescriptionRule) hasDescription(gexp *expression.GomegaExpression) bool {
	return gexp.HasDescription()
}
