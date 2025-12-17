package typecheck

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	gotypes "go/types"
	"runtime"
	"strings"
	"testing"
)

func TestImplementsTB(t *testing.T) {
	src := `
package main

var (
	NilInstance = any(nil)
)

type Error interface {
	Error(args ...any)
}

type Errorf interface {
	Errorf(fmt string, args ...any)
}

type Fatal interface {
	Fatal(args ...any)
}

type Fatalf interface {
	Fatalf(fmt string, args ...any)
}

type ErrorWrongArgs interface {
	Error(fmt string, args ...any)
}

type ErrorfWrongArgs interface {
	ErrorF(args ...any)
}

type FatalWrongArgs interface {
	Fatal(fmt string, args ...any)
}

type FatalfWrongArgs interface {
	Fatalf(args ...any)
}
`
	pkg := mustTypecheck(src)

	for name, expectResult := range map[string]bool{
		"NilInstance":     false,
		"Error":           true,
		"Errorf":          true,
		"Fatal":           true,
		"Fatalf":          true,
		"ErrorWrongArgs":  false,
		"ErrorfWrongArgs": false,
		"FatalWrongArgs":  false,
		"FatalfWrongArgs": false,
	} {
		t.Run(name, func(t *testing.T) {
			instance := pkg.Scope().Lookup(name)
			if instance == nil {
				t.Fatalf("%s not defined in source code", name)
			}
			actualResult := ImplementsTB(instance.Type())
			if actualResult != expectResult {
				t.Errorf("expected %v, got %v", expectResult, actualResult)
			}
		})
	}
}

// The following code was copied from https://cs.opensource.google/go/go/+/refs/tags/go1.25.6:src/go/types/api_test.go;l=29-74;drc=34b70684ba2fc8c5cba900e9abdfb874c1bd8c0e
// and slightly simplified.
//
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the https://cs.opensource.google/go/go/+/refs/tags/go1.25.6:LICENSE

func defaultImporter(fset *token.FileSet) gotypes.Importer {
	return importer.ForCompiler(fset, runtime.Compiler, nil)
}

func mustParse(fset *token.FileSet, src string) *ast.File {
	f, err := parser.ParseFile(fset, pkgName(src), src, parser.ParseComments)
	if err != nil {
		panic(err) // so we don't need to pass *testing.T
	}
	return f
}

func typecheck(src string) (*gotypes.Package, error) {
	fset := token.NewFileSet()
	f := mustParse(fset, src)
	conf := &gotypes.Config{
		Error:    func(error) {}, // collect all errors
		Importer: defaultImporter(fset),
	}
	return conf.Check(f.Name.Name, fset, []*ast.File{f}, nil)
}

func mustTypecheck(src string) *gotypes.Package {
	pkg, err := typecheck(src)
	if err != nil {
		panic(err) // so we don't need to pass *testing.T
	}
	return pkg
}

// pkgName extracts the package name from src, which must contain a package header.
func pkgName(src string) string {
	const kw = "package "
	if i := strings.Index(src, kw); i >= 0 {
		after := src[i+len(kw):]
		n := len(after)
		if i := strings.IndexAny(after, "\n\t ;/"); i >= 0 {
			n = i
		}
		return after[:n]
	}
	panic("missing package header: " + src)
}
