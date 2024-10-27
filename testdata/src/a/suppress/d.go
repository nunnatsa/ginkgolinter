package suppress

// ginkgo-linter:ignore-nil-assert-warning

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("suppress file", func() {
	var x *int
	It("should ignore nil warning", func() {
		Expect(x == nil).Should(Equal(true)) // want `ginkgo-linter: wrong boolean assertion\. Consider using .Expect\(x == nil\)\.Should\(BeTrue\(\)\). instead`
		Expect(x == nil).ShouldNot(BeTrue())
	})
})
