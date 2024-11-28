package issue_173

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("test if issue 173 was solved", func() {
	It("should respect the Error() method", func() {
		gomega.Expect(intErrFunc()).Error().To(gomega.BeNil()) // want `ginkgo-linter: wrong error assertion. Consider using .gomega\.Expect\(intErrFunc\(\)\)\.Error\(\)\.ToNot\(gomega\.HaveOccurred\(\)\). instead`
		gomega.Expect(intErrFunc()).Error().ToNot(gomega.HaveOccurred())
		gomega.Expect(intErrFunc()).Error().ToNot(gomega.Succeed())
		gomega.Expect(strErr()).Error().To(gomega.MatchError(gomega.ContainSubstring("error")))
	})
})
