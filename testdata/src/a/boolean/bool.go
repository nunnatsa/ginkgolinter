package boolean

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Equal(true/false)", func() {
	It("check Equal(true/false)", func() {
		t := true
		f := false
		Expect(t).To(Equal(true))                        // want `ginkgo-linter: wrong boolean assertion\nconsider using .Expect\(t\)\.To\(BeTrue\(\)\). instead`
		Expect(f).To(Equal(false))                       // want `ginkgo-linter: wrong boolean assertion\nconsider using .Expect\(f\)\.To\(BeFalse\(\)\). instead`
		ExpectWithOffset(2, t).Should(Not(Equal(false))) // want `ginkgo-linter: wrong boolean assertion\nconsider using .ExpectWithOffset\(2, t\)\.ShouldNot\(BeFalse\(\)\). instead`
	})
})
