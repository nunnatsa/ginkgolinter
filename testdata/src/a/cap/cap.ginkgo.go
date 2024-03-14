package cap

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("check cap", func() {
	It("should not allow expect cap", func() {
		slice := make([]int, 0, 10)
		gomega.Expect(cap(slice)).To(gomega.Equal(10))                 // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.To\(gomega\.HaveCap\(10\)\). instead`
		gomega.Expect(cap(slice)).ToNot(gomega.Equal(0))               // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.ToNot\(gomega\.HaveCap\(0\)\). instead`
		gomega.Expect(cap(slice)).ToNot(gomega.Equal(5))               // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.ToNot\(gomega\.HaveCap\(5\)\). instead`
		gomega.Expect(cap(slice)).ToNot(gomega.BeZero())               // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.ToNot\(gomega\.HaveCap\(0\)\). instead`
		gomega.Expect(cap(slice)).To(gomega.BeNumerically("==", 10))   // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.To\(gomega\.HaveCap\(10\)\). instead`
		gomega.Expect(cap(slice)).To(gomega.BeNumerically("!=", 0))    // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.ToNot\(gomega\.HaveCap\(0\)\). instead`
		gomega.Expect(cap(slice)).ToNot(gomega.BeNumerically("==", 0)) // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.ToNot\(gomega\.HaveCap\(0\)\). instead`
		gomega.Expect(cap(slice)).To(gomega.BeNumerically("!=", 5))    // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.ToNot\(gomega\.HaveCap\(5\)\). instead`
		gomega.Expect(cap(slice)).ToNot(gomega.BeNumerically("==", 5)) // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.ToNot\(gomega\.HaveCap\(5\)\). instead`
		gomega.Expect(cap(slice)).To(gomega.BeNumerically(">", 0))     // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.ToNot\(gomega\.HaveCap\(0\)\). instead`
		gomega.Expect(cap(slice)).To(gomega.BeNumerically(">=", 1))    // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.ToNot\(gomega\.HaveCap\(0\)\). instead`
		gomega.Expect(slice).To(gomega.BeEmpty())
		gomega.Expect(slice).To(gomega.HaveCap(10))
	})

	It("should not allow comparison with cap", func() {
		slice := make([]int, 0, 10)
		gomega.Expect(cap(slice) == 10).To(gomega.BeTrue())        // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.To\(gomega\.HaveCap\(10\)\). instead`
		gomega.Expect(cap(slice) == 10).To(gomega.Equal(true))     // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.To\(gomega\.HaveCap\(10\)\). instead`
		gomega.Expect(cap(slice) != 10).To(gomega.BeFalse())       // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.To\(gomega\.HaveCap\(10\)\). instead`
		gomega.Expect(cap(slice) != 10).To(gomega.Equal(false))    // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.To\(gomega\.HaveCap\(10\)\). instead`
		gomega.Expect(cap(slice) == 10).ToNot(gomega.BeFalse())    // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.To\(gomega\.HaveCap\(10\)\). instead`
		gomega.Expect(cap(slice) == 10).ToNot(gomega.Equal(false)) // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.To\(gomega\.HaveCap\(10\)\). instead`
		gomega.Expect(cap(slice) != 10).ToNot(gomega.BeTrue())     // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.To\(gomega\.HaveCap\(10\)\). instead`
		gomega.Expect(cap(slice) != 10).ToNot(gomega.Equal(true))  // want `ginkgo-linter: wrong cap assertion. Consider using .gomega\.Expect\(slice\)\.To\(gomega\.HaveCap\(10\)\). instead`
	})
})
