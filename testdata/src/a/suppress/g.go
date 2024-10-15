package suppress

// ginkgo-linter:ignore-err-assert-warning

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("suppress file", func() {
	var x error
	It("should ignore nil warning", func() {
		Expect(x == nil).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion. Consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
		Expect(x == nil).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion. Consider using .Expect\(x\)\.ShouldNot\(BeNil\(\)\). instead`
	})
})
