package pointerval

import (
	. "github.com/onsi/ginkgo/v2"
	g "github.com/onsi/gomega"
)

var _ = Describe("", Label("pointers1"), func() {

	It("pointer to var", func() {
		g.Expect(*p).Should(g.Equal(v))
		g.Expect(p).Should(g.HaveValue(g.Equal(v)))
		g.Expect(p).Should(g.Equal(v)) // want `ginkgo-linter: comparing a pointer to a value will always fail. consider using .g\.Expect\(p\)\.Should\(g\.HaveValue\(g\.Equal\(v\)\)\). instead`
	})

	It("pointer to const", func() {
		g.Expect(*p).Should(g.Equal(c))
		g.Expect(p).Should(g.HaveValue(g.Equal(c)))
		g.Expect(p).Should(g.Equal(c)) // want `ginkgo-linter: comparing a pointer to a value will always fail. consider using .g\.Expect\(p\)\.Should\(g\.HaveValue\(g\.Equal\(c\)\)\). instead`
	})

	It("function", func() {
		g.Expect(retPointer()).Should(g.HaveValue(g.Equal(c))) // valid
		g.Expect(*retPointer()).Should(g.Equal(c))             // valid
		g.Expect(retPointer()).Should(g.Equal(c))              // want `ginkgo-linter: comparing a pointer to a value will always fail. consider using .g\.Expect\(retPointer\(\)\)\.Should\(g\.HaveValue\(g\.Equal\(c\)\)\). instead`
	})

	Context("struct", func() {
		s1 := strct{field: c}
		s2 := &strct{field: c}
		s3 := strctPoint{field: &v}
		s4 := &strctPoint{field: &v}

		It("struct with non-pointer field", func() {
			g.Expect(s1.field).Should(g.Equal(c)) // valid
		})
		It("struct pointer with non-pointer field", func() {
			g.Expect(s2.field).Should(g.Equal(c)) // valid
		})
		It("struct with pointer field", func() {
			g.Expect(*s3.field).Should(g.Equal(c))             // valid
			g.Expect(s3.field).Should(g.HaveValue(g.Equal(c))) // valid
			g.Expect(s3.field).Should(g.Equal(c))              // want `ginkgo-linter: comparing a pointer to a value will always fail. consider using .g\.Expect\(s3\.field\)\.Should\(g\.HaveValue\(g\.Equal\(c\)\)\). instead`
		})
		It("pointer struct with pointer field", func() {
			g.Expect(*s4.field).Should(g.Equal(c))             // valid
			g.Expect(s4.field).Should(g.HaveValue(g.Equal(c))) // valid
			g.Expect(s4.field).Should(g.Equal(c))              // want `ginkgo-linter: comparing a pointer to a value will always fail. consider using .g\.Expect\(s4\.field\)\.Should\(g\.HaveValue\(g\.Equal\(c\)\)\). instead`
		})
	})

	It("do not add HaveValue for nil", func() {
		g.Expect(p).ShouldNot(g.Equal(nil)) // want `ginkgo-linter: wrong nil assertion; consider using .g\.Expect\(p\)\.ShouldNot\(g\.BeNil\(\)\). instead`
		g.Expect(p).ShouldNot(g.BeNil())    // valid
	})

	Context("boolean", func() {
		t := true
		f := false
		pt := &t
		pf := &f

		It("true", func() {
			g.Expect(pt).Should(g.Equal(true))     // want `ginkgo-linter: comparing a pointer to a value will always fail. consider using .g\.Expect\(pt\)\.Should\(g\.HaveValue\(g\.BeTrue\(\)\)\). instead`
			g.Expect(pt).ShouldNot(g.Equal(false)) // want `ginkgo-linter: comparing a pointer to a value will always fail. consider using .g\.Expect\(pt\)\.Should\(g\.HaveValue\(g\.BeTrue\(\)\)\). instead`
		})

		It("BeTrue", func() {
			g.Expect(pt).Should(g.BeTrue())     // want `ginkgo-linter: comparing a pointer to a value will always fail. consider using .g\.Expect\(pt\)\.Should\(g\.HaveValue\(g\.BeTrue\(\)\)\). instead`
			g.Expect(pt).ShouldNot(g.BeFalse()) // want `ginkgo-linter: comparing a pointer to a value will always fail. consider using .g\.Expect\(pt\)\.Should\(g\.HaveValue\(g\.BeTrue\(\)\)\). instead`
		})

		It("false", func() {
			g.Expect(pf).Should(g.Equal(false))   // want `ginkgo-linter: comparing a pointer to a value will always fail. consider using .g\.Expect\(pf\)\.Should\(g\.HaveValue\(g\.BeFalse\(\)\)\). instead`
			g.Expect(pf).ShouldNot(g.Equal(true)) // want `ginkgo-linter: comparing a pointer to a value will always fail. consider using .g\.Expect\(pf\)\.ShouldNot\(g\.HaveValue\(g\.BeTrue\(\)\)\). instead`
		})
	})
})
