package gomegaonly

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGomegaOnly_NewGomegaWithT(t *testing.T) {
	g := NewGomegaWithT(t)

	var err error
	g.Expect(err).ToNot(BeNil()) // want `ginkgo-linter: wrong error assertion\. Consider using .g\.Expect\(err\)\.To\(HaveOccurred\(\)\). instead`
	assert(g, err)
	assertWithT(g, err)
	assertGomegaWithT(g, err)
}

func assert(g Gomega, err error) {
	g.Expect(err).To(BeNil())               // want `ginkgo-linter: wrong error assertion\. Consider using .g\.Expect\(err\)\.ToNot\(HaveOccurred\(\)\). instead`
	g.Expect(err).WithOffset(1).To(BeNil()) // want `ginkgo-linter: wrong error assertion\. Consider using .g\.Expect\(err\)\.WithOffset\(1\)\.ToNot\(HaveOccurred\(\)\). instead`
}

func assertWithT(g *WithT, err error) {
	g.Expect(err).To(BeNil()) // want `ginkgo-linter: wrong error assertion\. Consider using .g\.Expect\(err\)\.ToNot\(HaveOccurred\(\)\). instead`
}

func assertGomegaWithT(g *GomegaWithT, err error) {
	g.Expect(err).To(BeNil()) // want `ginkgo-linter: wrong error assertion\. Consider using .g\.Expect\(err\)\.ToNot\(HaveOccurred\(\)\). instead`
}

func TestGomegaOnly_NewWithT(t *testing.T) {
	g := NewWithT(t)

	var err error
	g.Expect(err).ToNot(BeNil())               // want `ginkgo-linter: wrong error assertion\. Consider using .g\.Expect\(err\)\.To\(HaveOccurred\(\)\). instead`
	g.Expect(err).WithOffset(1).ToNot(BeNil()) // want `ginkgo-linter: wrong error assertion\. Consider using .g\.Expect\(err\)\.WithOffset\(1\)\.To\(HaveOccurred\(\)\). instead`
	assert(g, err)
}

func TestGomegaOnly_NewGomega(t *testing.T) {
	g := NewGomega(Fail)

	var err error
	g.Expect(err).ToNot(BeNil()) // want `ginkgo-linter: wrong error assertion\. Consider using .g\.Expect\(err\)\.To\(HaveOccurred\(\)\). instead`
	assert(g, err)
}

var _ = Describe("check gomega parameter", func() {
	Eventually(func(g Gomega) error {
		arr := []int{1, 2, 3}
		g.Expect(len(arr)).Should(Equal(3))         // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(arr\)\.Should\(HaveLen\(3\)\). instead`
		g.Expect(len(arr) == 3).Should(Equal(true)) // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(arr\)\.Should\(HaveLen\(3\)\). instead`
		return nil
	}).Should(Succeed())
})
