package ginkgolinter_test

import (
	"golang.org/x/tools/go/analysis"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/nunnatsa/ginkgolinter"
)

func TestGinkgoLenLinter(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), ginkgolinter.NewAnalyzer(), "a/len")
}

func TestGinkgoNilLinter(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), ginkgolinter.NewAnalyzer(), "a/nil")
}

func TestSuppress(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), ginkgolinter.NewAnalyzer(), "a/suppress")
}

func TestFlags_suppress_len(t *testing.T) {
	analyzerFunc := func() *analysis.Analyzer {
		a := ginkgolinter.NewAnalyzer()
		err := a.Flags.Set("suppress-len-assertion", "true")
		if err != nil {
			t.Error(err)
		}
		return a
	}

	a := analyzerFunc()
	analysistest.Run(t, analysistest.TestData(), a, "a/configlen")
}

func TestFlags_suppress_nil(t *testing.T) {
	analyzerFunc := func() *analysis.Analyzer {
		a := ginkgolinter.NewAnalyzer()
		err := a.Flags.Set("suppress-nil-assertion", "true")
		if err != nil {
			t.Error(err)
		}
		return a
	}

	a := analyzerFunc()
	analysistest.Run(t, analysistest.TestData(), a, "a/confignil")
}

func TestFlags_suppress_all(t *testing.T) {
	analyzerFunc := func() *analysis.Analyzer {
		a := ginkgolinter.NewAnalyzer()

		err := a.Flags.Set("suppress-len-assertion", "true")
		if err != nil {
			t.Error(err)
		}
		err = a.Flags.Set("suppress-nil-assertion", "true")
		if err != nil {
			t.Error(err)
		}

		return a
	}

	a := analyzerFunc()
	analysistest.Run(t, analysistest.TestData(), a, "a/configlen", "a/confignil")
}
