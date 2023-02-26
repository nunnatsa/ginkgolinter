package comparison

import (
	. "github.com/onsi/ginkgo/v2"
	g "github.com/onsi/gomega"
	"time"
)

var _ = Describe("remove comparison", func() {
	Context("equal/not equal cases", func() {
		It("should find comparison assertions", func() {
			g.Expect(exampleInt == 0).To(g.Equal(false)) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.ToNot\(g\.BeZero\(\)\). instead`
			g.Expect(0 == exampleInt).To(g.Equal(false)) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.ToNot\(g\.BeZero\(\)\). instead`
			g.Expect(exampleInt == 0).To(g.BeFalse())    // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.ToNot\(g\.BeZero\(\)\). instead`

			g.Expect(exampleInt == constZero).To(g.Equal(false)) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.ToNot\(g\.BeZero\(\)\). instead`
			g.Expect(constZero == exampleInt).To(g.Equal(false)) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.ToNot\(g\.BeZero\(\)\). instead`
			g.Expect(exampleInt == constZero).To(g.BeFalse())    // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.ToNot\(g\.BeZero\(\)\). instead`
			g.Expect(exampleInt != constZero).To(g.BeTrue())     // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.ToNot\(g\.BeZero\(\)\). instead`
			g.Expect(constZero != exampleInt).To(g.BeTrue())     // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.ToNot\(g\.BeZero\(\)\). instead`

			g.Expect(exampleInt).To(g.Equal(1))
			g.Expect(exampleInt == 1).To(g.Equal(true))  // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.Equal\(1\)\). instead`
			g.Expect(exampleInt != 1).To(g.Equal(false)) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.Equal\(1\)\). instead`

			g.Expect(1 == retNum()).To(g.BeTrue())     // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(retNum\(\)\)\.To\(g\.Equal\(1\)\). instead`
			g.Expect(retNum() == 1).To(g.Equal(true))  // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(retNum\(\)\)\.To\(g\.Equal\(1\)\). instead`
			g.Expect(retNum() != 1).To(g.Equal(false)) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(retNum\(\)\)\.To\(g\.Equal\(1\)\). instead`

			exampleInt = 0

			g.Expect(exampleInt == 0).To(g.Equal(true)) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeZero\(\)\). instead`
			g.Expect(exampleInt == 0).To(g.BeTrue())    // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeZero\(\)\). instead`
		})
		It("should find comparison assertions", func() {
			g.Expect(exampleInt == 1).To(g.BeTrue())  // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.Equal\(1\)\). instead`
			g.Expect(exampleInt != 1).To(g.BeFalse()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.Equal\(1\)\). instead`
		})
		It("should find comparison assertions", func() {
			g.Expect(exampleInt == constOne).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.Equal\(constOne\)\). instead`
		})

		It("should find comparison assertions", func() {
			g.Expect(exampleFloat32 == 1).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleFloat32\)\.To\(g\.Equal\(1\)\). instead`
		})
		It("should find comparison assertions", func() {
			g.Expect(exampleFloat32 == 1.0).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleFloat32\)\.To\(g\.Equal\(1.0\)\). instead`
		})
		It("should find comparison assertions", func() {
			g.Expect(exampleFloat32 == constOne).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleFloat32\)\.To\(g\.Equal\(constOne\)\). instead`
		})
		It("imported const", func() {
			g.Expect(time.Millisecond == 1000000).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(time.Millisecond\)\.To\(g\.Equal\(1000000\)\). instead`
		})
	})

	Context("non-equal comparisons", func() {
		It("check greater than", func() {
			g.Expect(exampleInt > constZero).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeNumerically\(">", constZero\)\). instead`
			g.Expect(constZero < exampleInt).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeNumerically\(">", constZero\)\). instead`
			g.Expect(exampleInt).To(g.BeNumerically(">", 0))
			g.Expect(exampleInt > 0).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeNumerically\(">", 0\)\). instead`
			g.Expect(0 < exampleInt).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeNumerically\(">", 0\)\). instead`

			g.Expect(retNum() > constZero).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(retNum\(\)\)\.To\(g\.BeNumerically\(">", constZero\)\). instead`
			g.Expect(constZero < retNum()).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(retNum\(\)\)\.To\(g\.BeNumerically\(">", constZero\)\). instead`
			g.Expect(retNum()).To(g.BeNumerically(">", 0))
			g.Expect(retNum() > 0).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(retNum\(\)\)\.To\(g\.BeNumerically\(">", 0\)\). instead`
			g.Expect(0 < retNum()).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(retNum\(\)\)\.To\(g\.BeNumerically\(">", 0\)\). instead`
		})

		It("check greater than or equal", func() {
			g.Expect(constZero <= exampleInt).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeNumerically\(">=", constZero\)\). instead`
			g.Expect(exampleInt >= constZero).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeNumerically\(">=", constZero\)\). instead`
			g.Expect(exampleInt).To(g.BeNumerically(">=", 0))
			g.Expect(exampleInt >= 0).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeNumerically\(">=", 0\)\). instead`
			g.Expect(0 <= exampleInt).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeNumerically\(">=", 0\)\). instead`
		})

		It("check less than", func() {
			g.Expect(exampleInt < constTwo).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeNumerically\("<", constTwo\)\). instead`
			g.Expect(constTwo > exampleInt).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeNumerically\("<", constTwo\)\). instead`
			g.Expect(exampleInt).To(g.BeNumerically("<", 2))
			g.Expect(exampleInt < 2).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeNumerically\("<", 2\)\). instead`
			g.Expect(2 > exampleInt).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeNumerically\("<", 2\)\). instead`
		})

		It("check less than or equal", func() {
			g.Expect(constTwo >= exampleInt).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeNumerically\("<=", constTwo\)\). instead`
			g.Expect(exampleInt <= constTwo).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeNumerically\("<=", constTwo\)\). instead`
			g.Expect(exampleInt).To(g.BeNumerically("<=", 2))
			g.Expect(exampleInt <= 2).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeNumerically\("<=", 2\)\). instead`
			g.Expect(2 >= exampleInt).To(g.BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .g\.Expect\(exampleInt\)\.To\(g\.BeNumerically\("<=", 2\)\). instead`
		})

	})
})
