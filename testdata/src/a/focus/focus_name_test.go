package focus

import (
	tester "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = tester.FDescribe("should warn", func() {
	tester.When("should ignore", func() {
		tester.It("should ignore", func() {
			Expect(len("abcd")).Should(Equal(4)) // want `ginkgo-linter: wrong length assertion; consider using .Expect\("abcd"\)\.Should\(HaveLen\(4\)\). instead`
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
	tester.FWhen("should warn", func() {
		tester.Context("should ignore", func() {
			tester.It("should ignore", func() {
				Expect("abcd").Should(HaveLen(4))
			})
		})

		tester.It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})

	tester.FContext("should warn", func() {
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
		tester.FIt("should warn", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})
})
