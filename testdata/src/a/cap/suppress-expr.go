package cap

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("check cap", func() {
	It("should not allow expect cap", func() {
		slice := make([]int, 0, 10)
		// ginkgo-linter:ignore-len-assert-warning
		Expect(cap(slice)).To(Equal(10))
		// ginkgo-linter:ignore-len-assert-warning
		Expect(cap(slice)).ToNot(BeZero())
		// ginkgo-linter:ignore-len-assert-warning
		Expect(cap(slice)).To(BeNumerically("==", 10))
		Expect(cap(slice)).To(Equal(10))               // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.To\(HaveCap\(10\)\). instead`
		Expect(cap(slice)).ToNot(BeZero())             // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.ToNot\(HaveCap\(0\)\). instead`
		Expect(cap(slice)).To(BeNumerically("==", 10)) // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.To\(HaveCap\(10\)\). instead`
	})

	It("should not allow comparison with cap", func() {
		slice := make([]int, 0, 10)
		// ginkgo-linter:ignore-len-assert-warning
		Expect(cap(slice) == 10).To(BeTrue()) // want `wrong comparison assertion. Consider using .Expect\(cap\(slice\)\)\.To\(Equal\(10\)\). instead`
		Expect(cap(slice) == 10).To(BeTrue()) // want `ginkgo-linter: wrong cap assertion. Consider using .Expect\(slice\)\.To\(HaveCap\(10\)\). instead`
	})
})
