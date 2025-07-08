package assertiondescription

import (
	. "github.com/onsi/ginkgo/v2"
	g "github.com/onsi/gomega"
)

var _ = Describe("assertion description tests", func() {
	It("should test missing descriptions", func() {
		// This should trigger warning when ForceAssertionDescription is enabled
		g.Expect("hello").To(g.Equal("hello")) // want `missing assertion description`

		// This should trigger warning - no description
		g.Expect(42).Should(g.Equal(42)) // want `missing assertion description`

		// This should trigger warning - no description
		g.Expect([]int{1, 2, 3}).To(g.HaveLen(3)) // want `missing assertion description`

		// This should NOT trigger warning - has description
		g.Expect(true).To(g.BeTrue(), "this should be true")
	})

	It("should test async assertions", func() {
		// Test g.Eventually without description
		g.Eventually(func() bool { return true }).Should(g.BeTrue())          // want `missing assertion description`
		g.Eventually(func() string { return "test" }).Should(g.Equal("test")) // want `missing assertion description`

		// Test g.Consistently without description
		g.Consistently(func() bool { return true }).Should(g.BeTrue()) // want `missing assertion description`

		// Test g.Eventually with g.Expect inside - should trigger warning on g.Expect
		g.Eventually(func() bool { // want `missing assertion description`
			g.Expect("inner").To(g.Equal("inner")) // want `missing assertion description`
			return true
		}).Should(g.BeTrue())

		// Test with gomega prefix
		g.Eventually(func() bool { return true }).Should(g.BeTrue()) // want `missing assertion description`

		// Test g.Eventually with multiple g.Expects inside
		g.Eventually(func() bool {
			g.Expect(42).To(g.Equal(42))  // want `missing assertion description`
			g.Expect(true).To(g.BeTrue()) // want `missing assertion description`
			return true
		}).Should(g.BeTrue(), "g.Eventually should succeed")

		// These should NOT trigger warning - have descriptions
		g.Eventually(func() bool { return true }).Should(g.BeTrue(), "g.Eventually should be true")
		g.Consistently(func() bool { return true }).Should(g.BeTrue(), "g.Consistently should be true")

		// g.Expect inside g.Eventually with descriptions - should NOT trigger warning
		g.Eventually(func() bool {
			g.Expect("inner with desc").To(g.Equal("inner with desc"), "inner assertion")
			return true
		}).Should(g.BeTrue(), "g.Eventually should succeed")
	})

	It("should test mixed cases", func() {
		// Mix of sync and async without descriptions
		g.Expect("sync").To(g.Equal("sync"))                                    // want `missing assertion description`
		g.Eventually(func() string { return "async" }).Should(g.Equal("async")) // want `missing assertion description`

		// Mix with descriptions
		g.Expect("sync with desc").To(g.Equal("sync with desc"), "sync assertion")
		g.Eventually(func() string { return "async with desc" }).Should(g.Equal("async with desc"), "async assertion")

		// Mix of dot import and named import
		g.Expect(42).To(g.Equal(42)) // want `missing assertion description`
		g.Expect(43).To(g.Equal(43)) // want `missing assertion description`
	})
})
