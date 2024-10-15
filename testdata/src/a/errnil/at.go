package errnil

import (
	"errors"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"a/mytypes"
)

func errFunc() error {
	return nil
}

func tupleFunc() (error, fmt.Stringer) {
	return nil, nil
}

func noErrorFunc() int {
	return 1
}

type t struct {
	err error
}

func (t) typeErrorFunc() error {
	return nil
}

func (t) typeTupleFunc() (error, int) {
	return errors.New("fake error"), 1
}

func (t) typeNoErrorFunc() int {
	return 1
}

const zero = 0

var _ = Describe("check Expect(err).To(BeNil())", func() {
	var err error
	tt := t{}

	It("check custom errors", func() {
		var arr []int
		Expect(zero == 0).To(BeTrue())               // want `ginkgo-linter: wrong comparison assertion. Consider using .Expect\(0\)\.To\(BeZero\(\)\). instead`
		Expect(len(arr) == 0.0).To(BeTrue())         // want `ginkgo-linter: wrong length assertion. Consider using .Expect\(arr\)\.To\(BeEmpty\(\)\). instead`
		Expect(len(arr) == zero).To(BeTrue())        // want `ginkgo-linter: wrong length assertion. Consider using .Expect\(arr\)\.To\(BeEmpty\(\)\). instead`
		Expect(mytypes.MyError == nil).To(BeFalse()) // want `ginkgo-linter: wrong error assertion. Consider using .Expect\(mytypes\.MyError\)\.To\(HaveOccurred\(\)\). instead`
		Expect(nil == mytypes.MyError).To(BeFalse()) // want `ginkgo-linter: wrong error assertion. Consider using .Expect\(mytypes\.MyError\)\.To\(HaveOccurred\(\)\). instead`

		e1 := &mytypes.MyErr{}
		Expect(nil == e1).To(BeFalse()) // want `ginkgo-linter: wrong error assertion. Consider using .Expect\(e1\)\.To\(HaveOccurred\(\)\). instead`
		Expect(e1 == nil).To(BeFalse()) // want `ginkgo-linter: wrong error assertion. Consider using .Expect\(e1\)\.To\(HaveOccurred\(\)\). instead`

		e2 := os.ErrNoDeadline
		Expect(e2 == nil).To(BeFalse()) // want `ginkgo-linter: wrong error assertion. Consider using .Expect\(e2\)\.To\(HaveOccurred\(\)\). instead`
		Expect(nil == e2).To(BeFalse()) // want `ginkgo-linter: wrong error assertion. Consider using .Expect\(e2\)\.To\(HaveOccurred\(\)\). instead`
	})

	It("check Expect(err).To(BeNil())", func() {
		Expect(errors.New("fake error")).To(Equal(nil)) // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(errors\.New\("fake error"\)\)\.To\(Succeed\(\)\). instead`
		Expect(err).To(BeNil())                         // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(err\)\.ToNot\(HaveOccurred\(\)\). instead`
		ExpectWithOffset(1, err).To(BeNil())            // want `ginkgo-linter: wrong error assertion\. Consider using .ExpectWithOffset\(1, err\)\.ToNot\(HaveOccurred\(\)\). instead`
		Expect(err).To(Not(BeNil()))                    // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(err\)\.To\(HaveOccurred\(\)\). instead`
		Expect(err).ToNot(BeNil())                      // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(err\)\.To\(HaveOccurred\(\)\). instead`
		Expect(err).WithOffset(1).ToNot(BeNil())        // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(err\)\.WithOffset\(1\)\.To\(HaveOccurred\(\)\). instead`

		Expect(tt.err).To(BeNil()) // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(tt\.err\)\.ToNot\(HaveOccurred\(\)\). instead`

		Expect(errFunc()).To(BeNil())      // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(errFunc\(\)\)\.To\(Succeed\(\)\). instead`
		Expect(errFunc()).To(Not(BeNil())) // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(errFunc\(\)\)\.ToNot\(Succeed\(\)\). instead`
		Expect(tupleFunc()).ToNot(BeNil()) // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(tupleFunc\(\)\)\.ToNot\(Succeed\(\)\). instead`
		Expect(noErrorFunc()).ToNot(Equal(1))

		Expect(tt.typeErrorFunc()).ToNot(BeNil()) // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(tt\.typeErrorFunc\(\)\)\.ToNot\(Succeed\(\)\). instead`
		Expect(tt.typeTupleFunc()).ToNot(BeNil()) // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(tt\.typeTupleFunc\(\)\)\.ToNot\(Succeed\(\)\). instead`
		Expect(tt.typeNoErrorFunc()).ToNot(Equal(1))
	})
	It("check Expect(err).To(Equal(nil))", func() {
		Expect(err).To(Equal(nil))                  // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(err\)\.ToNot\(HaveOccurred\(\)\). instead`
		Expect(err).To(Not(Equal(nil)))             // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(err\)\.To\(HaveOccurred\(\)\). instead`
		Expect(err).ToNot(Equal(nil))               // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(err\)\.To\(HaveOccurred\(\)\). instead`
		Expect(tt.err).To(Equal(nil))               // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(tt\.err\)\.ToNot\(HaveOccurred\(\)\). instead`
		Expect(tt.err).WithOffset(1).To(Equal(nil)) // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(tt\.err\)\.WithOffset\(1\)\.ToNot\(HaveOccurred\(\)\). instead`

		Expect(errFunc()).To(Equal(nil))      // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(errFunc\(\)\)\.To\(Succeed\(\)\). instead`
		Expect(errFunc()).To(Not(Equal(nil))) // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(errFunc\(\)\)\.ToNot\(Succeed\(\)\). instead`
		Expect(tupleFunc()).ToNot(Equal(nil)) // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(tupleFunc\(\)\)\.ToNot\(Succeed\(\)\). instead`
		Expect(noErrorFunc()).ToNot(Equal(1))

		Expect(tt.typeErrorFunc()).ToNot(Equal(nil)) // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(tt\.typeErrorFunc\(\)\)\.ToNot\(Succeed\(\)\). instead`
		Expect(tt.typeTupleFunc()).ToNot(Equal(nil)) // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(tt\.typeTupleFunc\(\)\)\.ToNot\(Succeed\(\)\). instead`
		Expect(tt.typeTupleFunc()).ToNot(HaveOccurred())
		Expect(tt.typeTupleFunc()).To(Succeed())
		Expect(tt.typeNoErrorFunc()).ToNot(Equal(1))
	})

	It("check err == nil", func() {
		Expect(err == nil).To(Equal(true))               // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(err\)\.ToNot\(HaveOccurred\(\)\). instead`
		Expect(err == nil).To(Equal(false))              // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(err\)\.To\(HaveOccurred\(\)\). instead`
		Expect(err != nil).To(Equal(true))               // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(err\)\.To\(HaveOccurred\(\)\). instead`
		Expect(err != nil).To(Equal(false))              // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(err\)\.ToNot\(HaveOccurred\(\)\). instead`
		Expect(nil == err).To(Equal(true))               // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(err\)\.ToNot\(HaveOccurred\(\)\). instead`
		Expect(nil == err).WithOffset(1).To(Equal(true)) // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(err\)\.WithOffset\(1\)\.ToNot\(HaveOccurred\(\)\). instead`
		Expect(nil == errFunc()).To(Equal(true))         // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(errFunc\(\)\)\.To\(Succeed\(\)\). instead`
		Expect(errFunc() != nil).To(Equal(true))         // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(errFunc\(\)\)\.ToNot\(Succeed\(\)\). instead`

		Expect(err == nil).To(BeTrue())                      // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(err\)\.ToNot\(HaveOccurred\(\)\). instead`
		Expect(err != nil).To(BeFalse())                     // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(err\)\.ToNot\(HaveOccurred\(\)\). instead`
		Expect(nil == err).To(BeTrue())                      // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(err\)\.ToNot\(HaveOccurred\(\)\). instead`
		Expect(nil == errFunc()).To(BeTrue())                // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(errFunc\(\)\)\.To\(Succeed\(\)\). instead`
		Expect(errFunc() != nil).To(BeFalse())               // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(errFunc\(\)\)\.To\(Succeed\(\)\). instead`
		Expect(errFunc() != nil).WithOffset(1).To(BeFalse()) // want `ginkgo-linter: wrong error assertion\. Consider using .Expect\(errFunc\(\)\)\.WithOffset\(1\)\.To\(Succeed\(\)\). instead`
	})
})
