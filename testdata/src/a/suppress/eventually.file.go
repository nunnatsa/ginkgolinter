package suppress

// ginkgo-linter:ignore-async-assert-warning

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("should suppress async assertions", func() {
	It("should suppress async assertions", func() {
		Eventually(slowInt()).Should(Equal(42))
	})

	It("should suppress async assertions, but not Eqaul(true)", func() {
		Eventually(slowBool()).Should(Equal(true)) // want `ginkgo-linter: wrong boolean assertion\. Consider using .Eventually\(slowBool\(\)\)\.Should\(BeTrue\(\)\). instead`
	})
})
