package nil

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func fNil() *string {
	return nil
}

func fNotNil() *string {
	val := "no nil"
	return &val
}

var _ = Describe("", func() {

	var x *int
	val := 3
	y := &val

	Context("test Expect", func() {
		Context("test Should", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Expect(x == nil).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
					Expect(nil == x).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y != nil).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil != y).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() == nil).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Expect(nil == fNil()).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil != fNotNil()).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Expect(x != nil).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
					Expect(nil != x).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y == nil).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil == y).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() != nil).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Expect(nil != fNil()).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil == fNotNil()).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Expect(x == nil).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
					Expect(nil == x).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y != nil).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil != y).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() == nil).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Expect(nil == fNil()).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil != fNotNil()).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Expect(x != nil).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
					Expect(nil != x).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y == nil).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil == y).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() != nil).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Expect(nil != fNil()).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil == fNotNil()).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
		})
		Context("test Should(Not())", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Expect(x == nil).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
					Expect(nil == x).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y != nil).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil != y).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() == nil).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Expect(nil == fNil()).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil != fNotNil()).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Expect(x != nil).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
					Expect(nil != x).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y == nil).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil == y).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() != nil).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Expect(nil != fNil()).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil == fNotNil()).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Expect(x == nil).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
					Expect(nil == x).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y != nil).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil != y).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() == nil).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Expect(nil == fNil()).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil != fNotNil()).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Expect(x != nil).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
					Expect(nil != x).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y == nil).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil == y).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() != nil).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Expect(nil != fNil()).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil == fNotNil()).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
		})

		Context("test To", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Expect(x == nil).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
					Expect(nil == x).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y != nil).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil != y).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() == nil).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Expect(nil == fNil()).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil != fNotNil()).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Expect(x != nil).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
					Expect(nil != x).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y == nil).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil == y).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() != nil).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Expect(nil != fNil()).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil == fNotNil()).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Expect(x == nil).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
					Expect(nil == x).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y != nil).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil != y).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() == nil).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Expect(nil == fNil()).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil != fNotNil()).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Expect(x != nil).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
					Expect(nil != x).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y == nil).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil == y).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() != nil).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Expect(nil != fNil()).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil == fNotNil()).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
		})
		Context("test To(Not())", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Expect(x == nil).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
					Expect(nil == x).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y != nil).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil != y).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() == nil).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Expect(nil == fNil()).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil != fNotNil()).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Expect(x != nil).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
					Expect(nil != x).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y == nil).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil == y).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() != nil).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Expect(nil != fNil()).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil == fNotNil()).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Expect(x == nil).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
					Expect(nil == x).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y != nil).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil != y).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() == nil).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Expect(nil == fNil()).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil != fNotNil()).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Expect(x != nil).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
					Expect(nil != x).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y == nil).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil == y).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() != nil).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Expect(nil != fNil()).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil == fNotNil()).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
		})

		Context("test ShouldNot", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Expect(x == nil).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
					Expect(nil == x).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y != nil).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil != y).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() == nil).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Expect(nil == fNil()).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil != fNotNil()).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Expect(x != nil).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
					Expect(nil != x).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y == nil).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil == y).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() != nil).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Expect(nil != fNil()).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil == fNotNil()).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Expect(x != nil).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
					Expect(nil != x).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y == nil).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil == y).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() != nil).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Expect(nil != fNil()).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil == fNotNil()).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Expect(x == nil).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
					Expect(nil == x).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y != nil).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil != y).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() == nil).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Expect(nil == fNil()).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Expect(nil != fNotNil()).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
		})

		Context("test NotTo", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Expect(x == nil).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
					Expect(nil == x).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y != nil).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil != y).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() == nil).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Expect(nil == fNil()).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil != fNotNil()).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Expect(x != nil).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
					Expect(nil != x).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y == nil).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.NotTo\(BeNil\(\)\). instead`
					Expect(nil == y).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.NotTo\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() != nil).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Expect(nil != fNil()).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.NotTo\(BeNil\(\)\). instead`
					Expect(nil == fNotNil()).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.NotTo\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Expect(x != nil).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
					Expect(nil != x).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y == nil).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.NotTo\(BeNil\(\)\). instead`
					Expect(nil == y).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.NotTo\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() != nil).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Expect(nil != fNil()).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.NotTo\(BeNil\(\)\). instead`
					Expect(nil == fNotNil()).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.NotTo\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Expect(x == nil).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
					Expect(nil == x).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y != nil).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil != y).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() == nil).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Expect(nil == fNil()).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil != fNotNil()).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
		})
		Context("test ToNot", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Expect(x == nil).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
					Expect(nil == x).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y != nil).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil != y).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() == nil).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Expect(nil == fNil()).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil != fNotNil()).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Expect(x != nil).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
					Expect(nil != x).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y == nil).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil == y).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() != nil).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Expect(nil != fNil()).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil == fNotNil()).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Expect(x != nil).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
					Expect(nil != x).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y == nil).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil == y).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() != nil).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Expect(nil != fNil()).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil == fNotNil()).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Expect(x == nil).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
					Expect(nil == x).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Expect(y != nil).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil != y).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Expect(fNil() == nil).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Expect(nil == fNil()).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Expect(nil != fNotNil()).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Expect\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
		})
	})

	Context("test ExpectWithOffset", func() {
		Context("test Should", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == x).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != y).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNil()).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNotNil()).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != x).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == y).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNil()).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNotNil()).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == x).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != y).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNil()).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNotNil()).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != x).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == y).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNil()).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNotNil()).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
		})
		Context("test Should(Not())", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == x).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != y).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNil()).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNotNil()).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != x).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == y).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNil()).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNotNil()).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == x).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != y).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNil()).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNotNil()).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != x).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == y).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNil()).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNotNil()).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
		})

		Context("test To", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == x).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != y).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNil()).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNotNil()).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != x).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == y).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNil()).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNotNil()).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == x).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != y).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNil()).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNotNil()).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != x).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == y).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNil()).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNotNil()).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
		})
		Context("test To(Not())", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == x).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != y).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNil()).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNotNil()).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != x).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == y).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNil()).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNotNil()).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == x).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != y).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNil()).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNotNil()).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != x).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == y).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNil()).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNotNil()).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
		})

		Context("test ShouldNot", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == x).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != y).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNil()).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNotNil()).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != x).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == y).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNil()).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNotNil()).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != x).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == y).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNil()).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNotNil()).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == x).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != y).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNil()).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNotNil()).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
		})

		Context("test NotTo", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == x).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != y).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNil()).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNotNil()).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != x).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.NotTo\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == y).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.NotTo\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNil()).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.NotTo\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNotNil()).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.NotTo\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != x).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.NotTo\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == y).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.NotTo\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNil()).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.NotTo\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNotNil()).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.NotTo\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == x).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != y).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNil()).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNotNil()).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
		})
		Context("test ToNot", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == x).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != y).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNil()).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNotNil()).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != x).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == y).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNil()).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNotNil()).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != x).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == y).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNil()).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNotNil()).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == x).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != y).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil == fNil()).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					ExpectWithOffset(1, nil != fNotNil()).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .ExpectWithOffset\(1, fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
		})
	})

	Context("test Ω", func() {
		Context("test Should", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Ω(x == nil).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
					Ω(nil == x).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y != nil).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil != y).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() == nil).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Ω(nil == fNil()).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil != fNotNil()).Should(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Ω(x != nil).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
					Ω(nil != x).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y == nil).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil == y).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() != nil).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Ω(nil != fNil()).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil == fNotNil()).Should(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Ω(x == nil).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
					Ω(nil == x).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y != nil).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil != y).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() == nil).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Ω(nil == fNil()).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil != fNotNil()).Should(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Ω(x != nil).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
					Ω(nil != x).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y == nil).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil == y).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() != nil).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Ω(nil != fNil()).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil == fNotNil()).Should(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
		})
		Context("test Should(Not())", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Ω(x == nil).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
					Ω(nil == x).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y != nil).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil != y).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() == nil).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Ω(nil == fNil()).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil != fNotNil()).Should(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Ω(x != nil).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
					Ω(nil != x).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y == nil).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil == y).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() != nil).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Ω(nil != fNil()).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil == fNotNil()).Should(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Ω(x == nil).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
					Ω(nil == x).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y != nil).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil != y).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() == nil).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Ω(nil == fNil()).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil != fNotNil()).Should(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Ω(x != nil).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
					Ω(nil != x).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y == nil).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil == y).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() != nil).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Ω(nil != fNil()).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil == fNotNil()).Should(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
		})

		Context("test To", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Ω(x == nil).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
					Ω(nil == x).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y != nil).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil != y).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() == nil).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Ω(nil == fNil()).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil != fNotNil()).To(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Ω(x != nil).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
					Ω(nil != x).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y == nil).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil == y).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() != nil).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Ω(nil != fNil()).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil == fNotNil()).To(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Ω(x == nil).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
					Ω(nil == x).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y != nil).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil != y).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() == nil).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Ω(nil == fNil()).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil != fNotNil()).To(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Ω(x != nil).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
					Ω(nil != x).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y == nil).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil == y).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() != nil).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Ω(nil != fNil()).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil == fNotNil()).To(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
		})
		Context("test To(Not())", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Ω(x == nil).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
					Ω(nil == x).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y != nil).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil != y).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() == nil).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Ω(nil == fNil()).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil != fNotNil()).To(Not(BeFalse())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Ω(x != nil).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
					Ω(nil != x).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y == nil).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil == y).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() != nil).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Ω(nil != fNil()).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil == fNotNil()).To(Not(BeTrue())) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Ω(x == nil).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
					Ω(nil == x).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y != nil).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil != y).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() == nil).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Ω(nil == fNil()).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil != fNotNil()).To(Not(Equal(false))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Ω(x != nil).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
					Ω(nil != x).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y == nil).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil == y).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() != nil).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Ω(nil != fNil()).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil == fNotNil()).To(Not(Equal(true))) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
		})

		Context("test ShouldNot", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Ω(x == nil).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
					Ω(nil == x).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y != nil).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil != y).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() == nil).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Ω(nil == fNil()).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil != fNotNil()).ShouldNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Ω(x != nil).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
					Ω(nil != x).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y == nil).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil == y).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() != nil).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Ω(nil != fNil()).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil == fNotNil()).ShouldNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Ω(x != nil).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
					Ω(nil != x).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y == nil).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil == y).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() != nil).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Ω(nil != fNil()).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil == fNotNil()).ShouldNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Ω(x == nil).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
					Ω(nil == x).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y != nil).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil != y).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ShouldNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() == nil).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
					Ω(nil == fNil()).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.Should\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
					Ω(nil != fNotNil()).ShouldNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ShouldNot\(BeNil\(\)\). instead`
				})
			})
		})

		Context("test NotTo", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Ω(x == nil).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
					Ω(nil == x).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y != nil).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil != y).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() == nil).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Ω(nil == fNil()).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil != fNotNil()).NotTo(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Ω(x != nil).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
					Ω(nil != x).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y == nil).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.NotTo\(BeNil\(\)\). instead`
					Ω(nil == y).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.NotTo\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() != nil).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Ω(nil != fNil()).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.NotTo\(BeNil\(\)\). instead`
					Ω(nil == fNotNil()).NotTo(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.NotTo\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Ω(x != nil).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
					Ω(nil != x).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y == nil).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.NotTo\(BeNil\(\)\). instead`
					Ω(nil == y).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.NotTo\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() != nil).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Ω(nil != fNil()).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.NotTo\(BeNil\(\)\). instead`
					Ω(nil == fNotNil()).NotTo(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.NotTo\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Ω(x == nil).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
					Ω(nil == x).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y != nil).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil != y).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() == nil).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Ω(nil == fNil()).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil != fNotNil()).NotTo(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
		})
		Context("test ToNot", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Ω(x == nil).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
					Ω(nil == x).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y != nil).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil != y).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() == nil).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Ω(nil == fNil()).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil != fNotNil()).ToNot(BeFalse()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Ω(x != nil).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
					Ω(nil != x).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y == nil).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil == y).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() != nil).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Ω(nil != fNil()).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil == fNotNil()).ToNot(BeTrue()) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Ω(x != nil).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
					Ω(nil != x).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y == nil).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil == y).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() != nil).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Ω(nil != fNil()).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil == fNotNil()).ToNot(Equal(true)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Ω(x == nil).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
					Ω(nil == x).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(x\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil value", func() {
					Ω(y != nil).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil != y).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(y\)\.ToNot\(BeNil\(\)\). instead`
				})
				It("test nil func", func() {
					Ω(fNil() == nil).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
					Ω(nil == fNil()).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNil\(\)\)\.To\(BeNil\(\)\). instead`
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
					Ω(nil != fNotNil()).ToNot(Equal(false)) // want `ginkgo-linter: wrong nil assertion; consider using .Ω\(fNotNil\(\)\)\.ToNot\(BeNil\(\)\). instead`
				})
			})
		})
	})
})
