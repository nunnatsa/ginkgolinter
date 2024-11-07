package noassersion

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("", func() {
	It("should not allow expecting without assertion", func() {
		Expect("aaa")                                                                           // want `ginkgo-linter: "Expect": missing assertion method\. Expected "To\(\)", "ToNot\(\)" or "NotTo\(\)"`
		Eventually(func() {}, 100*time.Millisecond, 10*time.Millisecond)                        // want `ginkgo-linter: "Eventually": missing assertion method\. Expected "Should\(\)" or "ShouldNot\(\)"`
		Eventually(func() {}).Within(100 * time.Millisecond).WithPolling(10 * time.Millisecond) // want `ginkgo-linter: "Eventually": missing assertion method\. Expected "Should\(\)" or "ShouldNot\(\)"`
		Consistently(func() bool { return true })                                               // want `ginkgo-linter: "Consistently": missing assertion method\. Expected "Should\(\)" or "ShouldNot\(\)"`
		EventuallyWithOffset(1, func(g Gomega) { g.Expect(true) }).WithTimeout(2 * time.Second) // want `ginkgo-linter: "EventuallyWithOffset": missing assertion method. Expected "Should\(\)" or "ShouldNot\(\)"` `ginkgo-linter: "Expect": missing assertion method. Expected "To\(\)", "ToNot\(\)" or "NotTo\(\)"`
		ConsistentlyWithOffset(2, func() bool { return true })                                  // want `ginkgo-linter: "ConsistentlyWithOffset": missing assertion method\. Expected "Should\(\)" or "ShouldNot\(\)"`
		Ω("omega")                                                                              // want `ginkgo-linter: "Ω": missing assertion method\. Expected "Should\(\)", "To\(\)", "ShouldNot\(\)", "ToNot\(\)" or "NotTo\(\)"`
	})
})
