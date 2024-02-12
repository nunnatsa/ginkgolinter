package pointers

import (
	. "github.com/onsi/ginkgo/v2"
	g "github.com/onsi/gomega"
)

var _ = Describe("", func() {
	It("", func() {
		val1, val2 := 8, 8
		p1, p2, p3 := &val1, &val1, &val2
		g.Expect(p1 == p2).To(g.BeTrue())  // want `ginkgo-linter: wrong comparison assertion\. Consider using .g\.Expect\(p1\)\.To\(g\.BeIdenticalTo\(p2\)\). instead`
		g.Expect(p1 == p3).To(g.BeFalse()) // want `ginkgo-linter: wrong comparison assertion\. Consider using .g\.Expect\(p1\)\.ToNot\(g\.BeIdenticalTo\(p3\)\). instead`
	})

	It("", func() {
		var (
			v1 = 8
			v2 = 8
		)
		t1 := &ttt{in: &inner{i: &v1}}
		t2 := &ttt{in: &inner{i: &v1}}
		t3 := &ttt{in: &inner{i: &v2}}

		g.Expect(t1 != t2).To(g.BeTrue())           // want `ginkgo-linter: wrong comparison assertion\. Consider using .g\.Expect\(t1\)\.ToNot\(g\.BeIdenticalTo\(t2\)\). instead`
		g.Expect(t1.in.i == t2.in.i).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion\. Consider using .g\.Expect\(t1\.in\.i\)\.To\(g\.BeIdenticalTo\(t2\.in\.i\)\). instead`
		g.Expect(t1.in.i != t3.in.i).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion\. Consider using .g\.Expect\(t1\.in\.i\)\.ToNot\(g\.BeIdenticalTo\(t3\.in\.i\)\). instead`

		t4 := GetBBB()
		g.Expect(t4 == t1).To(g.BeFalse()) // want `ginkgo-linter: wrong comparison assertion\. Consider using .g\.Expect\(t4\)\.ToNot\(g\.BeIdenticalTo\(t1\)\). instead`

	})
})
