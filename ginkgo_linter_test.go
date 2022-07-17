package ginkgolinter_test

import (
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/nunnatsa/ginkgolinter"
)

func TestGinkgoLenLinter(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), ginkgolinter.NewAnalyzer(), "a/len")
}

func TestGinkgoNilLinter(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), ginkgolinter.NewAnalyzer(), "a/nil")
}

func TestGinkgoEqualNilLinter(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), ginkgolinter.NewAnalyzer(), "a/equalnil")
}

func TestSuppress(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), ginkgolinter.NewAnalyzer(), "a/Suppress")
}

func TestFlags_suppress_len(t *testing.T) {
	analyzerFunc := func() *analysis.Analyzer {
		a := ginkgolinter.NewAnalyzer()
		err := a.Flags.Set("Suppress-len-assertion", "true")
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
		err := a.Flags.Set("Suppress-nil-assertion", "true")
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

		err := a.Flags.Set("Suppress-len-assertion", "true")
		if err != nil {
			t.Error(err)
		}
		err = a.Flags.Set("Suppress-nil-assertion", "true")
		if err != nil {
			t.Error(err)
		}

		return a
	}

	a := analyzerFunc()
	analysistest.Run(t, analysistest.TestData(), a, "a/configlen", "a/confignil")
}

func TestNoGinkgo(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), ginkgolinter.NewAnalyzer(), "a/noginkgo")
}

func TestNoDotImport(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), ginkgolinter.NewAnalyzer(), "a/nodotimport")
}
