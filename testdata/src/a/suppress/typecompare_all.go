package suppress

// ginkgo-linter:ignore-type-compare-warning

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("suppress file", func() {
	x := 5
	px := &x
	It("should ignore type compare warning", func() {
		Expect(x).Should(Equal(uint(5)))
		Expect(px).Should(HaveValue(Equal(uint(5))))
	})
})
