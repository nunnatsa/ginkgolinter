package reverseassertion_test

import (
	"testing"

	"github.com/nunnatsa/ginkgolinter/reverseassertion"
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
