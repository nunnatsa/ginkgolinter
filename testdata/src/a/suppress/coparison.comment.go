package suppress

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("should suppress comparison assertions", func() {
	It("should suppress comparison assertions", func() {
		Expect(len("abcd")).To(Equal(4)) // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\("abcd"\)\.To\(HaveLen\(4\)\). instead`
		str := "abcd"
		// ginkgo-linter:ignore-compare-assert-warning
		Expect("abcd" == str).To(BeTrue()) // no warning triggered
		Expect("abcd" == str).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion\. Consider using .Expect\(str\)\.To\(Equal\("abcd"\)\). instead`
	})
})
