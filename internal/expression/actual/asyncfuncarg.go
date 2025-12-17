package actual

import (
	gotypes "go/types"

	"github.com/nunnatsa/ginkgolinter/internal/gomegainfo"
	"github.com/nunnatsa/ginkgolinter/internal/typecheck"
)

func getAsyncFuncArg(sig *gotypes.Signature) ArgPayload {
	argType := FuncSigArgType
	if sig.Results().Len() == 1 {
		if typecheck.ImplementsError(sig.Results().At(0).Type().Underlying()) {
			argType |= ErrFuncActualArgType | ErrorTypeArgType
		}
	}

	if sig.Params().Len() > 0 {
		arg := sig.Params().At(0).Type()
		if sig.Results().Len() == 0 {
			if gomegainfo.IsGomegaType(arg) {
				argType |= FuncSigArgType | GomegaParamArgType
			}
			if implementsTB(arg) {
				argType |= FuncSigArgType | TBParamArgType
			}
		}
	}

	if sig.Results().Len() > 1 {
		argType |= FuncSigArgType | MultiRetsArgType
	}

	return &FuncSigArgPayload{argType: argType}
}

type FuncSigArgPayload struct {
	argType ArgType
}

func (f FuncSigArgPayload) ArgType() ArgType {
	return f.argType
}

// implementsTB checks if the argument type implements any of the methods in testing.TB which
// can be used to report test failures. Such a type is a potential alternative to a Gomega
// parameter in some Gomega wrappers.
func implementsTB(t gotypes.Type) bool {
	for _, interfaceType := range tbInterfaces {
		if gotypes.Implements(t, interfaceType) {
			return true
		}
	}
	return false
}

var tbInterfaces = []*gotypes.Interface{
	tbInterface("Error", false),
}

func tbInterface(name string, printf bool) *gotypes.Interface {
	var params []*gotypes.Var
	if printf {
		params = append(params, gotypes.NewVar(0, nil, "", gotypes.Typ[gotypes.String]))
	}
	params = append(params, gotypes.NewVar(0, nil, "", gotypes.NewSlice(gotypes.Universe.Lookup("any").Type())))
	signature := gotypes.NewSignatureType(nil, nil, nil,
		gotypes.NewTuple(params...),
		gotypes.NewTuple(),
		true,
	)
	method := gotypes.NewFunc(0, nil, name, signature)
	return gotypes.NewInterfaceType([]*gotypes.Func{method}, nil)
}
