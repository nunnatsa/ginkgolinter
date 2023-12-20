package issue_124

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"a/pkg/custommatcher"
)

var _ = Describe("test if issue 124 was solved", func() {
	It("should not have false positive for custom matchers in another packages", func() {
		Expect("a").To(custommatcher.CustomMatcher())
	})
})
