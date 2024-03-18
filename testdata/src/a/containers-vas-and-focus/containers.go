package vars_in_containers

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("When's", func() { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code. Consider to replace with "Describe"`
	FWhen("test FWhen", func() { // want `Focus container found. This is used only for local debug and should not be part of the actual source code. Consider to replace with "When"`
		var x = 1 // want `use BeforeEach\(\) to assign variable x`
		It("use x", func() {
			Expect(x).To(Equal(1))
		})
	})
})

var _ = FWhen("Context's", func() { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code. Consider to replace with "When"`
	FContext("test FContext", func() { // want `Focus container found. This is used only for local debug and should not be part of the actual source code. Consider to replace with "Context"`
		var x = 1 // want `use BeforeEach\(\) to assign variable x`
		It("use x", func() {
			Expect(x).To(Equal(1))
		})
	})
})

var _ = FContext("Describe's", func() { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code. Consider to replace with "Context"`
	FDescribe("test FDescribe", func() { // want `Focus container found. This is used only for local debug and should not be part of the actual source code. Consider to replace with "Describe"`
		var x = 1 // want `use BeforeEach\(\) to assign variable x`
		It("use x", func() {
			Expect(x).To(Equal(1))
		})
	})
})
