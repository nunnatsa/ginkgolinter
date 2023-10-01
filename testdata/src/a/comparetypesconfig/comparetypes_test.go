package comparetypes_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("compare different types", func() {
	It("find false positive check", func() {
		a := 5
		Expect(multi()).ShouldNot(Equal(4))
		Expect(a).ShouldNot(Equal(5))
		Expect(a).ShouldNot(Equal(uint(5)))
		Expect(fint64()).ShouldNot(Equal(uint64(5)))
		Expect(fmytype()).ShouldNot(Equal(uint64(5)))
		Expect(a).ShouldNot(Equal(int32(5)))
		Expect(a).ShouldNot(Equal(uint8(5)))
		Expect(a).ShouldNot(Equal(mytype(5)))
		Expect(5).ShouldNot(Equal(mytype(5)))

		b := int16(5)
		Expect(a).ShouldNot(Equal(b))

		c := mytype(5)
		Expect(a).ShouldNot(Equal(c))
	})

	It("compare interfaces", func() {
		var (
			a myinf = imp1(3)
			b myinf = imp2(3)
		)
		Expect(a).ShouldNot(Equal(b))
		Expect(a).ShouldNot(BeEquivalentTo(b))
		Expect(a).ShouldNot(BeIdenticalTo(b))
	})

	It("check HaveValue(Equal)", func() {
		a := 5
		pa := &a

		Expect(pa).Should(HaveValue(Equal(uint64(5))))
		Expect(a).Should(Not(Equal(uint64(5))))
		Expect(a).Should(And(Equal(uint64(5)), Not(Equal(int32(6)))))
		Expect(a).Should(Or(Equal(uint64(5)), Not(Equal(int32(6))), Not(Equal(int8(4)))))
		Expect(a).Should(WithTransform(func(i int) int { return i + 1 }, Equal(uint(6))))
	})
})
