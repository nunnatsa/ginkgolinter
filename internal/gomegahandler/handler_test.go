//nolint:staticcheck
package gomegahandler

import (
	"go/ast"
	gotypes "go/types"
	"testing"

	"golang.org/x/tools/go/analysis"
)

func TestGetGomegaHandler_dot(t *testing.T) {
	name := ast.NewIdent("test.go")
	file := &ast.File{
		Name: name,
		Imports: []*ast.ImportSpec{
			{
				Name: ast.NewIdent("."),
				Path: &ast.BasicLit{Value: `"github.com/onsi/gomega"`},
			},
		},
	}

	h := GetGomegaHandler(file, nil)
	if h == nil {
		t.Fatalf("should return dotHandler")
	}
	_, ok := h.(*dotHandler)
	if !ok {
		t.Error("should return dotHandler")
	}
}

func TestGetGomegaHandler_noname(t *testing.T) {
	name := ast.NewIdent("test.go")
	file := &ast.File{
		Name: name,
		Imports: []*ast.ImportSpec{
			{
				Path: &ast.BasicLit{Value: `"github.com/onsi/gomega"`},
			},
		},
	}

	h := GetGomegaHandler(file, nil)
	if h == nil {
		t.Fatalf("should return nameHandler")
	}
	n, ok := h.(*nameHandler)
	if !ok {
		t.Error("should return nameHandler")
	}

	if n.name != "gomega" {
		t.Errorf("import name should be `gomega`, but it's %s", n.name)
	}
}

func TestGetGomegaHandler_name(t *testing.T) {
	name := ast.NewIdent("test.go")
	file := &ast.File{
		Name: name,
		Imports: []*ast.ImportSpec{
			{
				Name: ast.NewIdent("name"),
				Path: &ast.BasicLit{Value: `"github.com/onsi/gomega"`},
			},
		},
	}

	h := GetGomegaHandler(file, nil)
	if h == nil {
		t.Fatalf("should return nameHandler")
	}
	n, ok := h.(*nameHandler)
	if !ok {
		t.Error("should return nameHandler")
	}

	if n.name != "name" {
		t.Errorf("import name should be `name`, but it's %s", n.name)
	}
}

func TestGetGomegaHandler_no_gomega(t *testing.T) {
	name := ast.NewIdent("test.go")
	file := &ast.File{
		Name: name,
		Imports: []*ast.ImportSpec{
			{
				Name: ast.NewIdent("."),
				Path: &ast.BasicLit{Value: `"github.com/onsi/ginkgo/v2"`},
			},
		},
	}

	h := GetGomegaHandler(file, nil)
	if h != nil {
		t.Fatalf("should return nil")
	}
}

const actualName = "actual"

var (
	gVarVar     = ast.NewIdent("g")
	gVarPointer = ast.NewIdent("g")
	noGomegaVar = ast.NewIdent("g")
)

func newGomegaPass() *analysis.Pass {
	return &analysis.Pass{
		TypesInfo: &gotypes.Info{
			Types: map[ast.Expr]gotypes.TypeAndValue{
				gVarVar: {
					Type: gotypes.NewNamed(gotypes.NewTypeName(0, gotypes.NewPackage(`github.com/onsi/gomega/internal`, ""), `Gomega`, &gotypes.Named{}), nil, nil),
				},
				gVarPointer: {
					Type: gotypes.NewPointer(gotypes.NewNamed(gotypes.NewTypeName(0, gotypes.NewPackage(`github.com/onsi/gomega/types`, ""), `Gomega`, &gotypes.Named{}), nil, nil)),
				},
				noGomegaVar: {
					Type: gotypes.NewPointer(gotypes.NewNamed(gotypes.NewTypeName(0, gotypes.NewPackage(`github.com/something/else`, ""), `somethingElse`, &gotypes.Named{}), nil, nil)),
				},
			},
		},
	}
}

func TestGomegaDotHandler_GetGomegaBasicInfo(t *testing.T) {
	h := &dotHandler{
		pass: newGomegaPass(),
	}

	for _, tc := range []struct {
		name              string
		exp               *ast.CallExpr
		expectedOK        bool
		expectedName      string
		expectedGomegaVar bool
	}{
		{
			name: "simple happy case",
			exp: &ast.CallExpr{
				Fun: ast.NewIdent(actualName),
			},
			expectedOK:        true,
			expectedName:      actualName,
			expectedGomegaVar: false,
		},
		{
			name: "non-ident func",
			exp: &ast.CallExpr{
				Fun: &ast.CallExpr{},
			},
			expectedOK:        false,
			expectedName:      "",
			expectedGomegaVar: false,
		},
		{
			name: "var happy case gomega var",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   gVarVar,
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:        true,
			expectedName:      actualName,
			expectedGomegaVar: true,
		},
		{
			name: "var happy case gomega pointer",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   gVarPointer,
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:        true,
			expectedName:      actualName,
			expectedGomegaVar: true,
		},
		{
			name: "non-gomega var",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   noGomegaVar,
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:        false,
			expectedName:      "",
			expectedGomegaVar: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			info, ok := h.GetGomegaBasicInfo(tc.exp)
			if ok != tc.expectedOK {
				t.Errorf(`expected ok = "%t" but it's "%t"'`, tc.expectedOK, ok)
			}

			if !ok {
				return
			}

			if info.MethodName != tc.expectedName {
				t.Errorf(`expected name = "%s" but it's "%s"'`, tc.expectedName, info.MethodName)
			}

			if info.UseGomegaVar != tc.expectedGomegaVar {
				t.Errorf(`expected gomega var to be %T, but it's %T'`, tc.expectedGomegaVar, info.UseGomegaVar)
			}
		})
	}
}

func TestGomegaNameHandler_GetGomegaBasicInfo(t *testing.T) {
	h := nameHandler{name: "gomega", pass: newGomegaPass()}

	for _, tc := range []struct {
		name              string
		exp               *ast.CallExpr
		expectedOK        bool
		expectedName      string
		expectedGomegaVar bool
	}{
		{
			name: "happy usecase",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("gomega"),
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:        true,
			expectedName:      actualName,
			expectedGomegaVar: false,
		},
		{
			name: "not a selector",
			exp: &ast.CallExpr{
				Fun: ast.NewIdent("name"),
			},
			expectedOK:        false,
			expectedName:      "",
			expectedGomegaVar: false,
		},
		{
			name: "CallExpr",
			exp: &ast.CallExpr{
				Fun: ast.NewIdent("name"),
			},
			expectedOK:        false,
			expectedName:      "",
			expectedGomegaVar: false,
		},
		{
			name: "no gomega",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("notgomega"),
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:        false,
			expectedName:      "",
			expectedGomegaVar: false,
		},
		{
			name: "gomega variable",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   gVarVar,
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:        true,
			expectedName:      actualName,
			expectedGomegaVar: true,
		},
		{
			name: "gomega pointer",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   gVarPointer,
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:        true,
			expectedName:      actualName,
			expectedGomegaVar: true,
		},
		{
			name: "gomega variable from non gomega function",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   noGomegaVar,
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:        false,
			expectedName:      "",
			expectedGomegaVar: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			info, ok := h.GetGomegaBasicInfo(tc.exp)
			if ok != tc.expectedOK {
				t.Errorf(`expected ok = "%t" but it's "%t"'`, tc.expectedOK, ok)
			}
			if !ok {
				return
			}
			if info.MethodName != tc.expectedName {
				t.Errorf(`expected name = "%s" but it's "%s"'`, tc.expectedName, info.MethodName)
			}
			if info.UseGomegaVar != tc.expectedGomegaVar {
				t.Errorf(`expected gomega var to be %T, but it's %T'`, tc.expectedGomegaVar, info.UseGomegaVar)
			}
		})
	}
}

func TestGomegaDotHandler_ReplaceFunction(t *testing.T) {
	h := dotHandler{}

	expr := &ast.CallExpr{
		Fun: ast.NewIdent("one"),
	}

	h.ReplaceFunction(expr, ast.NewIdent("two"))

	f, ok := expr.Fun.(*ast.Ident)
	if !ok {
		t.Error("should be ast.Ident")
	} else if f.Name != "two" {
		t.Error("the new function name should be 'two'")
	}
}

func TestGomegaNameHandler_ReplaceFunction(t *testing.T) {
	h := &nameHandler{name: "gomega"}

	expr := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			Sel: ast.NewIdent("one"),
		},
	}

	h.ReplaceFunction(expr, ast.NewIdent("two"))

	f, ok := expr.Fun.(*ast.SelectorExpr)
	if !ok {
		t.Error("should be ast.Ident")
	} else if f.Sel.Name != "two" {
		t.Error("the new function name should be 'two'")
	}
}
