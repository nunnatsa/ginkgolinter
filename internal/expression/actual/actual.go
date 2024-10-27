package actual

import (
	"go/ast"
	gotypes "go/types"

	"golang.org/x/tools/go/analysis"

	"github.com/nunnatsa/ginkgolinter/internal/gomegahandler"
)

const ( // gomega actual method names
	expect                 = "Expect"
	expectWithOffset       = "ExpectWithOffset"
	omega                  = "Ω"
	eventually             = "Eventually"
	eventuallyWithOffset   = "EventuallyWithOffset"
	consistently           = "Consistently"
	consistentlyWithOffset = "ConsistentlyWithOffset"
)

func IsActualMethod(name string) bool {
	_, found := funcOffsetMap[name]
	return found
}

type Actual struct {
	Orig         *ast.CallExpr
	Clone        *ast.CallExpr
	Arg          ArgPayload
	argType      gotypes.Type
	isAsync      bool
	asyncArg     *AsyncArg
	actualOffset int
}

func New(origExpr, cloneExpr *ast.CallExpr, orig *ast.CallExpr, clone *ast.CallExpr, pass *analysis.Pass, handler gomegahandler.Handler, timePkg string) (*Actual, bool) {
	funcName, ok := handler.GetActualFuncName(orig)
	if !ok {
		return nil, false
	}

	arg, actualOffset := getActualArgPayload(orig, clone, pass, funcName)
	if arg == nil {
		return nil, false
	}

	isAsyncExpr := isAsync(funcName)

	argType := pass.TypesInfo.TypeOf(orig.Args[actualOffset])

	if tpl, ok := argType.(*gotypes.Tuple); ok {
		if tpl.Len() > 0 {
			argType = tpl.At(0).Type()
		} else {
			argType = nil
		}
	}

	var asyncArg *AsyncArg
	if isAsyncExpr {
		asyncArg = newAsyncArg(origExpr, cloneExpr, orig, clone, argType, pass, actualOffset, timePkg)
	}

	return &Actual{
		Orig:         orig,
		Clone:        clone,
		Arg:          arg,
		argType:      argType,
		isAsync:      isAsyncExpr,
		asyncArg:     asyncArg,
		actualOffset: actualOffset,
	}, true
}

func NewNoAssertion(expr *ast.CallExpr, handler gomegahandler.Handler) *Actual {
	funcName, _ := handler.GetActualFuncName(expr)
	return &Actual{
		Arg: newNoAssertionActual(funcName),
	}
}

func (a *Actual) ReplaceActual(newArgs ast.Expr) {
	a.Clone.Args[a.actualOffset] = newArgs
}

func (a *Actual) ReplaceActualWithItsFirstArg() {
	firstArgOfArg := a.Clone.Args[a.actualOffset].(*ast.CallExpr).Args[0]
	a.ReplaceActual(firstArgOfArg)
}

func (a *Actual) IsAsync() bool {
	return a.isAsync
}

func (a *Actual) ArgGOType() gotypes.Type {
	return a.argType
}

func (a *Actual) GetAsyncArg() *AsyncArg {
	return a.asyncArg
}

func (a *Actual) AppendWithArgsMethod() {
	if a.asyncArg.fun != nil {
		if len(a.asyncArg.fun.Args) > 0 {
			actualOrigFunc := a.Clone.Fun
			actualOrigArgs := a.Clone.Args

			actualOrigArgs[a.actualOffset] = a.asyncArg.fun.Fun
			call := &ast.SelectorExpr{
				Sel: ast.NewIdent("WithArguments"),
				X: &ast.CallExpr{
					Fun:  actualOrigFunc,
					Args: actualOrigArgs,
				},
			}

			a.Clone.Fun = call
			a.Clone.Args = a.asyncArg.fun.Args
			a.Clone = a.Clone.Fun.(*ast.SelectorExpr).X.(*ast.CallExpr)
		} else {
			a.Clone.Args[a.actualOffset] = a.asyncArg.fun.Fun
		}
	}
}

func (a *Actual) GetActualArg() ast.Expr {
	return a.Clone.Args[a.actualOffset]
}

var asyncFuncSet = map[string]struct{}{
	eventually:             {},
	eventuallyWithOffset:   {},
	consistently:           {},
	consistentlyWithOffset: {},
}

func isAsync(name string) bool {
	_, ok := asyncFuncSet[name]
	return ok
}

var funcOffsetMap = map[string]int{
	expect:                 0,
	expectWithOffset:       1,
	omega:                  0,
	eventually:             0,
	eventuallyWithOffset:   1,
	consistently:           0,
	consistentlyWithOffset: 1,
}

func IsValidAsyncValueType(t gotypes.Type) bool {
	switch t.(type) {
	// allow functions that return function or channel.
	case *gotypes.Signature, *gotypes.Chan, *gotypes.Pointer:
		return true
	case *gotypes.Named:
		return IsValidAsyncValueType(t.Underlying())
	}

	return false
}
