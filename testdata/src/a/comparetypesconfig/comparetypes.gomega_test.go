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
		gomega.Expect(multi()).ShouldNot(gomega.Equal(4))
		gomega.Expect(a).ShouldNot(gomega.Equal(5))
		gomega.Expect(a).ShouldNot(gomega.Equal(uint(5)))
		gomega.Expect(fint64()).ShouldNot(gomega.Equal(uint64(5)))
		gomega.Expect(fmytype()).ShouldNot(gomega.Equal(uint64(5)))
		gomega.Expect(a).ShouldNot(gomega.Equal(int32(5)))
		gomega.Expect(a).ShouldNot(gomega.Equal(uint8(5)))
		gomega.Expect(a).ShouldNot(gomega.Equal(mytype(5)))
		gomega.Expect(5).ShouldNot(gomega.Equal(mytype(5)))

		b := int16(5)
		gomega.Expect(a).ShouldNot(gomega.Equal(b))

		c := mytype(5)
		gomega.Expect(a).ShouldNot(gomega.Equal(c))
	})

	It("compare interfaces", func() {
		var (
			a myinf = imp1(3)
			b myinf = imp2(3)
		)
		gomega.Expect(a).ShouldNot(gomega.Equal(b))
		gomega.Expect(a).ShouldNot(gomega.BeEquivalentTo(b))
		gomega.Expect(a).ShouldNot(gomega.BeIdenticalTo(b))
	})

	It("check HaveValue(Equal)", func() {
		a := 5
		pa := &a

		gomega.Expect(pa).Should(gomega.HaveValue(gomega.Equal(uint64(5))))
		gomega.Expect(a).Should(gomega.Not(gomega.Equal(uint64(5))))
		gomega.Expect(a).Should(gomega.And(gomega.Equal(uint64(5)), gomega.Not(gomega.Equal(int32(6)))))
		gomega.Expect(a).Should(gomega.Or(gomega.Equal(uint64(5)), gomega.Not(gomega.Equal(int32(6))), gomega.Not(gomega.Equal(int8(4)))))
		gomega.Expect(a).Should(gomega.WithTransform(func(i int) int { return i + 1 }, gomega.Equal(uint(6))))
	})
})
