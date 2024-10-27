package expression

import (
	"fmt"
	"github.com/nunnatsa/ginkgolinter/internal/gomegainfo"
	"go/ast"
	"go/token"

	"github.com/go-toolsmith/astcopy"
	"golang.org/x/tools/go/analysis"

	"github.com/nunnatsa/ginkgolinter/internal/expression/actual"
	"github.com/nunnatsa/ginkgolinter/internal/expression/matcher"
	"github.com/nunnatsa/ginkgolinter/internal/expression/value"
	"github.com/nunnatsa/ginkgolinter/internal/gomegahandler"
	"github.com/nunnatsa/ginkgolinter/internal/reverseassertion"
)

type GomegaExpression struct {
	Orig  *ast.CallExpr
	Clone *ast.CallExpr

	assertionFuncName     string
	origAssertionFuncName string
	actualFuncName        string

	isAsync bool

	Actual  *actual.Actual
	Matcher *matcher.Matcher

	handler gomegahandler.Handler
}

func New(origExpr *ast.CallExpr, pass *analysis.Pass, handler gomegahandler.Handler, timePkg string) (*GomegaExpression, bool) {
	actualMethodName, ok := handler.GetActualFuncName(origExpr)
	if !ok || !gomegainfo.IsActualMethod(actualMethodName) {
		return nil, false
	}

	origSel, ok := origExpr.Fun.(*ast.SelectorExpr)
	if !ok || !gomegainfo.IsAssertionFunc(origSel.Sel.Name) {
		return &GomegaExpression{
			Orig:           origExpr,
			actualFuncName: actualMethodName,
		}, true
	}

	exprClone := astcopy.CallExpr(origExpr)
	selClone := exprClone.Fun.(*ast.SelectorExpr)

	origActual := handler.GetActualExpr(origSel)
	if origActual == nil {
		return nil, false
	}

	actualClone := handler.GetActualExprClone(origSel, selClone)
	if actualClone == nil {
		return nil, false
	}

	actl, ok := actual.New(origExpr, exprClone, origActual, actualClone, pass, handler, timePkg)
	if !ok {
		return nil, false
	}

	origMatcher, ok := origExpr.Args[0].(*ast.CallExpr)
	if !ok {
		return nil, false
	}

	matcherClone := exprClone.Args[0].(*ast.CallExpr)

	mtchr, ok := matcher.New(origMatcher, matcherClone, pass, handler)
	if !ok {
		return nil, false
	}

	exprClone.Args[0] = mtchr.Clone

	gexp := &GomegaExpression{
		Orig:  origExpr,
		Clone: exprClone,

		assertionFuncName:     origSel.Sel.Name,
		origAssertionFuncName: origSel.Sel.Name,
		actualFuncName:        actualMethodName,

		isAsync: actl.IsAsync(),

		Actual:  actl,
		Matcher: mtchr,

		handler: handler,
	}

	if mtchr.ShouldReverseLogic() {
		gexp.ReverseAssertionFuncLogic()
	}

	return gexp, true
}

func (e *GomegaExpression) IsMissingAssertion() bool {
	return e.Matcher == nil
}

func (e *GomegaExpression) GetActualFuncName() string {
	if e == nil {
		return ""
	}
	return e.actualFuncName
}

func (e *GomegaExpression) GetAssertFuncName() string {
	if e == nil {
		return ""
	}
	return e.assertionFuncName
}

func (e *GomegaExpression) GetOrigAssertFuncName() string {
	if e == nil {
		return ""
	}
	return e.origAssertionFuncName
}

func (e *GomegaExpression) IsAsync() bool {
	return e.isAsync
}

func (e *GomegaExpression) ReverseAssertionFuncLogic() {
	assertionFunc := e.Clone.Fun.(*ast.SelectorExpr).Sel
	newName := reverseassertion.ChangeAssertionLogic(assertionFunc.Name)
	assertionFunc.Name = newName
	e.assertionFuncName = newName
}

func (e *GomegaExpression) ReplaceAssertionMethod(name string) {
	e.Clone.Fun.(*ast.SelectorExpr).Sel.Name = name
}

func (e *GomegaExpression) ReplaceMatcherFuncName(name string) {
	e.Matcher.ReplaceMatcherFuncName(name)
}

func (e *GomegaExpression) ReplaceMatcherArgs(newArgs []ast.Expr) {
	e.Matcher.ReplaceMatcherArgs(newArgs)
}

func (e *GomegaExpression) RemoveMatcherArgs() {
	e.Matcher.ReplaceMatcherArgs(nil)
}

func (e *GomegaExpression) ReplaceActual(newArg ast.Expr) {
	e.Actual.ReplaceActual(newArg)
}

func (e *GomegaExpression) ReplaceActualWithItsFirstArg() {
	e.Actual.ReplaceActualWithItsFirstArg()
}

func (e *GomegaExpression) replaceMathcerFuncNoArgs(name string) {
	e.Matcher.ReplaceMatcherFuncName(name)
	e.RemoveMatcherArgs()
}

func (e *GomegaExpression) SetMatcherBeZero() {
	e.replaceMathcerFuncNoArgs("BeZero")
}

func (e *GomegaExpression) SetMatcherBeEmpty() {
	e.replaceMathcerFuncNoArgs("BeEmpty")
}

func (e *GomegaExpression) SetLenNumericMatcher() {
	if m, ok := e.Matcher.GetMatcherInfo().(value.Valuer); ok && m.IsValueZero() {
		e.SetMatcherBeEmpty()
	} else {
		e.ReplaceMatcherFuncName("HaveLen")
		e.ReplaceMatcherArgs([]ast.Expr{m.GetValueExpr()})
	}
}

func (e *GomegaExpression) SetLenNumericActual() {
	if m, ok := e.Matcher.GetMatcherInfo().(value.Valuer); ok && m.IsValueZero() {
		e.SetMatcherBeEmpty()
	} else {
		e.ReplaceMatcherFuncName("HaveLen")
		e.ReplaceMatcherArgs([]ast.Expr{m.GetValueExpr()})
	}
}

func (e *GomegaExpression) SetMatcherLen(arg ast.Expr) {
	e.ReplaceMatcherFuncName("HaveLen")
	e.ReplaceMatcherArgs([]ast.Expr{arg})
}

func (e *GomegaExpression) SetMatcherCap(arg ast.Expr) {
	e.ReplaceMatcherFuncName("HaveCap")
	e.ReplaceMatcherArgs([]ast.Expr{arg})
}

func (e *GomegaExpression) SetMatcherCapZero() {
	e.ReplaceMatcherFuncName("HaveCap")
	e.ReplaceMatcherArgs([]ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "0"}})
}

func (e *GomegaExpression) SetMatcherSucceed() {
	e.replaceMathcerFuncNoArgs("Succeed")
}

func (e *GomegaExpression) SetMatcherHaveOccurred() {
	e.replaceMathcerFuncNoArgs("HaveOccurred")
}

func (e *GomegaExpression) SetMatcherBeNil() {
	e.replaceMathcerFuncNoArgs("BeNil")
}

func (e *GomegaExpression) SetMatcherBeTrue() {
	e.replaceMathcerFuncNoArgs("BeTrue")
}

func (e *GomegaExpression) SetMatcherBeFalse() {
	e.replaceMathcerFuncNoArgs("BeFalse")
}

func (e *GomegaExpression) SetMatcherHaveValue() {
	newMatcherExp := e.handler.GetNewWrapperMatcher("HaveValue", e.Matcher.Clone)
	e.Clone.Args[0] = newMatcherExp
	e.Matcher.Clone = newMatcherExp
}

func (e *GomegaExpression) SetMatcherEqual(arg ast.Expr) {
	e.ReplaceMatcherFuncName("Equal")
	e.ReplaceMatcherArgs([]ast.Expr{arg})
}

func (e *GomegaExpression) SetMatcherBeIdenticalTo(arg ast.Expr) {
	e.ReplaceMatcherFuncName("BeIdenticalTo")
	e.ReplaceMatcherArgs([]ast.Expr{arg})
}

func (e *GomegaExpression) SetMatcherBeNumerically(op token.Token, arg ast.Expr) {
	e.ReplaceMatcherFuncName("BeNumerically")
	e.ReplaceMatcherArgs([]ast.Expr{
		&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", op.String())},
		arg,
	})
}

func (e *GomegaExpression) IsNegativeAssertion() bool {
	return reverseassertion.IsNegativeLogic(e.assertionFuncName)
}
