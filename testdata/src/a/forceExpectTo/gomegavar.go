package forceExpectTo

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("gomega var", func() {
	It("in a valid Eventually", func() {
		Eventually(func(g Gomega) {
			g.Expect("a").Should(HaveLen(1))    // want `ginkgo-linter: must not use Expect with Should; consider using .g\.Expect\("a"\)\.To\(HaveLen\(1\)\). instead`
			g.Expect(len("a")).Should(Equal(1)) // want `ginkgo-linter: wrong length assertion; consider using .g\.Expect\("a"\)\.Should\(HaveLen\(1\)\). instead` `ginkgo-linter: must not use Expect with Should; consider using .g\.Expect\("a"\)\.To\(HaveLen\(1\)\). instead`
		}).Should(Succeed())
	})

	It("in an invalid Eventually", func() {
		Eventually(func(g Gomega) { // want `ginkgo-linter: "Eventually": missing assertion method\. Expected "Should\(\)" or "ShouldNot\(\)"`
			g.Expect("a").Should(HaveLen(1))    // want `ginkgo-linter: must not use Expect with Should; consider using .g\.Expect\("a"\)\.To\(HaveLen\(1\)\). instead`
			g.Expect(len("a")).Should(Equal(1)) // want `ginkgo-linter: wrong length assertion; consider using .g\.Expect\("a"\)\.Should\(HaveLen\(1\)\). instead` `ginkgo-linter: must not use Expect with Should; consider using .g\.Expect\("a"\)\.To\(HaveLen\(1\)\). instead`
		})
	})
})
