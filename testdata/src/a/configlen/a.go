package configlen

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("test data for the ginkgo-linter", func() {
	Context("test Expect", func() {
		Context("test Should", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).Should(Equal(4))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).Should(Equal(0))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).Should(BeZero())
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).Should(BeNumerically(">", 0))
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).Should(BeNumerically(">=", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).Should(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).Should(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).Should(BeNumerically("==", 11))
			})

			It("should suggest ShouldNot HaveLen", func() {
				Expect(len("abcd")).Should(BeNumerically("!=", 11))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).Should(BeNumerically("==", 0))
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).Should(BeNumerically("!=", 0))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).Should(Equal(len("1234")))
			})
		})
		Context("test Should(Not", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).Should(Not(Equal(4)))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).Should(Not(Equal(0)))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).Should(Not(BeZero()))
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).Should(Not(BeNumerically(">", 0)))
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).Should(Not(BeNumerically(">=", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).Should(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).Should(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).Should(Not(BeNumerically("==", 11)))
			})

			It("should suggest Should HaveLen", func() {
				Expect(len("abcd")).Should(Not(BeNumerically("!=", 11)))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).Should(Not(BeNumerically("==", 0)))
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).Should(Not(BeNumerically("!=", 0)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).Should(Not(Equal(len("12345"))))
			})
		})
		Context("test ShouldNot", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ShouldNot(Equal(4))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ShouldNot(Equal(0))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ShouldNot(BeZero())
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically(">", 0))
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically(">=", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically("==", 11))
			})

			It("should suggest Should HaveLen", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically("!=", 11))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically("==", 0))
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically("!=", 0))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ShouldNot(Equal(len("12345")))
			})
		})
		Context("test ShouldNot(Not", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ShouldNot(Not(Equal(4)))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ShouldNot(Not(Equal(0)))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ShouldNot(Not(BeZero()))
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically(">", 0)))
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically(">=", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically("==", 11)))
			})

			It("should suggest ShouldNot HaveLen", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically("!=", 11)))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically("==", 0)))
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically("!=", 0)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ShouldNot(Not(Equal(len("1234"))))
			})
		})
		Context("test To", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).To(Equal(4))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).To(Equal(0))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).To(BeZero())
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).To(BeNumerically(">", 0))
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).To(BeNumerically(">=", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).To(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).To(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).To(BeNumerically("==", 11))
			})

			It("should suggest ToNot HaveLen", func() {
				Expect(len("abcd")).To(BeNumerically("!=", 11))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).To(BeNumerically("==", 0))
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).To(BeNumerically("!=", 0))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).To(Equal(len("1234")))
			})
		})
		Context("test To(Not", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).To(Not(Equal(4)))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).To(Not(Equal(0)))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).To(Not(BeZero()))
			})

			It("should suggest To BeEmpty", func() {
				Expect(len("abcd")).To(Not(BeNumerically(">", 0)))
			})

			It("should suggest To BeEmpty", func() {
				Expect(len("abcd")).To(Not(BeNumerically(">=", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).To(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).To(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).To(Not(BeNumerically("==", 11)))
			})

			It("should suggest To HaveLen", func() {
				Expect(len("abcd")).To(Not(BeNumerically("!=", 11)))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).To(Not(BeNumerically("==", 0)))
			})

			It("should suggest To BeEmpty", func() {
				Expect(len("abcd")).To(Not(BeNumerically("!=", 0)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).To(Not(Equal(len("12345"))))
			})
		})
		Context("test ToNot", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ToNot(Equal(4))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ToNot(Equal(0))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ToNot(BeZero())
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ToNot(BeNumerically(">", 0))
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ToNot(BeNumerically(">=", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ToNot(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ToNot(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ToNot(BeNumerically("==", 11))
			})

			It("should suggest Should HaveLen", func() {
				Expect(len("abcd")).ToNot(BeNumerically("!=", 11))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).ToNot(BeNumerically("==", 0))
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ToNot(BeNumerically("!=", 0))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ToNot(Equal(len("12345")))
			})
		})
		Context("test ToNot(Not", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ToNot(Not(Equal(4)))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ToNot(Not(Equal(0)))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ToNot(Not(BeZero()))
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically(">", 0)))
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically(">=", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically("==", 11)))
			})

			It("should suggest ToNot HaveLen", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically("!=", 11)))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically("==", 0)))
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically("!=", 0)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ToNot(Not(Equal(len("1234"))))
			})
		})
		Context("test NotTo", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).NotTo(Equal(4))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).NotTo(Equal(0))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).NotTo(BeZero())
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).NotTo(BeNumerically(">", 0))
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).NotTo(BeNumerically(">=", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).NotTo(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).NotTo(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).NotTo(BeNumerically("==", 11))
			})

			It("should suggest Should HaveLen", func() {
				Expect(len("abcd")).NotTo(BeNumerically("!=", 11))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).NotTo(BeNumerically("==", 0))
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).NotTo(BeNumerically("!=", 0))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).NotTo(Equal(len("12345")))
			})
		})
		Context("test NotTo(Not", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).NotTo(Not(Equal(4)))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).NotTo(Not(Equal(0)))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).NotTo(Not(BeZero()))
			})

			It("should suggest NotTo BeEmpty", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically(">", 0)))
			})

			It("should suggest NotTo BeEmpty", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically(">=", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically("==", 11)))
			})

			It("should suggest NotTo HaveLen", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically("!=", 11)))
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically("==", 0)))
			})

			It("should suggest NotTo BeEmpty", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically("!=", 0)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).NotTo(Not(Equal(len("1234"))))
			})
		})
	})
	Context("test ExpectWithOffset", func() {
		Context("test Should", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).Should(Equal(4))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).Should(Equal(0))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).Should(BeZero())
			})

			It("should suggest ShouldNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).Should(BeNumerically(">", 0))
			})

			It("should suggest ShouldNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).Should(BeNumerically(">=", 1))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).Should(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).Should(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).Should(BeNumerically("==", 11))
			})

			It("should suggest ShouldNot HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).Should(BeNumerically("!=", 11))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).Should(BeNumerically("==", 0))
			})

			It("should suggest ShouldNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).Should(BeNumerically("!=", 0))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).Should(Equal(len("1234")))
			})
		})
		Context("test Should(Not", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(Equal(4)))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).Should(Not(Equal(0)))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).Should(Not(BeZero()))
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(BeNumerically(">", 0)))
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(BeNumerically(">=", 1)))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(BeNumerically("==", 11)))
			})

			It("should suggest Should HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(BeNumerically("!=", 11)))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(BeNumerically("==", 0)))
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(BeNumerically("!=", 0)))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(Equal(len("12345"))))
			})
		})
		Context("test ShouldNot", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Equal(4))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).ShouldNot(Equal(0))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).ShouldNot(BeZero())
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(BeNumerically(">", 0))
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(BeNumerically(">=", 1))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(BeNumerically("==", 11))
			})

			It("should suggest Should HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(BeNumerically("!=", 11))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(BeNumerically("==", 0))
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(BeNumerically("!=", 0))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Equal(len("12345")))
			})
		})
		Context("test ShouldNot(Not", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(Equal(4)))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).ShouldNot(Not(Equal(0)))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).ShouldNot(Not(BeZero()))
			})

			It("should suggest ShouldNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(BeNumerically(">", 0)))
			})

			It("should suggest ShouldNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(BeNumerically(">=", 1)))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(BeNumerically("==", 11)))
			})

			It("should suggest ShouldNot HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(BeNumerically("!=", 11)))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(BeNumerically("==", 0)))
			})

			It("should suggest ShouldNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(BeNumerically("!=", 0)))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(Equal(len("1234"))))
			})
		})
		Context("test To", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).To(Equal(4))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).To(Equal(0))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).To(BeZero())
			})

			It("should suggest ToNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).To(BeNumerically(">", 0))
			})

			It("should suggest ToNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).To(BeNumerically(">=", 1))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).To(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).To(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).To(BeNumerically("==", 11))
			})

			It("should suggest ToNot HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).To(BeNumerically("!=", 11))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).To(BeNumerically("==", 0))
			})

			It("should suggest ToNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).To(BeNumerically("!=", 0))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).To(Equal(len("1234")))
			})
		})
		Context("test To(Not", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(Equal(4)))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).To(Not(Equal(0)))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).To(Not(BeZero()))
			})

			It("should suggest To BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(BeNumerically(">", 0)))
			})

			It("should suggest To BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(BeNumerically(">=", 1)))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(BeNumerically("==", 11)))
			})

			It("should suggest To HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(BeNumerically("!=", 11)))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(BeNumerically("==", 0)))
			})

			It("should suggest To BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(BeNumerically("!=", 0)))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(Equal(len("12345"))))
			})
		})
		Context("test ToNot", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Equal(4))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).ToNot(Equal(0))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).ToNot(BeZero())
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(BeNumerically(">", 0))
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(BeNumerically(">=", 1))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(BeNumerically("==", 11))
			})

			It("should suggest Should HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(BeNumerically("!=", 11))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(BeNumerically("==", 0))
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(BeNumerically("!=", 0))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Equal(len("12345")))
			})
		})
		Context("test ToNot(Not", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(Equal(4)))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).ToNot(Not(Equal(0)))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).ToNot(Not(BeZero()))
			})

			It("should suggest ToNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(BeNumerically(">", 0)))
			})

			It("should suggest ToNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(BeNumerically(">=", 1)))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(BeNumerically("==", 11)))
			})

			It("should suggest ToNot HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(BeNumerically("!=", 11)))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(BeNumerically("==", 0)))
			})

			It("should suggest ToNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(BeNumerically("!=", 0)))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(Equal(len("1234"))))
			})
		})
		Context("test NotTo", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Equal(4))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).NotTo(Equal(0))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).NotTo(BeZero())
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(BeNumerically(">", 0))
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(BeNumerically(">=", 1))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(BeNumerically("==", 11))
			})

			It("should suggest Should HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(BeNumerically("!=", 11))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(BeNumerically("==", 0))
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(BeNumerically("!=", 0))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Equal(len("12345")))
			})
		})
		Context("test NotTo(Not", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(Equal(4)))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).NotTo(Not(Equal(0)))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).NotTo(Not(BeZero()))
			})

			It("should suggest NotTo BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(BeNumerically(">", 0)))
			})

			It("should suggest NotTo BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(BeNumerically(">=", 1)))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(BeNumerically("==", 11)))
			})

			It("should suggest NotTo HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(BeNumerically("!=", 11)))
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(BeNumerically("==", 0)))
			})

			It("should suggest NotTo BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(BeNumerically("!=", 0)))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(Equal(len("1234"))))
			})
		})
	})
	Context("test Ω", func() {
		Context("test Should", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).Should(Equal(4))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).Should(Equal(0))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).Should(BeZero())
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Ω(len("abcd")).Should(BeNumerically(">", 0))
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Ω(len("abcd")).Should(BeNumerically(">=", 1))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).Should(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).Should(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).Should(BeNumerically("==", 11))
			})

			It("should suggest ShouldNot HaveLen", func() {
				Ω(len("abcd")).Should(BeNumerically("!=", 11))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).Should(BeNumerically("==", 0))
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Ω(len("abcd")).Should(BeNumerically("!=", 0))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).Should(Equal(len("1234")))
			})
		})
		Context("test Should(Not", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).Should(Not(Equal(4)))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).Should(Not(Equal(0)))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).Should(Not(BeZero()))
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).Should(Not(BeNumerically(">", 0)))
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).Should(Not(BeNumerically(">=", 1)))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).Should(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).Should(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).Should(Not(BeNumerically("==", 11)))
			})

			It("should suggest Should HaveLen", func() {
				Ω(len("abcd")).Should(Not(BeNumerically("!=", 11)))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).Should(Not(BeNumerically("==", 0)))
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).Should(Not(BeNumerically("!=", 0)))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).Should(Not(Equal(len("12345"))))
			})
		})
		Context("test ShouldNot", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ShouldNot(Equal(4))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).ShouldNot(Equal(0))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).ShouldNot(BeZero())
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).ShouldNot(BeNumerically(">", 0))
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).ShouldNot(BeNumerically(">=", 1))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).ShouldNot(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).ShouldNot(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ShouldNot(BeNumerically("==", 11))
			})

			It("should suggest Should HaveLen", func() {
				Ω(len("abcd")).ShouldNot(BeNumerically("!=", 11))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).ShouldNot(BeNumerically("==", 0))
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).ShouldNot(BeNumerically("!=", 0))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ShouldNot(Equal(len("12345")))
			})
		})
		Context("test ShouldNot(Not", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ShouldNot(Not(Equal(4)))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).ShouldNot(Not(Equal(0)))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).ShouldNot(Not(BeZero()))
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Ω(len("abcd")).ShouldNot(Not(BeNumerically(">", 0)))
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Ω(len("abcd")).ShouldNot(Not(BeNumerically(">=", 1)))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).ShouldNot(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).ShouldNot(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ShouldNot(Not(BeNumerically("==", 11)))
			})

			It("should suggest ShouldNot HaveLen", func() {
				Ω(len("abcd")).ShouldNot(Not(BeNumerically("!=", 11)))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).ShouldNot(Not(BeNumerically("==", 0)))
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Ω(len("abcd")).ShouldNot(Not(BeNumerically("!=", 0)))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ShouldNot(Not(Equal(len("1234"))))
			})
		})
		Context("test To", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).To(Equal(4))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).To(Equal(0))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).To(BeZero())
			})

			It("should suggest ToNot BeEmpty", func() {
				Ω(len("abcd")).To(BeNumerically(">", 0))
			})

			It("should suggest ToNot BeEmpty", func() {
				Ω(len("abcd")).To(BeNumerically(">=", 1))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).To(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).To(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).To(BeNumerically("==", 11))
			})

			It("should suggest ToNot HaveLen", func() {
				Ω(len("abcd")).To(BeNumerically("!=", 11))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).To(BeNumerically("==", 0))
			})

			It("should suggest ToNot BeEmpty", func() {
				Ω(len("abcd")).To(BeNumerically("!=", 0))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).To(Equal(len("1234")))
			})
		})
		Context("test To(Not", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).To(Not(Equal(4)))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).To(Not(Equal(0)))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).To(Not(BeZero()))
			})

			It("should suggest To BeEmpty", func() {
				Ω(len("abcd")).To(Not(BeNumerically(">", 0)))
			})

			It("should suggest To BeEmpty", func() {
				Ω(len("abcd")).To(Not(BeNumerically(">=", 1)))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).To(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).To(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).To(Not(BeNumerically("==", 11)))
			})

			It("should suggest To HaveLen", func() {
				Ω(len("abcd")).To(Not(BeNumerically("!=", 11)))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).To(Not(BeNumerically("==", 0)))
			})

			It("should suggest To BeEmpty", func() {
				Ω(len("abcd")).To(Not(BeNumerically("!=", 0)))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).To(Not(Equal(len("12345"))))
			})
		})
		Context("test ToNot", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ToNot(Equal(4))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).ToNot(Equal(0))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).ToNot(BeZero())
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).ToNot(BeNumerically(">", 0))
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).ToNot(BeNumerically(">=", 1))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).ToNot(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).ToNot(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ToNot(BeNumerically("==", 11))
			})

			It("should suggest Should HaveLen", func() {
				Ω(len("abcd")).ToNot(BeNumerically("!=", 11))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).ToNot(BeNumerically("==", 0))
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).ToNot(BeNumerically("!=", 0))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ToNot(Equal(len("12345")))
			})
		})
		Context("test ToNot(Not", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ToNot(Not(Equal(4)))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).ToNot(Not(Equal(0)))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).ToNot(Not(BeZero()))
			})

			It("should suggest ToNot BeEmpty", func() {
				Ω(len("abcd")).ToNot(Not(BeNumerically(">", 0)))
			})

			It("should suggest ToNot BeEmpty", func() {
				Ω(len("abcd")).ToNot(Not(BeNumerically(">=", 1)))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).ToNot(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).ToNot(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ToNot(Not(BeNumerically("==", 11)))
			})

			It("should suggest ToNot HaveLen", func() {
				Ω(len("abcd")).ToNot(Not(BeNumerically("!=", 11)))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).ToNot(Not(BeNumerically("==", 0)))
			})

			It("should suggest ToNot BeEmpty", func() {
				Ω(len("abcd")).ToNot(Not(BeNumerically("!=", 0)))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ToNot(Not(Equal(len("1234"))))
			})
		})
		Context("test NotTo", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).NotTo(Equal(4))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).NotTo(Equal(0))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).NotTo(BeZero())
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).NotTo(BeNumerically(">", 0))
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).NotTo(BeNumerically(">=", 1))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).NotTo(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).NotTo(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).NotTo(BeNumerically("==", 11))
			})

			It("should suggest Should HaveLen", func() {
				Ω(len("abcd")).NotTo(BeNumerically("!=", 11))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).NotTo(BeNumerically("==", 0))
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).NotTo(BeNumerically("!=", 0))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).NotTo(Equal(len("12345")))
			})
		})
		Context("test NotTo(Not", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).NotTo(Not(Equal(4)))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).NotTo(Not(Equal(0)))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).NotTo(Not(BeZero()))
			})

			It("should suggest NotTo BeEmpty", func() {
				Ω(len("abcd")).NotTo(Not(BeNumerically(">", 0)))
			})

			It("should suggest NotTo BeEmpty", func() {
				Ω(len("abcd")).NotTo(Not(BeNumerically(">=", 1)))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).NotTo(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).NotTo(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).NotTo(Not(BeNumerically("==", 11)))
			})

			It("should suggest NotTo HaveLen", func() {
				Ω(len("abcd")).NotTo(Not(BeNumerically("!=", 11)))
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).NotTo(Not(BeNumerically("==", 0)))
			})

			It("should suggest NotTo BeEmpty", func() {
				Ω(len("abcd")).NotTo(Not(BeNumerically("!=", 0)))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).NotTo(Not(Equal(len("1234"))))
			})
		})
	})

	Context("check cap", func() {
		slice := make([]int, 0, 10)
		It("check cap", func() {
			Expect(cap(slice)).To(Equal(10))
			Expect(cap(slice)).ToNot(BeZero())
			Expect(cap(slice)).ToNot(Equal(5))
			Expect(cap(slice)).To(BeNumerically("==", 10))
			Expect(cap(slice)).ToNot(BeNumerically("!=", 10))
			Expect(cap(slice)).To(BeNumerically(">", 0))
			Expect(cap(slice)).To(BeNumerically(">=", 1))
		})
		It("check cap compare", func() {
			Expect(cap(slice) == 10).To(BeTrue())    // want `wrong comparison assertion. Consider using .Expect\(cap\(slice\)\).To\(Equal\(10\)\). instead`
			Expect(cap(slice) == 10).To(Equal(true)) // want `wrong comparison assertion. Consider using .Expect\(cap\(slice\)\).To\(Equal\(10\)\). instead`
			Expect(cap(slice) != 0).To(BeFalse())    // want `wrong comparison assertion. Consider using .Expect\(cap\(slice\)\).To\(BeZero\(\)\). instead`
			Expect(cap(slice) != 0).To(Equal(false)) // want `wrong comparison assertion. Consider using .Expect\(cap\(slice\)\).To\(BeZero\(\)\). instead`
		})
	})
})
