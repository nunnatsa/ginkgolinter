package to_and_not

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("test data for the ginkgo-linter", func() {
	It("Not and To for boolean assertions", func() {
		gomega.Expect(true).To(gomega.Not(gomega.BeFalse()))       // want `avoid double negative assertion. Consider using .gomega\.Expect\(true\)\.To\(gomega\.BeTrue\(\)\). instead`
		gomega.Expect(true).ToNot(gomega.Not(gomega.BeTrue()))     // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Expect\(true\)\.To\(gomega\.BeTrue\(\)\). instead`
		gomega.Expect(true).NotTo(gomega.Not(gomega.BeTrue()))     // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Expect\(true\)\.To\(gomega\.BeTrue\(\)\). instead`
		gomega.Expect(true).ShouldNot(gomega.Not(gomega.BeTrue())) // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Expect\(true\)\.Should\(gomega\.BeTrue\(\)\). instead`

		gomega.Eventually(func() bool { return true }).Should(gomega.Not(gomega.BeFalse()))   // want `avoid double negative assertion\. Consider using .gomega\.Eventually\(func\(\) bool { return true }\).Should\(gomega\.BeTrue\(\)\). instead`
		gomega.Eventually(func() bool { return true }).ShouldNot(gomega.Not(gomega.BeTrue())) // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Eventually\(func\(\) bool { return true }\)\.Should\(gomega\.BeTrue\(\)\). instead`
		gomega.Eventually(func() bool { return true }).To(gomega.Not(gomega.BeFalse()))       // want `avoid double negative assertion\. Consider using .gomega\.Eventually\(func\(\) bool { return true }\)\.To\(gomega\.BeTrue\(\)\). instead`
		gomega.Eventually(func() bool { return true }).ToNot(gomega.Not(gomega.BeTrue()))     // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Eventually\(func\(\) bool { return true }\)\.To\(gomega\.BeTrue\(\)\). instead`
		gomega.Eventually(func() bool { return true }).NotTo(gomega.Not(gomega.BeTrue()))     // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Eventually\(func\(\) bool { return true }\)\.To\(gomega\.BeTrue\(\)\). instead`
	})

	It("Not and To for other matchers", func() {
		s := []int{1, 2, 3}
		gomega.Expect(s).To(gomega.Not(gomega.BeEmpty()))                                      // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Expect\(s\)\.ToNot\(gomega\.BeEmpty\(\)\). instead`
		gomega.Expect(s).ToNot(gomega.Not(gomega.HaveLen(3)))                                  // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Expect\(s\)\.To\(gomega\.HaveLen\(3\)\). instead`
		gomega.Expect(s).To(gomega.Not(gomega.Not(gomega.HaveLen(3))))                         // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Expect\(s\)\.To\(gomega\.HaveLen\(3\)\). instead`
		gomega.Expect(s).To(gomega.Not(gomega.Not(gomega.Not(gomega.HaveLen(3)))))             // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Expect\(s\)\.ToNot\(gomega\.HaveLen\(3\)\). instead`
		gomega.Expect(s).To(gomega.Not(gomega.Not(gomega.Not(gomega.Not(gomega.HaveLen(3)))))) // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Expect\(s\)\.To\(gomega\.HaveLen\(3\)\). instead`
		gomega.Expect(s).NotTo(gomega.Not(gomega.HaveLen(3)))                                  // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Expect\(s\)\.To\(gomega\.HaveLen\(3\)\). instead`
		gomega.Expect(s).ShouldNot(gomega.Not(gomega.HaveLen(3)))                              // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Expect\(s\)\.Should\(gomega\.HaveLen\(3\)\). instead`

		gomega.Eventually(func() []int { return s }).Should(gomega.Not(gomega.HaveLen(3)))    // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Eventually\(func\(\) \[\]int { return s }\).ShouldNot\(gomega\.HaveLen\(3\)\). instead`
		gomega.Eventually(func() []int { return s }).ShouldNot(gomega.Not(gomega.HaveLen(3))) // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Eventually\(func\(\) \[\]int { return s }\).Should\(gomega\.HaveLen\(3\)\). instead`
		gomega.Eventually(func() []int { return s }).To(gomega.Not(gomega.HaveLen(3)))        // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Eventually\(func\(\) \[\]int { return s }\).ToNot\(gomega\.HaveLen\(3\)\). instead`
		gomega.Eventually(func() []int { return s }).ToNot(gomega.Not(gomega.HaveLen(3)))     // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Eventually\(func\(\) \[\]int { return s }\).To\(gomega\.HaveLen\(3\)\). instead`
		gomega.Eventually(func() []int { return s }).NotTo(gomega.Not(gomega.HaveLen(3)))     // want `simplify negation by removing the 'Not' matcher\. Consider using .gomega\.Eventually\(func\(\) \[\]int { return s }\).To\(gomega\.HaveLen\(3\)\). instead`
	})
})
