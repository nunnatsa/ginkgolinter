package configcompare

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("", func() {
	When("configured to suppress comparison assertion warning", func() {
		It("should not trigger warning", func() {
			abcd := "abcd"
			Expect("abcd" == abcd).To(BeTrue())
		})
	})
})
