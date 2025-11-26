package noassersion

import (
	"a/wrappers"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("", func() {
	It("should not allow expecting without assertion", func() {
		wrappers.MyEventually("aaa") // want `ginkgo-linter: "Expect": missing assertion method\. Expected "To\(\)", "ToNot\(\)" or "NotTo\(\)"`
	})
})
