package pointers

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type bbber interface {
	bbb()
}

type inner struct {
	i *int
}

type ttt struct {
	in *inner
}

func (*ttt) bbb() {}

func GetBBB() bbber {
	var i = 8
	return &ttt{in: &inner{i: &i}}
}

var _ = Describe("", func() {
	It("", func() {
		val1, val2 := 8, 8
		p1, p2, p3 := &val1, &val1, &val2
		Expect(p1 == p2).To(BeTrue())  // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(p1\)\.To\(BeIdenticalTo\(p2\)\). instead`
		Expect(p1 == p3).To(BeFalse()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(p1\)\.ToNot\(BeIdenticalTo\(p3\)\). instead`
	})

	It("", func() {
		var (
			v1 = 8
			v2 = 8
		)
		t1 := &ttt{in: &inner{i: &v1}}
		t2 := &ttt{in: &inner{i: &v1}}
		t3 := &ttt{in: &inner{i: &v2}}

		Expect(t1 != t2).To(BeTrue())           // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(t1\)\.ToNot\(BeIdenticalTo\(t2\)\). instead`
		Expect(t1.in.i == t2.in.i).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(t1.in.i\)\.To\(BeIdenticalTo\(t2.in.i\)\). instead`
		Expect(t1.in.i != t3.in.i).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(t1.in.i\)\.ToNot\(BeIdenticalTo\(t3.in.i\)\). instead`

		t4 := GetBBB()
		Expect(t4 == t1).To(BeFalse()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(t4\)\.ToNot\(BeIdenticalTo\(t1\)\). instead`
	})
})
