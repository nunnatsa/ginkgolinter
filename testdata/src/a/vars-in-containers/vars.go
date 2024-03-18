package vars_in_containers

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Human struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

var _ = Describe("vars in Describe", func() {
	var a = 1 // want `use BeforeEach\(\) to assign variable a`
	var (
		b = "testing" // want `use BeforeEach\(\) to assign variable b`
		c = &Human{   // want `use BeforeEach\(\) to assign variable c`
			Name: "someone",
			ID:   "123456789",
		}
		d     string
		valid int
		f1    = func(i int) {
			Expect(i).To(BeNumerically(">", 0))
		}

		g, h, valid2 int
	)

	if b == "testing" {
		d = b // want `use BeforeEach\(\) to assign variable d`
	} else {
		d = "something else" // want `use BeforeEach\(\) to assign variable d`
	}

	e := 1.23 // want `use BeforeEach\(\) to assign variable e`

	f2 := func(i int) {
		Expect(i).To(BeNumerically(">", 0))
	}

	g, h = 40, 41 // want `use BeforeEach\(\) to assign variable g` `use BeforeEach\(\) to assign variable h`

	BeforeEach(func() {
		valid = 42 // valid
	})

	It("just use the vars", func() {
		Expect(a).To(Equal(1))
		Expect(b).To(Equal("testing"))
		Expect(c).ToNot(BeNil())
		Expect(d).To(Equal("testing"))
		Expect(e).To(Equal(1.23))
		Expect(valid).To(Equal(42))
		Expect(valid2).To(Equal(42))
		Expect(g).To(Equal(40))
		Expect(h).To(Equal(41))
		f1(valid)
		f2(valid)
	})

	When("test When", func() {
		var x = 1 // want `use BeforeEach\(\) to assign variable x`
		var (
			y     = "testing" // want `use BeforeEach\(\) to assign variable y`
			human = &Human{   // want `use BeforeEach\(\) to assign variable human`
				Name: "someone",
				ID:   "123456789",
			}
		)

		a := 1.23 // want `use BeforeEach\(\) to assign variable a`

		It("just use the vars", func() {
			Expect(x).To(Equal(1))
			Expect(y).To(Equal("testing"))
			Expect(human).ToNot(BeNil())
			Expect(a).To(Equal(1.23))
		})
	})
})
