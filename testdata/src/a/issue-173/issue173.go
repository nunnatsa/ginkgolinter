package issue_173

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func intErrFunc() (int, error) {
	return 42, nil
}

func strErr() (string, error) {
	return "", errors.New("error")
}

var _ = Describe("test if issue 173 was solved", func() {
	It("should respect the Error() method", func() {
		Expect(intErrFunc()).Error().To(BeNil()) // want `ginkgo-linter: wrong error assertion. Consider using .Expect\(intErrFunc\(\)\)\.Error\(\)\.ToNot\(HaveOccurred\(\)\). instead`
		Expect(intErrFunc()).Error().ToNot(HaveOccurred())
		Expect(intErrFunc()).Error().ToNot(Succeed())
		Expect(strErr()).Error().To(MatchError(ContainSubstring("error")))
	})
})
