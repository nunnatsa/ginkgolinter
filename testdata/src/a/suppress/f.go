package suppress

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("suppress file", func() {
	var x = "abcd"
	It("should ignore length warning", func() {
		// ginkgo-linter:ignore-nil-assert-warning
		// ginkgo-linter:ignore-len-assert-warning
		// check that linter is suppressed when all flags are true
		Expect(len(x)).Should(Equal(4))
	})
})
