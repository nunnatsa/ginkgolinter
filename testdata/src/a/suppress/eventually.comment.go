package suppress

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func slowInt() int {
	time.Sleep(time.Second)
	return 42
}

func slowBool() bool {
	time.Sleep(time.Second)
	return true
}

var _ = Describe("should suppress async assertions", func() {
	It("should suppress async assertions", func() {
		Eventually(slowInt()).Should(Equal(42)) // want `ginkgo-linter: use a function call in Eventually. This actually checks nothing, because Eventually receives the function returned value, instead of function itself, and this value is never changed; consider using .Eventually\(slowInt\)\.Should\(Equal\(42\)\). instead`

		// ginkgo-linter:ignore-async-assert-warning
		Eventually(slowInt()).Should(Equal(42))
	})

	It("should suppress async assertions, but not Eqaul(true)", func() {
		Eventually(slowBool()).Should(Equal(true)) // want `ginkgo-linter: use a function call in Eventually. This actually checks nothing, because Eventually receives the function returned value, instead of function itself, and this value is never changed; consider using .Eventually\(slowBool\)\.Should\(BeTrue\(\)\). instead` `ginkgo-linter: wrong boolean assertion; consider using .Eventually\(slowBool\)\.Should\(BeTrue\(\)\). instead`

		// ginkgo-linter:ignore-async-assert-warning
		Eventually(slowBool()).Should(Equal(true)) // want `ginkgo-linter: wrong boolean assertion; consider using .Eventually\(slowBool\(\)\)\.Should\(BeTrue\(\)\). instead`
	})
})
