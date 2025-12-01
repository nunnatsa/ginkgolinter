package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/tools/txtar"
)

var (
	targetDir = flag.String("target-dir", "./tests/testdata", "target directory, where the txtar files are stored")
	sourceDir = flag.String("source-dir", "./testdata/src/a", "the path of the directory with the required go.mod and go.sum files")
)

func main() {
	dir, err := os.ReadDir(*targetDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read directory %q; %v", *targetDir, err)
		os.Exit(1)
	}

	for _, d := range dir {
		if d.IsDir() {
			continue
		}

		if !strings.HasSuffix(d.Name(), ".txtar") {
			continue
		}

		filePath := filepath.Join(*targetDir, d.Name())
		err = updateFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to update file %q; %v", filePath, err)
			os.Exit(1)
		}
	}

	fmt.Println("Successfully updated all txtar files in", *targetDir)
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s [flags]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
}

var moduleRegex = regexp.MustCompile(`(?m)^module [a-zA-Z0-9_]+$`)

func readSourceFiles() ([]byte, []byte, error) {
	goMod, err := os.ReadFile(filepath.Join(*sourceDir, "go.mod"))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read go.mod: %w", err)
	}

	goSum, err := os.ReadFile(filepath.Join(*sourceDir, "go.sum"))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read go.sum: %w", err)
	}

	return goMod, goSum, nil
}

func updateFile(filename string) error {
	goMod, goSum, err := readSourceFiles()
	if err != nil {
		return err
	}

	stat, err := os.Stat(filename)
	if err != nil {
		return fmt.Errorf("can't stat %q; %w", filename, err)
	}

	txtarFile, err := txtar.ParseFile(filename)
	if err != nil {
		return fmt.Errorf("can't parse %q; %w", filename, err)
	}

	for i, file := range txtarFile.Files {
		switch fileName := file.Name; fileName {
		case "go.mod":
			pkgLine := moduleRegex.Find(file.Data)
			if pkgLine == nil {
				return fmt.Errorf("can't find the module line in file %q", fileName)
			}
			txtarFile.Files[i].Data = moduleRegex.ReplaceAll(goMod, pkgLine)
		case "go.sum":
			txtarFile.Files[i].Data = goSum
		default:
			continue
		}
	}

	updated := txtar.Format(txtarFile)
	err = os.WriteFile(filename, updated, stat.Mode())
	if err != nil {
		return fmt.Errorf("can't update %q; %w", filename, err)
	}

	return nil
}
