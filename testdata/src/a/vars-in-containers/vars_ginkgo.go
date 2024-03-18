package vars_in_containers

import (
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = ginkgo.Describe("vars in Describe", func() {
	var x = 1 // want `use BeforeEach\(\) to assign variable x`
	var (
		y     = "testing" // want `use BeforeEach\(\) to assign variable y`
		human = &Human{   // want `use BeforeEach\(\) to assign variable human`
			Name: "someone",
			ID:   "123456789",
		}
		valid int
	)

	a := 1.23 // want `use BeforeEach\(\) to assign variable a`

	ginkgo.BeforeEach(func() {
		valid = 42 // valid
	})

	ginkgo.It("just use the vars", func() {
		Expect(x).To(Equal(1))
		Expect(y).To(Equal("testing"))
		Expect(human).ToNot(BeNil())
		Expect(a).To(Equal(1.23))
		Expect(valid).To(Equal(42))
	})

	ginkgo.When("test When", func() {
		var x = 1 // want `use BeforeEach\(\) to assign variable x`
		var (
			y     = "testing" // want `use BeforeEach\(\) to assign variable y`
			human = &Human{   // want `use BeforeEach\(\) to assign variable human`
				Name: "someone",
				ID:   "123456789",
			}
		)

		a := 1.23 // want `use BeforeEach\(\) to assign variable a`

		ginkgo.It("just use the vars", func() {
			v1 := 12      // may define vars within `It`
			var v2 = "34" // may define vars within `It`
			Expect(x).To(Equal(1))
			Expect(y).To(Equal("testing"))
			Expect(human).ToNot(BeNil())
			Expect(a).To(Equal(1.23))
			Expect(v1).To(Equal(12))
			Expect(v2).To(Equal("34"))
		})
	})
})
