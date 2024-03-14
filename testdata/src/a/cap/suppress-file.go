package cap

// ginkgo-linter:ignore-len-assert-warning

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

// ginkgo-linter:ignore-len-assert-warning

var _ = Describe("check cap", func() {
	It("should not allow expect cap", func() {
		slice := make([]int, 0, 10)
		gomega.Expect(cap(slice)).To(gomega.Equal(10))
		gomega.Expect(cap(slice)).ToNot(gomega.Equal(0))
		gomega.Expect(cap(slice)).ToNot(gomega.Equal(5))
		gomega.Expect(cap(slice)).ToNot(gomega.BeZero())
		gomega.Expect(cap(slice)).To(gomega.BeNumerically("==", 10))
		gomega.Expect(cap(slice)).To(gomega.BeNumerically("!=", 0))
		gomega.Expect(cap(slice)).ToNot(gomega.BeNumerically("==", 0))
		gomega.Expect(cap(slice)).To(gomega.BeNumerically("!=", 5))
		gomega.Expect(cap(slice)).ToNot(gomega.BeNumerically("==", 5))
		gomega.Expect(cap(slice)).To(gomega.BeNumerically(">", 0))
		gomega.Expect(cap(slice)).To(gomega.BeNumerically(">=", 1))
		gomega.Expect(slice).To(gomega.BeEmpty())
		gomega.Expect(slice).To(gomega.HaveCap(10))
	})

	It("should not allow comparison with cap", func() {
		slice := make([]int, 0, 10)
		gomega.Expect(cap(slice) == 10).To(gomega.BeTrue())        // want `wrong comparison assertion. Consider using .gomega\.Expect\(cap\(slice\)\)\.To\(gomega\.Equal\(10\)\). instead`
		gomega.Expect(cap(slice) == 10).To(gomega.Equal(true))     // want `wrong comparison assertion. Consider using .gomega\.Expect\(cap\(slice\)\)\.To\(gomega\.Equal\(10\)\). instead`
		gomega.Expect(cap(slice) != 10).To(gomega.BeFalse())       // want `wrong comparison assertion. Consider using .gomega\.Expect\(cap\(slice\)\)\.To\(gomega\.Equal\(10\)\). instead`
		gomega.Expect(cap(slice) != 10).To(gomega.Equal(false))    // want `wrong comparison assertion. Consider using .gomega\.Expect\(cap\(slice\)\)\.To\(gomega\.Equal\(10\)\). instead`
		gomega.Expect(cap(slice) == 10).ToNot(gomega.BeFalse())    // want `wrong comparison assertion. Consider using .gomega\.Expect\(cap\(slice\)\)\.To\(gomega\.Equal\(10\)\). instead`
		gomega.Expect(cap(slice) == 10).ToNot(gomega.Equal(false)) // want `wrong comparison assertion. Consider using .gomega\.Expect\(cap\(slice\)\)\.To\(gomega\.Equal\(10\)\). instead`
		gomega.Expect(cap(slice) != 10).ToNot(gomega.BeTrue())     // want `wrong comparison assertion. Consider using .gomega\.Expect\(cap\(slice\)\)\.To\(gomega\.Equal\(10\)\). instead`
		gomega.Expect(cap(slice) != 10).ToNot(gomega.Equal(true))  // want `wrong comparison assertion. Consider using .gomega\.Expect\(cap\(slice\)\)\.To\(gomega\.Equal\(10\)\). instead`
	})
})
