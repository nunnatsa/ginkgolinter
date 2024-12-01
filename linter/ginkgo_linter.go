package linter

import (
	"go/ast"

	"github.com/nunnatsa/ginkgolinter/gomegaanalyzer"

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

type GinkgoLinter struct {
	config         *types.Config
	gomegaAnalyzer *analysis.Analyzer
}

// NewGinkgoLinter return new ginkgolinter object
func NewGinkgoLinter(config *types.Config, gomegaAnalyzer *analysis.Analyzer) *GinkgoLinter {
	return &GinkgoLinter{
		config:         config,
		gomegaAnalyzer: gomegaAnalyzer,
	}
}

// Run is the main assertion function
func (l *GinkgoLinter) Run(pass *analysis.Pass) (any, error) {
	gomegaAnalyzerRes := pass.ResultOf[l.gomegaAnalyzer].(*gomegaanalyzer.Result)
	expRules := rules.GetRules(gomegaAnalyzerRes)

	for _, file := range pass.Files {
		gomegaResults, ok := gomegaAnalyzerRes.GetFileResult(file)
		if !ok {
			continue
		}

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
						goDeeper = ginkgoHndlr.HandleGinkgoSpecs(val, fileConfig, pass) || goDeeper
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
				if ginkgoHndlr.HandleGinkgoSpecs(assertionExp, config, pass) {
					return true
				}
			}

			// no more ginkgo checks. From here it's only gomega. So if there is no gomega handler, exit here.
			if gomegaHndlr == nil {
				return true
			}

			//gexp, ok := expression.New(assertionExp, pass, gomegaHndlr, "time")

			gexp, ok := gomegaResults[assertionExp]
			if !ok || gexp == nil {
				return true
			}

			reportBuilder := reports.NewBuilder(assertionExp, formatter.NewGoFmtFormatter(pass.Fset))
			return checkGomegaExpression(gexp, config, reportBuilder, pass, expRules)
		})
	}
	return nil, nil
}

func checkGomegaExpression(gexp *expression.GomegaExpression, config types.Config, reportBuilder *reports.Builder, pass *analysis.Pass, expRules rules.Rules) bool {
	goNested := false
	if rules.GetMissingAssertionRule().Apply(gexp, config, reportBuilder) {
		goNested = true
	} else {
		if gexp.IsAsync() {
			rules.GetAsyncRules().Apply(gexp, config, reportBuilder)
			goNested = true
		} else {
			expRules.Apply(gexp, config, reportBuilder)
		}
	}

	if reportBuilder.HasReport() {
		reportBuilder.SetFixOffer(gexp.GetClone())
		pass.Report(reportBuilder.Build())
	}

	return goNested
}
