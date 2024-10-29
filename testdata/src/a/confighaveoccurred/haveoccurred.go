package haveoccurred

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HaveOccurred", func() {
	It("HaveOccurred for non error types", func() {
		err := fmt.Errorf("err")
		Expect(err).To(HaveOccurred())
	})
	It("HaveOccurred for non error types", func() {
		var p *int
		Expect(p).To(HaveOccurred()) // want `ginkgo-linter: asserting a non-error type with HaveOccurred matcher`
	})

	It("HaveOccurred for non error types", func() {
		Expect(fmt.Errorf("err")).To(HaveOccurred()) // want `prefer using the Succeed matcher for error function, instead of HaveOccurred. Consider using .Expect\(fmt\.Errorf\("err"\)\)\.ToNot\(Succeed\(\)\). instead`
	})

})

var _ = Describe("Succeed", func() {
	It("Succeed for non error types", func() {
		err := fmt.Errorf("err")
		Expect(err).To(Succeed()) // want `prefer using the HaveOccurred matcher for non-function error value, instead of Succeed. Consider using .Expect\(err\)\.ToNot\(HaveOccurred\(\)\). instead`
	})
	It("Succeed for non error types", func() {
		var p *int
		Expect(p).To(Succeed()) // want `ginkgo-linter: asserting a non-error type with Succeed matcher`
	})

	It("Succeed for non error types", func() {
		Expect(fmt.Errorf("err")).To(Succeed())
	})

})
