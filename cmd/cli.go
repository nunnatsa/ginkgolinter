package main

import (
	"github.com/nunnatsa/ginkgolinter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(ginkgolinter.Analyzer)
}
