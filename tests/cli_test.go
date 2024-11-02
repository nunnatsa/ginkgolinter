package tests_test

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"

	"github.com/nunnatsa/ginkgolinter/cmd/ginkgolinter/cli"
)

func TestMain(m *testing.M) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if err = os.Setenv("HOME", path.Join(pwd, "cli-test")); err != nil {
		panic(err)
	}

	os.Exit(testscript.RunMain(m, map[string]func() int{
		"ginkgolinter": cli.Main,
	}))
}

func TestCli(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if !strings.HasSuffix(pwd, "tests") {
		pwd = path.Join(pwd, "tests")
	}

	testscript.Run(t, testscript.Params{
		Dir: path.Join(pwd, "testdata"),
	})
}
