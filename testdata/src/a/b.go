package a

var _ = Describe("Supress wrong length check", func() {
	Context("test ginkgo-linter:ignore-len-assert-warning", func() {
		It("should ignore length warning", func() {
			// ginkgo-linter:ignore-len-assert-warning
			Expect(len("abc")).Should(Equal(3))
			Expect(len("abc")).Should(Equal(3)) // want `ginkgo-linter: wrong length assertion; consider using .Expect\("abc"\)\.Should\(HaveLen\(3\)\). instead`
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
})
