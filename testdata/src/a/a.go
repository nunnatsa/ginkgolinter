package a

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var obj = struct{ Items []int }{Items: []int{1234}}

var _ = Describe("test data for the ginkgo-linter", func() {
	Context("test Expect", func() {
		Context("test Should", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).Should(Equal(4)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.Should\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).Should(Equal(0)) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).Should(BeZero()) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).Should(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).Should(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).Should(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).Should(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).Should(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.Should\(HaveLen\(11\)\). instead`
			})

			It("should suggest ShouldNot HaveLen", func() {
				Expect(len("abcd")).Should(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ShouldNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).Should(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).Should(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})
		})
		Context("test Should(Not", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).Should(Not(Equal(4))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ShouldNot\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).Should(Not(Equal(0))) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).Should(Not(BeZero())) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).Should(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).Should(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).Should(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).Should(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).Should(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ShouldNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest Should HaveLen", func() {
				Expect(len("abcd")).Should(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.Should\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).Should(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).Should(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})
		})
		Context("test ShouldNot", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ShouldNot(Equal(4)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ShouldNot\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ShouldNot(Equal(0)) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ShouldNot(BeZero()) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ShouldNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest Should HaveLen", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.Should\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})
		})
		Context("test ShouldNot(Not", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ShouldNot(Not(Equal(4))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.Should\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ShouldNot(Not(Equal(0))) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ShouldNot(Not(BeZero())) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.Should\(HaveLen\(11\)\). instead`
			})

			It("should suggest ShouldNot HaveLen", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ShouldNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})
		})
		Context("test To", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).To(Equal(4)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).To(Equal(0)) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).To(BeZero()) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).To(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).To(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).To(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).To(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).To(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest ToNot HaveLen", func() {
				Expect(len("abcd")).To(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).To(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).To(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})
		})
		Context("test To(Not", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).To(Not(Equal(4))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).To(Not(Equal(0))) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).To(Not(BeZero())) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest To BeEmpty", func() {
				Expect(len("abcd")).To(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest To BeEmpty", func() {
				Expect(len("abcd")).To(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).To(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).To(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).To(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest To HaveLen", func() {
				Expect(len("abcd")).To(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).To(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest To BeEmpty", func() {
				Expect(len("abcd")).To(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})
		})
		Context("test ToNot", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ToNot(Equal(4)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ToNot(Equal(0)) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ToNot(BeZero()) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ToNot(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ToNot(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ToNot(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ToNot(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ToNot(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest Should HaveLen", func() {
				Expect(len("abcd")).ToNot(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).ToNot(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ToNot(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})
		})
		Context("test ToNot(Not", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ToNot(Not(Equal(4))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ToNot(Not(Equal(0))) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ToNot(Not(BeZero())) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest ToNot HaveLen", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})
		})
		Context("test NotTo", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).NotTo(Equal(4)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.NotTo\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).NotTo(Equal(0)) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.NotTo\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).NotTo(BeZero()) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.NotTo\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).NotTo(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).NotTo(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).NotTo(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).NotTo(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).NotTo(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.NotTo\(HaveLen\(11\)\). instead`
			})

			It("should suggest Should HaveLen", func() {
				Expect(len("abcd")).NotTo(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).NotTo(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.NotTo\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).NotTo(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})
		})
		Context("test NotTo(Not", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).NotTo(Not(Equal(4))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).NotTo(Not(Equal(0))) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).NotTo(Not(BeZero())) // want `ginkgo-linter: wrong length check; consider using .Expect\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest NotTo BeEmpty", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest NotTo BeEmpty", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest NotTo HaveLen", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest NotTo BeEmpty", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length check; consider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})
		})
	})
})
