package haveoccurred

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("HaveOccurred", func() {
	It("HaveOccurred for non error types", func() {
		err := fmt.Errorf("err")
		gomega.Expect(err).To(gomega.HaveOccurred())
	})
	It("HaveOccurred for non error types", func() {
		var p *int
		gomega.Expect(p).To(gomega.HaveOccurred()) // want `ginkgo-linter: asserting a non-error type with HaveOccurred matcher`
	})

	It("HaveOccurred for non error types", func() {
		gomega.Expect(fmt.Errorf("err")).To(gomega.HaveOccurred()) // want `prefer using the Succeed matcher for error function, instead of HaveOccurred. Consider using .gomega\.Expect\(fmt\.Errorf\("err"\)\)\.ToNot\(gomega\.Succeed\(\)\). instead`
	})

})

var _ = Describe("Succeed", func() {
	It("Succeed for non error types", func() {
		err := fmt.Errorf("err")
		gomega.Expect(err).To(gomega.Succeed()) // want `prefer using the HaveOccurred matcher for non-function error value, instead of Succeed. Consider using .gomega\.Expect\(err\)\.ToNot\(gomega\.HaveOccurred\(\)\). instead`
	})
	It("Succeed for non error types", func() {
		var p *int
		gomega.Expect(p).To(gomega.Succeed()) // want `ginkgo-linter: asserting a non-error type with Succeed matcher`
	})

	It("Succeed for non error types", func() {
		gomega.Expect(fmt.Errorf("err")).To(gomega.Succeed())
	})

})
