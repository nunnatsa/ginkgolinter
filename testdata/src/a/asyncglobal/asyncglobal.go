package asyncglobal

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Checkginkgo", func() {

	It("", func() {
		Eventually(func(g Gomega) {
			Expect(42).To(Equal(nil))          // want `multiple issues:.*using a global Expect in an async assertion at.*; wrong nil assertion. Consider using .Expect\(42\)\.To\(BeNil\(\)\). instead`
			g.Expect(len("abcd")).To(Equal(4)) // want `wrong length assertion`
		}).Should(Succeed())
	})

	It("", func() {
		Eventually(func(g Gomega) {
			Ω(42).To(Equal(42))      // want `using a global Ω in an async assertion at`
			Expect(42).To(Equal(42)) // want `using a global Expect in an async assertion at`
			g.Expect(42).To(Equal(42))
		}).Should(Succeed())
	})

	It("", func(ctx context.Context) {
		i := 42
		Eventually(func(ctx context.Context) *int {
			var err error
			Expect(err).ToNot(HaveOccurred()) // want `using a global Expect in an async assertion`
			return &i
		}).WithTimeout(60 * time.Second).WithPolling(time.Second).WithContext(ctx).Should(HaveValue(Equal(42)))

		Eventually(func(g Gomega) {
			Expect(42).To(Equal(42)) // want `using a global Expect in an async assertion`
			g.Expect(42).To(Equal(42))
		}).Should(Succeed())
	})

	It("", func() {
		Eventually(helperFuncFromAsync).Should(Succeed())
	})

	It("should not fail; no async", func() {
		helperFuncFromIt()
	})
})

func helperFuncFromAsync(g Gomega) {
	Expect(42).To(Equal(42)) // want `using a global Expect in an async assertion`
	g.Expect(42).To(Equal(42))
}

func helperFuncFromIt() {
	Expect(42).To(Equal(42))
}
