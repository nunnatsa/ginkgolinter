package a

// ginkgo-linter:ignore-length-warning

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("suppress file", func() {
	It("should ignore length warning", func() {
		Expect(len("abc")).Should(Equal(3))
		Expect(len("abc")).ShouldNot(BeZero())
	})
})
