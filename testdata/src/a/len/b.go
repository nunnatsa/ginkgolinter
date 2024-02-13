package len

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("", func() {
	It("", func() {
		s := []int{1, 2, 3, 4}
		Expect(len(s) == 4).Should(Equal(true))               // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\(s\)\.Should\(HaveLen\(4\)\). instead`
		Expect(4 == len(s)).Should(Equal(true))               // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\(s\)\.Should\(HaveLen\(4\)\). instead`
		Expect(len(s) == 4).WithOffset(1).Should(Equal(true)) // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\(s\)\.WithOffset\(1\)\.Should\(HaveLen\(4\)\). instead`
		Expect(4 == len(s)).WithOffset(1).Should(Equal(true)) // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\(s\)\.WithOffset\(1\)\.Should\(HaveLen\(4\)\). instead`
		s2 := make([]int, 4)
		copy(s2, s)
		Expect(len(s2) == len(s)).Should(Equal(true)) // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\(s2\)\.Should\(HaveLen\(len\(s\)\)\). instead`
		Expect(len(s) == 4).Should(Not(BeFalse()))    // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\(s\)\.Should\(HaveLen\(4\)\). instead`
	})
})
