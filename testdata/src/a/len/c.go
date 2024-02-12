package len

import (
	. "github.com/onsi/ginkgo/v2"
	g "github.com/onsi/gomega"
)

var _ = Describe("", func() {
	It("", func() {
		s := []int{1, 2, 3, 4}
		g.Expect(len(s) == 4).Should(g.Equal(true))               // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(s\)\.Should\(g\.HaveLen\(4\)\). instead`
		g.Expect(4 == len(s)).Should(g.Equal(true))               // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(s\)\.Should\(g\.HaveLen\(4\)\). instead`
		g.Expect(4 == len(s)).Should(g.BeTrue())                  // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(s\)\.Should\(g\.HaveLen\(4\)\). instead`
		g.Expect(len(s) == 4).WithOffset(1).Should(g.Equal(true)) // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(s\)\.WithOffset\(1\)\.Should\(g\.HaveLen\(4\)\). instead`
		g.Expect(4 == len(s)).WithOffset(1).Should(g.Equal(true)) // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(s\)\.WithOffset\(1\)\.Should\(g\.HaveLen\(4\)\). instead`
	})
	It("", func() {
		s := []int{1, 2, 3, 4}
		g.Expect(len(s) != 4).Should(g.Equal(false))                 // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(s\)\.Should\(g\.HaveLen\(4\)\). instead`
		g.Expect(4 != len(s)).ShouldNot(g.Equal(true))               // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(s\)\.Should\(g\.HaveLen\(4\)\). instead`
		g.Expect(4 != len(s)).Should(g.BeFalse())                    // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(s\)\.Should\(g\.HaveLen\(4\)\). instead`
		g.Expect(len(s) != 4).WithOffset(1).Should(g.Equal(false))   // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(s\)\.WithOffset\(1\)\.Should\(g\.HaveLen\(4\)\). instead`
		g.Expect(4 != len(s)).WithOffset(1).ShouldNot(g.Equal(true)) // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(s\)\.WithOffset\(1\)\.Should\(g\.HaveLen\(4\)\). instead`
		g.Expect(len(s) == 5).WithOffset(1).Should(g.Equal(false))   // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(s\)\.WithOffset\(1\)\.ShouldNot\(g\.HaveLen\(5\)\). instead`
	})
})
