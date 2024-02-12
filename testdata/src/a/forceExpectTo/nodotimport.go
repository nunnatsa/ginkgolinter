package forceExpectTo

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("simple case", func() {
	It("just replace assertion method", func() {
		gomega.Expect("a").Should(gomega.HaveLen(1))            // want `ginkgo-linter: must not use Expect with Should; consider using .gomega\.Expect\("a"\)\.To\(gomega\.HaveLen\(1\)\). instead`
		gomega.Expect("a").ShouldNot(gomega.BeEmpty())          // want `ginkgo-linter: must not use Expect with ShouldNot; consider using .gomega\.Expect\("a"\)\.ToNot\(gomega\.BeEmpty\(\)\). instead`
		gomega.Expect("a").Should(gomega.Not(gomega.BeEmpty())) // want `ginkgo-linter: must not use Expect with ShouldNot; consider using .gomega\.Expect\("a"\)\.ToNot\(gomega\.BeEmpty\(\)\). instead`
	})

	It("mix with len assertion", func() {
		gomega.Expect(len("a")).Should(gomega.Equal(1))    // want `ginkgo-linter: wrong length assertion; consider using .gomega\.Expect\("a"\)\.Should\(gomega\.HaveLen\(1\)\). instead` `ginkgo-linter: must not use Expect with Should; consider using .gomega\.Expect\("a"\)\.To\(gomega\.HaveLen\(1\)\). instead`
		gomega.Expect(len("a")).ShouldNot(gomega.Equal(0)) // want `ginkgo-linter: wrong length assertion; consider using .gomega\.Expect\("a"\)\.ShouldNot\(gomega\.BeEmpty\(\)\). instead` `ginkgo-linter: must not use Expect with ShouldNot; consider using .gomega\.Expect\("a"\)\.ToNot\(gomega\.BeEmpty\(\)\). instead`
	})
})
