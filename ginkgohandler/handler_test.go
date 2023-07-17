package ginkgohandler

import (
	"go/ast"
	"testing"
)

var imports = []string{`"github.com/onsi/ginkgo"`, `"github.com/onsi/ginkgo/v2"`}

func TestGetGinkgoHandler_dot(t *testing.T) {
	for _, imp := range imports {
		t.Run("TestGetGinkgoHandler_dot: "+imp, func(tt *testing.T) {
			name := ast.NewIdent("test.go")
			file := &ast.File{
				Name: name,
				Imports: []*ast.ImportSpec{
					{
						Name: ast.NewIdent("."),
						Path: &ast.BasicLit{Value: imp},
					},
				},
			}

			h := GetGinkgoHandler(file)
			if h == nil {
				tt.Fatalf("should return dotHandler")
			}
			_, ok := h.(dotHandler)
			if !ok {
				tt.Error("should return dotHandler")
			}
		})
	}
}

func TestGetGinkgoHandler_noname(t *testing.T) {
	for _, imp := range imports {
		t.Run("TestGetGinkgoHandler_noname: "+imp, func(tt *testing.T) {
			name := ast.NewIdent("test.go")
			file := &ast.File{
				Name: name,
				Imports: []*ast.ImportSpec{
					{
						Path: &ast.BasicLit{Value: imp},
					},
				},
			}

			h := GetGinkgoHandler(file)
			if h == nil {
				tt.Fatalf("should return nameHandler")
			}
			n, ok := h.(nameHandler)
			if !ok {
				tt.Error("should return nameHandler")
			}

			if string(n) != "ginkgo" {
				tt.Errorf("import name should be `ginkgo`, but it's %s", string(n))
			}

		})
	}
}

func TestGetGinkgoHandler_name(t *testing.T) {
	for _, imp := range imports {
		t.Run("TestGetGinkgoHandler_name: "+imp, func(tt *testing.T) {
			name := ast.NewIdent("test.go")
			file := &ast.File{
				Name: name,
				Imports: []*ast.ImportSpec{
					{
						Name: ast.NewIdent("name"),
						Path: &ast.BasicLit{Value: `"github.com/onsi/ginkgo"`},
					},
				},
			}

			h := GetGinkgoHandler(file)
			if h == nil {
				tt.Fatalf("should return nameHandler")
			}
			n, ok := h.(nameHandler)
			if !ok {
				tt.Error("should return nameHandler")
			}

			if string(n) != "name" {
				tt.Errorf("import name should be `name`, but it's %s", string(n))
			}
		})
	}
}

func TestGetGinkgoHandler_no_ginkgo(t *testing.T) {
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

	h := GetGinkgoHandler(file)
	if h != nil {
		t.Fatalf("should return nil")
	}
}
