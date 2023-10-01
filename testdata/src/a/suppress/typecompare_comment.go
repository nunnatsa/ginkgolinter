package suppress

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("suppress file", func() {
	x := 5
	px := &x
	It("should ignore type compare warning", func() {
		// ginkgo-linter:ignore-type-compare-warning
		Expect(x).Should(Equal(uint(5)))
		Expect(px).Should(HaveValue(Equal(uint(5)))) // want `ginkgo-linter: use Equal with different types: Comparing int with uint; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
	})
})
