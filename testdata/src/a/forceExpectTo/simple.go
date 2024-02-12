package forceExpectTo

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("simple case", func() {
	It("just replace assertion method", func() {
		Expect("a").Should(HaveLen(1))     // want `ginkgo-linter: must not use Expect with Should\. Consider using .Expect\("a"\)\.To\(HaveLen\(1\)\). instead`
		Expect("a").ShouldNot(BeEmpty())   // want `ginkgo-linter: must not use Expect with ShouldNot\. Consider using .Expect\("a"\)\.ToNot\(BeEmpty\(\)\). instead`
		Expect("a").Should(Not(BeEmpty())) // want `ginkgo-linter: must not use Expect with Should\. Consider using .Expect\("a"\)\.ToNot\(BeEmpty\(\)\). instead`
		Ω("a").Should(HaveLen(1))
		Ω("a").ShouldNot(BeEmpty())
	})

	It("mix with len assertion", func() {
		Expect(len("a")).Should(Equal(1))    // want `ginkgo-linter: multiple issues: must not use Expect with Should; wrong length assertion\. Consider using .Expect\("a"\)\.To\(HaveLen\(1\)\). instead`
		Expect(len("a")).ShouldNot(Equal(0)) // want `ginkgo-linter: multiple issues: must not use Expect with ShouldNot; wrong length assertion\. Consider using .Expect\("a"\)\.ToNot\(BeEmpty\(\)\). instead`
	})
})
