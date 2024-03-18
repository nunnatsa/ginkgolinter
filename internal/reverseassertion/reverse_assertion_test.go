package reverseassertion_test

import (
	"go/token"
	"testing"

	"github.com/nunnatsa/ginkgolinter/internal/reverseassertion"
)

func TestChangeAssertionLogic(t *testing.T) {
	for _, assertionFunc := range []string{"To", "ToNot", "Should", "ShouldNot"} {
		rev := reverseassertion.ChangeAssertionLogic(assertionFunc)
		if rev == assertionFunc {
			t.Errorf("reversed function %s should be different than the original one %s", rev, assertionFunc)
		}

		revRev := reverseassertion.ChangeAssertionLogic(rev)
		if rev == revRev {
			t.Errorf("reversed function %s should be different than the double reversed one %s", rev, revRev)
		}

		if assertionFunc != revRev {
			t.Errorf("original function %s should be the same as the double reversed one %s", assertionFunc, revRev)
		}
	}
}

func TestChangeAssertionLogicWithNotTo(t *testing.T) {
	rev := reverseassertion.ChangeAssertionLogic("NotTo")
	if rev != "To" {
		t.Errorf("reverced function of NotTo should be NotTo, but it's %s", rev)
	}

	revRev := reverseassertion.ChangeAssertionLogic(rev)

	if revRev != "ToNot" {
		t.Errorf("reverced function of To should be ToNot, but it's %s", revRev)
	}
}

func TestChangeAssertionLogicWithUnknown(t *testing.T) {
	rev := reverseassertion.ChangeAssertionLogic("unknown")
	if rev != "unknown" {
		t.Errorf("reverced function of unknown should be the same, but it's %s", rev)
	}
}

func TestChangeCompareOperator(t *testing.T) {
	for _, tc := range []struct {
		testName string
		checked  token.Token
		expected token.Token
	}{
		{"check less than", token.LSS, token.GTR},
		{"check greater than", token.GTR, token.LSS},
		{"check less than or equal to", token.LEQ, token.GEQ},
		{"check greater than or equal to", token.GEQ, token.LEQ},
		{"check non-supported token", token.ILLEGAL, token.ILLEGAL},
	} {
		t.Run(tc.testName, func(tt *testing.T) {
			if reverseassertion.ChangeCompareOperator(tc.checked) != tc.expected {
				tt.Errorf("expected %v to become %v", tc.checked, tc.expected)
			}
		})
	}
}

func TestIsNegativeLogic(t *testing.T) {
	for _, funcName := range []string{"ToNot", "NotTo", "ShouldNot"} {
		t.Run("check "+funcName, func(tt *testing.T) {
			if !reverseassertion.IsNegativeLogic(funcName) {
				tt.Errorf("%s should be a negative function", funcName)
			}
		})
	}

	for _, methodName := range []string{"To", "Should", "Unsupported"} {
		t.Run("check "+methodName, func(tt *testing.T) {
			if reverseassertion.IsNegativeLogic(methodName) {
				tt.Errorf("%s should be a positive method", methodName)
			}
		})
	}
}
