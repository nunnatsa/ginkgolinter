package focus

// ginkgo-linter:ignore-focus-container-warning

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("should warn", func() {
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
	FWhen("should warn", func() {
		Context("should ignore", func() {
			It("should ignore", func() {
				Expect("abcd").Should(HaveLen(4))
			})
		})

		It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})

	FContext("should warn", func() {
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
		FIt("should warn", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})
})
