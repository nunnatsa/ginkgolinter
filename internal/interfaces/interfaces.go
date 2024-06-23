package interfaces

import (
	"go/token"
	"go/types"
	gotypes "go/types"
	"strings"
)

var (
	errorType         *gotypes.Interface
	gomegaMatcherType *gotypes.Interface
)

func init() {
	errorType = gotypes.Universe.Lookup("error").Type().Underlying().(*gotypes.Interface)
	gomegaMatcherType = generateTheGomegaMatcherInfType()
}

// generateTheGomegaMatcherInfType generates a types.Interface instance that represents the
// GomegaMatcher interface.
// The original code is (copied from https://github.com/nunnatsa/ginkgolinter/blob/8fdd05eee922578d4699f49d267001c01e0b9f1e/testdata/src/a/vendor/github.com/onsi/gomega/types/types.go)
//
//	type GomegaMatcher interface {
//		Match(actual interface{}) (success bool, err error)
//		FailureMessage(actual interface{}) (message string)
//		NegatedFailureMessage(actual interface{}) (message string)
//	}
func generateTheGomegaMatcherInfType() *gotypes.Interface {
	err := gotypes.Universe.Lookup("error").Type()
	bl := gotypes.Typ[gotypes.Bool]
	str := gotypes.Typ[gotypes.String]
	anyType := gotypes.Universe.Lookup("any").Type()

	return gotypes.NewInterfaceType([]*gotypes.Func{
		// Match(actual interface{}) (success bool, err error)
		gotypes.NewFunc(token.NoPos, nil, "Match", gotypes.NewSignatureType(
			nil, nil, nil,
			gotypes.NewTuple(
				gotypes.NewVar(token.NoPos, nil, "actual", anyType),
			),
			gotypes.NewTuple(
				gotypes.NewVar(token.NoPos, nil, "", bl),
				gotypes.NewVar(token.NoPos, nil, "", err),
			), false),
		),
		// FailureMessage(actual interface{}) (message string)
		gotypes.NewFunc(token.NoPos, nil, "FailureMessage", gotypes.NewSignatureType(
			nil, nil, nil,
			gotypes.NewTuple(
				gotypes.NewVar(token.NoPos, nil, "", anyType),
			),
			gotypes.NewTuple(
				gotypes.NewVar(token.NoPos, nil, "", str),
			),
			false),
		),
		//NegatedFailureMessage(actual interface{}) (message string)
		gotypes.NewFunc(token.NoPos, nil, "NegatedFailureMessage", gotypes.NewSignatureType(
			nil, nil, nil,
			gotypes.NewTuple(
				gotypes.NewVar(token.NoPos, nil, "", anyType),
			),
			gotypes.NewTuple(
				gotypes.NewVar(token.NoPos, nil, "", str),
			),
			false),
		),
	}, nil)
}

func ImplementsError(t gotypes.Type) bool {
	return gotypes.Implements(t, errorType)
}

func ImplementsGomegaMatcher(t gotypes.Type) bool {
	return t != nil && gotypes.Implements(t, gomegaMatcherType)
}

// Note: we cannot check for an argument which _implements_ gomega.Gomega without adding a Go
// module dependency on gomega.  This is because the gomega.Gomega interface definition references
// gomega.AsyncAssertion, whose methods return gomega.AsyncAssertion. Because Go does not have
// interface covariance or contravariance, any "local copy" of gomega.AsyncAssertion cannot
// be satisified by any actual `gomega.AsyncAssertion` implementation, as the methods do not return
// local.AsyncAssertion but rather gomega.AsyncAssertion.
//
// Also, Gomega probably doesn't even accept an argument whose type implements the interface, but
// rather whose type _is_ the interface. So this check should suffice.
func IsGomega(t gotypes.Type) bool {
	named, isNamed := t.(*gotypes.Named)
	if !isNamed {
		return false
	}

	obj := named.Obj()

	if obj.Name() != "Gomega" {
		return false
	}

	return isPackageSymbol(named.Obj(), "github.com/onsi/gomega/types", "Gomega")
}

func isPackageSymbol(obj types.Object, pkgPath, name string) bool {
	if obj.Name() != name {
		return false
	}

	vendorPieces := strings.Split(pkgPath, "/vendor/")

	return pkgPath == vendorPieces[len(vendorPieces)-1]
}
