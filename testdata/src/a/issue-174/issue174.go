package issue_171

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func errFunc() func() error {
	return func() error {
		return errors.New("error")
	}
}

var _ = Describe("test if issue 174 was solved", func() {
	It("should not trigger", func() {
		Eventually(errFunc()).Should(MatchError(ContainSubstring("error")))
	})
})
