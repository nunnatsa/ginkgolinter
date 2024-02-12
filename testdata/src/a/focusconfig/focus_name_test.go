package focus

import (
	tester "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = tester.FDescribe("should warn", func() { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code\. Consider to replace with "Describe"`
	tester.When("should ignore", func() {
		tester.It("should ignore", func() {
			Expect(len("abcd")).Should(Equal(4)) // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\("abcd"\)\.Should\(HaveLen\(4\)\). instead`
		})
	})
	tester.Context("should ignore", func() {
		tester.It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})

	tester.It("should ignore", func() {
		Expect("abcd").Should(HaveLen(4))
	})
})

var _ = tester.Describe("should ignore", func() {
	tester.FWhen("should warn", func() { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code\. Consider to replace with "When"`
		tester.Context("should ignore", func() {
			tester.It("should ignore", func() {
				Expect("abcd").Should(HaveLen(4))
			})
		})

		tester.It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})

	tester.FContext("should warn", func() { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code\. Consider to replace with "Context"`
		tester.Context("should ignore", func() {
			tester.It("should ignore", func() {
				Expect("abcd").Should(HaveLen(4))
			})
		})

		tester.It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})

	tester.Context("ignore", func() {
		tester.FIt("should warn", func() { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code\. Consider to replace with "It"`
			Expect("abcd").Should(HaveLen(4))
		})
	})
})
