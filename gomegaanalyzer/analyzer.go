package gomegaanalyzer

import (
	"github.com/nunnatsa/ginkgolinter/internal/expression"
	"github.com/nunnatsa/ginkgolinter/internal/gomegahandler"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"reflect"
)

type Result map[ast.Expr]*expression.GomegaExpression

func New() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:       "gomegaAnalyzer",
		Doc:        "gather all gomega expressions",
		Run:        run,
		ResultType: reflect.TypeOf(Result{}),
	}
}

func run(pass *analysis.Pass) (any, error) {
	expressions := make(Result)
	for _, file := range pass.Files {
		gomegaHndlr := gomegahandler.GetGomegaHandler(file, pass)
		if gomegaHndlr == nil {
			continue
		}

		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			expr, ok := expression.New(call, pass, gomegaHndlr, getTimePkg(file))
			if !ok {
				return true
			}

			expressions[call] = expr

			return true
		})
	}

	return expressions, nil
}

func getTimePkg(file *ast.File) string {
	timePkg := "time"
	for _, imp := range file.Imports {
		if imp.Path.Value == `"time"` && imp.Name != nil {
			timePkg = imp.Name.Name
		}
	}

	return timePkg
}
