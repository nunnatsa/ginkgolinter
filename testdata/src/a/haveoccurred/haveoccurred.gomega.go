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
		gomega.Expect(fmt.Errorf("err")).To(gomega.HaveOccurred())
	})

})

var _ = Describe("Succeed", func() {
	It("Succeed for non error types", func() {
		err := fmt.Errorf("err")
		gomega.Expect(err).To(gomega.Succeed())
	})
	It("Succeed for non error types", func() {
		var p *int
		gomega.Expect(p).To(gomega.Succeed()) // want `ginkgo-linter: asserting a non-error type with Succeed matcher`
	})

	It("Succeed for non error types", func() {
		gomega.Expect(fmt.Errorf("err")).To(gomega.Succeed())
	})

})
