package ginkgolinter_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/nunnatsa/ginkgolinter"
)

func TestGinkgoLenLinter(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), ginkgolinter.Analyzer, "a/len")
}

func TestGinkgoNilLinter(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), ginkgolinter.Analyzer, "a/nil")
}

func TestSuppress(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), ginkgolinter.Analyzer, "a/suppress")
}
