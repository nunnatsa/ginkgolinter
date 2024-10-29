package errnil

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("check Expect(err).To(BeNil())", func() {
	var err error
	tt := t{}

	It("check Expect(err).To(BeNil())", func() {
		gomega.Expect(errors.New("fake error")).To(gomega.Equal(nil)) // want `wrong nil assertion. Consider using .gomega\.Expect\(errors\.New\("fake error"\)\)\.To\(gomega\.BeNil\(\)\). instead`
		gomega.Expect(err).To(gomega.BeNil())
		gomega.ExpectWithOffset(1, err).To(gomega.BeNil())
		gomega.Expect(err).To(gomega.Not(gomega.BeNil()))
		gomega.Expect(err).ToNot(gomega.BeNil())

		gomega.Expect(tt.err).To(gomega.BeNil())

		gomega.Expect(errFunc()).To(gomega.BeNil())
		gomega.Expect(errFunc()).To(gomega.Not(gomega.BeNil()))
		gomega.Expect(tupleFunc()).ToNot(gomega.BeNil())
		gomega.Expect(noErrorFunc()).ToNot(gomega.Equal(1))

		gomega.Expect(tt.typeErrorFunc()).ToNot(gomega.BeNil())
		gomega.Expect(tt.typeTupleFunc()).ToNot(gomega.BeNil())
		gomega.Expect(tt.typeNoErrorFunc()).ToNot(gomega.Equal(1))
	})

	It("check err == nil", func() {
		gomega.Expect(err == nil).To(gomega.BeTrue())        // want `ginkgo-linter: wrong nil assertion\. Consider using .gomega\.Expect\(err\)\.To\(gomega\.BeNil\(\)\). instead`
		gomega.Expect(err != nil).To(gomega.BeFalse())       // want `ginkgo-linter: wrong nil assertion\. Consider using .gomega\.Expect\(err\)\.To\(gomega\.BeNil\(\)\). instead`
		gomega.Expect(nil == err).To(gomega.BeTrue())        // want `ginkgo-linter: wrong nil assertion\. Consider using .gomega\.Expect\(err\)\.To\(gomega\.BeNil\(\)\). instead`
		gomega.Expect(nil == errFunc()).To(gomega.BeTrue())  // want `ginkgo-linter: wrong nil assertion\. Consider using .gomega\.Expect\(errFunc\(\)\)\.To\(gomega\.BeNil\(\)\). instead`
		gomega.Expect(errFunc() != nil).To(gomega.BeFalse()) // want `ginkgo-linter: wrong nil assertion\. Consider using .gomega\.Expect\(errFunc\(\)\)\.To\(gomega\.BeNil\(\)\). instead`
	})
})
