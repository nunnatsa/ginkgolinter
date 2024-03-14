package cap

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("mix len and cap", func() {
	slice := make([]int, 0)
	It("compare len with cap", func() {
		Expect(cap(slice) == len(slice)).To(BeTrue()) // want `wrong cap assertion. Consider using .Expect\(slice\)\.To\(HaveCap\(len\(slice\)\)\). instead`
		Expect(len(slice) == cap(slice)).To(BeTrue()) // want `wrong length assertion. Consider using .Expect\(slice\)\.To\(HaveLen\(cap\(slice\)\)\). instead`
	})
})
