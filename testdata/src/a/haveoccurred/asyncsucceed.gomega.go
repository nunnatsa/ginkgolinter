package haveoccurred

import (
	"context"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("check async with Success", func() {

	localRetOneErr := func() error {
		return fmt.Errorf("err")
	}

	localRetValAndErr := func() (int, error) {
		return 42, fmt.Errorf("err")
	}

	It("should not trigger warnings for valid use cases", func(ctx context.Context) {
		gomega.Eventually(func(g gomega.Gomega) {
			g.Expect(true).To(gomega.BeTrue())
		}, 500*time.Microsecond, 50*time.Millisecond).Should(gomega.Succeed())

		gomega.Eventually(ctx, func(g gomega.Gomega) {
			g.Expect(true).To(gomega.BeTrue())
		}, 500*time.Microsecond, 50*time.Millisecond).
			WithContext(ctx).
			Should(gomega.Succeed())

		gomega.EventuallyWithOffset(1, ctx, func(g gomega.Gomega) {
			g.Expect(true).To(gomega.BeTrue())
		}, 500*time.Microsecond, 50*time.Millisecond).
			WithContext(ctx).
			Should(gomega.Succeed())

		gomega.Eventually(func() error {
			return nil
		}).Should(gomega.Succeed())

		gomega.Eventually(retOneErr).ShouldNot(gomega.Succeed())
		gomega.Eventually(localRetOneErr).Should(gomega.Succeed())
	})

	It("should check async assertions with Succeed matcher", func(ctx context.Context) {
		gomega.Eventually(ctx, func() int { // want `ginkgo-linter: Success matcher only support a single error value, or function with Gomega as its first parameter`
			return 42
		}).Should(gomega.Succeed())

		gomega.Eventually(retValAndErr).ShouldNot(gomega.Succeed())                       // want `ginkgo-linter: Success matcher does not support multiple values`
		gomega.Eventually(localRetValAndErr).Should(gomega.Succeed())                     // want `ginkgo-linter: Success matcher does not support multiple values`
		gomega.Eventually(ctx, localRetValAndErr).Should(gomega.Succeed())                // want `ginkgo-linter: Success matcher does not support multiple values`
		gomega.Consistently(ctx, localRetValAndErr).Should(gomega.Succeed())              // want `ginkgo-linter: Success matcher does not support multiple values`
		gomega.ConsistentlyWithOffset(1, ctx, localRetValAndErr).Should(gomega.Succeed()) // want `ginkgo-linter: Success matcher does not support multiple values`
	})
})
