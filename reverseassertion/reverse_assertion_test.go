package reverseassertion

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestReverseAssertion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reverse Assertion Suite")
}

var _ = Describe("test ChangeAssertionLogic", func() {
	It("should return the reverse function", func() {
		for _, assertionFunc := range []string{"To", "ToNot", "Should", "ShouldNot"} {
			rev := ChangeAssertionLogic(assertionFunc)
			Expect(rev).ShouldNot(Equal(assertionFunc))
			revRev := ChangeAssertionLogic(rev)
			Expect(revRev).ShouldNot(Equal(rev))
			Expect(revRev).Should(Equal(assertionFunc))
		}
	})

	It("reverse of 'NotTo' should be 'To'", func() {
		rev := ChangeAssertionLogic("NotTo")
		Expect(rev).Should(Equal("To"))
		revRev := ChangeAssertionLogic(rev)
		Expect(revRev).Should(Equal("ToNot"))
	})

	It("should do nothing if the assertion is unknown", func() {
		Expect(ChangeAssertionLogic("unknown")).Should(Equal("unknown"))
	})
})
