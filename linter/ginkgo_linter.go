package linter

import (
	"fmt"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis"

	"github.com/nunnatsa/ginkgolinter/internal/expression"
	"github.com/nunnatsa/ginkgolinter/internal/formatter"
	"github.com/nunnatsa/ginkgolinter/internal/ginkgohandler"
	"github.com/nunnatsa/ginkgolinter/internal/gomegahandler"
	"github.com/nunnatsa/ginkgolinter/internal/reports"
	"github.com/nunnatsa/ginkgolinter/internal/rules"
	"github.com/nunnatsa/ginkgolinter/types"
)

// The ginkgolinter enforces standards of using ginkgo and gomega.
//
// For more details, look at the README.md file

const (
	linterName            = "ginkgo-linter"
	focusContainerFound   = linterName + ": Focus container found. This is used only for local debug and should not be part of the actual source code. Consider to replace with %q"
	focusSpecFound        = linterName + ": Focus spec found. This is used only for local debug and should not be part of the actual source code. Consider to remove it"
	useBeforeEachTemplate = "use BeforeEach() to assign variable %s"
)

type GinkgoLinter struct {
	config *types.Config
}

// NewGinkgoLinter return new ginkgolinter object
func NewGinkgoLinter(config *types.Config) *GinkgoLinter {
	return &GinkgoLinter{
		config: config,
	}
}

// Run is the main assertion function
func (l *GinkgoLinter) Run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		fileConfig := l.config.Clone()

		cm := ast.NewCommentMap(pass.Fset, file, file.Comments)

		fileConfig.UpdateFromFile(cm)

		gomegaHndlr := gomegahandler.GetGomegaHandler(file, pass)
		ginkgoHndlr := ginkgohandler.GetGinkgoHandler(file)

		if gomegaHndlr == nil && ginkgoHndlr == nil { // no gomega or ginkgo imports => no use in gomega in this file; nothing to do here
			continue
		}

		ast.Inspect(file, func(n ast.Node) bool {
			if ginkgoHndlr != nil {
				goDeeper := false
				spec, ok := n.(*ast.ValueSpec)
				if ok {
					for _, val := range spec.Values {
						goDeeper = handleGinkgoSpecs(val, fileConfig, pass, ginkgoHndlr) || goDeeper
					}
				}
				if goDeeper {
					return true
				}
			}

			stmt, ok := n.(*ast.ExprStmt)
			if !ok {
				return true
			}

			// search for function calls
			assertionExp, ok := stmt.X.(*ast.CallExpr)
			if !ok {
				return true
			}

			config := fileConfig.Clone()
			if comments, ok := cm[stmt]; ok {
				config.UpdateFromComment(comments)
			}

			if ginkgoHndlr != nil {
				if handleGinkgoSpecs(assertionExp, config, pass, ginkgoHndlr) {
					return true
				}
			}

			// no more ginkgo checks. From here it's only gomega. So if there is no gomega handler, exit here.
			if gomegaHndlr == nil {
				return true
			}

			gexp, ok := expression.New(assertionExp, pass, gomegaHndlr, getTimePkg(file))
			if !ok || gexp == nil {
				return true
			}

			return checkGomegaExpression(pass, config, gexp)
		})
	}
	return nil, nil
}

func handleGinkgoSpecs(expr ast.Expr, config types.Config, pass *analysis.Pass, ginkgoHndlr ginkgohandler.Handler) bool {
	goDeeper := false
	if exp, ok := expr.(*ast.CallExpr); ok {
		if bool(config.ForbidFocus) && checkFocusContainer(pass, ginkgoHndlr, exp) {
			goDeeper = true
		}

		if bool(config.ForbidSpecPollution) && checkAssignmentsInContainer(pass, ginkgoHndlr, exp) {
			goDeeper = true
		}
	}
	return goDeeper
}

func checkAssignmentsInContainer(pass *analysis.Pass, ginkgoHndlr ginkgohandler.Handler, exp *ast.CallExpr) bool {
	foundSomething := false
	if ginkgoHndlr.IsWrapContainer(exp) {
		for _, arg := range exp.Args {
			if fn, ok := arg.(*ast.FuncLit); ok {
				if fn.Body != nil {
					if checkAssignments(pass, fn.Body.List) {
						foundSomething = true
					}
					break
				}
			}
		}
	}

	return foundSomething
}

func checkAssignments(pass *analysis.Pass, list []ast.Stmt) bool {
	foundSomething := false
	for _, stmt := range list {
		switch st := stmt.(type) {
		case *ast.DeclStmt:
			if gen, ok := st.Decl.(*ast.GenDecl); ok {
				if gen.Tok != token.VAR {
					continue
				}
				for _, spec := range gen.Specs {
					if valSpec, ok := spec.(*ast.ValueSpec); ok {
						if checkAssignmentsValues(pass, valSpec.Names, valSpec.Values) {
							foundSomething = true
						}
					}
				}
			}

		case *ast.AssignStmt:
			for i, val := range st.Rhs {
				if !is[*ast.FuncLit](val) {
					if id, isIdent := st.Lhs[i].(*ast.Ident); isIdent && id.Name != "_" {
						reportNoFix(pass, id.Pos(), useBeforeEachTemplate, id.Name)
						foundSomething = true
					}
				}
			}

		case *ast.IfStmt:
			if st.Body != nil {
				if checkAssignments(pass, st.Body.List) {
					foundSomething = true
				}
			}
			if st.Else != nil {
				if block, isBlock := st.Else.(*ast.BlockStmt); isBlock {
					if checkAssignments(pass, block.List) {
						foundSomething = true
					}
				}
			}
		}
	}

	return foundSomething
}

func checkAssignmentsValues(pass *analysis.Pass, names []*ast.Ident, values []ast.Expr) bool {
	foundSomething := false
	for i, val := range values {
		if !is[*ast.FuncLit](val) {
			reportNoFix(pass, names[i].Pos(), useBeforeEachTemplate, names[i].Name)
			foundSomething = true
		}
	}

	return foundSomething
}

func checkFocusContainer(pass *analysis.Pass, ginkgoHndlr ginkgohandler.Handler, exp *ast.CallExpr) bool {
	foundFocus := false
	isFocus, id := ginkgoHndlr.GetFocusContainerName(exp)
	if isFocus {
		reportNewName(pass, id, id.Name[1:], focusContainerFound, id.Name)
		foundFocus = true
	}

	if id != nil && ginkgohandler.IsContainer(id.Name) {
		for _, arg := range exp.Args {
			if ginkgoHndlr.IsFocusSpec(arg) {
				reportNoFix(pass, arg.Pos(), focusSpecFound)
				foundFocus = true
			} else if callExp, ok := arg.(*ast.CallExpr); ok {
				if checkFocusContainer(pass, ginkgoHndlr, callExp) { // handle table entries
					foundFocus = true
				}
			}
		}
	}

	return foundFocus
}

func checkGomegaExpression(pass *analysis.Pass, config types.Config, gexp *expression.GomegaExpression) bool {
	reportBuilder := reports.NewBuilder(gexp.Orig, formatter.NewGoFmtFormatter(pass.Fset))

	goNested := false
	if rules.GetMissingAssertionRule().Apply(gexp, config, reportBuilder) {
		goNested = true
	} else {
		if gexp.IsAsync() {
			rules.GetAsyncRules().Apply(gexp, config, reportBuilder)
			goNested = true
		} else {
			rules.GetRules().Apply(gexp, config, reportBuilder)
		}
	}

	if reportBuilder.HasReport() {
		reportBuilder.SetFixOffer(gexp.Clone)
		pass.Report(reportBuilder.Build())
	}

	return goNested
}

func reportNewName(pass *analysis.Pass, id *ast.Ident, newName string, messageTemplate, oldExpr string) {
	pass.Report(analysis.Diagnostic{
		Pos:     id.Pos(),
		Message: fmt.Sprintf(messageTemplate, newName),
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: fmt.Sprintf("should replace %s with %s", oldExpr, newName),
				TextEdits: []analysis.TextEdit{
					{
						Pos:     id.Pos(),
						End:     id.End(),
						NewText: []byte(newName),
					},
				},
			},
		},
	})
}

func reportNoFix(pass *analysis.Pass, pos token.Pos, message string, args ...any) {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	pass.Report(analysis.Diagnostic{
		Pos:     pos,
		Message: message,
	})
}

func is[T any](x any) bool {
	_, matchType := x.(T)
	return matchType
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
