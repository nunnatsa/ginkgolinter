# contributing Guide

Thanks for contributing to the ginkgolinter project, and helping us improve it!

## Reporting a bugs
Please let us know if you have any issue using the ginkgolinter. Under the "issues" tab, click on the "New Issue" button, then click on the "Get Started" button in the "Bug Report" section. Fill the details and submit the new issue. 

## Enhancement Suggestion
We want to get better. If you have an idea how to do that, please open an Enhancement Suggestion issue. Under the "issues" tab, click on the "New Issue" button, then click on the "Get Started" button in the "Enhancement Suggestion" section. Fill the details and submit the new issue.

## Build the linter executable (CLI)
Just run `make build`

## Contribute Code
In order to contribute code changes, fork this repository into your own organization, commit the code changes to a new git branch, then push the changes and create a PR.

### Code Structure
The code is using the `go/ast` package and the `golang.org/x/tools/go/analysis` module to implement the logic.

The main logic of the linter is placed in the `ginkgoLinter::run method`, in the [ginkgo_linter.go](../ginkgo_linter.go) file. 

The ginkgo related code is under internal/ginkgohandler.

The gomega related code is larger and spread over several packages:
```
...
internal/
    ...
    + gomegahandler - holds logic to get expression information, regardless if the gomega import is done with or without dot.
    + gomegainfo - general utilities for the gomega handling code
    + expression - gomega expression types and code. Each gomega expression is modeled into these types, to later be processed by the rules
        + actual - types to model the actual part of the gomega assertion expression
        + matcher - types to model the matcher part of the gomega assertion expression
        + value - models a value in several actual and matcher types
    + rules - the gomega linter rules. A rule receives a gomega expression and trigger a linter error if needed. 
    ...
... 
```
### Testing
Any functional code change must be tested. Please make sure to add relevant tests that demonstrate that the code changes are working. Run the full test suite to make sure that the code changes didn't break the linter functionality.

#### Unit tests
The main unit tests are built with two parts. The testdata and the actual test logic. The test data is placed under [testdata/src/a](../testdata/src/a). Add a new directory for each new feature, or update an existing file.

The content of the test data is a gomega and ginkgo test files. This may be confusing: This linter domain is testing, so the testdata are test files, but these files should not be running. the content of this files may not make any logical sense. But the files must be valid, compilable golang files.

For each line that should trigger a linter warning, add a comment with the following format
```
// want `<expected linter message>`
```

for example:
```go
Expect(t).To(Equal(true)) // want `ginkgo-linter: wrong boolean assertion; consider using .Expect\(t\)\.To\(BeTrue\(\)\). instead`
```

The `want` comment is working with a limited regex. This is why you'll have to escape parentheses and dots. there is no way to escape the "`" character if it is part of the linter message. replace it with a dot.

The second part of the unit test is the test logic itself. We're using the `golang.org/x/tools/go/analysis/analysistest` package for that. The `analysistest.Run` function, runs the linter against the test data and compares the linter output with the `want` comments.

There are two parametrized unit tests in the [ginkgo_linter_test.go](../ginkgo_linter_test.go) file. The `TestAllUseCases` function just run the `analysistest.Run` function against the testdata package, and the `TestFlags` also modify the linter configurations (by setting command line flags) before running the tests.

To add a test for a new testdata package, add a new object to the test parameters in `TestAllUseCases`; e.g.
```go
		{
			testName: "find wrong length assertion",
			testData: "a/len",
		},
```

If you added a new flag, or need to test your changes with an existing flag, add a new object to the test parameters in `TestFlags`. The `flags` field is a slice of strings, containing the flag list to set to `"true"`.

To run the unit tests, do:
```shell
make unit-test
```

#### CLI test
The CLI test is done using [testscript](https://github.com/rogpeppe/go-internal). Meaning, we build a text base isolated 
environment for each test, with all the required files.

These test files are located under tests/testdata, and contains a test script and the required files within the same
txtar file.

The purpose of these tests is to check the CLI itself. We don't want to test everything there, but only the CLI and its 
command line parameters. 

If you made a change with the flags, make sure to check this change in the CLI test. For regular change, use the
unit-tests.

To run the CLI test, run 
```shell
make test-cli
```

To run both unit test and CLI test, run
```shell
make test
```
