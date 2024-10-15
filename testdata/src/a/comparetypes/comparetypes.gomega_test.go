package comparetypes_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func fint64() int64 {
	return 5
}

type mytype int

func fmytype() mytype {
	return 5
}

func multi() (int64, error) {
	return 4, nil
}

type myinf interface {
	doSomething()
}

type imp1 int

func (imp1) doSomething() {
	fmt.Println("doing something")
}

type imp2 int

func (imp2) doSomething() {
	fmt.Println("doing something else")
}

var _ = Describe("compare different types", func() {
	It("find false positive check", func() {
		a := 5
		gomega.Expect(multi()).ShouldNot(gomega.Equal(4)) // want `ginkgo-linter: use Equal with different types: Comparing int64 with int; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		gomega.Expect(a).ShouldNot(gomega.Equal(5))
		gomega.Expect(a).ShouldNot(gomega.Equal(uint(5)))           // want `ginkgo-linter: use Equal with different types: Comparing int with uint; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		gomega.Expect(fint64()).ShouldNot(gomega.Equal(uint64(5)))  // want `ginkgo-linter: use Equal with different types: Comparing int64 with uint64; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		gomega.Expect(fmytype()).ShouldNot(gomega.Equal(uint64(5))) // want `ginkgo-linter: use Equal with different types: Comparing a/comparetypes_test\.mytype with uint64; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		gomega.Expect(a).ShouldNot(gomega.Equal(int32(5)))          // want `ginkgo-linter: use Equal with different types: Comparing int with int32; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		gomega.Expect(a).ShouldNot(gomega.Equal(uint8(5)))          // want `ginkgo-linter: use Equal with different types: Comparing int with uint8; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		gomega.Expect(a).ShouldNot(gomega.Equal(mytype(5)))         // want `ginkgo-linter: use Equal with different types: Comparing int with a/comparetypes_test\.mytype; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		gomega.Expect(5).ShouldNot(gomega.Equal(mytype(5)))         // want `ginkgo-linter: use Equal with different types: Comparing int with a/comparetypes_test\.mytype; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`

		b := int16(5)
		gomega.Expect(a).ShouldNot(gomega.Equal(b)) // want `ginkgo-linter: use Equal with different types: Comparing int with int16; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`

		c := mytype(5)
		gomega.Expect(a).ShouldNot(gomega.Equal(c)) // want `ginkgo-linter: use Equal with different types: Comparing int with a/comparetypes_test\.mytype; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
	})

	It("compare interfaces", func() {
		var (
			a myinf = imp1(3)
			b myinf = imp2(3)
		)
		gomega.Expect(a).ShouldNot(gomega.Equal(b)) // this is not good, but not an error. Should have kind of warning here.
		gomega.Expect(a).ShouldNot(gomega.BeEquivalentTo(b))
		gomega.Expect(a).ShouldNot(gomega.BeIdenticalTo(b)) // this is not good, but not an error. Should have kind of warning here.
	})

	It("check HaveValue(Equal)", func() {
		a := 5
		pa := &a

		gomega.Expect(pa).Should(gomega.HaveValue(gomega.Equal(uint64(5))))                                                                // want `ginkgo-linter: use Equal with different types: Comparing int with uint64; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		gomega.Expect(a).Should(gomega.Not(gomega.Equal(uint64(5))))                                                                       // want `ginkgo-linter: use Equal with different types: Comparing int with uint64; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		gomega.Expect(a).Should(gomega.And(gomega.Equal(uint64(5)), gomega.Not(gomega.Equal(int32(6)))))                                   // want `ginkgo-linter: multiple issues: use Equal with different types: Comparing int with uint64; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\); use Equal with different types: Comparing int with int32; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		gomega.Expect(a).Should(gomega.Or(gomega.Equal(uint64(5)), gomega.Not(gomega.Equal(int32(6))), gomega.Not(gomega.Equal(int8(4))))) // want `ginkgo-linter: multiple issues: use Equal with different types: Comparing int with uint64; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\); use Equal with different types: Comparing int with int32; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\); use Equal with different types: Comparing int with int8; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		gomega.Expect(a).Should(gomega.WithTransform(func(i int) int { return i + 1 }, gomega.Equal(uint(6))))                             // want `ginkgo-linter: use Equal with different types: Comparing int with uint; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
	})

	It("test WithTransform", func() {
		a := uint(5)
		gomega.Expect(uint(5)).Should(gomega.WithTransform(func(i uint) int { return int(i) }, gomega.Equal(5)))
		gomega.Expect(uint(5)).Should(gomega.WithTransform(func(i uint) uint64 { return uint64(i) }, gomega.Equal(5))) // want `ginkgo-linter: use Equal with different types: Comparing uint64 with int; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		gomega.Expect(a).Should(gomega.WithTransform(func(i uint) int { return int(i) }, gomega.Equal(5)))
		gomega.Expect(a).Should(gomega.WithTransform(uint2int, gomega.Equal(5)))
		gomega.Expect(a).Should(gomega.WithTransform(uint2int, gomega.Equal(uint64(5))))             // want `ginkgo-linter: use Equal with different types: Comparing int with uint64; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		gomega.Expect(a).Should(gomega.WithTransform(uint2int, gomega.Not(gomega.Equal(uint64(5))))) // want `ginkgo-linter: use Equal with different types: Comparing int with uint64; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		gomega.Expect(5).Should(gomega.WithTransform(func(i int) myinf { return imp1(i) }, gomega.Equal(imp1(5))))
	})

	It("issue 115", func() {
		gomega.Expect([]int{42, 23}).Should(gomega.WithTransform(func(v []int) []string {
			ret := make([]string, 0, len(v))
			for _, i := range v {
				ret = append(ret, fmt.Sprintf("%v", i))
			}
			return ret
		}, gomega.ContainElement("42")))
	})
})
