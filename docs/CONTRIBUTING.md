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
The main logic of the linter is placed in the `ginkgoLinter::run method`, in the [ginkgo_linter.go](../ginkgo_linter.go) file. 

The code is using the `go/ast` package and the `golang.org/x/tools/go/analysis` module to implement the logic.

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
go test ./...
```

#### Functional tests
There is a limited functional test, in [.github/workflows/sanity.yml](../.github/workflows/sanity.yml) running as a github action. The test builds the ginkgolinter executable and runs it against the test data (**Note**: the `want` comments are ignored by the actual linter. They are only working with the `analysistest.Run` function). 

The test counts the number of linter warnings with different combinations of the commandline flags.

Make sure to update the expected numbers to match your changes.