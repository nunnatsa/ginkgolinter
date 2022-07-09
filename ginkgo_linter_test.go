package ginkgolinter

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestGinkgoLinter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo Linter Suite")
}

//func TestGinkgoLinter(t *testing.T) {
//	testdata := analysistest.TestData()
//
//	os.Setenv("GOPATH", "/home/nunnatsa/go")
//	analysistest.Run(t, testdata, Analyzer, "a")
//}

var _ = Describe("", func() {
	Context("test a", func() {
		It("Test all", func() {
			testdata := analysistest.TestData()

			analysistest.Run(GinkgoT(), testdata, Analyzer, "a")
		})
	})
})
