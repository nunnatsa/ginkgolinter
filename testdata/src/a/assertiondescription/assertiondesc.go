package assertiondescription

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("assertion description tests", func() {
	It("should test missing descriptions", func() {
		// This should trigger warning when ForceAssertionDescription is enabled
		Expect("hello").To(Equal("hello")) // want `missing assertion description`

		// This should trigger warning - no description
		Expect(42).Should(Equal(42)) // want `missing assertion description`

		// This should trigger warning - no description
		Expect([]int{1, 2, 3}).To(HaveLen(3)) // want `missing assertion description`

		// This should NOT trigger warning - has description
		Expect(true).To(BeTrue(), "this should be true")
	})

	It("should test async assertions", func() {
		// Test Eventually without description
		Eventually(func() bool { return true }).Should(BeTrue())          // want `missing assertion description`
		Eventually(func() string { return "test" }).Should(Equal("test")) // want `missing assertion description`

		// Test Consistently without description
		Consistently(func() bool { return true }).Should(BeTrue()) // want `missing assertion description`

		// Test Eventually with Expect inside - should trigger warning on Expect
		Eventually(func() bool { // want `missing assertion description`
			Expect("inner").To(Equal("inner")) // want `missing assertion description`
			return true
		}).Should(BeTrue())

		// Test Eventually with multiple Expects inside
		Eventually(func() bool {
			Expect(42).To(Equal(42))  // want `missing assertion description`
			Expect(true).To(BeTrue()) // want `missing assertion description`
			return true
		}).Should(BeTrue(), "eventually should succeed")

		// These should NOT trigger warning - have descriptions
		Eventually(func() bool { return true }).Should(BeTrue(), "eventually should be true")
		Consistently(func() bool { return true }).Should(BeTrue(), "consistently should be true")

		// Expect inside Eventually with descriptions - should NOT trigger warning
		Eventually(func() bool {
			Expect("inner with desc").To(Equal("inner with desc"), "inner assertion")
			return true
		}).Should(BeTrue(), "eventually should succeed")
	})

	It("should test mixed cases", func() {
		// Mix of sync and async without descriptions
		Expect("sync").To(Equal("sync"))                                    // want `missing assertion description`
		Eventually(func() string { return "async" }).Should(Equal("async")) // want `missing assertion description`

		// Mix with descriptions
		Expect("sync with desc").To(Equal("sync with desc"), "sync assertion")
		Eventually(func() string { return "async with desc" }).Should(Equal("async with desc"), "async assertion")

		// Mix of dot import and named import
		Expect(42).To(Equal(42)) // want `missing assertion description`
	})
})
