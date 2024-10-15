package errnil

import (
	"a/mytypes"
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
		e1 := &mytypes.MyErr{}
		Expect(nil == e1).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion. Consider using .Expect\(e1\)\.ToNot\(BeNil\(\)\). instead`
		Expect(e1 == nil).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion. Consider using .Expect\(e1\)\.ToNot\(BeNil\(\)\). instead`

		Expect(errors.New("fake error")).To(Equal(nil))
		Expect(err).To(BeNil())
		ExpectWithOffset(1, err).To(BeNil())
		Expect(err).To(Not(BeNil()))
		Expect(err).ToNot(BeNil())

		Expect(tt.err).To(BeNil())

		Expect(errFunc()).To(BeNil())
		Expect(errFunc()).To(Not(BeNil()))
		Expect(tupleFunc()).ToNot(BeNil())
		Expect(noErrorFunc()).ToNot(Equal(1))

		Expect(tt.typeErrorFunc()).ToNot(BeNil())
		Expect(tt.typeTupleFunc()).ToNot(BeNil())
		Expect(tt.typeNoErrorFunc()).ToNot(Equal(1))
	})
	It("check Expect(err).To(Equal(nil))", func() {
		Expect(err).To(Equal(nil))
		Expect(err).To(Not(Equal(nil)))
		Expect(err).ToNot(Equal(nil))
		Expect(tt.err).To(Equal(nil))

		Expect(errFunc()).To(Equal(nil))
		Expect(errFunc()).To(Not(Equal(nil)))
		Expect(tupleFunc()).ToNot(Equal(nil))
		Expect(noErrorFunc()).ToNot(Equal(1))

		Expect(tt.typeErrorFunc()).ToNot(Equal(nil))
		Expect(tt.typeTupleFunc()).ToNot(Equal(nil))
		Expect(tt.typeTupleFunc()).ToNot(HaveOccurred())
		Expect(tt.typeTupleFunc()).To(Succeed())
		Expect(tt.typeNoErrorFunc()).ToNot(Equal(1))
	})

	It("check err == nil", func() {
		Expect(err == nil).To(Equal(true))       // want `ginkgo-linter: wrong nil assertion\. Consider using .Expect\(err\)\.To\(BeNil\(\)\). instead`
		Expect(err == nil).To(Equal(false))      // want `ginkgo-linter: wrong nil assertion\. Consider using .Expect\(err\)\.ToNot\(BeNil\(\)\). instead`
		Expect(err != nil).To(Equal(true))       // want `ginkgo-linter: wrong nil assertion\. Consider using .Expect\(err\)\.ToNot\(BeNil\(\)\). instead`
		Expect(err != nil).To(Equal(false))      // want `ginkgo-linter: wrong nil assertion\. Consider using .Expect\(err\)\.To\(BeNil\(\)\). instead`
		Expect(nil == err).To(Equal(true))       // want `ginkgo-linter: wrong nil assertion\. Consider using .Expect\(err\)\.To\(BeNil\(\)\). instead`
		Expect(nil == errFunc()).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion\. Consider using .Expect\(errFunc\(\)\)\.To\(BeNil\(\)\). instead`
		Expect(errFunc() != nil).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion\. Consider using .Expect\(errFunc\(\)\)\.ToNot\(BeNil\(\)\). instead`

		Expect(err == nil).To(BeTrue())        // want `ginkgo-linter: wrong nil assertion\. Consider using .Expect\(err\)\.To\(BeNil\(\)\). instead`
		Expect(err != nil).To(BeFalse())       // want `ginkgo-linter: wrong nil assertion\. Consider using .Expect\(err\)\.To\(BeNil\(\)\). instead`
		Expect(nil == err).To(BeTrue())        // want `ginkgo-linter: wrong nil assertion\. Consider using .Expect\(err\)\.To\(BeNil\(\)\). instead`
		Expect(nil == errFunc()).To(BeTrue())  // want `ginkgo-linter: wrong nil assertion\. Consider using .Expect\(errFunc\(\)\)\.To\(BeNil\(\)\). instead`
		Expect(errFunc() != nil).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion\. Consider using .Expect\(errFunc\(\)\)\.To\(BeNil\(\)\). instead`
	})
})
