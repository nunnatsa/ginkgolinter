package rules

import (
	"strconv"

	"github.com/nunnatsa/ginkgolinter/config"
	"github.com/nunnatsa/ginkgolinter/internal/expression"
	"github.com/nunnatsa/ginkgolinter/internal/gomegainfo"
	"github.com/nunnatsa/ginkgolinter/internal/reports"
)

const missingAssertionMessage = `%smissing assertion method. Expected %s`

type MissingAssertionRule struct{}

func (r MissingAssertionRule) isApplied(gexp *expression.GomegaExpression) bool {
	return gexp.IsMissingAssertion()
}

// MissingAssertionRule checks if the assertion method is missing. In this case, the test does not make any assertion.
// This is mostly relevant for the async actual methods, that tend to be longer, and so harder to spot the missing assertion
// by just reading the test code.
//
// Examples:
//
//		// Bad:
//		Expect(x)
//		Eventually(func() error {
//			return nil
//	 	})
//
//		// Good:
//		Expect(x).To(Equal(42))
//		Eventually(func() error {
//			return nil
//		}).Should(Succeed())
func (r MissingAssertionRule) Apply(gexp *expression.GomegaExpression, _ config.Config, reportBuilder *reports.Builder) bool {
	if !r.isApplied(gexp) {
		return false
	}

	actualMethodName := gexp.GetActualFuncName()
	allowed := gomegainfo.GetAllowedAssertionMethods(actualMethodName)
	var prefix string
	if actualMethodName != "" {
		prefix = strconv.Quote(actualMethodName) + ": "
	}
	if allowed == "" {
		allowed = `one of "To/NotTo/To" (for Expect assertions) or "Should/ShouldNot" (for Eventually/Consistently assertions)`
	}
	reportBuilder.AddIssue(false, missingAssertionMessage, prefix, allowed)

	return true
}
