package gomegaonly

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestGomegaOnlyWithName_NewGomegaWithT(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	var err error
	g.Expect(err).ToNot(gomega.BeNil()) // want `ginkgo-linter: wrong error assertion\. Consider using .g\.Expect\(err\)\.To\(gomega\.HaveOccurred\(\)\). instead`

	assertWithName(g, err)
	assertWithNameWithT(g, err)
	assertWithNameGomegaWithT(g, err)
}

func TestGomegaOnlyWithName_NewWithT(t *testing.T) {
	g := gomega.NewWithT(t)

	var err error
	g.Expect(err).ToNot(gomega.BeNil()) // want `ginkgo-linter: wrong error assertion\. Consider using .g\.Expect\(err\)\.To\(gomega\.HaveOccurred\(\)\). instead`
}

func TestGomegaOnlyWithName_NewGomega(t *testing.T) {
	g := gomega.NewGomega(Fail)

	var err error
	g.Expect(err).ToNot(gomega.BeNil()) // want `ginkgo-linter: wrong error assertion\. Consider using .g\.Expect\(err\)\.To\(gomega\.HaveOccurred\(\)\). instead`
}

func assertWithName(g gomega.Gomega, err error) {
	g.Expect(err).To(gomega.BeNil()) // want `ginkgo-linter: wrong error assertion\. Consider using .g\.Expect\(err\)\.ToNot\(gomega\.HaveOccurred\(\)\). instead`
}

func assertWithNameWithT(g *gomega.WithT, err error) {
	g.Expect(err).To(gomega.BeNil()) // want `ginkgo-linter: wrong error assertion\. Consider using .g\.Expect\(err\)\.ToNot\(gomega\.HaveOccurred\(\)\). instead`
}

func assertWithNameGomegaWithT(g *gomega.WithT, err error) {
	g.Expect(err).To(gomega.BeNil()) // want `ginkgo-linter: wrong error assertion\. Consider using .g\.Expect\(err\)\.ToNot\(gomega\.HaveOccurred\(\)\). instead`
}

var _ = Describe("check gomega parameter", func() {
	gomega.Eventually(func(g gomega.Gomega) error {
		arr := []int{1, 2, 3}
		g.Expect(len(arr)).Should(gomega.Equal(3)) // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(arr\)\.Should\(gomega\.HaveLen\(3\)\). instead`
		return nil
	}).Should(gomega.Succeed())
})
