package cap

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("check cap", func() {
	It("should not allow expect cap", func() {
		slice := make([]int, 0, 10)
		Expect(cap(slice)).To(Equal(10))                 // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.To\(HaveCap\(10\)\). instead`
		Expect(cap(slice)).ToNot(Equal(0))               // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.ToNot\(HaveCap\(0\)\). instead`
		Expect(cap(slice)).ToNot(Equal(5))               // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.ToNot\(HaveCap\(5\)\). instead`
		Expect(cap(slice)).ToNot(BeZero())               // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.ToNot\(HaveCap\(0\)\). instead`
		Expect(cap(slice)).To(BeNumerically("==", 10))   // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.To\(HaveCap\(10\)\). instead`
		Expect(cap(slice)).To(BeNumerically("!=", 0))    // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.ToNot\(HaveCap\(0\)\). instead`
		Expect(cap(slice)).ToNot(BeNumerically("==", 0)) // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.ToNot\(HaveCap\(0\)\). instead`
		Expect(cap(slice)).To(BeNumerically("!=", 5))    // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.ToNot\(HaveCap\(5\)\). instead`
		Expect(cap(slice)).ToNot(BeNumerically("==", 5)) // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.ToNot\(HaveCap\(5\)\). instead`
		Expect(cap(slice)).To(BeNumerically(">", 0))     // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.ToNot\(HaveCap\(0\)\). instead`
		Expect(cap(slice)).To(BeNumerically(">=", 1))    // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.ToNot\(HaveCap\(0\)\). instead`
		Expect(slice).To(BeEmpty())
		Expect(slice).To(HaveCap(10))
	})

	It("should not allow comparison with cap", func() {
		slice := make([]int, 0, 10)
		Expect(cap(slice) == 10).To(BeTrue())        // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.To\(HaveCap\(10\)\). instead`
		Expect(cap(slice) == 10).To(Equal(true))     // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.To\(HaveCap\(10\)\). instead`
		Expect(cap(slice) != 10).To(BeFalse())       // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.To\(HaveCap\(10\)\). instead`
		Expect(cap(slice) != 10).To(Equal(false))    // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.To\(HaveCap\(10\)\). instead`
		Expect(cap(slice) == 10).ToNot(BeFalse())    // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.To\(HaveCap\(10\)\). instead`
		Expect(cap(slice) == 10).ToNot(Equal(false)) // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.To\(HaveCap\(10\)\). instead`
		Expect(cap(slice) != 10).ToNot(BeTrue())     // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.To\(HaveCap\(10\)\). instead`
		Expect(cap(slice) != 10).ToNot(Equal(true))  // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.To\(HaveCap\(10\)\). instead`
		Expect(10 != cap(slice)).ToNot(Equal(true))  // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.To\(HaveCap\(10\)\). instead`
	})
})
