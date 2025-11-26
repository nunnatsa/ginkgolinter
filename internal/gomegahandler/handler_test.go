//nolint:staticcheck
package gomegahandler

import (
	"errors"
	"fmt"
	"go/ast"
	"os"
	"path"

	// gotypes "go/types"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

// run sets up a directory with the given code in main.go such that
// importing gomega/ginkgo works, then invokes the callback for that file.
//
// The callback then can run checks and report failures through errors.
//
// This approach is necessary because stubbing a full analysis.Pass
// instance is hard.
func run(t *testing.T, code string, cb func(file *ast.File, pass *analysis.Pass) error) {
	tmpDir := t.TempDir()
	srcDir := path.Join(tmpDir, "src")
	if err := os.Mkdir(srcDir, 0700); err != nil {
		t.Fatalf("creating src dir failed: %v", err)
	}

	if err := os.CopyFS(srcDir, os.DirFS("../../testdata/src/a/vendor")); err != nil {
		t.Fatalf("populating temp dir with vendored source failed: %v", err)
	}

	testFileName := path.Join(tmpDir, "main.go")
	if err := os.WriteFile(testFileName, []byte(code), 0600); err != nil {
		t.Fatalf("populating temp dir with test file code: %v", err)
	}

	analyzer := &analysis.Analyzer{
		Name: "ginkgolinter",
		Doc:  "test linter",
		Run: func(pass *analysis.Pass) (any, error) {
			for _, file := range pass.Files {
				if file.Name.Name == "main" {
					return nil, cb(file, pass)
				}
			}
			return nil, errors.New("main file not found")
		},
	}
	analysistest.Run(t, tmpDir, analyzer, ".")
}

func TestGetGomegaHandler(t *testing.T) {
	for name, tt := range map[string]struct {
		code         string
		expectGomega bool
		// expectInfo contains the expected info for each call expression,
		// in depth-first preorder.
		expectInfos []*GomegaBasicInfo
	}{
		"dot": {
			code: `package main
import (
	. "github.com/onsi/gomega"
	gtypes "github.com/onsi/gomega/types"
)

func other() interface{foo ()} {
	return nil
}

func myExpect(actual any) Assertion {
	return Expect(actual)
}

func myEventually(actual any) AsyncAssertion {
	return Eventually(actual)
}

func myEqual(expect any) gtypes.GomegaMatcher {
	return Equal(expect)
}

func main() {
	Expect(1)
	Expect(1).Error()
	NewWithT(nil).Expect(1)
	Ω(1)
	Eventually(0)
	Consistently(0)
	Consistently(0).WithTimeout(0).WithPolling(0)

	other()
	other().foo()

	myExpect(1)
	myEventually(1)

	Equal(0)
	myEqual(0)
}
`,
			expectGomega: true,
			expectInfos: []*GomegaBasicInfo{
				{MethodName: "Expect", RootCallType: SyncAssertionCall},
				{MethodName: "Expect", RootCallType: SyncAssertionCall, HasErrorMethod: true},
				{MethodName: "Expect", RootCallType: SyncAssertionCall},
				{MethodName: "Ω", RootCallType: SyncAssertionCall},
				{MethodName: "Eventually", RootCallType: AsyncAssertionCall},
				{MethodName: "Consistently", RootCallType: AsyncAssertionCall},
				{MethodName: "Consistently", RootCallType: AsyncAssertionCall},

				nil,
				nil,

				{MethodName: "myExpect", RootCallType: SyncAssertionCall},
				{MethodName: "myEventually", RootCallType: AsyncAssertionCall},

				{MethodName: "Equal", RootCallType: MatcherCall},
				{MethodName: "myEqual", RootCallType: MatcherCall},
			},
		},

		"noname": {
			code: `package main
import (
	"github.com/onsi/gomega"
)

func main() {
	gomega.Expect(1)
	gomega.Expect(1).Error()
	gomega.NewWithT(nil).Expect(1)
	gomega.Ω(1)
	gomega.Eventually(0)
	gomega.Consistently(0)
	gomega.Consistently(0).WithTimeout(0).WithPolling(0)
	gomega.Equal(0)
}
`,
			expectGomega: true,
			expectInfos: []*GomegaBasicInfo{
				{MethodName: "Expect", RootCallType: SyncAssertionCall},
				{MethodName: "Expect", RootCallType: SyncAssertionCall, HasErrorMethod: true},
				{MethodName: "Expect", RootCallType: SyncAssertionCall},
				{MethodName: "Ω", RootCallType: SyncAssertionCall},
				{MethodName: "Eventually", RootCallType: AsyncAssertionCall},
				{MethodName: "Consistently", RootCallType: AsyncAssertionCall},
				{MethodName: "Consistently", RootCallType: AsyncAssertionCall},
				{MethodName: "Equal", RootCallType: MatcherCall},
			},
		},

		"name": {
			code: `package main
import (
	g "github.com/onsi/gomega"
)

func main() {
	g.Expect(1)
	g.Expect(1).Error()
	g.NewWithT(nil).Expect(1)
	g.Ω(1)
	g.Eventually(0)
	g.Consistently(0)
	g.Consistently(0).WithTimeout(0).WithPolling(0)
	g.Equal(0)
}
`,
			expectGomega: true,
			expectInfos: []*GomegaBasicInfo{
				{MethodName: "Expect", RootCallType: SyncAssertionCall},
				{MethodName: "Expect", RootCallType: SyncAssertionCall, HasErrorMethod: true},
				{MethodName: "Expect", RootCallType: SyncAssertionCall},
				{MethodName: "Ω", RootCallType: SyncAssertionCall},
				{MethodName: "Eventually", RootCallType: AsyncAssertionCall},
				{MethodName: "Consistently", RootCallType: AsyncAssertionCall},
				{MethodName: "Consistently", RootCallType: AsyncAssertionCall},
				{MethodName: "Equal", RootCallType: MatcherCall},
			},
		},

		"none": {
			code: `package main
`,
			expectGomega: false,
		},
	} {
		t.Run(name, func(t *testing.T) {
			run(t, tt.code, func(file *ast.File, pass *analysis.Pass) error {
				h := GetGomegaHandler(file, pass)
				if (h != nil) != tt.expectGomega {
					return fmt.Errorf("expectNoGomega %v, got handler %#v", tt.expectGomega, h)
				}
				if h == nil {
					return nil
				}

				infoIndex := 0
				for n := range ast.Preorder(file) {
					stmt, ok := n.(*ast.ExprStmt)
					if !ok {
						continue
					}
					call, ok := stmt.X.(*ast.CallExpr)
					if !ok {
						continue
					}

					if infoIndex >= len(tt.expectInfos) {
						return fmt.Errorf("unexpected call statement %#v", call)
					}
					actualInfo := h.GetGomegaBasicInfo(call)
					expectInfo := tt.expectInfos[infoIndex]
					currentIndex := infoIndex
					infoIndex++
					if (actualInfo == nil) != (expectInfo == nil) {
						return fmt.Errorf("#%d: expected %#v, actual %#v", currentIndex, expectInfo, actualInfo)
					}
					if actualInfo == nil {
						continue
					}

					// Cannot compare RootCall, but it has to be set.
					if actualInfo.RootCall == nil {
						return fmt.Errorf("#%d: missing RootCall in %#v", currentIndex, actualInfo)
					}
					actualInfo.RootCall = nil
					if *actualInfo != *expectInfo {
						return fmt.Errorf("#%d: expected %#v, actual %#v", currentIndex, expectInfo, actualInfo)
					}
				}
				if infoIndex != len(tt.expectInfos) {
					return fmt.Errorf("wanted %d call expressions, got %d", len(tt.expectInfos), infoIndex)
				}
				return nil
			})
		})
	}
}

func TestGomegaDotHandler_ReplaceFunction(t *testing.T) {
	h := &Handler{name: "" /* no name, dot import! */}

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
	h := &Handler{name: "gomega"}

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
