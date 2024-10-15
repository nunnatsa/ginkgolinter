package suppress

// ginkgo-linter:ignore-len-assert-warning
// ginkgo-linter:ignore-nil-assert-warning
// ginkgo-linter:ignore-err-assert-warning
// ginkgo-linter:ignore-compare-assert-warning
// ginkgo-linter:ignore-async-assert-warning
// Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna
// aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
// Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint
// occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna
// aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
// Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint
// occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
var _ = Describe("suppress file", func() {
	It("should ignore length warning", func() {
		Expect(len("abcd")).Should(Equal(4))
		Expect(len("abcd")).ShouldNot(BeZero())
	})
	It("should ignore nil warning", func() {
		var x *int
		Expect(x == nil).Should(Equal(true))
		Expect(x == nil).ShouldNot(BeTrue())
	})
	It("should ignore error warning", func() {
		var x error
		Expect(x == nil).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion. Consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
		Expect(x == nil).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion. Consider using .Expect\(x\)\.ShouldNot\(BeNil\(\)\). instead`
	})
})
