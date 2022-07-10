package ginkgolinter_test

import (
	"os"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/nunnatsa/ginkgolinter"
)

func TestGinkgoLinter(t *testing.T) {
	testdata := analysistest.TestData()

	os.Setenv("GOPATH", "/home/nunnatsa/go")
	analysistest.Run(t, testdata, ginkgolinter.Analyzer, "a")
}
