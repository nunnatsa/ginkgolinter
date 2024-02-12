package pointerval

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const c = "a"

type strct struct {
	field string
}

type strctPoint struct {
	field *string
}

func retPointer() *string {
	v := c
	return &v
}

type str string

func (s *str) String() string {
	if s == nil {
		return ""
	}
	return string(*s)
}

type strIfs struct {
	field fmt.Stringer
}

func retStringer() fmt.Stringer {
	v := str(c)
	return &v
}

var (
	v = c
	p = &v
)
var _ = Describe("", Label("pointers1"), func() {

	It("pointer to var", func() {
		Expect(*p).Should(Equal(v))
		Expect(p).Should(HaveValue(Equal(v)))
		Expect(&v).Should(Equal(c)) // want `ginkgo-linter: comparing a pointer to a value will always fail\. Consider using .Expect\(&v\)\.Should\(HaveValue\(Equal\(c\)\)\). instead`
		Expect(p).Should(Equal(&v))
		Expect(p).Should(Equal(v)) // want `ginkgo-linter: comparing a pointer to a value will always fail\. Consider using .Expect\(p\)\.Should\(HaveValue\(Equal\(v\)\)\). instead`
	})

	It("with description", func() {
		Expect(p).Should(Equal(v), "%d", 3) // want `ginkgo-linter: comparing a pointer to a value will always fail\. Consider using .Expect\(p\)\.Should\(HaveValue\(Equal\(v\)\), "%d", 3\). instead`
	})

	It("pointer to const", func() {
		Expect(*p).Should(Equal(c))
		Expect(p).Should(HaveValue(Equal(c)))
		Expect(p).Should(Equal(c)) // want `ginkgo-linter: comparing a pointer to a value will always fail\. Consider using .Expect\(p\)\.Should\(HaveValue\(Equal\(c\)\)\). instead`
	})

	It("function", func() {
		Expect(retPointer()).Should(HaveValue(Equal(c))) // valid
		Expect(*retPointer()).Should(Equal(c))           // valid
		Expect(retPointer()).Should(Equal(c))            // want `ginkgo-linter: comparing a pointer to a value will always fail\. Consider using .Expect\(retPointer\(\)\)\.Should\(HaveValue\(Equal\(c\)\)\). instead`
	})

	Context("struct", func() {
		s1 := strct{field: c}
		s2 := &strct{field: c}
		s3 := strctPoint{field: &v}
		s4 := &strctPoint{field: &v}

		It("struct with non-pointer field", func() {
			Expect(s1.field).Should(Equal(c)) // valid
		})
		It("struct pointer with non-pointer field", func() {
			Expect(s2.field).Should(Equal(c)) // valid
		})
		It("struct with pointer field", func() {
			Expect(*s3.field).Should(Equal(c))           // valid
			Expect(s3.field).Should(HaveValue(Equal(c))) // valid
			Expect(s3.field).Should(Equal(c))            // want `ginkgo-linter: comparing a pointer to a value will always fail\. Consider using .Expect\(s3\.field\)\.Should\(HaveValue\(Equal\(c\)\)\). instead`
			Expect(s3.field).Should(BeIdenticalTo(p))
			Expect(s3.field).ShouldNot(BeIdenticalTo(retStringer())) // want `ginkgo-linter: use BeIdenticalTo with different types: Comparing \*string with fmt\.Stringer; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of BeIdenticalTo\(\)`
		})
		It("pointer struct with pointer field", func() {
			Expect(*s4.field).Should(Equal(c))           // valid
			Expect(s4.field).Should(HaveValue(Equal(c))) // valid
			Expect(s4.field).Should(Equal(c))            // want `ginkgo-linter: comparing a pointer to a value will always fail\. Consider using .Expect\(s4\.field\)\.Should\(HaveValue\(Equal\(c\)\)\). instead`
		})
	})

	Context("struct with interface", func() {
		sv := str(v)
		s1 := strIfs{field: &sv}
		s2 := &strIfs{field: &sv}

		It("struct with pointer field", func() {
			Expect(s1.field).Should(HaveValue(Equal(c))) // want `ginkgo-linter: use Equal with different types: Comparing fmt\.Stringer with string; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
			Expect(s1.field).Should(Equal(c))            // want `ginkgo-linter: use Equal with different types: Comparing fmt\.Stringer with string; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
			Expect(s1.field).Should(BeIdenticalTo(sv))   // want `ginkgo-linter: use BeIdenticalTo with different types: Comparing fmt\.Stringer with a/pointerval\.str; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of BeIdenticalTo\(\)`
		})
		It("pointer struct with pointer field", func() {
			Expect(s2.field).Should(HaveValue(Equal(c))) // want `ginkgo-linter: use Equal with different types: Comparing fmt\.Stringer with string; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
			Expect(s2.field).Should(Equal(c))            // want `ginkgo-linter: use Equal with different types: Comparing fmt\.Stringer with string; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of Equal\(\)`
		})
		It("compare to interface", func() {
			Expect(&sv).Should(Equal(retStringer()))
		})
	})

	Context("negative", func() {
		var n *string
		Expect(n).ShouldNot(Equal(c)) // want `ginkgo-linter: comparing a pointer to a value will always fail\. Consider using .Expect\(n\)\.ShouldNot\(HaveValue\(Equal\(c\)\)\). instead`
		Expect(n).To(Not(Equal(c)))   // want `ginkgo-linter: comparing a pointer to a value will always fail\. Consider using .Expect\(n\)\.ToNot\(HaveValue\(Equal\(c\)\)\). instead`
	})

	It("do not add HaveValue for nil", func() {
		Expect(p).ShouldNot(Equal(nil)) // want `ginkgo-linter: wrong nil assertion\. Consider using .Expect\(p\)\.ShouldNot\(BeNil\(\)\). instead`
		Expect(p).ShouldNot(BeNil())    // valid
	})
	Context("boolean", func() {
		t := true
		f := false
		pt := &t
		pf := &f

		It("true", func() {
			Expect(pt).Should(Equal(true))     // want `ginkgo-linter: multiple issues: wrong boolean assertion; comparing a pointer to a value will always fail\. Consider using .Expect\(pt\)\.Should\(HaveValue\(BeTrue\(\)\)\). instead`
			Expect(pt).ShouldNot(Equal(false)) // want `ginkgo-linter: multiple issues: wrong boolean assertion; comparing a pointer to a value will always fail\. Consider using .Expect\(pt\)\.Should\(HaveValue\(BeTrue\(\)\)\). instead`
		})

		It("BeTrue", func() {
			Expect(pt).Should(BeTrue())     // want `ginkgo-linter: comparing a pointer to a value will always fail\. Consider using .Expect\(pt\)\.Should\(HaveValue\(BeTrue\(\)\)\). instead`
			Expect(pt).ShouldNot(BeFalse()) // want `ginkgo-linter: multiple issues: avoid double negative assertion; comparing a pointer to a value will always fail\. Consider using .Expect\(pt\)\.Should\(HaveValue\(BeTrue\(\)\)\). instead`
		})

		It("false", func() {
			Expect(pf).Should(Equal(false))   // want `ginkgo-linter: multiple issues: wrong boolean assertion; comparing a pointer to a value will always fail\. Consider using .Expect\(pf\)\.Should\(HaveValue\(BeFalse\(\)\)\). instead`
			Expect(pf).ShouldNot(Equal(true)) // want `ginkgo-linter: multiple issues: wrong boolean assertion; comparing a pointer to a value will always fail\. Consider using .Expect\(pf\)\.ShouldNot\(HaveValue\(BeTrue\(\)\)\). instead`
		})

		It("Not with boolean", func() {
			Expect(pf).Should(Not(Equal(true)))  // want `ginkgo-linter: multiple issues: wrong boolean assertion; comparing a pointer to a value will always fail\. Consider using .Expect\(pf\)\.ShouldNot\(HaveValue\(BeTrue\(\)\)\). instead`
			Expect(pt).Should(Not(Equal(false))) // want `ginkgo-linter: multiple issues: wrong boolean assertion; comparing a pointer to a value will always fail\. Consider using .Expect\(pt\)\.Should\(HaveValue\(BeTrue\(\)\)\). instead`
		})
	})

	Context("Other matchers", func() {
		x := float64(5)
		y := x
		px1 := &x
		px2 := &x
		It("BeEquivalentTo", func() {
			Expect(px1).Should(BeEquivalentTo(5))      // want `ginkgo-linter: comparing a pointer to a value will always fail\. Consider using .Expect\(px1\)\.Should\(HaveValue\(BeEquivalentTo\(5\)\)\). instead`
			Expect(px1).Should(BeEquivalentTo(y))      // want `ginkgo-linter: comparing a pointer to a value will always fail\. Consider using .Expect\(px1\)\.Should\(HaveValue\(BeEquivalentTo\(y\)\)\). instead`
			Expect(px1).Should(BeEquivalentTo(&x))     // valid - compare two pointers
			Expect(px1).Should(BeEquivalentTo(px2))    // valid - compare two pointers
			Expect(px1).ShouldNot(BeEquivalentTo(nil)) // valid
		})
		It("BeIdenticalTo", func() {
			Expect(px1).ShouldNot(BeIdenticalTo(5))       // want `ginkgo-linter: comparing a pointer to a value will always fail\. Consider using .Expect\(px1\)\.ShouldNot\(HaveValue\(BeIdenticalTo\(5\)\)\). instead`
			Expect(px1).ShouldNot(BeIdenticalTo(px2))     // valid
			Expect(px1).ShouldNot(BeIdenticalTo(nil))     // want `ginkgo-linter: use BeIdenticalTo with different types: Comparing \*float64 with untyped nil; either change the expected value type if possible, or use the BeEquivalentTo\(\) matcher, instead of BeIdenticalTo\(\)`
			Expect(px1).Should(BeIdenticalTo(float64(5))) // want `ginkgo-linter: comparing a pointer to a value will always fail\. Consider using .Expect\(px1\)\.Should\(HaveValue\(BeIdenticalTo\(float64\(5\)\)\)\). instead`
		})
		It("", func() {
			Expect(px1).Should(BeNumerically(">", 4))    // want `ginkgo-linter: comparing a pointer to a value will always fail\. Consider using .Expect\(px1\)\.Should\(HaveValue\(BeNumerically\(">", 4\)\)\). instead`
			Expect(px1).ShouldNot(BeNumerically(">", 6)) // want `ginkgo-linter: comparing a pointer to a value will always fail\. Consider using .Expect\(px1\)\.ShouldNot\(HaveValue\(BeNumerically\(">", 6\)\)\). instead`
		})
	})
})
