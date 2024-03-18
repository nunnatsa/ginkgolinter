package gomegahandler

import (
	"go/ast"
	"go/token"
	"testing"
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

	h := GetGomegaHandler(file)
	if h == nil {
		t.Fatalf("should return dotHandler")
	}
	_, ok := h.(dotHandler)
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

	h := GetGomegaHandler(file)
	if h == nil {
		t.Fatalf("should return nameHandler")
	}
	n, ok := h.(nameHandler)
	if !ok {
		t.Error("should return nameHandler")
	}

	if string(n) != "gomega" {
		t.Errorf("import name should be `gomega`, but it's %s", string(n))
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

	h := GetGomegaHandler(file)
	if h == nil {
		t.Fatalf("should return nameHandler")
	}
	n, ok := h.(nameHandler)
	if !ok {
		t.Error("should return nameHandler")
	}

	if string(n) != "name" {
		t.Errorf("import name should be `name`, but it's %s", string(n))
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

	h := GetGomegaHandler(file)
	if h != nil {
		t.Fatalf("should return nil")
	}
}

const actualName = "actual"

func TestGomegaDotHandler_GetActualFuncName(t *testing.T) {
	h := dotHandler{}

	for _, tc := range []struct {
		name         string
		exp          *ast.CallExpr
		expectedOK   bool
		expectedName string
	}{
		{
			name: "simple happy case",
			exp: &ast.CallExpr{
				Fun: ast.NewIdent(actualName),
			},
			expectedOK:   true,
			expectedName: actualName,
		},
		{
			name: "non-ident func",
			exp: &ast.CallExpr{
				Fun: &ast.CallExpr{},
			},
			expectedOK:   false,
			expectedName: "",
		},
		{
			name:         "var happy case - NewGomega",
			exp:          getVar("NewGomega"),
			expectedOK:   true,
			expectedName: actualName,
		},
		{
			name:         "var happy case - NewWithT",
			exp:          getVar("NewWithT"),
			expectedOK:   true,
			expectedName: actualName,
		},
		{
			name:         "var happy case - NewGomegaWithT",
			exp:          getVar("NewGomegaWithT"),
			expectedOK:   true,
			expectedName: actualName,
		},
		{
			name:         "var error case - no gomega function",
			exp:          getVar("SomethingElse"),
			expectedOK:   false,
			expectedName: "",
		},
		{
			name:         "field happy case - Gomega",
			exp:          getField("Gomega"),
			expectedOK:   true,
			expectedName: actualName,
		},
		{
			name:         "field error case - non Gomega",
			exp:          getField("SomethingElse"),
			expectedOK:   false,
			expectedName: "",
		},
		{
			name:         "field happy case - star - WithT",
			exp:          getStarField("WithT"),
			expectedOK:   true,
			expectedName: actualName,
		},
		{
			name:         "field happy case - star - GomegaWithT",
			exp:          getStarField("GomegaWithT"),
			expectedOK:   true,
			expectedName: actualName,
		},
		{
			name:         "field error case - star - non Gomega type",
			exp:          getStarField("SomethingElse"),
			expectedOK:   false,
			expectedName: "",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			name, ok := h.GetActualFuncName(tc.exp)
			if ok != tc.expectedOK {
				t.Errorf(`expected ok = "%t" but it's "%t"'`, tc.expectedOK, ok)
			}
			if name != tc.expectedName {
				t.Errorf(`expected name = "%s" but it's "%s"'`, tc.expectedName, name)
			}
		})
	}
}

func TestGomegaNameHandler_GetActualFuncName(t *testing.T) {
	h := nameHandler("gomega")

	for _, tc := range []struct {
		name         string
		exp          *ast.CallExpr
		expectedOK   bool
		expectedName string
	}{
		{
			name: "happy usecase",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("gomega"),
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:   true,
			expectedName: actualName,
		},
		{
			name: "not a selector",
			exp: &ast.CallExpr{
				Fun: ast.NewIdent("name"),
			},
			expectedOK:   false,
			expectedName: "",
		},
		{
			name: "CallExpr",
			exp: &ast.CallExpr{
				Fun: ast.NewIdent("name"),
			},
			expectedOK:   false,
			expectedName: "",
		},
		{
			name: "no gomega",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("notgomega"),
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:   false,
			expectedName: "",
		},
		{
			name: "gomega variable from NewGomega",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "g",
						Obj:  getGomagaVarObject("gomega", "NewGomega"),
					},
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:   true,
			expectedName: actualName,
		},
		{
			name: "gomega variable from NewWithT",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "g",
						Obj:  getGomagaVarObject("gomega", "NewWithT"),
					},
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:   true,
			expectedName: actualName,
		},
		{
			name: "gomega variable from NewGomegaWithT",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "g",
						Obj:  getGomagaVarObject("gomega", "NewGomegaWithT"),
					},
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:   true,
			expectedName: actualName,
		},
		{
			name: "gomega variable from non gomega function",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "g",
						Obj:  getGomagaVarObject("gomega", "SomethingElse"),
					},
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:   false,
			expectedName: "",
		},
		{
			name: "gomega variable from non gomega package",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "g",
						Obj:  getGomagaVarObject("SomethingElse", "NewGomegaWithT"),
					},
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:   false,
			expectedName: "",
		},
		{
			name: "gomega field Gomega",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "g",
						Obj:  getGomagaVarObjectForSelectorField("gomega", "Gomega"),
					},
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:   true,
			expectedName: actualName,
		},
		{
			name: "gomega variable from *WithT",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "g",
						Obj:  getGomagaVarObjectForStarField("gomega", "WithT"),
					},
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:   true,
			expectedName: actualName,
		},
		{
			name: "gomega variable from *GomegaWithT",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "g",
						Obj:  getGomagaVarObjectForStarField("gomega", "GomegaWithT"),
					},
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:   true,
			expectedName: actualName,
		},
		{
			name: "gomega variable from non gomega field - selector",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "g",
						Obj:  getGomagaVarObjectForSelectorField("gomega", "SomethingElse"),
					},
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:   false,
			expectedName: "",
		},
		{
			name: "gomega variable from non gomega field - star",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "g",
						Obj:  getGomagaVarObjectForStarField("gomega", "SomethingElse"),
					},
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:   false,
			expectedName: "",
		},
		{
			name: "non-gomega variable from gomega field - selector",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "g",
						Obj:  getGomagaVarObjectForSelectorField("somethingElse", "Gomega"),
					},
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:   false,
			expectedName: "",
		},
		{
			name: "non gomega variable from non gomega field - star",
			exp: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: "g",
						Obj:  getGomagaVarObjectForStarField("somethingElse", "WithT"),
					},
					Sel: ast.NewIdent(actualName),
				},
			},
			expectedOK:   false,
			expectedName: "",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			name, ok := h.GetActualFuncName(tc.exp)
			if ok != tc.expectedOK {
				t.Errorf(`expected ok = "%t" but it's "%t"'`, tc.expectedOK, ok)
			}
			if name != tc.expectedName {
				t.Errorf(`expected name = "%s" but it's "%s"'`, tc.expectedName, name)
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
	h := nameHandler("gomega")

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

func TestGetGomegaHandler_getFieldType(t *testing.T) {
	for _, tc := range []struct {
		testName     string
		h            Handler
		field        *ast.Field
		expectedName string
	}{
		{
			testName: "dotHandler: Ident",
			h:        dotHandler{},
			field: &ast.Field{
				Type: ast.NewIdent("name"),
			},
			expectedName: "name",
		},
		{
			testName: "dotHandler: Star",
			h:        dotHandler{},
			field: &ast.Field{
				Type: &ast.StarExpr{X: ast.NewIdent("name")},
			},
			expectedName: "name",
		},
		{
			testName: "nameHandler: SelectorExpr: var name == handler name",
			h:        nameHandler("g"),
			field: &ast.Field{
				Type: &ast.SelectorExpr{
					X:   ast.NewIdent("g"),
					Sel: ast.NewIdent("name"),
				},
			},
			expectedName: "name",
		},
		{
			testName: "nameHandler: SelectorExpr: var name != handler name",
			h:        nameHandler("g"),
			field: &ast.Field{
				Type: &ast.SelectorExpr{
					X:   ast.NewIdent("not_g"),
					Sel: ast.NewIdent("name"),
				},
			},
			expectedName: "",
		},
		{
			testName: "nameHandler: Star: var name == handler name",
			h:        nameHandler("g"),
			field: &ast.Field{
				Type: &ast.StarExpr{
					X: &ast.SelectorExpr{
						X:   ast.NewIdent("g"),
						Sel: ast.NewIdent("name"),
					},
				},
			},
			expectedName: "name",
		},
		{
			testName: "nameHandler: Star: var name != handler name",
			h:        nameHandler("g"),
			field: &ast.Field{
				Type: &ast.StarExpr{
					X: &ast.SelectorExpr{
						X:   ast.NewIdent("name"),
						Sel: ast.NewIdent("not_g"),
					},
				},
			},
			expectedName: "",
		},
	} {
		t.Run(tc.testName, func(t *testing.T) {
			name := tc.h.getFieldType(tc.field)
			if name != tc.expectedName {
				t.Errorf(`should return "%s", but returned "%s"`, tc.expectedName, name)
			}
		})
	}
}

func getVar(funcName string) *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X: &ast.Ident{
				Name: "g",
				Obj: &ast.Object{
					Kind: ast.Var,
					Decl: &ast.AssignStmt{
						Tok: token.DEFINE,
						Rhs: []ast.Expr{
							&ast.CallExpr{
								Fun: ast.NewIdent(funcName),
							},
						},
					},
				},
			},
			Sel: ast.NewIdent(actualName),
		},
	}
}

func getStarField(typeName string) *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X: &ast.Ident{
				Name: "g",
				Obj: &ast.Object{
					Kind: ast.Var,
					Decl: &ast.Field{
						Type: &ast.StarExpr{
							X: ast.NewIdent(typeName),
						},
					},
				},
			},
			Sel: ast.NewIdent(actualName),
		},
	}
}

func getField(typeName string) *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X: &ast.Ident{
				Name: "g",
				Obj: &ast.Object{
					Kind: ast.Var,
					Decl: &ast.Field{
						Type: ast.NewIdent(typeName),
					},
				},
			},
			Sel: ast.NewIdent(actualName),
		},
	}
}

func getGomagaVarObject(pkg, funcName string) *ast.Object {
	return &ast.Object{
		Kind: ast.Var,
		Decl: &ast.AssignStmt{
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   ast.NewIdent(pkg),
						Sel: ast.NewIdent(funcName),
					},
				},
			},
		},
	}
}

func getGomagaVarObjectForSelectorField(pkg, typeName string) *ast.Object {
	return &ast.Object{
		Kind: ast.Var,
		Decl: &ast.Field{
			Type: &ast.SelectorExpr{
				X:   ast.NewIdent(pkg),
				Sel: ast.NewIdent(typeName),
			},
		},
	}
}

func getGomagaVarObjectForStarField(pkg, typeName string) *ast.Object {
	return &ast.Object{
		Kind: ast.Var,
		Decl: &ast.Field{
			Type: &ast.StarExpr{
				X: &ast.SelectorExpr{
					X:   ast.NewIdent(pkg),
					Sel: ast.NewIdent(typeName),
				},
			},
		},
	}
}
