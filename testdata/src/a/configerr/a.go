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
		Expect(err == nil).To(Equal(true))
		Expect(err == nil).To(Equal(false))
		Expect(err != nil).To(Equal(true))
		Expect(err != nil).To(Equal(false))
		Expect(nil == err).To(Equal(true))
		Expect(nil == errFunc()).To(Equal(true))
		Expect(errFunc() != nil).To(Equal(true))

		Expect(err == nil).To(BeTrue())
		Expect(err != nil).To(BeFalse())
		Expect(nil == err).To(BeTrue())
		Expect(nil == errFunc()).To(BeTrue())
		Expect(errFunc() != nil).To(BeFalse())
	})
})
