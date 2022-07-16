package nodotimport

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("no dot import", func() {
	ginkgo.It("should trigger length warning", func() {
		const length = 3
		gomega.Expect(len("abc")).Should(gomega.Equal(3))               // want `ginkgo-linter: wrong length assertion; consider using .gomega\.Expect\("abc"\)\.Should\(gomega\.HaveLen\(3\)\). instead`
		gomega.Expect(len("abc")).Should(gomega.Not(gomega.Equal(4)))   // want `ginkgo-linter: wrong length assertion; consider using .gomega\.Expect\("abc"\)\.ShouldNot\(gomega\.HaveLen\(4\)\). instead`
		gomega.Expect(len("abc")).Should(gomega.Equal(length))          // want `ginkgo-linter: wrong length assertion; consider using .gomega\.Expect\("abc"\)\.Should\(gomega\.HaveLen\(length\)\). instead`
		gomega.Expect(len("")).Should(gomega.Equal(0))                  // want `ginkgo-linter: wrong length assertion; consider using .gomega\.Expect\(""\)\.Should\(gomega\.BeEmpty\(\)\). instead`
		gomega.Expect(len("abc")).ShouldNot(gomega.BeZero())            // want `ginkgo-linter: wrong length assertion; consider using .gomega\.Expect\("abc"\)\.ShouldNot\(gomega\.BeEmpty\(\)\). instead`
		gomega.Expect(len("abc")).Should(gomega.BeNumerically(">", 0))  // want `ginkgo-linter: wrong length assertion; consider using .gomega\.Expect\("abc"\)\.ShouldNot\(gomega\.BeEmpty\(\)\). instead`
		gomega.Expect(len("abc")).Should(gomega.BeNumerically("==", 3)) // want `ginkgo-linter: wrong length assertion; consider using .gomega\.Expect\("abc"\)\.Should\(gomega\.HaveLen\(3\)\). instead`
	})
	ginkgo.It("should trigger nil warning", func() {
		var x *int
		gomega.Expect(x == nil).Should(gomega.Equal(true))           // want `ginkgo-linter: wrong nil assertion; consider using .gomega\.Expect\(x\)\.Should\(gomega\.beNil\(\)\). instead`
		gomega.Expect(x == nil).ShouldNot(gomega.Equal(false))       // want `ginkgo-linter: wrong nil assertion; consider using .gomega\.Expect\(x\)\.Should\(gomega\.beNil\(\)\). instead`
		gomega.Expect(nil == x).Should(gomega.BeTrue())              // want `ginkgo-linter: wrong nil assertion; consider using .gomega\.Expect\(x\)\.Should\(gomega\.beNil\(\)\). instead`
		gomega.Expect(x == nil).ShouldNot(gomega.BeFalse())          // want `ginkgo-linter: wrong nil assertion; consider using .gomega\.Expect\(x\)\.Should\(gomega\.beNil\(\)\). instead`
		gomega.Expect(x == nil).Should(gomega.Not(gomega.BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .gomega\.Expect\(x\)\.Should\(gomega\.beNil\(\)\). instead`
	})
})
