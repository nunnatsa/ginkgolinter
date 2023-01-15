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

func TestGinkgoEqualBooleanLinter(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), ginkgolinter.NewAnalyzer(), "a/boolean")
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

func TestFlags_suppress_err(t *testing.T) {
	analyzerFunc := func() *analysis.Analyzer {
		a := ginkgolinter.NewAnalyzer()
		err := a.Flags.Set("suppress-err-assertion", "true")
		if err != nil {
			t.Error(err)
		}
		return a
	}

	a := analyzerFunc()
	analysistest.Run(t, analysistest.TestData(), a, "a/configerr")
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

		err = a.Flags.Set("suppress-err-assertion", "true")
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

func TestErrNil(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), ginkgolinter.NewAnalyzer(), "a/errnil")
}

func TestGomegaOnly(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), ginkgolinter.NewAnalyzer(), "a/gomegaonly")
}
