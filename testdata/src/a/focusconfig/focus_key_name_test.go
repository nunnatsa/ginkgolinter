package focus

import (
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = ginkgo.Describe("focused Describe", ginkgo.Focus, func() { // want `ginkgo-linter: Focus spec found. This is used only for local debug and should not be part of the actual source code, consider to remove it`
	ginkgo.When("focused when", ginkgo.Focus, func() { // want `ginkgo-linter: Focus spec found. This is used only for local debug and should not be part of the actual source code, consider to remove it`
		ginkgo.Context("focused Context", ginkgo.Focus, func() { // want `ginkgo-linter: Focus spec found. This is used only for local debug and should not be part of the actual source code, consider to remove it`
			ginkgo.It("focused It", ginkgo.Focus, func() { // want `ginkgo-linter: Focus spec found. This is used only for local debug and should not be part of the actual source code, consider to remove it`
				Expect(len("1234")).Should(Equal(4)) // want `ginkgo-linter: wrong length assertion; consider using .Expect\("1234"\)\.Should\(HaveLen\(4\)\). instead`
			})
		})
	})
})
