package focus

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("should warn", func() { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code\. Consider to replace with "Describe"`
	When("should ignore", func() {
		It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})
	Context("should ignore", func() {
		It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})

	It("should ignore", func() {
		Expect("abcd").Should(HaveLen(4))
	})
})

var _ = Describe("should ignore", func() {
	FWhen("should warn", func() { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code\. Consider to replace with "When"`
		Context("should ignore", func() {
			It("should ignore", func() {
				Expect("abcd").Should(HaveLen(4))
			})
		})

		It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})

	FContext("should warn", func() { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code\. Consider to replace with "Context"`
		Context("should ignore", func() {
			It("should ignore", func() {
				Expect("abcd").Should(HaveLen(4))
			})
		})

		It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})

	Context("ignore", func() {
		FIt("should warn", func() { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code\. Consider to replace with "It"`
			Expect("abcd").Should(HaveLen(4))
		})
	})
})

var _ = Describe("suppress", func() {
	// ginkgo-linter:ignore-focus-container-warning
	FContext("supress", func() {
		// ginkgo-linter:ignore-focus-container-warning
		FWhen("suppress", func() {
			// ginkgo-linter:ignore-focus-container-warning
			FIt("suppress", func() {
				Expect(len("abcd")).Should(Equal(4)) // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\("abcd"\)\.Should\(HaveLen\(4\)\). instead`
			})
		})
	})
})
