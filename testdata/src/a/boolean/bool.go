package boolean

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("boolean warnings", func() {
	t := true
	f := false

	It("check Equal(true/false)", func() {
		Expect(t).To(Equal(true))                        // want `ginkgo-linter: wrong boolean assertion\. Consider using .Expect\(t\)\.To\(BeTrue\(\)\). instead`
		Expect(f).To(Equal(false))                       // want `ginkgo-linter: wrong boolean assertion\. Consider using .Expect\(f\)\.To\(BeFalse\(\)\). instead`
		ExpectWithOffset(2, t).Should(Not(Equal(false))) // want `ginkgo-linter: wrong boolean assertion\. Consider using .ExpectWithOffset\(2, t\)\.Should\(BeTrue\(\)\). instead`
		ExpectWithOffset(2, t).ShouldNot(Equal(false))   // want `ginkgo-linter: wrong boolean assertion\. Consider using .ExpectWithOffset\(2, t\)\.Should\(BeTrue\(\)\). instead`
		Expect(t).WithOffset(2).ToNot(Equal(false))      // want `ginkgo-linter: wrong boolean assertion\. Consider using .Expect\(t\)\.WithOffset\(2\)\.To\(BeTrue\(\)\). instead`
		Expect(t).NotTo(Equal(false))                    // want `ginkgo-linter: wrong boolean assertion\. Consider using .Expect\(t\)\.To\(BeTrue\(\)\). instead`
		Expect(f).WithOffset(1).ToNot(Equal(true))       // want `ginkgo-linter: wrong boolean assertion\. Consider using .Expect\(f\)\.WithOffset\(1\)\.ToNot\(BeTrue\(\)\). instead`
	})

	It("check double negative", func() {
		Expect(t).ToNot(BeFalse())                   // want `ginkgo-linter: avoid double negative assertion\. Consider using .Expect\(t\)\.To\(BeTrue\(\)\). instead`
		Expect(t).NotTo(BeFalse())                   // want `ginkgo-linter: avoid double negative assertion\. Consider using .Expect\(t\)\.To\(BeTrue\(\)\). instead`
		Expect(t).WithOffset(2).ShouldNot(BeFalse()) // want `ginkgo-linter: avoid double negative assertion\. Consider using .Expect\(t\)\.WithOffset\(2\)\.Should\(BeTrue\(\)\). instead`
	})
})
