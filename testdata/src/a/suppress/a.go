package suppress

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Supress wrong length check", func() {
	Context("test ginkgo-linter:ignore-len-assert-warning", func() {
		It("should ignore length warning", func() {
			// ginkgo-linter:ignore-len-assert-warning
			Expect(len("abc")).Should(Equal(3))
			Expect(len("abc")).Should(Equal(3)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abc"\)\.Should\(HaveLen\(3\)\). instead`
			Expect("123").To(HaveLen(3))
			/*

				Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna
				aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
				Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint
				occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

				ginkgo-linter:ignore-len-assert-warning

				Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna
				aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
				Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint
				occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
			*/
			Expect(len("abc")).Should(Equal(3))
		})
	})

	Context("test ginkgo-linter:ignore-nil-assert-warning", func() {
		var x *int
		It("should ignore length warning", func() {
			// ginkgo-linter:ignore-nil-assert-warning
			Expect(x == nil).Should(BeTrue())
			Expect(x == nil).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion\nconsider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
			Expect(x).To(BeNil())
			/*

				Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna
				aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
				Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint
				occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

				ginkgo-linter:ignore-nil-assert-warning

				Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna
				aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
				Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint
				occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
			*/
			Expect(x == nil).Should(BeTrue())
		})
	})

	Context("test ginkgo-linter:ignore-err-assert-warning", func() {
		var x error
		It("should ignore error warning", func() {
			// ginkgo-linter:ignore-err-assert-warning
			Expect(x).To(BeNil())
			Expect(x == nil).Should(BeTrue()) // want `ginkgo-linter: wrong error assertion\nconsider using .Expect\(x\)\.ShouldNot\(HaveOccurred\(\)\). instead`
			// ginkgo-linter:ignore-err-assert-warning
			Expect(x == nil).Should(BeTrue())
			/*

				Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna
				aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
				Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint
				occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

				ginkgo-linter:ignore-err-assert-warning

				Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna
				aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
				Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint
				occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
			*/
			Expect(x == nil).Should(BeTrue())
		})
	})
})
