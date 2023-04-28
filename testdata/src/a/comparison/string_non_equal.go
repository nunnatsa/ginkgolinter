package comparison

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("string comparison", func() {
	It("should not trigger warning for string non-equal comparison ", func() {
		a, b := "aaa", "bbb"
		Expect(a < b).To(BeTrue())
		Expect(a < "bbb").To(BeTrue())
	})
})
