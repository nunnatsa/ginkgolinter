package focus

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("focused Describe", Focus, func() { // want `ginkgo-linter: Focus spec found. This is used only for local debug and should not be part of the actual source code\. Consider to remove it`
	When("focused when", Focus, func() { // want `ginkgo-linter: Focus spec found. This is used only for local debug and should not be part of the actual source code\. Consider to remove it`
		Context("focused Context", Focus, func() { // want `ginkgo-linter: Focus spec found. This is used only for local debug and should not be part of the actual source code\. Consider to remove it`
			It("focused It", Focus, func() { // want `ginkgo-linter: Focus spec found. This is used only for local debug and should not be part of the actual source code\. Consider to remove it`
				Expect(len("1234")).Should(Equal(4)) // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\("1234"\)\.Should\(HaveLen\(4\)\). instead`
			})
		})
	})
})
