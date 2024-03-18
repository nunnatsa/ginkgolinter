package vars_in_containers

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("When's", func() {
	const shouldNotTriggerWarning = 1
	const (
		shouldNotTriggerWarningEither = 2
	)

	When("test When", func() {
		const shouldNotTriggerWarning = 3
		const (
			shouldNotTriggerWarningEither = 4
		)

		var x = 1 // want `use BeforeEach\(\) to assign variable x`

		It("use x", func() {
			Expect(x).To(Equal(1))
		})
	})
	FWhen("test FWhen", func() {
		var x = 1 // want `use BeforeEach\(\) to assign variable x`
		It("use x", func() {
			Expect(x).To(Equal(1))
		})
	})
	PWhen("test PWhen", func() {
		var x = 1 // want `use BeforeEach\(\) to assign variable x`
		It("use x", func() {
			Expect(x).To(Equal(1))
		})
	})
	XWhen("test XWhen", func() {
		var x = 1 // want `use BeforeEach\(\) to assign variable x`
		It("use x", func() {
			Expect(x).To(Equal(1))
		})
	})
})

var _ = When("Context's", func() {
	Context("test Context", func() {
		var x = 1 // want `use BeforeEach\(\) to assign variable x`
		It("use x", func() {
			Expect(x).To(Equal(1))
		})
	})
	FContext("test FContext", func() {
		var x = 1 // want `use BeforeEach\(\) to assign variable x`
		It("use x", func() {
			Expect(x).To(Equal(1))
		})
	})
	PContext("test PContext", func() {
		var x = 1 // want `use BeforeEach\(\) to assign variable x`
		It("use x", func() {
			Expect(x).To(Equal(1))
		})
	})
	XContext("test XContext", func() {
		var x = 1 // want `use BeforeEach\(\) to assign variable x`
		It("use x", func() {
			Expect(x).To(Equal(1))
		})
	})
})

var _ = Context("Describe's", func() {
	Describe("test Describe", func() {
		var x = 1 // want `use BeforeEach\(\) to assign variable x`
		It("use x", func() {
			Expect(x).To(Equal(1))
		})
	})
	FDescribe("test FDescribe", func() {
		var x = 1 // want `use BeforeEach\(\) to assign variable x`
		It("use x", func() {
			Expect(x).To(Equal(1))
		})
	})
	PDescribe("test PDescribe", func() {
		var x = 1 // want `use BeforeEach\(\) to assign variable x`
		It("use x", func() {
			Expect(x).To(Equal(1))
		})
	})
	XDescribe("test XDescribe", func() {
		var x = 1 // want `use BeforeEach\(\) to assign variable x`
		It("use x", func() {
			Expect(x).To(Equal(1))
		})
	})
})
