package gomegaanalyzer

import (
	"go/ast"
	"go/token"
	"reflect"
	"sort"

	"github.com/nunnatsa/ginkgolinter/internal/asyncfuncpos"

	"golang.org/x/tools/go/analysis"

	"github.com/nunnatsa/ginkgolinter/internal/expression"
	"github.com/nunnatsa/ginkgolinter/internal/gomegahandler"
)

type FileResult map[ast.Expr]*expression.GomegaExpression

type Result struct {
	expressions map[*ast.File]FileResult
	asyncScope  asyncfuncpos.Poses
}

func (r Result) GetFileResults() map[*ast.File]FileResult {
	return r.expressions
}

func (r Result) GetFileResult(file *ast.File) (FileResult, bool) {
	fr, ok := r.expressions[file]
	return fr, ok
}

func (r Result) InAsyncScope(pos token.Pos) (token.Pos, bool) {
	if fncPos, found := r.asyncScope.Contains(pos); found {
		return fncPos.CallPos, true
	}

	return token.NoPos, false
}

func New() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:       "gomegaAnalyzer",
		Doc:        "gather all gomega expressions",
		Run:        run,
		ResultType: reflect.TypeOf(&Result{}),
	}
}

func run(pass *analysis.Pass) (any, error) {
	expressions := make(map[*ast.File]FileResult)
	var funcPoss asyncfuncpos.Poses

	for _, file := range pass.Files {
		gomegaHndlr := gomegahandler.GetGomegaHandler(file, pass)
		if gomegaHndlr == nil {
			continue
		}

		fileResult := make(FileResult)

		ast.Inspect(file, func(n ast.Node) bool {
			stmt, ok := n.(*ast.ExprStmt)
			if !ok {
				return true
			}

			call, ok := stmt.X.(*ast.CallExpr)
			if !ok {
				return true
			}

			expr, ok := expression.New(call, pass, gomegaHndlr, getTimePkg(file))
			if !ok {
				return true
			}

			fileResult[call] = expr

			if expr.IsAsync() {
				if arg := expr.GetAsyncActualArg(); arg != nil && arg.HasAsyncFuncPos() {
					funcPoss = append(funcPoss, arg.AsyncFucPos())
				}
			}

			return true
		})

		expressions[file] = fileResult
	}

	sort.Sort(funcPoss)

	return &Result{expressions: expressions, asyncScope: funcPoss}, nil
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
