package boolean

import (
	"a/wrappers"

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

		// Wrappers which are named exactly like their wrapped methods are treated like the wrapped method.
		// The suggestion is just a guess, but perhaps not a bad one.
		Expect(t).WithOffset(2).ToNot(wrappers.Equal(false)) // want `ginkgo-linter: wrong boolean assertion\. Consider using .Expect\(t\)\.WithOffset\(2\)\.To\(wrappers.BeTrue\(\)\). instead`
		Expect(t).WithOffset(2).ToNot(wrappers.MyEqual(false))
	})

	It("check double negative", func() {
		Expect(t).ToNot(BeFalse())                   // want `ginkgo-linter: avoid double negative assertion\. Consider using .Expect\(t\)\.To\(BeTrue\(\)\). instead`
		Expect(t).NotTo(BeFalse())                   // want `ginkgo-linter: avoid double negative assertion\. Consider using .Expect\(t\)\.To\(BeTrue\(\)\). instead`
		Expect(t).WithOffset(2).ShouldNot(BeFalse()) // want `ginkgo-linter: avoid double negative assertion\. Consider using .Expect\(t\)\.WithOffset\(2\)\.Should\(BeTrue\(\)\). instead`

		wrappers.MyExpect(t).ToNot(BeFalse())               // want `ginkgo-linter: avoid double negative assertion\. Consider using .wrappers\.MyExpect\(t\)\.To\(BeTrue\(\)\). instead`
		wrappers.MyExpect(t).WithOffset(2).ToNot(BeFalse()) // want `ginkgo-linter: avoid double negative assertion\. Consider using .wrappers\.MyExpect\(t\)\.WithOffset\(2\)\.To\(BeTrue\(\)\). instead`
	})
})
