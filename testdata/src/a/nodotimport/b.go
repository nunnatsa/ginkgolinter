package nodotimport

import (
	g1 "github.com/onsi/ginkgo/v2"
	g2 "github.com/onsi/gomega"
)

var _ = g1.Describe("no dot import", func() {
	g1.It("should ignore length warning", func() {
		const length = 3
		g2.Expect(len("abc")).Should(g2.Equal(3))                             // want `ginkgo-linter: wrong length assertion; consider using .g2\.Expect\("abc"\)\.Should\(g2\.HaveLen\(3\)\). instead`
		g2.Expect(len("abc")).Should(g2.Equal(length))                        // want `ginkgo-linter: wrong length assertion; consider using .g2\.Expect\("abc"\)\.Should\(g2\.HaveLen\(length\)\). instead`
		g2.Expect(len("")).Should(g2.Equal(0))                                // want `ginkgo-linter: wrong length assertion; consider using .g2\.Expect\(""\)\.Should\(g2\.BeEmpty\(\)\). instead`
		g2.Expect(len("abc")).ShouldNot(g2.BeZero())                          // want `ginkgo-linter: wrong length assertion; consider using .g2\.Expect\("abc"\)\.ShouldNot\(g2\.BeEmpty\(\)\). instead`
		g2.Expect(len("abc")).Should(g2.BeNumerically(">", 0))                // want `ginkgo-linter: wrong length assertion; consider using .g2\.Expect\("abc"\)\.ShouldNot\(g2\.BeEmpty\(\)\). instead`
		g2.Expect(len("abc")).Should(g2.BeNumerically("==", 3))               // want `ginkgo-linter: wrong length assertion; consider using .g2\.Expect\("abc"\)\.Should\(g2\.HaveLen\(3\)\). instead`
		g2.Expect(len("abc")).WithOffset(1).Should(g2.BeNumerically("==", 3)) // want `ginkgo-linter: wrong length assertion; consider using .g2\.Expect\("abc"\)\.WithOffset\(1\)\.Should\(g2\.HaveLen\(3\)\). instead`
	})
	g1.It("should trigger nil warning", func() {
		var x *int
		g2.Expect(x == nil).Should(g2.Equal(true))       // want `ginkgo-linter: wrong nil assertion; consider using .g2\.Expect\(x\)\.Should\(g2\.BeNil\(\)\). instead`
		g2.Expect(x == nil).ShouldNot(g2.Equal(false))   // want `ginkgo-linter: wrong nil assertion; consider using .g2\.Expect\(x\)\.Should\(g2\.BeNil\(\)\). instead`
		g2.Expect(nil == x).Should(g2.BeTrue())          // want `ginkgo-linter: wrong nil assertion; consider using .g2\.Expect\(x\)\.Should\(g2\.BeNil\(\)\). instead`
		g2.Expect(x == nil).ShouldNot(g2.BeFalse())      // want `ginkgo-linter: wrong nil assertion; consider using .g2\.Expect\(x\)\.Should\(g2\.BeNil\(\)\). instead`
		g2.Expect(x == nil).Should(g2.Not(g2.BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .g2\.Expect\(x\)\.Should\(g2\.BeNil\(\)\). instead`
	})
})
