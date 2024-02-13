package suppress

// ginkgo-linter:ignore-compare-assert-warning

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("should suppress comparison assertions", func() {
	It("should suppress comparison assertions", func() {
		Expect(len("abcd")).To(Equal(4)) // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\("abcd"\)\.To\(HaveLen\(4\)\). instead`
		str := "abcd"
		Expect("abcd" == str).To(BeTrue()) // no warning triggered
	})
})
