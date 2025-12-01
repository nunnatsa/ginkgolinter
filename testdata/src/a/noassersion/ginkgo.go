package noassersion

import (
	"time"

	"a/wrappers"

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
		wrappers.MyExpect("aaa")                                                                // want `ginkgo-linter: "MyExpect": missing assertion method\. Expected one of "To/NotTo/ToNot" \(for Expect assertions\) or "Should/ShouldNot" \(for Eventually/Consistently assertions\)`
		wrappers.MyEventually("aaa")                                                            // want `ginkgo-linter: "MyEventually": missing assertion method\. Expected one of "To/NotTo/ToNot" \(for Expect assertions\) or "Should/ShouldNot" \(for Eventually/Consistently assertions\)`
		wrappers.MyConsistently("aaa")                                                          // want `ginkgo-linter: "MyConsistently": missing assertion method\. Expected one of "To/NotTo/ToNot" \(for Expect assertions\) or "Should/ShouldNot" \(for Eventually/Consistently assertions\)`
		wrappers.NewAssertion("aaa")                                                            // want `ginkgo-linter: "NewAssertion": missing assertion method\. Expected one of "To/NotTo/ToNot" \(for Expect assertions\) or "Should/ShouldNot" \(for Eventually/Consistently assertions\)`
		wrappers.NewAsyncAssertion("aaa")                                                       // want `ginkgo-linter: "NewAsyncAssertion": missing assertion method\. Expected one of "To/NotTo/ToNot" \(for Expect assertions\) or "Should/ShouldNot" \(for Eventually/Consistently assertions\)`

		func() Assertion { // want `ginkgo-linter: "": missing assertion method\. Expected one of "To/NotTo/ToNot" \(for Expect assertions\) or "Should/ShouldNot" \(for Eventually/Consistently assertions\)`
			return Expect("aaa")
		}()

		Eventually(func() error { // want `ginkgo-linter: "Eventually": missing assertion method\. Expected "Should\(\)" or "ShouldNot\(\)"`
			return nil
		}).WithTimeout(500 * time.Millisecond).WithPolling(10 * time.Millisecond)
	})
})
