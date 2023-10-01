package comparetypes_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("compare different types", func() {
	It("find false positive check", func() {
		a := 5
		Expect(multi()).ShouldNot(Equal(4)) // want `ginkgo-linter: use Equal with different types: Comparing int64 with int; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		Expect(a).ShouldNot(Equal(5))
		Expect(a).ShouldNot(Equal(uint(5)))           // want `ginkgo-linter: use Equal with different types: Comparing int with uint; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		Expect(fint64()).ShouldNot(Equal(uint64(5)))  // want `ginkgo-linter: use Equal with different types: Comparing int64 with uint64; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		Expect(fmytype()).ShouldNot(Equal(uint64(5))) // want `ginkgo-linter: use Equal with different types: Comparing a/comparetypes_test.mytype with uint64; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		Expect(a).ShouldNot(Equal(int32(5)))          // want `ginkgo-linter: use Equal with different types: Comparing int with int32; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		Expect(a).ShouldNot(Equal(uint8(5)))          // want `ginkgo-linter: use Equal with different types: Comparing int with uint8; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		Expect(a).ShouldNot(Equal(mytype(5)))         // want `ginkgo-linter: use Equal with different types: Comparing int with a/comparetypes_test.mytype; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		Expect(5).ShouldNot(Equal(mytype(5)))         // want `ginkgo-linter: use Equal with different types: Comparing int with a/comparetypes_test.mytype; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`

		b := int16(5)
		Expect(a).ShouldNot(Equal(b)) // want `ginkgo-linter: use Equal with different types: Comparing int with int16; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`

		c := mytype(5)
		Expect(a).ShouldNot(Equal(c)) // want `ginkgo-linter: use Equal with different types: Comparing int with a/comparetypes_test.mytype; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
	})

	It("compare interfaces", func() {
		var (
			a myinf = imp1(3)
			b myinf = imp2(3)
		)
		Expect(a).ShouldNot(Equal(b)) // this is not good, but not an error. Should have kind of warning here.
		Expect(a).ShouldNot(BeEquivalentTo(b))
		Expect(a).ShouldNot(BeIdenticalTo(b)) // this is not good, but not an error. Should have kind of warning here.
	})

	It("check HaveValue(Equal)", func() {
		a := 5
		pa := &a

		Expect(pa).Should(HaveValue(Equal(uint64(5))))                                    // want `ginkgo-linter: use Equal with different types: Comparing int with uint64; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		Expect(a).Should(Not(Equal(uint64(5))))                                           // want `ginkgo-linter: use Equal with different types: Comparing int with uint64; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		Expect(a).Should(And(Equal(uint64(5)), Not(Equal(int32(6)))))                     // want `ginkgo-linter: use Equal with different types: Comparing int with uint64; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)` `ginkgo-linter: use Equal with different types: Comparing int with int32; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		Expect(a).Should(Or(Equal(uint64(5)), Not(Equal(int32(6))), Not(Equal(int8(4))))) // want `ginkgo-linter: use Equal with different types: Comparing int with uint64; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)` `ginkgo-linter: use Equal with different types: Comparing int with int32; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)` `ginkgo-linter: use Equal with different types: Comparing int with int8; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		Expect(a).Should(WithTransform(func(i int) int { return i + 1 }, Equal(uint(6)))) // want `ginkgo-linter: use Equal with different types: Comparing int with uint; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
	})

	It("compare interface with implementations", func() {
		var (
			a myinf = imp1(3)
			b       = imp1(3)
		)
		Expect(a).Should(Equal(b))
		Expect(b).Should(Equal(a))
	})
})
