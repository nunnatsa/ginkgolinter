package matcher

import (
	"go/ast"
	"go/constant"
	"go/token"
	gotypes "go/types"

	"golang.org/x/tools/go/analysis"

	"github.com/nunnatsa/ginkgolinter/internal/expression/value"
	"github.com/nunnatsa/ginkgolinter/internal/gomegahandler"
	"github.com/nunnatsa/ginkgolinter/internal/interfaces"
)

type Type uint64

const (
	Unspecified Type = 1 << iota
	EqualMatcherType
	BeZeroMatcherType
	BeEmptyMatcherType
	BeTrueMatcherType
	BeFalseMatcherType
	BeNumericallyMatcherType
	HaveLenZeroMatcherType
	BeEquivalentToMatcherType
	BeIdenticalToMatcherType
	BeNilMatcherType
	MatchErrorMatcherType
	MultipleMatcherMatherType
	HaveValueMatherType
	WithTransformMatherType
	EqualBoolValueMatcherType
	EqualValueMatcherType
	HaveOccurredMatcherType
	SucceedMatcherType
	EqualNilMatcherType

	BoolValueFalse
	BoolValueTrue

	OrMatherType
	AndMatherType

	ErrMatchWithErr
	ErrMatchWithErrFunc
	ErrMatchWithString
	ErrMatchWithMatcher

	EqualZero
	GreaterThanZero
)

func (t Type) Is(other Type) bool {
	return t&other != 0
}

type Info interface {
	Type() Type
	MatcherName() string
}

func getMatcherInfo(orig, clone *ast.CallExpr, matcherName string, pass *analysis.Pass, handler gomegahandler.Handler) Info {
	switch matcherName {
	case equal:
		return newEqualMatcher(orig.Args[0], clone.Args[0], pass)

	case beZero:
		return &BeZeroMatcher{}

	case beEmpty:
		return &BeEmptyMatcher{}

	case beTrue:
		return &BeTrueMatcher{}

	case beFalse:
		return &BeFalseMatcher{}

	case beNil:
		return &BeNilMatcher{}

	case beNumerically:
		if len(orig.Args) == 2 {
			return newBeNumericallyMatcher(orig.Args[0], orig.Args[1], clone.Args[1], pass)
		}

	case haveLen:
		if value.GetValuer(orig.Args[0], clone.Args[0], pass).IsValueZero() {
			return &HaveLenZeroMatcher{}
		}

	case beEquivalentTo:
		return &BeEquivalentToMatcher{
			Value: value.New(orig.Args[0], clone.Args[0], pass),
		}

	case beIdenticalTo:
		return &BeIdenticalToMatcher{
			Value: value.New(orig.Args[0], clone.Args[0], pass),
		}

	case matchError:
		return newMatchErrorMatcher(orig.Args, pass)

	case haveValue:
		if nestedMatcher, ok := getNestedMatcher(orig, clone, 0, pass, handler); ok {
			return &HaveValueMatcher{
				nested: nestedMatcher,
			}
		}

	case withTransform:
		if nestedMatcher, ok := getNestedMatcher(orig, clone, 1, pass, handler); ok {
			return newWithTransformMatcher(orig.Args[0], nestedMatcher, pass)
		}

	case or, and:
		matcherType := MultipleMatcherMatherType
		if matcherName == or {
			matcherType |= OrMatherType
		} else {
			matcherType |= AndMatherType
		}

		if m, ok := newMultipleMatchersMatcher(matcherType, orig.Args, clone.Args, pass, handler); ok {
			return m
		}

	case succeed:
		return &SucceedMatcher{}

	case haveOccurred:
		return &HaveOccurredMatcher{}

	}

	return &UnspecifiedMatcher{matcherName: matcherName}
}

func getNestedMatcher(orig, clone *ast.CallExpr, offset int, pass *analysis.Pass, handler gomegahandler.Handler) (*Matcher, bool) {
	if origNested, ok := orig.Args[offset].(*ast.CallExpr); ok {
		cloneNested := clone.Args[offset].(*ast.CallExpr)

		return New(origNested, cloneNested, pass, handler)
	}

	return nil, false
}

func newMatchErrorMatcher(args []ast.Expr, pass *analysis.Pass) MatchErrorMatcher {
	numArgs := len(args)
	if value.IsExprError(pass, args[0]) {
		return &MatchErrorMatcherWithErr{numArgs: numArgs}
	}

	t := pass.TypesInfo.TypeOf(args[0])
	if IsString(args[0], pass) {
		return &MatchErrorMatcherWithString{numArgs: numArgs}
	}

	if interfaces.ImplementsGomegaMatcher(t) {
		return &MatchErrorMatcherWithMatcher{numArgs: numArgs}
	}

	if isFuncErrBool(t) {
		isString := false
		if numArgs > 1 {
			t2 := pass.TypesInfo.TypeOf(args[1])
			isString = gotypes.Identical(t2, gotypes.Typ[gotypes.String])
		}
		return &MatchErrorMatcherWithErrFunc{numArgs: numArgs, secondArgIsString: isString}
	}

	return &InvalidMatchErrorMatcher{numArgs: numArgs}
}

type UnspecifiedMatcher struct {
	matcherName string
}

func (UnspecifiedMatcher) Type() Type {
	return Unspecified
}

func (u UnspecifiedMatcher) MatcherName() string {
	return u.matcherName
}

func newEqualMatcher(orig, clone ast.Expr, pass *analysis.Pass) Info {
	t := pass.TypesInfo.Types[orig]

	if t.Value != nil {
		if t.Value.Kind() == constant.Bool {
			if t.Value.String() == "true" {
				return &EqualTrueMatcher{}
			}
			return &EqualFalseMatcher{}
		}
	}

	if value.IsNil(orig, pass) {
		return &EqualNilMatcher{
			gotype: pass.TypesInfo.TypeOf(orig),
		}
	}

	val := value.GetValuer(orig, clone, pass)

	return &EqualMatcher{
		val: val,
	}
}

type EqualMatcher struct {
	val value.Valuer
}

func (EqualMatcher) Type() Type {
	return EqualMatcherType
}

func (EqualMatcher) MatcherName() string {
	return equal
}

func (m EqualMatcher) GetValue() constant.Value {
	return m.val.GetValue()
}

func (m EqualMatcher) GetType() gotypes.Type {
	return m.val.GetType()
}

func (m EqualMatcher) GetValueExpr() ast.Expr {
	return m.val.GetValueExpr()
}

func (m EqualMatcher) IsValueZero() bool {
	return m.val.IsValueZero()
}

func (m EqualMatcher) IsValueInt() bool {
	return m.val.IsValueInt()
}

func (m EqualMatcher) IsValueNumeric() bool {
	return m.val.IsValueNumeric()
}

func (m EqualMatcher) IsError() bool {
	return m.val.IsError()
}

func (m EqualMatcher) IsFunc() bool {
	return m.val.IsFunc()
}

func (m EqualMatcher) IsInterface() bool {
	return m.val.IsInterface()
}

func (m EqualMatcher) IsPointer() bool {
	return m.val.IsPointer()
}

type BeIdenticalToMatcher struct {
	value.Value
}

func (BeIdenticalToMatcher) Type() Type {
	return BeIdenticalToMatcherType
}

func (BeIdenticalToMatcher) MatcherName() string {
	return beIdenticalTo
}

type BeEquivalentToMatcher struct {
	value.Value
}

func (BeEquivalentToMatcher) Type() Type {
	return BeEquivalentToMatcherType
}

func (BeEquivalentToMatcher) MatcherName() string {
	return beEquivalentTo
}

type EqualNilMatcher struct {
	gotype gotypes.Type
}

func (EqualNilMatcher) Type() Type {
	return EqualNilMatcherType | EqualMatcherType | EqualValueMatcherType
}

func (EqualNilMatcher) MatcherName() string {
	return equal
}

func (n EqualNilMatcher) GetType() gotypes.Type {
	return n.gotype
}

type EqualTrueMatcher struct{}

func (EqualTrueMatcher) Type() Type {
	return EqualMatcherType | EqualBoolValueMatcherType | BoolValueTrue
}

func (EqualTrueMatcher) MatcherName() string {
	return equal
}

type EqualFalseMatcher struct{}

func (EqualFalseMatcher) Type() Type {
	return EqualMatcherType | EqualBoolValueMatcherType | BoolValueFalse
}

func (EqualFalseMatcher) MatcherName() string {
	return equal
}

type BeZeroMatcher struct{}

func (BeZeroMatcher) Type() Type {
	return BeZeroMatcherType
}

func (BeZeroMatcher) MatcherName() string {
	return beZero
}

type BeEmptyMatcher struct{}

func (BeEmptyMatcher) Type() Type {
	return BeEmptyMatcherType
}

func (BeEmptyMatcher) MatcherName() string {
	return beEmpty
}

type BeTrueMatcher struct{}

func (BeTrueMatcher) Type() Type {
	return BeTrueMatcherType | BoolValueTrue
}

func (BeTrueMatcher) MatcherName() string {
	return beTrue
}

type BeFalseMatcher struct{}

func (BeFalseMatcher) Type() Type {
	return BeFalseMatcherType | BoolValueFalse
}

func (BeFalseMatcher) MatcherName() string {
	return beFalse
}

type HaveLenZeroMatcher struct{}

func (HaveLenZeroMatcher) Type() Type {
	return HaveLenZeroMatcherType
}

func (HaveLenZeroMatcher) MatcherName() string {
	return haveLen
}

type BeNilMatcher struct{}

func (BeNilMatcher) Type() Type {
	return BeNilMatcherType
}

func (BeNilMatcher) MatcherName() string {
	return beNil
}

type BeNumericallyMatcher struct {
	op      token.Token
	value   value.Valuer
	argType Type
}

var compareOps = map[string]token.Token{
	`"=="`: token.EQL,
	`"<"`:  token.LSS,
	`">"`:  token.GTR,
	`"="`:  token.ASSIGN,
	`"!="`: token.NEQ,
	`"<="`: token.LEQ,
	`">="`: token.GEQ,
}

func getCompareOp(opExp ast.Expr) token.Token {
	basic, ok := opExp.(*ast.BasicLit)
	if !ok {
		return token.ILLEGAL
	}
	if basic.Kind != token.STRING {
		return token.ILLEGAL
	}

	if tk, ok := compareOps[basic.Value]; ok {
		return tk
	}

	return token.ILLEGAL
}

func newBeNumericallyMatcher(opExp, orig, clone ast.Expr, pass *analysis.Pass) Info {
	op := getCompareOp(opExp)
	if op == token.ILLEGAL {
		return &UnspecifiedMatcher{
			matcherName: beNumerically,
		}
	}

	val := value.GetValuer(orig, clone, pass)
	argType := BeNumericallyMatcherType

	if val.IsValueNumeric() {
		if v := val.GetValue().String(); v == "0" {
			switch op {
			case token.EQL:
				argType |= EqualZero

			case token.NEQ, token.GTR:
				argType |= GreaterThanZero
			}
		} else if v == "1" && op == token.GEQ {
			argType |= GreaterThanZero
		}
	}

	return &BeNumericallyMatcher{
		op:      op,
		value:   val,
		argType: argType,
	}
}

func (m BeNumericallyMatcher) Type() Type {
	return m.argType
}

func (BeNumericallyMatcher) MatcherName() string {
	return beNumerically
}

func (m BeNumericallyMatcher) GetValueExpr() ast.Expr {
	return m.value.GetValueExpr()
}

func (m BeNumericallyMatcher) GetValue() constant.Value {
	return m.value.GetValue()
}

func (m BeNumericallyMatcher) GetType() gotypes.Type {
	return m.value.GetType()
}

func (m BeNumericallyMatcher) GetOp() token.Token {
	return m.op
}

func (m BeNumericallyMatcher) IsValueZero() bool {
	return m.value.IsValueZero()
}

func (m BeNumericallyMatcher) IsValueInt() bool {
	return m.value.IsValueInt()
}

func (m BeNumericallyMatcher) IsValueNumeric() bool {
	return m.value.IsValueNumeric()
}

func (m BeNumericallyMatcher) IsError() bool {
	return m.value.IsError()
}

func (m BeNumericallyMatcher) IsFunc() bool {
	return m.value.IsFunc()
}

func (m BeNumericallyMatcher) IsInterface() bool {
	return m.value.IsInterface()
}

func (m BeNumericallyMatcher) IsPointer() bool {
	return m.value.IsPointer()
}

type MultipleMatchersMatcher struct {
	matherType Type
	matchers   []*Matcher
}

func (m *MultipleMatchersMatcher) Type() Type {
	return m.matherType
}

func (m *MultipleMatchersMatcher) MatcherName() string {
	if m.matherType.Is(OrMatherType) {
		return or
	}
	return and
}

func newMultipleMatchersMatcher(matherType Type, orig, clone []ast.Expr, pass *analysis.Pass, handler gomegahandler.Handler) (*MultipleMatchersMatcher, bool) {
	matchers := make([]*Matcher, len(orig))

	for i := range orig {
		nestedOrig, ok := orig[i].(*ast.CallExpr)
		if !ok {
			return nil, false
		}

		m, ok := New(nestedOrig, clone[i].(*ast.CallExpr), pass, handler)
		if !ok {
			return nil, false
		}

		m.reverseLogic = false

		matchers[i] = m
	}

	return &MultipleMatchersMatcher{
		matherType: matherType,
		matchers:   matchers,
	}, true
}

func (m *MultipleMatchersMatcher) Len() int {
	return len(m.matchers)
}

func (m *MultipleMatchersMatcher) At(i int) *Matcher {
	if i >= len(m.matchers) {
		panic("index out of range")
	}

	return m.matchers[i]
}

type MatchErrorMatcher interface {
	Info
	AllowedNumArgs() int
	NumArgs() int
}

type InvalidMatchErrorMatcher struct {
	firstAgr ast.Expr
	numArgs  int
}

func (m *InvalidMatchErrorMatcher) Type() Type {
	return MatchErrorMatcherType
}

func (m *InvalidMatchErrorMatcher) MatcherName() string {
	return matchError
}

func (m *InvalidMatchErrorMatcher) AllowedNumArgs() int {
	return 1
}

func (m *InvalidMatchErrorMatcher) NumArgs() int {
	return m.numArgs
}

func (m *InvalidMatchErrorMatcher) GetValueExpr() ast.Expr {
	return m.firstAgr
}

type MatchErrorMatcherWithErr struct {
	numArgs int
}

func (m *MatchErrorMatcherWithErr) Type() Type {
	return MatchErrorMatcherType | ErrMatchWithErr
}

func (m *MatchErrorMatcherWithErr) MatcherName() string {
	return matchError
}

func (m *MatchErrorMatcherWithErr) AllowedNumArgs() int {
	return 1
}

func (m *MatchErrorMatcherWithErr) NumArgs() int {
	return m.numArgs
}

type MatchErrorMatcherWithErrFunc struct {
	numArgs           int
	secondArgIsString bool
}

func (m *MatchErrorMatcherWithErrFunc) Type() Type {
	return MatchErrorMatcherType | ErrMatchWithErrFunc
}

func (m *MatchErrorMatcherWithErrFunc) MatcherName() string {
	return matchError
}

func (m *MatchErrorMatcherWithErrFunc) AllowedNumArgs() int {
	return 2
}

func (m *MatchErrorMatcherWithErrFunc) NumArgs() int {
	return m.numArgs
}

func (m *MatchErrorMatcherWithErrFunc) IsSecondArgString() bool {
	return m.secondArgIsString
}

type MatchErrorMatcherWithString struct {
	numArgs int
}

func (m *MatchErrorMatcherWithString) Type() Type {
	return MatchErrorMatcherType | ErrMatchWithString
}

func (m *MatchErrorMatcherWithString) MatcherName() string {
	return matchError
}

func (m *MatchErrorMatcherWithString) AllowedNumArgs() int {
	return 1
}

func (m *MatchErrorMatcherWithString) NumArgs() int {
	return m.numArgs
}

type MatchErrorMatcherWithMatcher struct {
	numArgs int
}

func (m *MatchErrorMatcherWithMatcher) Type() Type {
	return MatchErrorMatcherType | ErrMatchWithMatcher
}

func (m *MatchErrorMatcherWithMatcher) MatcherName() string {
	return matchError
}

func (m *MatchErrorMatcherWithMatcher) AllowedNumArgs() int {
	return 1
}

func (m *MatchErrorMatcherWithMatcher) NumArgs() int {
	return m.numArgs
}

type HaveValueMatcher struct {
	nested *Matcher
}

func (m *HaveValueMatcher) Type() Type {
	return HaveValueMatherType
}
func (m *HaveValueMatcher) MatcherName() string {
	return haveValue
}

func (m *HaveValueMatcher) GetNested() *Matcher {
	return m.nested
}

type WithTransformMatcher struct {
	funcType gotypes.Type
	nested   *Matcher
}

func newWithTransformMatcher(fun ast.Expr, nested *Matcher, pass *analysis.Pass) *WithTransformMatcher {
	funcType := pass.TypesInfo.TypeOf(fun)
	if sig, ok := funcType.(*gotypes.Signature); ok && sig.Results().Len() > 0 {
		funcType = sig.Results().At(0).Type()
	}
	return &WithTransformMatcher{
		funcType: funcType,
		nested:   nested,
	}
}

func (m *WithTransformMatcher) Type() Type {
	return WithTransformMatherType
}
func (m *WithTransformMatcher) MatcherName() string {
	return withTransform
}

func (m *WithTransformMatcher) GetNested() *Matcher {
	return m.nested
}

func (m *WithTransformMatcher) GetFuncType() gotypes.Type {
	return m.funcType
}

// isFuncErrBool checks if a function is with the signature `func(error) bool`
func isFuncErrBool(t gotypes.Type) bool {
	sig, ok := t.(*gotypes.Signature)
	if !ok {
		return false
	}
	if sig.Params().Len() != 1 || sig.Results().Len() != 1 {
		return false
	}

	if !interfaces.ImplementsError(sig.Params().At(0).Type()) {
		return false
	}

	b, ok := sig.Results().At(0).Type().(*gotypes.Basic)
	if ok && b.Name() == "bool" && b.Info() == gotypes.IsBoolean && b.Kind() == gotypes.Bool {
		return true
	}

	return false
}

func IsString(exp ast.Expr, pass *analysis.Pass) bool {
	t := pass.TypesInfo.TypeOf(exp)
	return gotypes.Identical(t, gotypes.Typ[gotypes.String])
}

type HaveOccurredMatcher struct{}

func (m *HaveOccurredMatcher) Type() Type {
	return HaveOccurredMatcherType
}
func (m *HaveOccurredMatcher) MatcherName() string {
	return haveOccurred
}

type SucceedMatcher struct{}

func (m *SucceedMatcher) Type() Type {
	return SucceedMatcherType
}
func (m *SucceedMatcher) MatcherName() string {
	return succeed
}
