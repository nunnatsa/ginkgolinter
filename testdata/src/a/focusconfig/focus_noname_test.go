package focus

import (
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = ginkgo.FDescribe("should warn", func() {
	ginkgo.When("should ignore", func() {
		ginkgo.It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})
	ginkgo.Context("should ignore", func() {
		ginkgo.It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})

	ginkgo.It("should ignore", func() {
		Expect("abcd").Should(HaveLen(4))
	})
})

var _ = ginkgo.Describe("should ignore", func() {
	ginkgo.FWhen("should warn", func() {
		ginkgo.Context("should ignore", func() {
			ginkgo.It("should ignore", func() {
				Expect("abcd").Should(HaveLen(4))
			})
		})

		ginkgo.It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})

	ginkgo.FContext("should warn", func() {
		ginkgo.Context("should ignore", func() {
			ginkgo.It("should ignore", func() {
				Expect("abcd").Should(HaveLen(4))
			})
		})

		ginkgo.It("should ignore", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})

	ginkgo.Context("ignore", func() {
		ginkgo.FIt("should warn", func() {
			Expect("abcd").Should(HaveLen(4))
		})
	})
})
