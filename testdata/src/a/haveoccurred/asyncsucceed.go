package haveoccurred

import (
	"context"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func retNonErr() int {
	return 42
}

func retOneErr() error {
	return fmt.Errorf("err")
}

func retValAndErr() (int, error) {
	return 42, fmt.Errorf("err")
}

func retFuncRetErr() func() error {
	return func() error { return fmt.Errorf("err") }
}

var _ = Describe("check async with Success", func() {

	localRetOneErr := func() error {
		return fmt.Errorf("err")
	}

	localRetValAndErr := func() (int, error) {
		return 42, fmt.Errorf("err")
	}

	It("should not trigger warnings for valid use cases", func(ctx context.Context) {
		Eventually(func(g Gomega) {
			g.Expect(true).To(BeTrue())
		}, 500*time.Microsecond, 50*time.Millisecond).Should(Succeed())

		Eventually(ctx, func(g Gomega) {
			g.Expect(true).To(BeTrue())
		}, 500*time.Microsecond, 50*time.Millisecond).
			WithContext(ctx).
			Should(Succeed())

		EventuallyWithOffset(1, ctx, func(g Gomega) {
			g.Expect(true).To(BeTrue())
		}, 500*time.Microsecond, 50*time.Millisecond).
			WithContext(ctx).
			Should(Succeed())

		Eventually(func() error {
			return nil
		}).Should(Succeed())

		Eventually(retOneErr).ShouldNot(Succeed())
		Eventually(localRetOneErr).Should(Succeed())
		Eventually(ctx, localRetOneErr).Should(Succeed())
	})

	It("should check async assertions with Succeed matcher", func(ctx context.Context) {
		Eventually(ctx, func() int { // want `ginkgo-linter: Success matcher only support a single error value, or function with Gomega as its first parameter`
			return 42
		}).Should(Succeed())

		Eventually(retFuncRetErr).ShouldNot(Succeed())                      // want `ginkgo-linter: Success matcher only support a single error value, or function with Gomega as its first parameter`
		Eventually(retValAndErr).ShouldNot(Succeed())                       // want `ginkgo-linter: Success matcher does not support multiple values`
		Eventually(localRetValAndErr).Should(Succeed())                     // want `ginkgo-linter: Success matcher does not support multiple values`
		Eventually(ctx, localRetValAndErr).Should(Succeed())                // want `ginkgo-linter: Success matcher does not support multiple values`
		Consistently(ctx, localRetValAndErr).Should(Succeed())              // want `ginkgo-linter: Success matcher does not support multiple values`
		ConsistentlyWithOffset(1, ctx, localRetValAndErr).Should(Succeed()) // want `ginkgo-linter: Success matcher does not support multiple values`
	})

	It("valid inline func", func(ctx context.Context) {
		Eventually(func(g Gomega, ctx context.Context) {
			g.Expect(true).To(BeTrue())
		}).WithContext(ctx).WithPolling(time.Microsecond).WithTimeout(10 * time.Millisecond).Should(Succeed())
	})

	It("invalid inline funcs: non-error ret value", func(ctx context.Context) {
		Eventually(func(ctx context.Context) int { // want `ginkgo-linter: Success matcher only support a single error value, or function with Gomega as its first parameter`
			return 42
		}).WithContext(ctx).WithPolling(time.Microsecond).WithTimeout(10 * time.Millisecond).Should(Succeed())
	})

	It("invalid inline funcs: non-error ret value + Gomega var", func(ctx context.Context) {
		Eventually(func(g Gomega, ctx context.Context) int { // want `ginkgo-linter: Success matcher only support a single error value, or function with Gomega as its first parameter`
			g.Expect(true) // want `ginkgo-linter: "Expect": missing assertion method. Expected "To\(\)", "ToNot\(\)" or "NotTo\(\)"`
			return 42
		}).WithContext(ctx).WithPolling(time.Microsecond).WithTimeout(10 * time.Millisecond).Should(Succeed())
	})

	It("invalid inline funcs: multi ret values", func(ctx context.Context) {
		Eventually(func() (int, error) { // want `ginkgo-linter: Success matcher does not support multiple values`
			return 42, nil
		}).WithContext(ctx).WithPolling(time.Microsecond).WithTimeout(10 * time.Millisecond).Should(Succeed())
	})

	It("invalid inline funcs: multi ret values + Gomega var", func(ctx context.Context) {
		Eventually(retValAndErr).ShouldNot(Succeed()) // want `ginkgo-linter: Success matcher does not support multiple values`

		Eventually(func(g Gomega) (int, error) { // want `ginkgo-linter: Success matcher does not support multiple values`
			g.Expect(true).To(BeTrue())
			return 42, nil
		}).WithContext(ctx).WithPolling(time.Microsecond).WithTimeout(10 * time.Millisecond).Should(Succeed())
	})
})
