package sprintfchecker

import (
	"go/ast"
	"go/token"
)

type Checker struct {
	importName string
}

func NewSprintfChecker(file *ast.File) *Checker {
	for _, imp := range file.Imports {
		if imp.Path.Kind == token.STRING && imp.Path.Value == `"fmt"` {
			name := "fmt"
			if imp.Name != nil {
				name = imp.Name.Name
			}

			return &Checker{
				importName: name,
			}
		}
	}
	return nil
}

func (s *Checker) Check(assertionExp *ast.CallExpr) bool {
	if s == nil {
		return false
	}

	if len(assertionExp.Args) != 2 {
		return false
	}

	if call, ok := assertionExp.Args[1].(*ast.CallExpr); ok {
		if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
			if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == s.importName && sel.Sel.Name == "Sprintf" {
				args := make([]ast.Expr, len(call.Args)+1)
				args[0] = assertionExp.Args[0]
				copy(args[1:], call.Args)
				assertionExp.Args = args

				return true
			}
		}
	}

	return false
}
