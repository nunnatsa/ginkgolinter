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

	testscript.Main(m, map[string]func(){
		"ginkgolinter": cli.Main,
	})
}

func TestCli(t *testing.T) {
	if os.Getenv("RUN_CLI_TESTS") != "true" {
		t.Skipf(`Skipping CLI tests. These tests are not unit tests, but more like integration tests.
The tests will run when the RUN_CLI_TESTS environment variable is set to "true"`)
	}
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
