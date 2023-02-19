package len

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("test data for the ginkgo-linter", func() {
	Context("test Expect", func() {
		Context("test Should", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).Should(Equal(4)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.Should\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).Should(Equal(0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).Should(BeZero()) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).Should(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).Should(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).Should(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).Should(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).Should(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.Should\(HaveLen\(11\)\). instead`
			})

			It("should suggest ShouldNot HaveLen", func() {
				Expect(len("abcd")).Should(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ShouldNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).Should(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).Should(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).Should(Equal(len("1234"))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.Should\(HaveLen\(len\("1234"\)\)\). instead`
			})
		})
		Context("test Should(Not", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).Should(Not(Equal(4))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ShouldNot\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).Should(Not(Equal(0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).Should(Not(BeZero())) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).Should(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).Should(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).Should(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).Should(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).Should(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ShouldNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest Should HaveLen", func() {
				Expect(len("abcd")).Should(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.Should\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).Should(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).Should(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).Should(Not(Equal(len("12345")))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ShouldNot\(HaveLen\(len\("12345"\)\)\). instead`
			})
		})
		Context("test ShouldNot", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ShouldNot(Equal(4)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ShouldNot\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ShouldNot(Equal(0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ShouldNot(BeZero()) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ShouldNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest Should HaveLen", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.Should\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ShouldNot(Equal(len("12345"))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ShouldNot\(HaveLen\(len\("12345"\)\)\). instead`
			})
		})
		Context("test ShouldNot(Not", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ShouldNot(Not(Equal(4))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.Should\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ShouldNot(Not(Equal(0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ShouldNot(Not(BeZero())) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.Should\(HaveLen\(11\)\). instead`
			})

			It("should suggest ShouldNot HaveLen", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ShouldNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Expect(len("abcd")).ShouldNot(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ShouldNot(Not(Equal(len("1234")))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.Should\(HaveLen\(len\("1234"\)\)\). instead`
			})
		})
		Context("test To", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).To(Equal(4)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).To(Equal(0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).To(BeZero()) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).To(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).To(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).To(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).To(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).To(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest ToNot HaveLen", func() {
				Expect(len("abcd")).To(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).To(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).To(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).To(Equal(len("1234"))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(HaveLen\(len\("1234"\)\)\). instead`
			})
		})
		Context("test To(Not", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).To(Not(Equal(4))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).To(Not(Equal(0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).To(Not(BeZero())) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest To BeEmpty", func() {
				Expect(len("abcd")).To(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest To BeEmpty", func() {
				Expect(len("abcd")).To(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).To(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).To(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).To(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest To HaveLen", func() {
				Expect(len("abcd")).To(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).To(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest To BeEmpty", func() {
				Expect(len("abcd")).To(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).To(Not(Equal(len("12345")))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(HaveLen\(len\("12345"\)\)\). instead`
			})
		})
		Context("test ToNot", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ToNot(Equal(4)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ToNot(Equal(0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ToNot(BeZero()) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ToNot(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ToNot(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ToNot(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ToNot(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ToNot(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest Should HaveLen", func() {
				Expect(len("abcd")).ToNot(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).ToNot(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).ToNot(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ToNot(Equal(len("12345"))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(HaveLen\(len\("12345"\)\)\). instead`
			})
		})
		Context("test ToNot(Not", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ToNot(Not(Equal(4))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ToNot(Not(Equal(0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).ToNot(Not(BeZero())) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest ToNot HaveLen", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Expect(len("abcd")).ToNot(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).ToNot(Not(Equal(len("1234")))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\).To\(HaveLen\(len\("1234"\)\)\). instead`
			})
		})
		Context("test NotTo", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).NotTo(Equal(4)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.NotTo\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).NotTo(Equal(0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.NotTo\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).NotTo(BeZero()) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.NotTo\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).NotTo(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).NotTo(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).NotTo(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).NotTo(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).NotTo(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.NotTo\(HaveLen\(11\)\). instead`
			})

			It("should suggest Should HaveLen", func() {
				Expect(len("abcd")).NotTo(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).NotTo(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.NotTo\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Expect(len("abcd")).NotTo(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).NotTo(Equal(len("12345"))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.NotTo\(HaveLen\(len\("12345"\)\)\). instead`
			})
		})
		Context("test NotTo(Not", func() {
			It("should suggest HaveLen", func() {
				Expect(len("abcd")).NotTo(Not(Equal(4))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).NotTo(Not(Equal(0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("")).NotTo(Not(BeZero())) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest NotTo BeEmpty", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest NotTo BeEmpty", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest NotTo HaveLen", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest NotTo BeEmpty", func() {
				Expect(len("abcd")).NotTo(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Expect(len("abcd")).NotTo(Not(Equal(len("1234")))) // want `ginkgo-linter: wrong length assertion\nconsider using .Expect\("abcd"\)\.To\(HaveLen\(len\("1234"\)\)\). instead`
			})
		})
	})
	Context("test ExpectWithOffset", func() {
		Context("test Should", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).Should(Equal(4)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.Should\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).Should(Equal(0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).Should(BeZero()) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).Should(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).Should(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).Should(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).Should(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).Should(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.Should\(HaveLen\(11\)\). instead`
			})

			It("should suggest ShouldNot HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).Should(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ShouldNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).Should(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).Should(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).Should(Equal(len("1234"))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.Should\(HaveLen\(len\("1234"\)\)\). instead`
			})
		})
		Context("test Should(Not", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(Equal(4))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ShouldNot\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).Should(Not(Equal(0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).Should(Not(BeZero())) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ShouldNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest Should HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.Should\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).Should(Not(Equal(len("12345")))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ShouldNot\(HaveLen\(len\("12345"\)\)\). instead`
			})
		})
		Context("test ShouldNot", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Equal(4)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ShouldNot\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).ShouldNot(Equal(0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).ShouldNot(BeZero()) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ShouldNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest Should HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.Should\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Equal(len("12345"))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ShouldNot\(HaveLen\(len\("12345"\)\)\). instead`
			})
		})
		Context("test ShouldNot(Not", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(Equal(4))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.Should\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).ShouldNot(Not(Equal(0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).ShouldNot(Not(BeZero())) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.Should\(HaveLen\(11\)\). instead`
			})

			It("should suggest ShouldNot HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ShouldNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ShouldNot(Not(Equal(len("1234")))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.Should\(HaveLen\(len\("1234"\)\)\). instead`
			})
		})
		Context("test To", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).To(Equal(4)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).To(Equal(0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).To(BeZero()) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).To(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).To(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).To(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).To(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).To(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest ToNot HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).To(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).To(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).To(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).To(Equal(len("1234"))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(HaveLen\(len\("1234"\)\)\). instead`
			})
		})
		Context("test To(Not", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(Equal(4))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).To(Not(Equal(0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).To(Not(BeZero())) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest To BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest To BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest To HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest To BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).To(Not(Equal(len("12345")))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(HaveLen\(len\("12345"\)\)\). instead`
			})
		})
		Context("test ToNot", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Equal(4)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).ToNot(Equal(0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).ToNot(BeZero()) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest Should HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Equal(len("12345"))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(HaveLen\(len\("12345"\)\)\). instead`
			})
		})
		Context("test ToNot(Not", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(Equal(4))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).ToNot(Not(Equal(0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).ToNot(Not(BeZero())) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest ToNot HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).ToNot(Not(Equal(len("1234")))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\).To\(HaveLen\(len\("1234"\)\)\). instead`
			})
		})
		Context("test NotTo", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Equal(4)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.NotTo\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).NotTo(Equal(0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.NotTo\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).NotTo(BeZero()) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.NotTo\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.NotTo\(HaveLen\(11\)\). instead`
			})

			It("should suggest Should HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.NotTo\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Equal(len("12345"))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.NotTo\(HaveLen\(len\("12345"\)\)\). instead`
			})
		})
		Context("test NotTo(Not", func() {
			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(Equal(4))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).NotTo(Not(Equal(0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("")).NotTo(Not(BeZero())) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, ""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest NotTo BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest NotTo BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest NotTo HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest NotTo BeEmpty", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				ExpectWithOffset(12, len("abcd")).NotTo(Not(Equal(len("1234")))) // want `ginkgo-linter: wrong length assertion\nconsider using .ExpectWithOffset\(12, "abcd"\)\.To\(HaveLen\(len\("1234"\)\)\). instead`
			})
		})
	})
	Context("test Ω", func() {
		Context("test Should", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).Should(Equal(4)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.Should\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).Should(Equal(0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).Should(BeZero()) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Ω(len("abcd")).Should(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Ω(len("abcd")).Should(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).Should(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).Should(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).Should(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.Should\(HaveLen\(11\)\). instead`
			})

			It("should suggest ShouldNot HaveLen", func() {
				Ω(len("abcd")).Should(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ShouldNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).Should(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Ω(len("abcd")).Should(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).Should(Equal(len("1234"))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.Should\(HaveLen\(len\("1234"\)\)\). instead`
			})
		})
		Context("test Should(Not", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).Should(Not(Equal(4))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ShouldNot\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).Should(Not(Equal(0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).Should(Not(BeZero())) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).Should(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).Should(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).Should(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).Should(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).Should(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ShouldNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest Should HaveLen", func() {
				Ω(len("abcd")).Should(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.Should\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).Should(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).Should(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).Should(Not(Equal(len("12345")))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ShouldNot\(HaveLen\(len\("12345"\)\)\). instead`
			})
		})
		Context("test ShouldNot", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ShouldNot(Equal(4)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ShouldNot\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).ShouldNot(Equal(0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).ShouldNot(BeZero()) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).ShouldNot(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).ShouldNot(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).ShouldNot(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).ShouldNot(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ShouldNot(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ShouldNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest Should HaveLen", func() {
				Ω(len("abcd")).ShouldNot(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.Should\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).ShouldNot(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).ShouldNot(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ShouldNot(Equal(len("12345"))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ShouldNot\(HaveLen\(len\("12345"\)\)\). instead`
			})
		})
		Context("test ShouldNot(Not", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ShouldNot(Not(Equal(4))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.Should\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).ShouldNot(Not(Equal(0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).ShouldNot(Not(BeZero())) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Ω(len("abcd")).ShouldNot(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Ω(len("abcd")).ShouldNot(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).ShouldNot(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).ShouldNot(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ShouldNot(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.Should\(HaveLen\(11\)\). instead`
			})

			It("should suggest ShouldNot HaveLen", func() {
				Ω(len("abcd")).ShouldNot(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ShouldNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).ShouldNot(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.Should\(BeEmpty\(\)\). instead`
			})

			It("should suggest ShouldNot BeEmpty", func() {
				Ω(len("abcd")).ShouldNot(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ShouldNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ShouldNot(Not(Equal(len("1234")))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.Should\(HaveLen\(len\("1234"\)\)\). instead`
			})
		})
		Context("test To", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).To(Equal(4)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).To(Equal(0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).To(BeZero()) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Ω(len("abcd")).To(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Ω(len("abcd")).To(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).To(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).To(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).To(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest ToNot HaveLen", func() {
				Ω(len("abcd")).To(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).To(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Ω(len("abcd")).To(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).To(Equal(len("1234"))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(HaveLen\(len\("1234"\)\)\). instead`
			})
		})
		Context("test To(Not", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).To(Not(Equal(4))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).To(Not(Equal(0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).To(Not(BeZero())) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest To BeEmpty", func() {
				Ω(len("abcd")).To(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest To BeEmpty", func() {
				Ω(len("abcd")).To(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).To(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).To(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).To(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest To HaveLen", func() {
				Ω(len("abcd")).To(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).To(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest To BeEmpty", func() {
				Ω(len("abcd")).To(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).To(Not(Equal(len("12345")))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(HaveLen\(len\("12345"\)\)\). instead`
			})
		})
		Context("test ToNot", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ToNot(Equal(4)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).ToNot(Equal(0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).ToNot(BeZero()) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).ToNot(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).ToNot(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).ToNot(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).ToNot(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ToNot(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest Should HaveLen", func() {
				Ω(len("abcd")).ToNot(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).ToNot(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).ToNot(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ToNot(Equal(len("12345"))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(HaveLen\(len\("12345"\)\)\). instead`
			})
		})
		Context("test ToNot(Not", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ToNot(Not(Equal(4))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).ToNot(Not(Equal(0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).ToNot(Not(BeZero())) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Ω(len("abcd")).ToNot(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Ω(len("abcd")).ToNot(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).ToNot(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).ToNot(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ToNot(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest ToNot HaveLen", func() {
				Ω(len("abcd")).ToNot(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).ToNot(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest ToNot BeEmpty", func() {
				Ω(len("abcd")).ToNot(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).ToNot(Not(Equal(len("1234")))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\).To\(HaveLen\(len\("1234"\)\)\). instead`
			})
		})
		Context("test NotTo", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).NotTo(Equal(4)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.NotTo\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).NotTo(Equal(0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.NotTo\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).NotTo(BeZero()) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.NotTo\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).NotTo(BeNumerically(">", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).NotTo(BeNumerically(">=", 1)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).NotTo(BeNumerically(">", 1))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).NotTo(BeNumerically(">=", 11))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).NotTo(BeNumerically("==", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.NotTo\(HaveLen\(11\)\). instead`
			})

			It("should suggest Should HaveLen", func() {
				Ω(len("abcd")).NotTo(BeNumerically("!=", 11)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).NotTo(BeNumerically("==", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.NotTo\(BeEmpty\(\)\). instead`
			})

			It("should suggest Should BeEmpty", func() {
				Ω(len("abcd")).NotTo(BeNumerically("!=", 0)) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).NotTo(Equal(len("12345"))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.NotTo\(HaveLen\(len\("12345"\)\)\). instead`
			})
		})
		Context("test NotTo(Not", func() {
			It("should suggest HaveLen", func() {
				Ω(len("abcd")).NotTo(Not(Equal(4))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(HaveLen\(4\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).NotTo(Not(Equal(0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("")).NotTo(Not(BeZero())) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\(""\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest NotTo BeEmpty", func() {
				Ω(len("abcd")).NotTo(Not(BeNumerically(">", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest NotTo BeEmpty", func() {
				Ω(len("abcd")).NotTo(Not(BeNumerically(">=", 1))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).NotTo(Not(BeNumerically(">", 1)))
			})

			It("should ignore other lengths", func() {
				Ω(len("abcd")).NotTo(Not(BeNumerically(">=", 11)))
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).NotTo(Not(BeNumerically("==", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(HaveLen\(11\)\). instead`
			})

			It("should suggest NotTo HaveLen", func() {
				Ω(len("abcd")).NotTo(Not(BeNumerically("!=", 11))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(HaveLen\(11\)\). instead`
			})

			It("should suggest BeEmpty", func() {
				Ω(len("abcd")).NotTo(Not(BeNumerically("==", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(BeEmpty\(\)\). instead`
			})

			It("should suggest NotTo BeEmpty", func() {
				Ω(len("abcd")).NotTo(Not(BeNumerically("!=", 0))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.ToNot\(BeEmpty\(\)\). instead`
			})

			It("should suggest HaveLen", func() {
				Ω(len("abcd")).NotTo(Not(Equal(len("1234")))) // want `ginkgo-linter: wrong length assertion\nconsider using .Ω\("abcd"\)\.To\(HaveLen\(len\("1234"\)\)\). instead`
			})
		})
	})
	Context("valid len(x) cases", func() {
		Expect(len("abcd")).Should(BeElementOf([]int{1, 2, 3, 4, 5}))
	})
})
