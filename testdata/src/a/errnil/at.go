package errnil

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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

var _ = Describe("check Expect(err).To(BeNil())", func() {
	var err error
	tt := t{}

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
