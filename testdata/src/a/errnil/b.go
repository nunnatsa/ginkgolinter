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
		gomega.Expect(errors.New("fake error")).To(gomega.Equal(nil)) // want `ginkgo-linter: wrong error assertion\nconsider using .gomega\.Expect\(errors\.New\("fake error"\)\)\.To\(gomega\.Succeed\(\)\). instead`
		gomega.Expect(err).To(gomega.BeNil())                         // want `ginkgo-linter: wrong error assertion\nconsider using .gomega\.Expect\(err\)\.ToNot\(gomega\.HaveOccurred\(\)\). instead`
		gomega.ExpectWithOffset(1, err).To(gomega.BeNil())            // want `ginkgo-linter: wrong error assertion\nconsider using .gomega\.ExpectWithOffset\(1, err\)\.ToNot\(gomega\.HaveOccurred\(\)\). instead`
		gomega.Expect(err).To(gomega.Not(gomega.BeNil()))             // want `ginkgo-linter: wrong error assertion\nconsider using .gomega\.Expect\(err\)\.To\(gomega\.HaveOccurred\(\)\). instead`
		gomega.Expect(err).ToNot(gomega.BeNil())                      // want `ginkgo-linter: wrong error assertion\nconsider using .gomega\.Expect\(err\)\.To\(gomega\.HaveOccurred\(\)\). instead`

		gomega.Expect(tt.err).To(gomega.BeNil()) // want `ginkgo-linter: wrong error assertion\nconsider using .gomega\.Expect\(tt\.err\)\.ToNot\(gomega\.HaveOccurred\(\)\). instead`

		gomega.Expect(errFunc()).To(gomega.BeNil())             // want `ginkgo-linter: wrong error assertion\nconsider using .gomega\.Expect\(errFunc\(\)\)\.To\(gomega\.Succeed\(\)\). instead`
		gomega.Expect(errFunc()).To(gomega.Not(gomega.BeNil())) // want `ginkgo-linter: wrong error assertion\nconsider using .gomega\.Expect\(errFunc\(\)\)\.ToNot\(gomega\.Succeed\(\)\). instead`
		gomega.Expect(tupleFunc()).ToNot(gomega.BeNil())        // want `ginkgo-linter: wrong error assertion\nconsider using .gomega\.Expect\(tupleFunc\(\)\)\.ToNot\(gomega\.Succeed\(\)\). instead`
		gomega.Expect(noErrorFunc()).ToNot(gomega.Equal(1))

		gomega.Expect(tt.typeErrorFunc()).ToNot(gomega.BeNil()) // want `ginkgo-linter: wrong error assertion\nconsider using .gomega\.Expect\(tt\.typeErrorFunc\(\)\)\.ToNot\(gomega\.Succeed\(\)\). instead`
		gomega.Expect(tt.typeTupleFunc()).ToNot(gomega.BeNil()) // want `ginkgo-linter: wrong error assertion\nconsider using .gomega\.Expect\(tt\.typeTupleFunc\(\)\)\.ToNot\(gomega\.Succeed\(\)\). instead`
		gomega.Expect(tt.typeNoErrorFunc()).ToNot(gomega.Equal(1))
	})

	It("check err == nil", func() {
		gomega.Expect(err == nil).To(gomega.BeTrue())        // want `ginkgo-linter: wrong error assertion\nconsider using .gomega\.Expect\(err\)\.ToNot\(gomega\.HaveOccurred\(\)\). instead`
		gomega.Expect(err != nil).To(gomega.BeFalse())       // want `ginkgo-linter: wrong error assertion\nconsider using .gomega\.Expect\(err\)\.ToNot\(gomega\.HaveOccurred\(\)\). instead`
		gomega.Expect(nil == err).To(gomega.BeTrue())        // want `ginkgo-linter: wrong error assertion\nconsider using .gomega\.Expect\(err\)\.ToNot\(gomega\.HaveOccurred\(\)\). instead`
		gomega.Expect(nil == errFunc()).To(gomega.BeTrue())  // want `ginkgo-linter: wrong error assertion\nconsider using .gomega\.Expect\(errFunc\(\)\)\.To\(gomega\.Succeed\(\)\). instead`
		gomega.Expect(errFunc() != nil).To(gomega.BeFalse()) // want `ginkgo-linter: wrong error assertion\nconsider using .gomega\.Expect\(errFunc\(\)\)\.To\(gomega\.Succeed\(\)\). instead`
	})
})
