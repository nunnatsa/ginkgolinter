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
		{
			testName: "function call in Eventually",
			testData: "a/eventually",
		},
		{
			testName: "compare pointer to value",
			testData: "a/pointerval",
		},
		{
			testName: "no assertion",
			testData: "a/noassersion",
		},
		{
			testName: "focus",
			testData: "a/focus",
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
		flags    map[string]string
	}{
		{
			testName: "test the suppress-len-assertion flag",
			testData: []string{"a/configlen"},
			flags:    map[string]string{"suppress-len-assertion": "true"},
		},
		{
			testName: "test the suppress-nil-assertion flag",
			testData: []string{"a/confignil"},
			flags:    map[string]string{"suppress-nil-assertion": "true"},
		},
		{
			testName: "test the suppress-err-assertion flag",
			testData: []string{"a/configerr"},
			flags:    map[string]string{"suppress-err-assertion": "true"},
		},
		{
			testName: "test the suppress-compare-assertion flag",
			testData: []string{"a/configcompare"},
			flags:    map[string]string{"suppress-compare-assertion": "true"},
		},
		{
			testName: "test the allow-havelen-0 flag",
			testData: []string{"a/havelen0config"},
			flags:    map[string]string{"allow-havelen-0": "true"},
		},
		{
			testName: "test the suppress-async-assertion flag",
			testData: []string{"a/asyncconfig"},
			flags:    map[string]string{"suppress-async-assertion": "true"},
		},
		{
			testName: "test the forbid-focus-container flag",
			testData: []string{"a/focusconfig"},
			flags:    map[string]string{"forbid-focus-container": "true"},
		},
	} {
		t.Run(tc.testName, func(tt *testing.T) {
			analyzer := ginkgolinter.NewAnalyzer()
			for flag, value := range tc.flags {
				err := analyzer.Flags.Set(flag, value)
				if err != nil {
					tt.Errorf(`failed to set the "%s" flag; %v`, flag, err)
					return
				}
			}
			analysistest.Run(tt, analysistest.TestData(), analyzer, tc.testData...)
		})
	}
}
