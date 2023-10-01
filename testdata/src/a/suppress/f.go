package suppress

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("suppress file", func() {
	var x = "abcd"
	It("should ignore length warning", func() {
		// ginkgo-linter:ignore-nil-assert-warning
		// ginkgo-linter:ignore-len-assert-warning
		// ginkgo-linter:ignore-err-assert-warning
		// ginkgo-linter:ignore-compare-assert-warning
		// ginkgo-linter:ignore-async-assert-warning
		// check that linter is suppressed when all flags are true
		Expect(len(x)).Should(Equal(4))
	})

	It("should ignore equal nil waring", func() {
		var a *int = nil
		b := 3
		c := &b

		// ginkgo-linter:ignore-nil-assert-warning
		Expect(a).To(Equal(nil)) // want `ginkgo-linter: use Equal with different types: Comparing \*int with untyped nil; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		// ginkgo-linter:ignore-nil-assert-warning
		Expect(c).To(Not(Equal(nil))) // want `ginkgo-linter: use Equal with different types: Comparing \*int with untyped nil; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
	})
})
