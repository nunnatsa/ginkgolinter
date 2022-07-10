package ginkgolinter_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/nunnatsa/ginkgolinter"
)

func TestGinkgoLinter(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), ginkgolinter.Analyzer, "a")
}
