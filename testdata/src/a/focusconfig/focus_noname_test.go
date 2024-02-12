package focus

import (
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = ginkgo.FDescribe("should warn", func() { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code\. Consider to replace with "Describe"`
	ginkgo.When("should ignore", func() {
		ginkgo.It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})
	ginkgo.Context("should ignore", func() {
		ginkgo.It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})

	ginkgo.It("should ignore", func() {
		Expect("abcd").Should(HaveLen(4))
	})
})

var _ = ginkgo.Describe("should ignore", func() {
	ginkgo.FWhen("should warn", func() { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code\. Consider to replace with "When"`
		ginkgo.Context("should ignore", func() {
			ginkgo.It("should ignore", func() {
				Expect("abcd").Should(HaveLen(4))
			})
		})

		ginkgo.It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})

	ginkgo.FContext("should warn", func() { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code\. Consider to replace with "Context"`
		ginkgo.Context("should ignore", func() {
			ginkgo.It("should ignore", func() {
				Expect("abcd").Should(HaveLen(4))
			})
		})

		ginkgo.It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})

	ginkgo.Context("ignore", func() {
		ginkgo.FIt("should warn", func() { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code\. Consider to replace with "It"`
			Expect("abcd").Should(HaveLen(4))
		})
	})
})
