package gomegaonly

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	gom "github.com/onsi/gomega"
)

func TestGomegaOnlyWithModName_NewGomegaWithT(t *testing.T) {
	g := gom.NewGomegaWithT(t)

	var err error
	g.Expect(err).ToNot(gom.BeNil()) // want `ginkgo-linter: wrong error assertion; consider using .g\.Expect\(err\)\.To\(gom.HaveOccurred\(\)\). instead`

	assertWithModName(g, err)
	assertWithNameModWithT(g, err)
	assertWithModNameGomegaWithT(g, err)
}

func TestGomegaOnlyWithModName_NewWithT(t *testing.T) {
	g := gom.NewWithT(t)

	var err error
	g.Expect(err).ToNot(gom.BeNil()) // want `ginkgo-linter: wrong error assertion; consider using .g\.Expect\(err\)\.To\(gom.HaveOccurred\(\)\). instead`
}

func TestGomegaOnlyWithModName_NewGomega(t *testing.T) {
	g := gom.NewGomega(Fail)

	var err error
	g.Expect(err).ToNot(gom.BeNil()) // want `ginkgo-linter: wrong error assertion; consider using .g\.Expect\(err\)\.To\(gom.HaveOccurred\(\)\). instead`
}

func assertWithModName(g gom.Gomega, err error) {
	g.Expect(err).To(gom.BeNil()) // want `ginkgo-linter: wrong error assertion; consider using .g\.Expect\(err\)\.ToNot\(gom\.HaveOccurred\(\)\). instead`
}

func assertWithNameModWithT(g *gom.WithT, err error) {
	g.Expect(err).To(gom.BeNil()) // want `ginkgo-linter: wrong error assertion; consider using .g\.Expect\(err\)\.ToNot\(gom\.HaveOccurred\(\)\). instead`
}

func assertWithModNameGomegaWithT(g *gom.WithT, err error) {
	g.Expect(err).To(gom.BeNil()) // want `ginkgo-linter: wrong error assertion; consider using .g\.Expect\(err\)\.ToNot\(gom\.HaveOccurred\(\)\). instead`
}

var _ = Describe("check gomega parameter", func() {
	gom.Eventually(func(g gom.Gomega) error {
		arr := []int{1, 2, 3}
		g.Expect(len(arr)).Should(gom.Equal(3)) // want `ginkgo-linter: wrong length assertion; consider using .g\.Expect\(arr\)\.Should\(gom\.HaveLen\(3\)\). instead`
		return nil
	}).Should(gom.Succeed())
})
