package noassertionindirect

import (
	"a/wrappers"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("", func() {
	It("should not allow expecting without assertion", func() {
		wrappers.Expect("aaa")            // want `ginkgo-linter: "Expect": missing assertion method\. Expected "To\(\)", "ToNot\(\)" or "NotTo\(\)"`
		wrappers.MyExpect("aaa")          // want `ginkgo-linter: "MyExpect": missing assertion method\. Expected one of "To/NotTo/ToNot" \(for Expect assertions\) or "Should/ShouldNot" \(for Eventually/Consistently assertions\)`
		wrappers.MyEventually("aaa")      // want `ginkgo-linter: "MyEventually": missing assertion method\. Expected one of "To/NotTo/ToNot" \(for Expect assertions\) or "Should/ShouldNot" \(for Eventually/Consistently assertions\)`
		wrappers.MyConsistently("aaa")    // want `ginkgo-linter: "MyConsistently": missing assertion method\. Expected one of "To/NotTo/ToNot" \(for Expect assertions\) or "Should/ShouldNot" \(for Eventually/Consistently assertions\)`
		wrappers.NewAssertion("aaa")      // want `ginkgo-linter: "NewAssertion": missing assertion method\. Expected one of "To/NotTo/ToNot" \(for Expect assertions\) or "Should/ShouldNot" \(for Eventually/Consistently assertions\)`
		wrappers.NewAsyncAssertion("aaa") // want `ginkgo-linter: "NewAsyncAssertion": missing assertion method\. Expected one of "To/NotTo/ToNot" \(for Expect assertions\) or "Should/ShouldNot" \(for Eventually/Consistently assertions\)`
	})
})
