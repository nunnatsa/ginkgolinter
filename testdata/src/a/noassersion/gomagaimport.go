package noassersion

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	g "github.com/onsi/gomega"
)

var _ = Describe("", func() {
	It("should not allow expecting without assertion", func() {
		g.Expect("bbbb") // want `ginkgo-linter: "Expect": missing assertion method\. Expected "To\(\)", "ToNot\(\)" or "NotTo\(\)"`

		g.Eventually( // want `ginkgo-linter: "Eventually": missing assertion method\. Expected "Should\(\)" or "ShouldNot\(\)"`
			func(gg g.Gomega) {
				gg.Expect(nil) // want `ginkgo-linter: "Expect": missing assertion method\. Expected "To\(\)", "ToNot\(\)" or "NotTo\(\)"`
			},
		).Within(time.Second).
			WithPolling(time.Millisecond * 100).
			WithArguments(1).
			WithContext(context.Background())

		g.Consistently(func(gg g.Gomega) { gg.Expect(nil) }).Within(time.Second).WithPolling(time.Millisecond * 100).WithArguments(1).WithContext(context.Background())              // want `ginkgo-linter: "Consistently": missing assertion method. Expected "Should\(\)" or "ShouldNot\(\)"` `ginkgo-linter: "Expect": missing assertion method\. Expected "To\(\)", "ToNot\(\)" or "NotTo\(\)"`
		g.EventuallyWithOffset(1, func(gg g.Gomega) { gg.Expect(nil) }).Within(time.Second).WithPolling(time.Millisecond * 100).WithArguments(1).WithContext(context.Background())   // want `ginkgo-linter: "EventuallyWithOffset": missing assertion method. Expected "Should\(\)" or "ShouldNot\(\)"` `ginkgo-linter: "Expect": missing assertion method\. Expected "To\(\)", "ToNot\(\)" or "NotTo\(\)"`
		g.ConsistentlyWithOffset(1, func(gg g.Gomega) { gg.Expect(nil) }).Within(time.Second).WithPolling(time.Millisecond * 100).WithArguments(1).WithContext(context.Background()) // want `ginkgo-linter: "ConsistentlyWithOffset": missing assertion method. Expected "Should\(\)" or "ShouldNot\(\)"` `ginkgo-linter: "Expect": missing assertion method. Expected "To\(\)", "ToNot\(\)" or "NotTo\(\)"`
	})
})
