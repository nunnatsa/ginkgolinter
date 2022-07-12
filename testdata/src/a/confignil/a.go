package confignil

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
					Expect(x == nil).Should(BeTrue())
					Expect(nil == x).Should(BeTrue())
				})
				It("test non-nil value", func() {
					Expect(y != nil).Should(BeTrue())
					Expect(nil != y).Should(BeTrue())
				})
				It("test nil func", func() {
					Expect(fNil() == nil).Should(BeTrue())
					Expect(nil == fNil()).Should(BeTrue())
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).Should(BeTrue())
					Expect(nil != fNotNil()).Should(BeTrue())
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Expect(x != nil).Should(BeFalse())
					Expect(nil != x).Should(BeFalse())
				})
				It("test non-nil value", func() {
					Expect(y == nil).Should(BeFalse())
					Expect(nil == y).Should(BeFalse())
				})
				It("test nil func", func() {
					Expect(fNil() != nil).Should(BeFalse())
					Expect(nil != fNil()).Should(BeFalse())
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).Should(BeFalse())
					Expect(nil == fNotNil()).Should(BeFalse())
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Expect(x == nil).Should(Equal(true))
					Expect(nil == x).Should(Equal(true))
				})
				It("test non-nil value", func() {
					Expect(y != nil).Should(Equal(true))
					Expect(nil != y).Should(Equal(true))
				})
				It("test nil func", func() {
					Expect(fNil() == nil).Should(Equal(true))
					Expect(nil == fNil()).Should(Equal(true))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).Should(Equal(true))
					Expect(nil != fNotNil()).Should(Equal(true))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Expect(x != nil).Should(Equal(false))
					Expect(nil != x).Should(Equal(false))
				})
				It("test non-nil value", func() {
					Expect(y == nil).Should(Equal(false))
					Expect(nil == y).Should(Equal(false))
				})
				It("test nil func", func() {
					Expect(fNil() != nil).Should(Equal(false))
					Expect(nil != fNil()).Should(Equal(false))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).Should(Equal(false))
					Expect(nil == fNotNil()).Should(Equal(false))
				})
			})
		})
		Context("test Should(Not())", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Expect(x == nil).Should(Not(BeFalse()))
					Expect(nil == x).Should(Not(BeFalse()))
				})
				It("test non-nil value", func() {
					Expect(y != nil).Should(Not(BeFalse()))
					Expect(nil != y).Should(Not(BeFalse()))
				})
				It("test nil func", func() {
					Expect(fNil() == nil).Should(Not(BeFalse()))
					Expect(nil == fNil()).Should(Not(BeFalse()))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).Should(Not(BeFalse()))
					Expect(nil != fNotNil()).Should(Not(BeFalse()))
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Expect(x != nil).Should(Not(BeTrue()))
					Expect(nil != x).Should(Not(BeTrue()))
				})
				It("test non-nil value", func() {
					Expect(y == nil).Should(Not(BeTrue()))
					Expect(nil == y).Should(Not(BeTrue()))
				})
				It("test nil func", func() {
					Expect(fNil() != nil).Should(Not(BeTrue()))
					Expect(nil != fNil()).Should(Not(BeTrue()))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).Should(Not(BeTrue()))
					Expect(nil == fNotNil()).Should(Not(BeTrue()))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Expect(x == nil).Should(Not(Equal(false)))
					Expect(nil == x).Should(Not(Equal(false)))
				})
				It("test non-nil value", func() {
					Expect(y != nil).Should(Not(Equal(false)))
					Expect(nil != y).Should(Not(Equal(false)))
				})
				It("test nil func", func() {
					Expect(fNil() == nil).Should(Not(Equal(false)))
					Expect(nil == fNil()).Should(Not(Equal(false)))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).Should(Not(Equal(false)))
					Expect(nil != fNotNil()).Should(Not(Equal(false)))
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Expect(x != nil).Should(Not(Equal(true)))
					Expect(nil != x).Should(Not(Equal(true)))
				})
				It("test non-nil value", func() {
					Expect(y == nil).Should(Not(Equal(true)))
					Expect(nil == y).Should(Not(Equal(true)))
				})
				It("test nil func", func() {
					Expect(fNil() != nil).Should(Not(Equal(true)))
					Expect(nil != fNil()).Should(Not(Equal(true)))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).Should(Not(Equal(true)))
					Expect(nil == fNotNil()).Should(Not(Equal(true)))
				})
			})
		})

		Context("test To", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Expect(x == nil).To(BeTrue())
					Expect(nil == x).To(BeTrue())
				})
				It("test non-nil value", func() {
					Expect(y != nil).To(BeTrue())
					Expect(nil != y).To(BeTrue())
				})
				It("test nil func", func() {
					Expect(fNil() == nil).To(BeTrue())
					Expect(nil == fNil()).To(BeTrue())
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).To(BeTrue())
					Expect(nil != fNotNil()).To(BeTrue())
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Expect(x != nil).To(BeFalse())
					Expect(nil != x).To(BeFalse())
				})
				It("test non-nil value", func() {
					Expect(y == nil).To(BeFalse())
					Expect(nil == y).To(BeFalse())
				})
				It("test nil func", func() {
					Expect(fNil() != nil).To(BeFalse())
					Expect(nil != fNil()).To(BeFalse())
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).To(BeFalse())
					Expect(nil == fNotNil()).To(BeFalse())
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Expect(x == nil).To(Equal(true))
					Expect(nil == x).To(Equal(true))
				})
				It("test non-nil value", func() {
					Expect(y != nil).To(Equal(true))
					Expect(nil != y).To(Equal(true))
				})
				It("test nil func", func() {
					Expect(fNil() == nil).To(Equal(true))
					Expect(nil == fNil()).To(Equal(true))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).To(Equal(true))
					Expect(nil != fNotNil()).To(Equal(true))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Expect(x != nil).To(Equal(false))
					Expect(nil != x).To(Equal(false))
				})
				It("test non-nil value", func() {
					Expect(y == nil).To(Equal(false))
					Expect(nil == y).To(Equal(false))
				})
				It("test nil func", func() {
					Expect(fNil() != nil).To(Equal(false))
					Expect(nil != fNil()).To(Equal(false))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).To(Equal(false))
					Expect(nil == fNotNil()).To(Equal(false))
				})
			})
		})
		Context("test To(Not())", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Expect(x == nil).To(Not(BeFalse()))
					Expect(nil == x).To(Not(BeFalse()))
				})
				It("test non-nil value", func() {
					Expect(y != nil).To(Not(BeFalse()))
					Expect(nil != y).To(Not(BeFalse()))
				})
				It("test nil func", func() {
					Expect(fNil() == nil).To(Not(BeFalse()))
					Expect(nil == fNil()).To(Not(BeFalse()))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).To(Not(BeFalse()))
					Expect(nil != fNotNil()).To(Not(BeFalse()))
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Expect(x != nil).To(Not(BeTrue()))
					Expect(nil != x).To(Not(BeTrue()))
				})
				It("test non-nil value", func() {
					Expect(y == nil).To(Not(BeTrue()))
					Expect(nil == y).To(Not(BeTrue()))
				})
				It("test nil func", func() {
					Expect(fNil() != nil).To(Not(BeTrue()))
					Expect(nil != fNil()).To(Not(BeTrue()))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).To(Not(BeTrue()))
					Expect(nil == fNotNil()).To(Not(BeTrue()))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Expect(x == nil).To(Not(Equal(false)))
					Expect(nil == x).To(Not(Equal(false)))
				})
				It("test non-nil value", func() {
					Expect(y != nil).To(Not(Equal(false)))
					Expect(nil != y).To(Not(Equal(false)))
				})
				It("test nil func", func() {
					Expect(fNil() == nil).To(Not(Equal(false)))
					Expect(nil == fNil()).To(Not(Equal(false)))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).To(Not(Equal(false)))
					Expect(nil != fNotNil()).To(Not(Equal(false)))
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Expect(x != nil).To(Not(Equal(true)))
					Expect(nil != x).To(Not(Equal(true)))
				})
				It("test non-nil value", func() {
					Expect(y == nil).To(Not(Equal(true)))
					Expect(nil == y).To(Not(Equal(true)))
				})
				It("test nil func", func() {
					Expect(fNil() != nil).To(Not(Equal(true)))
					Expect(nil != fNil()).To(Not(Equal(true)))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).To(Not(Equal(true)))
					Expect(nil == fNotNil()).To(Not(Equal(true)))
				})
			})
		})

		Context("test ShouldNot", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Expect(x == nil).ShouldNot(BeFalse())
					Expect(nil == x).ShouldNot(BeFalse())
				})
				It("test non-nil value", func() {
					Expect(y != nil).ShouldNot(BeFalse())
					Expect(nil != y).ShouldNot(BeFalse())
				})
				It("test nil func", func() {
					Expect(fNil() == nil).ShouldNot(BeFalse())
					Expect(nil == fNil()).ShouldNot(BeFalse())
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).ShouldNot(BeFalse())
					Expect(nil != fNotNil()).ShouldNot(BeFalse())
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Expect(x != nil).ShouldNot(BeTrue())
					Expect(nil != x).ShouldNot(BeTrue())
				})
				It("test non-nil value", func() {
					Expect(y == nil).ShouldNot(BeTrue())
					Expect(nil == y).ShouldNot(BeTrue())
				})
				It("test nil func", func() {
					Expect(fNil() != nil).ShouldNot(BeTrue())
					Expect(nil != fNil()).ShouldNot(BeTrue())
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).ShouldNot(BeTrue())
					Expect(nil == fNotNil()).ShouldNot(BeTrue())
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Expect(x != nil).ShouldNot(Equal(true))
					Expect(nil != x).ShouldNot(Equal(true))
				})
				It("test non-nil value", func() {
					Expect(y == nil).ShouldNot(Equal(true))
					Expect(nil == y).ShouldNot(Equal(true))
				})
				It("test nil func", func() {
					Expect(fNil() != nil).ShouldNot(Equal(true))
					Expect(nil != fNil()).ShouldNot(Equal(true))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).ShouldNot(Equal(true))
					Expect(nil == fNotNil()).ShouldNot(Equal(true))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Expect(x == nil).ShouldNot(Equal(false))
					Expect(nil == x).ShouldNot(Equal(false))
				})
				It("test non-nil value", func() {
					Expect(y != nil).ShouldNot(Equal(false))
					Expect(nil != y).ShouldNot(Equal(false))
				})
				It("test nil func", func() {
					Expect(fNil() == nil).ShouldNot(Equal(false))
					Expect(nil == fNil()).ShouldNot(Equal(false))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).ShouldNot(Equal(false))
					Expect(nil != fNotNil()).ShouldNot(Equal(false))
				})
			})
		})

		Context("test NotTo", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Expect(x == nil).NotTo(BeFalse())
					Expect(nil == x).NotTo(BeFalse())
				})
				It("test non-nil value", func() {
					Expect(y != nil).NotTo(BeFalse())
					Expect(nil != y).NotTo(BeFalse())
				})
				It("test nil func", func() {
					Expect(fNil() == nil).NotTo(BeFalse())
					Expect(nil == fNil()).NotTo(BeFalse())
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).NotTo(BeFalse())
					Expect(nil != fNotNil()).NotTo(BeFalse())
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Expect(x != nil).NotTo(BeTrue())
					Expect(nil != x).NotTo(BeTrue())
				})
				It("test non-nil value", func() {
					Expect(y == nil).NotTo(BeTrue())
					Expect(nil == y).NotTo(BeTrue())
				})
				It("test nil func", func() {
					Expect(fNil() != nil).NotTo(BeTrue())
					Expect(nil != fNil()).NotTo(BeTrue())
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).NotTo(BeTrue())
					Expect(nil == fNotNil()).NotTo(BeTrue())
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Expect(x != nil).NotTo(Equal(true))
					Expect(nil != x).NotTo(Equal(true))
				})
				It("test non-nil value", func() {
					Expect(y == nil).NotTo(Equal(true))
					Expect(nil == y).NotTo(Equal(true))
				})
				It("test nil func", func() {
					Expect(fNil() != nil).NotTo(Equal(true))
					Expect(nil != fNil()).NotTo(Equal(true))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).NotTo(Equal(true))
					Expect(nil == fNotNil()).NotTo(Equal(true))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Expect(x == nil).NotTo(Equal(false))
					Expect(nil == x).NotTo(Equal(false))
				})
				It("test non-nil value", func() {
					Expect(y != nil).NotTo(Equal(false))
					Expect(nil != y).NotTo(Equal(false))
				})
				It("test nil func", func() {
					Expect(fNil() == nil).NotTo(Equal(false))
					Expect(nil == fNil()).NotTo(Equal(false))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).NotTo(Equal(false))
					Expect(nil != fNotNil()).NotTo(Equal(false))
				})
			})
		})
		Context("test ToNot", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Expect(x == nil).ToNot(BeFalse())
					Expect(nil == x).ToNot(BeFalse())
				})
				It("test non-nil value", func() {
					Expect(y != nil).ToNot(BeFalse())
					Expect(nil != y).ToNot(BeFalse())
				})
				It("test nil func", func() {
					Expect(fNil() == nil).ToNot(BeFalse())
					Expect(nil == fNil()).ToNot(BeFalse())
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).ToNot(BeFalse())
					Expect(nil != fNotNil()).ToNot(BeFalse())
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Expect(x != nil).ToNot(BeTrue())
					Expect(nil != x).ToNot(BeTrue())
				})
				It("test non-nil value", func() {
					Expect(y == nil).ToNot(BeTrue())
					Expect(nil == y).ToNot(BeTrue())
				})
				It("test nil func", func() {
					Expect(fNil() != nil).ToNot(BeTrue())
					Expect(nil != fNil()).ToNot(BeTrue())
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).ToNot(BeTrue())
					Expect(nil == fNotNil()).ToNot(BeTrue())
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Expect(x != nil).ToNot(Equal(true))
					Expect(nil != x).ToNot(Equal(true))
				})
				It("test non-nil value", func() {
					Expect(y == nil).ToNot(Equal(true))
					Expect(nil == y).ToNot(Equal(true))
				})
				It("test nil func", func() {
					Expect(fNil() != nil).ToNot(Equal(true))
					Expect(nil != fNil()).ToNot(Equal(true))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() == nil).ToNot(Equal(true))
					Expect(nil == fNotNil()).ToNot(Equal(true))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Expect(x == nil).ToNot(Equal(false))
					Expect(nil == x).ToNot(Equal(false))
				})
				It("test non-nil value", func() {
					Expect(y != nil).ToNot(Equal(false))
					Expect(nil != y).ToNot(Equal(false))
				})
				It("test nil func", func() {
					Expect(fNil() == nil).ToNot(Equal(false))
					Expect(nil == fNil()).ToNot(Equal(false))
				})
				It("test non-nil func", func() {
					Expect(fNotNil() != nil).ToNot(Equal(false))
					Expect(nil != fNotNil()).ToNot(Equal(false))
				})
			})
		})
	})

	Context("test ExpectWithOffset", func() {
		Context("test Should", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).Should(BeTrue())
					ExpectWithOffset(1, nil == x).Should(BeTrue())
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).Should(BeTrue())
					ExpectWithOffset(1, nil != y).Should(BeTrue())
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).Should(BeTrue())
					ExpectWithOffset(1, nil == fNil()).Should(BeTrue())
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).Should(BeTrue())
					ExpectWithOffset(1, nil != fNotNil()).Should(BeTrue())
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).Should(BeFalse())
					ExpectWithOffset(1, nil != x).Should(BeFalse())
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).Should(BeFalse())
					ExpectWithOffset(1, nil == y).Should(BeFalse())
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).Should(BeFalse())
					ExpectWithOffset(1, nil != fNil()).Should(BeFalse())
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).Should(BeFalse())
					ExpectWithOffset(1, nil == fNotNil()).Should(BeFalse())
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).Should(Equal(true))
					ExpectWithOffset(1, nil == x).Should(Equal(true))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).Should(Equal(true))
					ExpectWithOffset(1, nil != y).Should(Equal(true))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).Should(Equal(true))
					ExpectWithOffset(1, nil == fNil()).Should(Equal(true))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).Should(Equal(true))
					ExpectWithOffset(1, nil != fNotNil()).Should(Equal(true))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).Should(Equal(false))
					ExpectWithOffset(1, nil != x).Should(Equal(false))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).Should(Equal(false))
					ExpectWithOffset(1, nil == y).Should(Equal(false))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).Should(Equal(false))
					ExpectWithOffset(1, nil != fNil()).Should(Equal(false))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).Should(Equal(false))
					ExpectWithOffset(1, nil == fNotNil()).Should(Equal(false))
				})
			})
		})
		Context("test Should(Not())", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).Should(Not(BeFalse()))
					ExpectWithOffset(1, nil == x).Should(Not(BeFalse()))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).Should(Not(BeFalse()))
					ExpectWithOffset(1, nil != y).Should(Not(BeFalse()))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).Should(Not(BeFalse()))
					ExpectWithOffset(1, nil == fNil()).Should(Not(BeFalse()))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).Should(Not(BeFalse()))
					ExpectWithOffset(1, nil != fNotNil()).Should(Not(BeFalse()))
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).Should(Not(BeTrue()))
					ExpectWithOffset(1, nil != x).Should(Not(BeTrue()))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).Should(Not(BeTrue()))
					ExpectWithOffset(1, nil == y).Should(Not(BeTrue()))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).Should(Not(BeTrue()))
					ExpectWithOffset(1, nil != fNil()).Should(Not(BeTrue()))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).Should(Not(BeTrue()))
					ExpectWithOffset(1, nil == fNotNil()).Should(Not(BeTrue()))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).Should(Not(Equal(false)))
					ExpectWithOffset(1, nil == x).Should(Not(Equal(false)))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).Should(Not(Equal(false)))
					ExpectWithOffset(1, nil != y).Should(Not(Equal(false)))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).Should(Not(Equal(false)))
					ExpectWithOffset(1, nil == fNil()).Should(Not(Equal(false)))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).Should(Not(Equal(false)))
					ExpectWithOffset(1, nil != fNotNil()).Should(Not(Equal(false)))
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).Should(Not(Equal(true)))
					ExpectWithOffset(1, nil != x).Should(Not(Equal(true)))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).Should(Not(Equal(true)))
					ExpectWithOffset(1, nil == y).Should(Not(Equal(true)))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).Should(Not(Equal(true)))
					ExpectWithOffset(1, nil != fNil()).Should(Not(Equal(true)))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).Should(Not(Equal(true)))
					ExpectWithOffset(1, nil == fNotNil()).Should(Not(Equal(true)))
				})
			})
		})

		Context("test To", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).To(BeTrue())
					ExpectWithOffset(1, nil == x).To(BeTrue())
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).To(BeTrue())
					ExpectWithOffset(1, nil != y).To(BeTrue())
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).To(BeTrue())
					ExpectWithOffset(1, nil == fNil()).To(BeTrue())
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).To(BeTrue())
					ExpectWithOffset(1, nil != fNotNil()).To(BeTrue())
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).To(BeFalse())
					ExpectWithOffset(1, nil != x).To(BeFalse())
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).To(BeFalse())
					ExpectWithOffset(1, nil == y).To(BeFalse())
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).To(BeFalse())
					ExpectWithOffset(1, nil != fNil()).To(BeFalse())
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).To(BeFalse())
					ExpectWithOffset(1, nil == fNotNil()).To(BeFalse())
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).To(Equal(true))
					ExpectWithOffset(1, nil == x).To(Equal(true))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).To(Equal(true))
					ExpectWithOffset(1, nil != y).To(Equal(true))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).To(Equal(true))
					ExpectWithOffset(1, nil == fNil()).To(Equal(true))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).To(Equal(true))
					ExpectWithOffset(1, nil != fNotNil()).To(Equal(true))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).To(Equal(false))
					ExpectWithOffset(1, nil != x).To(Equal(false))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).To(Equal(false))
					ExpectWithOffset(1, nil == y).To(Equal(false))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).To(Equal(false))
					ExpectWithOffset(1, nil != fNil()).To(Equal(false))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).To(Equal(false))
					ExpectWithOffset(1, nil == fNotNil()).To(Equal(false))
				})
			})
		})
		Context("test To(Not())", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).To(Not(BeFalse()))
					ExpectWithOffset(1, nil == x).To(Not(BeFalse()))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).To(Not(BeFalse()))
					ExpectWithOffset(1, nil != y).To(Not(BeFalse()))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).To(Not(BeFalse()))
					ExpectWithOffset(1, nil == fNil()).To(Not(BeFalse()))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).To(Not(BeFalse()))
					ExpectWithOffset(1, nil != fNotNil()).To(Not(BeFalse()))
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).To(Not(BeTrue()))
					ExpectWithOffset(1, nil != x).To(Not(BeTrue()))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).To(Not(BeTrue()))
					ExpectWithOffset(1, nil == y).To(Not(BeTrue()))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).To(Not(BeTrue()))
					ExpectWithOffset(1, nil != fNil()).To(Not(BeTrue()))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).To(Not(BeTrue()))
					ExpectWithOffset(1, nil == fNotNil()).To(Not(BeTrue()))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).To(Not(Equal(false)))
					ExpectWithOffset(1, nil == x).To(Not(Equal(false)))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).To(Not(Equal(false)))
					ExpectWithOffset(1, nil != y).To(Not(Equal(false)))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).To(Not(Equal(false)))
					ExpectWithOffset(1, nil == fNil()).To(Not(Equal(false)))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).To(Not(Equal(false)))
					ExpectWithOffset(1, nil != fNotNil()).To(Not(Equal(false)))
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).To(Not(Equal(true)))
					ExpectWithOffset(1, nil != x).To(Not(Equal(true)))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).To(Not(Equal(true)))
					ExpectWithOffset(1, nil == y).To(Not(Equal(true)))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).To(Not(Equal(true)))
					ExpectWithOffset(1, nil != fNil()).To(Not(Equal(true)))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).To(Not(Equal(true)))
					ExpectWithOffset(1, nil == fNotNil()).To(Not(Equal(true)))
				})
			})
		})

		Context("test ShouldNot", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).ShouldNot(BeFalse())
					ExpectWithOffset(1, nil == x).ShouldNot(BeFalse())
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).ShouldNot(BeFalse())
					ExpectWithOffset(1, nil != y).ShouldNot(BeFalse())
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).ShouldNot(BeFalse())
					ExpectWithOffset(1, nil == fNil()).ShouldNot(BeFalse())
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).ShouldNot(BeFalse())
					ExpectWithOffset(1, nil != fNotNil()).ShouldNot(BeFalse())
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).ShouldNot(BeTrue())
					ExpectWithOffset(1, nil != x).ShouldNot(BeTrue())
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).ShouldNot(BeTrue())
					ExpectWithOffset(1, nil == y).ShouldNot(BeTrue())
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).ShouldNot(BeTrue())
					ExpectWithOffset(1, nil != fNil()).ShouldNot(BeTrue())
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).ShouldNot(BeTrue())
					ExpectWithOffset(1, nil == fNotNil()).ShouldNot(BeTrue())
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).ShouldNot(Equal(true))
					ExpectWithOffset(1, nil != x).ShouldNot(Equal(true))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).ShouldNot(Equal(true))
					ExpectWithOffset(1, nil == y).ShouldNot(Equal(true))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).ShouldNot(Equal(true))
					ExpectWithOffset(1, nil != fNil()).ShouldNot(Equal(true))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).ShouldNot(Equal(true))
					ExpectWithOffset(1, nil == fNotNil()).ShouldNot(Equal(true))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).ShouldNot(Equal(false))
					ExpectWithOffset(1, nil == x).ShouldNot(Equal(false))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).ShouldNot(Equal(false))
					ExpectWithOffset(1, nil != y).ShouldNot(Equal(false))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).ShouldNot(Equal(false))
					ExpectWithOffset(1, nil == fNil()).ShouldNot(Equal(false))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).ShouldNot(Equal(false))
					ExpectWithOffset(1, nil != fNotNil()).ShouldNot(Equal(false))
				})
			})
		})

		Context("test NotTo", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).NotTo(BeFalse())
					ExpectWithOffset(1, nil == x).NotTo(BeFalse())
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).NotTo(BeFalse())
					ExpectWithOffset(1, nil != y).NotTo(BeFalse())
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).NotTo(BeFalse())
					ExpectWithOffset(1, nil == fNil()).NotTo(BeFalse())
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).NotTo(BeFalse())
					ExpectWithOffset(1, nil != fNotNil()).NotTo(BeFalse())
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).NotTo(BeTrue())
					ExpectWithOffset(1, nil != x).NotTo(BeTrue())
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).NotTo(BeTrue())
					ExpectWithOffset(1, nil == y).NotTo(BeTrue())
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).NotTo(BeTrue())
					ExpectWithOffset(1, nil != fNil()).NotTo(BeTrue())
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).NotTo(BeTrue())
					ExpectWithOffset(1, nil == fNotNil()).NotTo(BeTrue())
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).NotTo(Equal(true))
					ExpectWithOffset(1, nil != x).NotTo(Equal(true))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).NotTo(Equal(true))
					ExpectWithOffset(1, nil == y).NotTo(Equal(true))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).NotTo(Equal(true))
					ExpectWithOffset(1, nil != fNil()).NotTo(Equal(true))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).NotTo(Equal(true))
					ExpectWithOffset(1, nil == fNotNil()).NotTo(Equal(true))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).NotTo(Equal(false))
					ExpectWithOffset(1, nil == x).NotTo(Equal(false))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).NotTo(Equal(false))
					ExpectWithOffset(1, nil != y).NotTo(Equal(false))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).NotTo(Equal(false))
					ExpectWithOffset(1, nil == fNil()).NotTo(Equal(false))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).NotTo(Equal(false))
					ExpectWithOffset(1, nil != fNotNil()).NotTo(Equal(false))
				})
			})
		})
		Context("test ToNot", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).ToNot(BeFalse())
					ExpectWithOffset(1, nil == x).ToNot(BeFalse())
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).ToNot(BeFalse())
					ExpectWithOffset(1, nil != y).ToNot(BeFalse())
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).ToNot(BeFalse())
					ExpectWithOffset(1, nil == fNil()).ToNot(BeFalse())
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).ToNot(BeFalse())
					ExpectWithOffset(1, nil != fNotNil()).ToNot(BeFalse())
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).ToNot(BeTrue())
					ExpectWithOffset(1, nil != x).ToNot(BeTrue())
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).ToNot(BeTrue())
					ExpectWithOffset(1, nil == y).ToNot(BeTrue())
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).ToNot(BeTrue())
					ExpectWithOffset(1, nil != fNil()).ToNot(BeTrue())
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).ToNot(BeTrue())
					ExpectWithOffset(1, nil == fNotNil()).ToNot(BeTrue())
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x != nil).ToNot(Equal(true))
					ExpectWithOffset(1, nil != x).ToNot(Equal(true))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y == nil).ToNot(Equal(true))
					ExpectWithOffset(1, nil == y).ToNot(Equal(true))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() != nil).ToNot(Equal(true))
					ExpectWithOffset(1, nil != fNil()).ToNot(Equal(true))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() == nil).ToNot(Equal(true))
					ExpectWithOffset(1, nil == fNotNil()).ToNot(Equal(true))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					ExpectWithOffset(1, x == nil).ToNot(Equal(false))
					ExpectWithOffset(1, nil == x).ToNot(Equal(false))
				})
				It("test non-nil value", func() {
					ExpectWithOffset(1, y != nil).ToNot(Equal(false))
					ExpectWithOffset(1, nil != y).ToNot(Equal(false))
				})
				It("test nil func", func() {
					ExpectWithOffset(1, fNil() == nil).ToNot(Equal(false))
					ExpectWithOffset(1, nil == fNil()).ToNot(Equal(false))
				})
				It("test non-nil func", func() {
					ExpectWithOffset(1, fNotNil() != nil).ToNot(Equal(false))
					ExpectWithOffset(1, nil != fNotNil()).ToNot(Equal(false))
				})
			})
		})
	})

	Context("test Ω", func() {
		Context("test Should", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Ω(x == nil).Should(BeTrue())
					Ω(nil == x).Should(BeTrue())
				})
				It("test non-nil value", func() {
					Ω(y != nil).Should(BeTrue())
					Ω(nil != y).Should(BeTrue())
				})
				It("test nil func", func() {
					Ω(fNil() == nil).Should(BeTrue())
					Ω(nil == fNil()).Should(BeTrue())
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).Should(BeTrue())
					Ω(nil != fNotNil()).Should(BeTrue())
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Ω(x != nil).Should(BeFalse())
					Ω(nil != x).Should(BeFalse())
				})
				It("test non-nil value", func() {
					Ω(y == nil).Should(BeFalse())
					Ω(nil == y).Should(BeFalse())
				})
				It("test nil func", func() {
					Ω(fNil() != nil).Should(BeFalse())
					Ω(nil != fNil()).Should(BeFalse())
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).Should(BeFalse())
					Ω(nil == fNotNil()).Should(BeFalse())
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Ω(x == nil).Should(Equal(true))
					Ω(nil == x).Should(Equal(true))
				})
				It("test non-nil value", func() {
					Ω(y != nil).Should(Equal(true))
					Ω(nil != y).Should(Equal(true))
				})
				It("test nil func", func() {
					Ω(fNil() == nil).Should(Equal(true))
					Ω(nil == fNil()).Should(Equal(true))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).Should(Equal(true))
					Ω(nil != fNotNil()).Should(Equal(true))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Ω(x != nil).Should(Equal(false))
					Ω(nil != x).Should(Equal(false))
				})
				It("test non-nil value", func() {
					Ω(y == nil).Should(Equal(false))
					Ω(nil == y).Should(Equal(false))
				})
				It("test nil func", func() {
					Ω(fNil() != nil).Should(Equal(false))
					Ω(nil != fNil()).Should(Equal(false))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).Should(Equal(false))
					Ω(nil == fNotNil()).Should(Equal(false))
				})
			})
		})
		Context("test Should(Not())", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Ω(x == nil).Should(Not(BeFalse()))
					Ω(nil == x).Should(Not(BeFalse()))
				})
				It("test non-nil value", func() {
					Ω(y != nil).Should(Not(BeFalse()))
					Ω(nil != y).Should(Not(BeFalse()))
				})
				It("test nil func", func() {
					Ω(fNil() == nil).Should(Not(BeFalse()))
					Ω(nil == fNil()).Should(Not(BeFalse()))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).Should(Not(BeFalse()))
					Ω(nil != fNotNil()).Should(Not(BeFalse()))
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Ω(x != nil).Should(Not(BeTrue()))
					Ω(nil != x).Should(Not(BeTrue()))
				})
				It("test non-nil value", func() {
					Ω(y == nil).Should(Not(BeTrue()))
					Ω(nil == y).Should(Not(BeTrue()))
				})
				It("test nil func", func() {
					Ω(fNil() != nil).Should(Not(BeTrue()))
					Ω(nil != fNil()).Should(Not(BeTrue()))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).Should(Not(BeTrue()))
					Ω(nil == fNotNil()).Should(Not(BeTrue()))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Ω(x == nil).Should(Not(Equal(false)))
					Ω(nil == x).Should(Not(Equal(false)))
				})
				It("test non-nil value", func() {
					Ω(y != nil).Should(Not(Equal(false)))
					Ω(nil != y).Should(Not(Equal(false)))
				})
				It("test nil func", func() {
					Ω(fNil() == nil).Should(Not(Equal(false)))
					Ω(nil == fNil()).Should(Not(Equal(false)))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).Should(Not(Equal(false)))
					Ω(nil != fNotNil()).Should(Not(Equal(false)))
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Ω(x != nil).Should(Not(Equal(true)))
					Ω(nil != x).Should(Not(Equal(true)))
				})
				It("test non-nil value", func() {
					Ω(y == nil).Should(Not(Equal(true)))
					Ω(nil == y).Should(Not(Equal(true)))
				})
				It("test nil func", func() {
					Ω(fNil() != nil).Should(Not(Equal(true)))
					Ω(nil != fNil()).Should(Not(Equal(true)))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).Should(Not(Equal(true)))
					Ω(nil == fNotNil()).Should(Not(Equal(true)))
				})
			})
		})

		Context("test To", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Ω(x == nil).To(BeTrue())
					Ω(nil == x).To(BeTrue())
				})
				It("test non-nil value", func() {
					Ω(y != nil).To(BeTrue())
					Ω(nil != y).To(BeTrue())
				})
				It("test nil func", func() {
					Ω(fNil() == nil).To(BeTrue())
					Ω(nil == fNil()).To(BeTrue())
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).To(BeTrue())
					Ω(nil != fNotNil()).To(BeTrue())
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Ω(x != nil).To(BeFalse())
					Ω(nil != x).To(BeFalse())
				})
				It("test non-nil value", func() {
					Ω(y == nil).To(BeFalse())
					Ω(nil == y).To(BeFalse())
				})
				It("test nil func", func() {
					Ω(fNil() != nil).To(BeFalse())
					Ω(nil != fNil()).To(BeFalse())
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).To(BeFalse())
					Ω(nil == fNotNil()).To(BeFalse())
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Ω(x == nil).To(Equal(true))
					Ω(nil == x).To(Equal(true))
				})
				It("test non-nil value", func() {
					Ω(y != nil).To(Equal(true))
					Ω(nil != y).To(Equal(true))
				})
				It("test nil func", func() {
					Ω(fNil() == nil).To(Equal(true))
					Ω(nil == fNil()).To(Equal(true))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).To(Equal(true))
					Ω(nil != fNotNil()).To(Equal(true))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Ω(x != nil).To(Equal(false))
					Ω(nil != x).To(Equal(false))
				})
				It("test non-nil value", func() {
					Ω(y == nil).To(Equal(false))
					Ω(nil == y).To(Equal(false))
				})
				It("test nil func", func() {
					Ω(fNil() != nil).To(Equal(false))
					Ω(nil != fNil()).To(Equal(false))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).To(Equal(false))
					Ω(nil == fNotNil()).To(Equal(false))
				})
			})
		})
		Context("test To(Not())", func() {
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Ω(x == nil).To(Not(BeFalse()))
					Ω(nil == x).To(Not(BeFalse()))
				})
				It("test non-nil value", func() {
					Ω(y != nil).To(Not(BeFalse()))
					Ω(nil != y).To(Not(BeFalse()))
				})
				It("test nil func", func() {
					Ω(fNil() == nil).To(Not(BeFalse()))
					Ω(nil == fNil()).To(Not(BeFalse()))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).To(Not(BeFalse()))
					Ω(nil != fNotNil()).To(Not(BeFalse()))
				})
			})
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Ω(x != nil).To(Not(BeTrue()))
					Ω(nil != x).To(Not(BeTrue()))
				})
				It("test non-nil value", func() {
					Ω(y == nil).To(Not(BeTrue()))
					Ω(nil == y).To(Not(BeTrue()))
				})
				It("test nil func", func() {
					Ω(fNil() != nil).To(Not(BeTrue()))
					Ω(nil != fNil()).To(Not(BeTrue()))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).To(Not(BeTrue()))
					Ω(nil == fNotNil()).To(Not(BeTrue()))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Ω(x == nil).To(Not(Equal(false)))
					Ω(nil == x).To(Not(Equal(false)))
				})
				It("test non-nil value", func() {
					Ω(y != nil).To(Not(Equal(false)))
					Ω(nil != y).To(Not(Equal(false)))
				})
				It("test nil func", func() {
					Ω(fNil() == nil).To(Not(Equal(false)))
					Ω(nil == fNil()).To(Not(Equal(false)))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).To(Not(Equal(false)))
					Ω(nil != fNotNil()).To(Not(Equal(false)))
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Ω(x != nil).To(Not(Equal(true)))
					Ω(nil != x).To(Not(Equal(true)))
				})
				It("test non-nil value", func() {
					Ω(y == nil).To(Not(Equal(true)))
					Ω(nil == y).To(Not(Equal(true)))
				})
				It("test nil func", func() {
					Ω(fNil() != nil).To(Not(Equal(true)))
					Ω(nil != fNil()).To(Not(Equal(true)))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).To(Not(Equal(true)))
					Ω(nil == fNotNil()).To(Not(Equal(true)))
				})
			})
		})

		Context("test ShouldNot", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Ω(x == nil).ShouldNot(BeFalse())
					Ω(nil == x).ShouldNot(BeFalse())
				})
				It("test non-nil value", func() {
					Ω(y != nil).ShouldNot(BeFalse())
					Ω(nil != y).ShouldNot(BeFalse())
				})
				It("test nil func", func() {
					Ω(fNil() == nil).ShouldNot(BeFalse())
					Ω(nil == fNil()).ShouldNot(BeFalse())
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).ShouldNot(BeFalse())
					Ω(nil != fNotNil()).ShouldNot(BeFalse())
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Ω(x != nil).ShouldNot(BeTrue())
					Ω(nil != x).ShouldNot(BeTrue())
				})
				It("test non-nil value", func() {
					Ω(y == nil).ShouldNot(BeTrue())
					Ω(nil == y).ShouldNot(BeTrue())
				})
				It("test nil func", func() {
					Ω(fNil() != nil).ShouldNot(BeTrue())
					Ω(nil != fNil()).ShouldNot(BeTrue())
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).ShouldNot(BeTrue())
					Ω(nil == fNotNil()).ShouldNot(BeTrue())
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Ω(x != nil).ShouldNot(Equal(true))
					Ω(nil != x).ShouldNot(Equal(true))
				})
				It("test non-nil value", func() {
					Ω(y == nil).ShouldNot(Equal(true))
					Ω(nil == y).ShouldNot(Equal(true))
				})
				It("test nil func", func() {
					Ω(fNil() != nil).ShouldNot(Equal(true))
					Ω(nil != fNil()).ShouldNot(Equal(true))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).ShouldNot(Equal(true))
					Ω(nil == fNotNil()).ShouldNot(Equal(true))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Ω(x == nil).ShouldNot(Equal(false))
					Ω(nil == x).ShouldNot(Equal(false))
				})
				It("test non-nil value", func() {
					Ω(y != nil).ShouldNot(Equal(false))
					Ω(nil != y).ShouldNot(Equal(false))
				})
				It("test nil func", func() {
					Ω(fNil() == nil).ShouldNot(Equal(false))
					Ω(nil == fNil()).ShouldNot(Equal(false))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).ShouldNot(Equal(false))
					Ω(nil != fNotNil()).ShouldNot(Equal(false))
				})
			})
		})

		Context("test NotTo", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Ω(x == nil).NotTo(BeFalse())
					Ω(nil == x).NotTo(BeFalse())
				})
				It("test non-nil value", func() {
					Ω(y != nil).NotTo(BeFalse())
					Ω(nil != y).NotTo(BeFalse())
				})
				It("test nil func", func() {
					Ω(fNil() == nil).NotTo(BeFalse())
					Ω(nil == fNil()).NotTo(BeFalse())
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).NotTo(BeFalse())
					Ω(nil != fNotNil()).NotTo(BeFalse())
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Ω(x != nil).NotTo(BeTrue())
					Ω(nil != x).NotTo(BeTrue())
				})
				It("test non-nil value", func() {
					Ω(y == nil).NotTo(BeTrue())
					Ω(nil == y).NotTo(BeTrue())
				})
				It("test nil func", func() {
					Ω(fNil() != nil).NotTo(BeTrue())
					Ω(nil != fNil()).NotTo(BeTrue())
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).NotTo(BeTrue())
					Ω(nil == fNotNil()).NotTo(BeTrue())
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Ω(x != nil).NotTo(Equal(true))
					Ω(nil != x).NotTo(Equal(true))
				})
				It("test non-nil value", func() {
					Ω(y == nil).NotTo(Equal(true))
					Ω(nil == y).NotTo(Equal(true))
				})
				It("test nil func", func() {
					Ω(fNil() != nil).NotTo(Equal(true))
					Ω(nil != fNil()).NotTo(Equal(true))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).NotTo(Equal(true))
					Ω(nil == fNotNil()).NotTo(Equal(true))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Ω(x == nil).NotTo(Equal(false))
					Ω(nil == x).NotTo(Equal(false))
				})
				It("test non-nil value", func() {
					Ω(y != nil).NotTo(Equal(false))
					Ω(nil != y).NotTo(Equal(false))
				})
				It("test nil func", func() {
					Ω(fNil() == nil).NotTo(Equal(false))
					Ω(nil == fNil()).NotTo(Equal(false))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).NotTo(Equal(false))
					Ω(nil != fNotNil()).NotTo(Equal(false))
				})
			})
		})
		Context("test ToNot", func() {
			Context("test BeFalse", func() {
				It("test nil value", func() {
					Ω(x == nil).ToNot(BeFalse())
					Ω(nil == x).ToNot(BeFalse())
				})
				It("test non-nil value", func() {
					Ω(y != nil).ToNot(BeFalse())
					Ω(nil != y).ToNot(BeFalse())
				})
				It("test nil func", func() {
					Ω(fNil() == nil).ToNot(BeFalse())
					Ω(nil == fNil()).ToNot(BeFalse())
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).ToNot(BeFalse())
					Ω(nil != fNotNil()).ToNot(BeFalse())
				})
			})
			Context("test BeTrue", func() {
				It("test nil value", func() {
					Ω(x != nil).ToNot(BeTrue())
					Ω(nil != x).ToNot(BeTrue())
				})
				It("test non-nil value", func() {
					Ω(y == nil).ToNot(BeTrue())
					Ω(nil == y).ToNot(BeTrue())
				})
				It("test nil func", func() {
					Ω(fNil() != nil).ToNot(BeTrue())
					Ω(nil != fNil()).ToNot(BeTrue())
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).ToNot(BeTrue())
					Ω(nil == fNotNil()).ToNot(BeTrue())
				})
			})
			Context("test Equal(true)", func() {
				It("test nil value", func() {
					Ω(x != nil).ToNot(Equal(true))
					Ω(nil != x).ToNot(Equal(true))
				})
				It("test non-nil value", func() {
					Ω(y == nil).ToNot(Equal(true))
					Ω(nil == y).ToNot(Equal(true))
				})
				It("test nil func", func() {
					Ω(fNil() != nil).ToNot(Equal(true))
					Ω(nil != fNil()).ToNot(Equal(true))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() == nil).ToNot(Equal(true))
					Ω(nil == fNotNil()).ToNot(Equal(true))
				})
			})
			Context("test Equal(false)", func() {
				It("test nil value", func() {
					Ω(x == nil).ToNot(Equal(false))
					Ω(nil == x).ToNot(Equal(false))
				})
				It("test non-nil value", func() {
					Ω(y != nil).ToNot(Equal(false))
					Ω(nil != y).ToNot(Equal(false))
				})
				It("test nil func", func() {
					Ω(fNil() == nil).ToNot(Equal(false))
					Ω(nil == fNil()).ToNot(Equal(false))
				})
				It("test non-nil func", func() {
					Ω(fNotNil() != nil).ToNot(Equal(false))
					Ω(nil != fNotNil()).ToNot(Equal(false))
				})
			})
		})
	})
})
