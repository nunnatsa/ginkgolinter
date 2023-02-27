package ginkgolinter_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/nunnatsa/ginkgolinter"
)

func TestAllUseCases(t *testing.T) {
	for _, tc := range []struct {
		testName string
		testData string
	}{
		{
			testName: "find wrong length assertion",
			testData: "a/len",
		},
		{
			testName: "find HaveLen(0)",
			testData: "a/havelen0",
		},
		{
			testName: "find wrong nil assertion",
			testData: "a/nil",
		},
		{
			testName: "find Equal(nil)",
			testData: "a/equalnil",
		},
		{
			testName: "find Equal() with a boolean value",
			testData: "a/boolean",
		},
		{
			testName: "check suppress comments",
			testData: "a/suppress",
		},
		{
			testName: "check no gomega import",
			testData: "a/noginkgo",
		},
		{
			testName: "check no dot-import",
			testData: "a/nodotimport",
		},
		{
			testName: "find assertions that compare errors to nil",
			testData: "a/errnil",
		},
		{
			testName: "check gomaga without ginkgo",
			testData: "a/gomegaonly",
		},
		{
			testName: "comparison",
			testData: "a/comparison",
		},
		{
			testName: "pointers",
			testData: "a/pointers",
		},
	} {
		t.Run(tc.testName, func(tt *testing.T) {
			analysistest.Run(tt, analysistest.TestData(), ginkgolinter.NewAnalyzer(), tc.testData)
		})
	}
}

func TestFlags(t *testing.T) {
	for _, tc := range []struct {
		testName string
		testData []string
		flags    []string
	}{
		{
			testName: "test the suppress-len-assertion flag",
			testData: []string{"a/configlen"},
			flags:    []string{"suppress-len-assertion"},
		},
		{
			testName: "test the suppress-nil-assertion flag",
			testData: []string{"a/confignil"},
			flags:    []string{"suppress-nil-assertion"},
		},
		{
			testName: "test the suppress-err-assertion flag",
			testData: []string{"a/configerr"},
			flags:    []string{"suppress-err-assertion"},
		},
		{
			testName: "test the suppress-compare-assertion flag",
			testData: []string{"a/configcompare"},
			flags:    []string{"suppress-compare-assertion"},
		},
		{
			testName: "test the allow-havelen-0 flag",
			testData: []string{"a/havelen0config"},
			flags:    []string{"allow-havelen-0"},
		},
		{
			testName: "supress all",
			testData: []string{"a/configlen", "a/confignil", "a/configerr", "a/havelen0config"},
			flags:    []string{"suppress-len-assertion", "suppress-nil-assertion", "suppress-err-assertion", "suppress-compare-assertion"},
		},
	} {
		t.Run(tc.testName, func(tt *testing.T) {
			analyzer := ginkgolinter.NewAnalyzer()
			for _, flag := range tc.flags {
				err := analyzer.Flags.Set(flag, "true")
				if err != nil {
					tt.Errorf(`failed to set the "%s" flag; %v`, flag, err)
					return
				}
			}
			analysistest.Run(tt, analysistest.TestData(), analyzer, tc.testData...)
		})
	}
}
