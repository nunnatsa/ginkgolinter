package focus

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("focused Describe", Focus, func() {
	When("focused when", Focus, func() {
		Context("focused Context", Focus, func() {
			It("focused It", Focus, func() {
				Expect(len("1234")).Should(Equal(4)) // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\("1234"\)\.Should\(HaveLen\(4\)\). instead`
			})
		})
	})
})
