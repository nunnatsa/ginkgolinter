package sprintfmsg

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSprintfMsg(t *testing.T) {
	g := NewWithT(t)
	s := "hello"
	g.Expect(s).To(Equal("hello"), `the error is that "%s" is not equal "%s"`, s, "hello")
	g.Expect(s).To(Equal("hello"), fmt.Sprintf(`the error is that "%s" is not equal "%s"`, s, "hello")) // want `ginkgo-linter: the assertion method is a format function. no need to use Sprintf\nconsider using .g\.Expect\(s\)\.To\(Equal\("hello"\), .the error is that "%s" is not equal "%s"., s, "hello"\). instead`

	//both
	g.Expect(len(s)).To(Equal(5), fmt.Sprintf(`the error is that "%s" is not with length of %d`, s, 5)) // want `ginkgo-linter: wrong length assertion\nalso, the assertion method is a format function. no need to use Sprintf\nconsider using .g\.Expect\(s\)\.To\(HaveLen\(5\), .the error is that "%s" is not with length of %d., s, 5\). instead`
}

var _ = Describe("test with ginkgo", func() {
	It("should warn about Sprintf in assertion description", func() {
		s := "a"
		Expect(s).Should(HaveLen(1), fmt.Sprintf("'%s' should have len of 1", s)) // want `ginkgo-linter: the assertion method is a format function\. no need to use Sprintf\nconsider using .Expect\(s\)\.Should\(HaveLen\(1\), "'%s' should have len of 1", s\). instead`
	})

	It("should warn about length assertion + Sprintf in assertion description", func() {
		s := "a"
		Expect(len(s)).Should(Equal(1), fmt.Sprintf("'%s' should have len of 1", s)) // want `ginkgo-linter: wrong length assertion\nalso, the assertion method is a format function\. no need to use Sprintf\nconsider using .Expect\(s\)\.Should\(HaveLen\(1\), "'%s' should have len of 1", s\). instead`
	})

	It("should warn about nil assertion + Sprintf in assertion description", func() {
		var s *int
		Expect(s == nil).Should(Equal(true), fmt.Sprintf("'%v' should be nil", s)) // want `ginkgo-linter: wrong nil assertion\nalso, the assertion method is a format function\. no need to use Sprintf\nconsider using .Expect\(s\)\.Should\(BeNil\(\), "'%v' should be nil", s\). instead`
	})

	It("should warn about error assertion + Sprintf in assertion description", func() {
		var err error
		Expect(err).To(BeNil(), fmt.Sprintf("%v error should be nil", err)) // want `ginkgo-linter: wrong error assertion\nalso, the assertion method is a format function\. no need to use Sprintf\nconsider using .Expect\(err\)\.ToNot\(HaveOccurred\(\), "%v error should be nil", err\). instead`
	})
})
