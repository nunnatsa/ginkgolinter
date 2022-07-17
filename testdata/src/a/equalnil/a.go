package equalnil

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Check Equal(nil)", func() {
	It("should trigger warning if comparing to nil", func() {
		var x *int
		var y = 5
		var py = &y
		Expect(x).Should(Equal(nil))                    // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
		Expect(py).Should(Not(Equal(nil)))              // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(py\)\.ShouldNot\(BeNil\(\)\). instead`
		ExpectWithOffset(1, py).Should(Not(Equal(nil))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, py\)\.ShouldNot\(BeNil\(\)\). instead`
	})
})
