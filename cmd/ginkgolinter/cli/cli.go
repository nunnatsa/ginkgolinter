package cli

import (
	"fmt"
	"os"
	"runtime"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/nunnatsa/ginkgolinter"
	"github.com/nunnatsa/ginkgolinter/version"
)

func Main() int {
	if len(os.Args) == 2 && os.Args[1] == "version" {
		fmt.Printf("ginkgolinter version: %s\n", version.Version())
		fmt.Printf("git hash:             %s\n", version.GitHash())
		fmt.Printf("go version:           %s\n", runtime.Version())
		os.Exit(0)
	}

	singlechecker.Main(ginkgolinter.NewAnalyzer())

	return 0
}
