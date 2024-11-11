package issue_171

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("test if issue 171 was solved", func() {
	It("should not trigger warning for nil function vars", func() {
		f := func(string) error { return nil }

		f = nil

		Expect(f).To(BeNil())
	})
})
