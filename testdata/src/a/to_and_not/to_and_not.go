package to_and_not

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("test data for the ginkgo-linter", func() {
	It("Not and To for boolean assertions", func() {
		Expect(true).To(Not(BeFalse()))       // want `avoid double negative assertion. Consider using .Expect\(true\)\.To\(BeTrue\(\)\). instead`
		Expect(true).ToNot(Not(BeTrue()))     // want `simplify negation by removing the 'Not' matcher\. Consider using .Expect\(true\)\.To\(BeTrue\(\)\). instead`
		Expect(true).NotTo(Not(BeTrue()))     // want `simplify negation by removing the 'Not' matcher\. Consider using .Expect\(true\)\.To\(BeTrue\(\)\). instead`
		Expect(true).ShouldNot(Not(BeTrue())) // want `simplify negation by removing the 'Not' matcher\. Consider using .Expect\(true\)\.Should\(BeTrue\(\)\). instead`

		Eventually(func() bool { return true }).Should(Not(BeFalse()))   // want `avoid double negative assertion\. Consider using .Eventually\(func\(\) bool { return true }\).Should\(BeTrue\(\)\). instead`
		Eventually(func() bool { return true }).ShouldNot(Not(BeTrue())) // want `simplify negation by removing the 'Not' matcher\. Consider using .Eventually\(func\(\) bool { return true }\)\.Should\(BeTrue\(\)\). instead`
		Eventually(func() bool { return true }).To(Not(BeFalse()))       // want `avoid double negative assertion\. Consider using .Eventually\(func\(\) bool { return true }\)\.To\(BeTrue\(\)\). instead`
		Eventually(func() bool { return true }).ToNot(Not(BeTrue()))     // want `simplify negation by removing the 'Not' matcher\. Consider using .Eventually\(func\(\) bool { return true }\)\.To\(BeTrue\(\)\). instead`
		Eventually(func() bool { return true }).NotTo(Not(BeTrue()))     // want `simplify negation by removing the 'Not' matcher\. Consider using .Eventually\(func\(\) bool { return true }\)\.To\(BeTrue\(\)\). instead`
	})

	It("Not and To for other matchers", func() {
		s := []int{1, 2, 3}
		Expect(s).To(Not(BeEmpty()))                 // want `simplify negation by removing the 'Not' matcher\. Consider using .Expect\(s\)\.ToNot\(BeEmpty\(\)\). instead`
		Expect(s).ToNot(Not(HaveLen(3)))             // want `simplify negation by removing the 'Not' matcher\. Consider using .Expect\(s\)\.To\(HaveLen\(3\)\). instead`
		Expect(s).To(Not(Not(HaveLen(3))))           // want `simplify negation by removing the 'Not' matcher\. Consider using .Expect\(s\)\.To\(HaveLen\(3\)\). instead`
		Expect(s).To(Not(Not(Not(HaveLen(3)))))      // want `simplify negation by removing the 'Not' matcher\. Consider using .Expect\(s\)\.ToNot\(HaveLen\(3\)\). instead`
		Expect(s).To(Not(Not(Not(Not(HaveLen(3)))))) // want `simplify negation by removing the 'Not' matcher\. Consider using .Expect\(s\)\.To\(HaveLen\(3\)\). instead`
		Expect(s).NotTo(Not(HaveLen(3)))             // want `simplify negation by removing the 'Not' matcher\. Consider using .Expect\(s\)\.To\(HaveLen\(3\)\). instead`
		Expect(s).ShouldNot(Not(HaveLen(3)))         // want `simplify negation by removing the 'Not' matcher\. Consider using .Expect\(s\)\.Should\(HaveLen\(3\)\). instead`

		Eventually(func() []int { return s }).Should(Not(HaveLen(3)))    // want `simplify negation by removing the 'Not' matcher\. Consider using .Eventually\(func\(\) \[\]int { return s }\).ShouldNot\(HaveLen\(3\)\). instead`
		Eventually(func() []int { return s }).ShouldNot(Not(HaveLen(3))) // want `simplify negation by removing the 'Not' matcher\. Consider using .Eventually\(func\(\) \[\]int { return s }\).Should\(HaveLen\(3\)\). instead`
		Eventually(func() []int { return s }).To(Not(HaveLen(3)))        // want `simplify negation by removing the 'Not' matcher\. Consider using .Eventually\(func\(\) \[\]int { return s }\).ToNot\(HaveLen\(3\)\). instead`
		Eventually(func() []int { return s }).ToNot(Not(HaveLen(3)))     // want `simplify negation by removing the 'Not' matcher\. Consider using .Eventually\(func\(\) \[\]int { return s }\).To\(HaveLen\(3\)\). instead`
		Eventually(func() []int { return s }).NotTo(Not(HaveLen(3)))     // want `simplify negation by removing the 'Not' matcher\. Consider using .Eventually\(func\(\) \[\]int { return s }\).To\(HaveLen\(3\)\). instead`
	})
})
