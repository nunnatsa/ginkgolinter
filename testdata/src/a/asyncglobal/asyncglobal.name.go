package asyncglobal

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("Checkginkgo", func() {

	It("", func() {
		gomega.Eventually(func(g gomega.Gomega) {
			gomega.Expect(42).To(gomega.Equal(nil))   // want `multiple issues: using a global Expect in an async assertion at.*; wrong nil assertion. Consider using .gomega\.Expect\(42\)\.To\(gomega\.BeNil\(\)\). instead`
			g.Expect(len("abcd")).To(gomega.Equal(4)) // want `wrong length assertion`
		}).Should(gomega.Succeed())
	})

	It("", func() {
		gomega.Eventually(func(g gomega.Gomega) {
			gomega.Ω(42).To(gomega.Equal(42))      // want `using a global Ω in an async assertion at`
			gomega.Expect(42).To(gomega.Equal(42)) // want `using a global Expect in an async assertion at`
			g.Expect(42).To(gomega.Equal(42))
		}).Should(gomega.Succeed())
	})

	It("", func(ctx context.Context) {
		i := 42
		gomega.Eventually(func(ctx context.Context) *int {
			var err error
			gomega.Expect(err).ToNot(gomega.HaveOccurred()) // want `using a global Expect in an async assertion at`
			return &i
		}).WithTimeout(60 * time.Second).WithPolling(time.Second).WithContext(ctx).Should(gomega.HaveValue(gomega.Equal(42)))

		gomega.Eventually(func(g gomega.Gomega) {
			gomega.Expect(42).To(gomega.Equal(42)) // want `using a global Expect in an async assertion at`
			g.Expect(42).To(gomega.Equal(42))
		}).Should(gomega.Succeed())
	})

	It("", func() {
		gomega.Eventually(helperFuncFromAsync).Should(gomega.Succeed())
	})
})
