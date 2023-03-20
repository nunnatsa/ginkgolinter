package comparison

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var exampleInt = 1

const constZero = 0
const constOne = 1
const constTwo = 2
const constStr = "abcd"

var exampleFloat32 float32 = 1

const constFloat32 float32 = 1

var err error

func retNum() int {
	return 1
}

var _ = Describe("remove comparison", func() {
	Context("equal/not equal cases", func() {
		It("should find comparison assertions", func() {
			Expect(exampleInt == 0).To(Equal(false)) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\).ToNot\(BeZero\(\)\). instead`
			Expect(0 == exampleInt).To(Equal(false)) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\).ToNot\(BeZero\(\)\). instead`
			Expect(exampleInt == 0).To(BeFalse())    // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\).ToNot\(BeZero\(\)\). instead`

			Expect(exampleInt == constZero).To(Equal(false)) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\).ToNot\(BeZero\(\)\). instead`
			Expect(constZero == exampleInt).To(Equal(false)) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\).ToNot\(BeZero\(\)\). instead`
			Expect(exampleInt == constZero).To(BeFalse())    // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\).ToNot\(BeZero\(\)\). instead`

			Expect(exampleInt).To(Equal(1))
			Expect(exampleInt == 1).To(Equal(true))  // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(Equal\(1\)\). instead`
			Expect(exampleInt != 1).To(Equal(false)) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(Equal\(1\)\). instead`

			Expect(1 == retNum()).To(BeTrue())     // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(retNum\(\)\)\.To\(Equal\(1\)\). instead`
			Expect(retNum() == 1).To(Equal(true))  // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(retNum\(\)\)\.To\(Equal\(1\)\). instead`
			Expect(retNum() != 1).To(Equal(false)) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(retNum\(\)\)\.To\(Equal\(1\)\). instead`

			exampleInt = 0

			Expect(exampleInt == 0).To(Equal(true))            // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeZero\(\)\). instead`
			Expect(exampleInt == 0).To(BeTrue())               // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeZero\(\)\). instead`
			Expect(exampleInt == 0).WithOffset(5).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.WithOffset\(5\)\.To\(BeZero\(\)\). instead`
		})
		It("should find comparison assertions", func() {
			Expect(exampleInt == 1).To(BeTrue())  // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(Equal\(1\)\). instead`
			Expect(exampleInt != 1).To(BeFalse()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(Equal\(1\)\). instead`
		})
		It("should find comparison assertions", func() {
			Expect(exampleInt == constOne).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(Equal\(constOne\)\). instead`
		})

		It("should find comparison assertions", func() {
			Expect(exampleFloat32 == 1).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleFloat32\)\.To\(Equal\(1\)\). instead`
		})
		It("should find comparison assertions", func() {
			Expect(exampleFloat32 == 1.0).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleFloat32\)\.To\(Equal\(1.0\)\). instead`
		})
		It("should find comparison assertions", func() {
			Expect(exampleFloat32 == constOne).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleFloat32\)\.To\(Equal\(constOne\)\). instead`
		})
		It("imported const", func() {
			Expect(time.Millisecond == 1000000).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(time.Millisecond\)\.To\(Equal\(1000000\)\). instead`
		})
		It("string const", func() {
			var s = "abcd"
			Expect(s == constStr).Should(BeTrue())    // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(s\)\.Should\(Equal\(constStr\)\). instead`
			Expect(constStr == s).Should(Equal(true)) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(s\)\.Should\(Equal\(constStr\)\). instead`
		})
	})

	Context("less than", func() {
		It("check greater than", func() {
			Expect(exampleInt > constZero).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeNumerically\(">", constZero\)\). instead`
			Expect(constZero < exampleInt).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeNumerically\(">", constZero\)\). instead`
			Expect(exampleInt).To(BeNumerically(">", 0))
			Expect(exampleInt > 0).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeNumerically\(">", 0\)\). instead`
			Expect(0 < exampleInt).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeNumerically\(">", 0\)\). instead`

			Expect(retNum() > constZero).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(retNum\(\)\)\.To\(BeNumerically\(">", constZero\)\). instead`
			Expect(constZero < retNum()).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(retNum\(\)\)\.To\(BeNumerically\(">", constZero\)\). instead`
			Expect(retNum()).To(BeNumerically(">", 0))
			Expect(retNum() > 0).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(retNum\(\)\)\.To\(BeNumerically\(">", 0\)\). instead`
			Expect(0 < retNum()).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(retNum\(\)\)\.To\(BeNumerically\(">", 0\)\). instead`
		})

		It("check greater than or equal", func() {
			Expect(constZero <= exampleInt).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeNumerically\(">=", constZero\)\). instead`
			Expect(exampleInt >= constZero).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeNumerically\(">=", constZero\)\). instead`
			Expect(exampleInt).To(BeNumerically(">=", 0))
			Expect(exampleInt >= 0).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeNumerically\(">=", 0\)\). instead`
			Expect(0 <= exampleInt).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeNumerically\(">=", 0\)\). instead`
		})

		It("check less than", func() {
			Expect(exampleInt < constTwo).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeNumerically\("<", constTwo\)\). instead`
			Expect(constTwo > exampleInt).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeNumerically\("<", constTwo\)\). instead`
			Expect(exampleInt).To(BeNumerically("<", 2))
			Expect(exampleInt < 2).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeNumerically\("<", 2\)\). instead`
			Expect(2 > exampleInt).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeNumerically\("<", 2\)\). instead`
		})

		It("check less than or equal", func() {
			Expect(constTwo >= exampleInt).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeNumerically\("<=", constTwo\)\). instead`
			Expect(exampleInt <= constTwo).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeNumerically\("<=", constTwo\)\). instead`
			Expect(exampleInt).To(BeNumerically("<=", 2))
			Expect(exampleInt <= 2).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeNumerically\("<=", 2\)\). instead`
			Expect(2 >= exampleInt).To(BeTrue()) // want `ginkgo-linter: wrong comparison assertion; consider using .Expect\(exampleInt\)\.To\(BeNumerically\("<=", 2\)\). instead`
		})

	})
})
