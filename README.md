# ginkgo-linter

This is a golang linter to check usage of the ginkgo and gomega packages.

ginkgo is a testing framework and gomega is its assertion package.

## Linter Checks
### Wrong Length checks
The linter finds usage of the golang built-in `len` function, and then all kind of matchers, while there are already gomega matchers for these usecases.

There are several wrong patterns:
```go
Expect(len(x)).To(Equal(0)) // should be Expect(x).To(BeEmpty())
Expect(len(x)).To(BeZero()) // should be Expect(x).To(BeEmpty())
Expect(len(x)).To(BeNumeric(">", 0)) // should be Expect(x).ToNot(BeEmpty())
Expect(len(x)).To(BeNumeric(">=", 1)) // should be Expect(x).ToNot(BeEmpty())
Expect(len(x)).To(BeNumeric("==", 0)) // should be Expect(x).To(BeEmpty())
Expect(len(x)).To(BeNumeric("!=", 0)) // should be Expect(x).ToNot(BeEmpty())

Expect(len(x)).To(Equal(1)) // should be Expect(x).To(HaveLen(1))
Expect(len(x)).To(BeNumeric("==", 2)) // should be Expect(x).To(HaveLen(2))
Expect(len(x)).To(BeNumeric("!=", 3)) // should be Expect(x).ToNot(HaveLen(3))
```

The linter supports the `Expect`, `ExpectWithOffset` and the `Ω` functions, and the `Should`, `ShouldNot`, `To`, `ToNot` and `NotTo` assertion functions.

It also supports the embedded `Not()` function; e.g.

`Ω(len(x)).Should(Not(Equal(4)))` => `Ω(x).ShouldNot(HaveLen(4))`

Or even (double negative):

`Ω(len(x)).To(Not(BeNumeric(">", 0)))` => `Ω(x).To(BeEmpty())`

The output of the linter,when finding issues, looks like this:
```
./testdata/src/a/a.go:14:5: ginkgo-linter: wrong length check; consider using `Expect("abcd").Should(HaveLen(4))` instead
./testdata/src/a/a.go:18:5: ginkgo-linter: wrong length check; consider using `Expect("").Should(BeEmpty())` instead
./testdata/src/a/a.go:22:5: ginkgo-linter: wrong length check; consider using `Expect("").Should(BeEmpty())` instead
```

## Suppress the linter
To suppress the wrong length check warning, add a comment with (only)
`ginkgo-linter:supressLengthCheckWarning`. There are two options to use this comment:
1. If the comment is at the top of the file, supress the warning for the whole file; e.g.:
   ```go
   package mypackage
   
   // ginkgo-linter:ignore-length-warning
   
   import (
       . "github.com/onsi/ginkgo/v2"
       . "github.com/onsi/gomega"
   )
   
   var _ = Describe("my test", func() {
        It("should do something", func() {
            Expect(len("abc")).Should(Equal(3)) // nothing in this file will trigger the warning
        })
   })
   ```
2. If the comment is before a wrong length check expression, the warning is suppressed for this expression only; for example:
   ```go
   It("should test something", func() {
       // ginkgo-linter:ignore-length-warning
       Expect(len("abc")).Should(Equal(3)) // this line will not trigger the warning
       Expect(len("abc")).Should(Equal(3)) // this line will trigger the warning
   }
   ```