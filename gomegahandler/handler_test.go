package gomegahandler

import (
	"go/ast"
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

func TestGomegaDotHandler_GetActualFuncName(t *testing.T) {
	h := dotHandler{}

	expr := &ast.CallExpr{
		Fun: ast.NewIdent("actual"),
	}

	name, ok := h.GetActualFuncName(expr)
	if !ok {
		t.Error("Should return valid value")
	} else if name != "actual" {
		t.Error("should return 'actual'")
	}

	expr = &ast.CallExpr{
		Fun: &ast.CallExpr{},
	}

	name, ok = h.GetActualFuncName(expr)
	if ok {
		t.Error("Should return empty response")
	}

	if len(name) > 0 {
		t.Error("Should return empty response")
	}
}

func TestGomegaNameHandler_GetActualFuncName(t *testing.T) {
	h := nameHandler("gomega")

	expr := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   ast.NewIdent("gomega"),
			Sel: ast.NewIdent("actual"),
		},
	}

	name, ok := h.GetActualFuncName(expr)
	if !ok {
		t.Error("Should return valid value")
	} else if name != "actual" {
		t.Error("should return 'actual'")
	}

	expr = &ast.CallExpr{
		Fun: ast.NewIdent("name"),
	}

	name, ok = h.GetActualFuncName(expr)
	if ok {
		t.Error("Should return empty response")
	} else if len(name) > 0 {
		t.Error("Should return empty response")
	}

	expr = &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.CallExpr{},
			Sel: ast.NewIdent("actual"),
		},
	}

	name, ok = h.GetActualFuncName(expr)
	if ok {
		t.Error("Should return empty response")
	} else if len(name) > 0 {
		t.Error("Should return empty response")
	}

	expr = &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   ast.NewIdent("notgomega"),
			Sel: ast.NewIdent("actual"),
		},
	}

	name, ok = h.GetActualFuncName(expr)
	if ok {
		t.Error("Should return empty response")
	} else if len(name) > 0 {
		t.Error("Should return empty response")
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
