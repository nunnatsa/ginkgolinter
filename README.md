# ginkgo-linter

This is a golang linter to check usage of the ginkgo and gomega packages.

ginkgo is a testing framework and gomega is its assertion package.

## usage
```shell
ginkgo-linter [-fix] .
```
Use the `-fix` flag to apply the fix suggestions to the source code.

## Linter Checks
The linter checks the `Expect`, `ExpectWithOffset` and the `Ω` "actual" functions, with the `Should`, `ShouldNot`, `To`, `ToNot` and `NotTo` assertion functions.

It also supports the embedded `Not()` matcher

### Wrong Length Assertion
The linter finds assertion of the golang built-in `len` function, with all kind of matchers, while there are already gomega matchers for these usecases; We want to assert the item, rather than its length.

There are several wrong patterns:
```go
Expect(len(x)).To(Equal(0)) // should be: Expect(x).To(BeEmpty())
Expect(len(x)).To(BeZero()) // should be: Expect(x).To(BeEmpty())
Expect(len(x)).To(BeNumeric(">", 0)) // should be: Expect(x).ToNot(BeEmpty())
Expect(len(x)).To(BeNumeric(">=", 1)) // should be: Expect(x).ToNot(BeEmpty())
Expect(len(x)).To(BeNumeric("==", 0)) // should be: Expect(x).To(BeEmpty())
Expect(len(x)).To(BeNumeric("!=", 0)) // should be: Expect(x).ToNot(BeEmpty())

Expect(len(x)).To(Equal(1)) // should be: Expect(x).To(HaveLen(1))
Expect(len(x)).To(BeNumeric("==", 2)) // should be: Expect(x).To(HaveLen(2))
Expect(len(x)).To(BeNumeric("!=", 3)) // should be: Expect(x).ToNot(HaveLen(3))
```

It also supports the embedded `Not()` matcher; e.g.

`Ω(len(x)).Should(Not(Equal(4)))` => `Ω(x).ShouldNot(HaveLen(4))`

Or even (double negative):

`Ω(len(x)).To(Not(BeNumeric(">", 0)))` => `Ω(x).To(BeEmpty())`

The output of the linter,when finding issues, looks like this:
```
./testdata/src/a/a.go:14:5: ginkgo-linter: wrong length assertion; consider using `Expect("abcd").Should(HaveLen(4))` instead
./testdata/src/a/a.go:18:5: ginkgo-linter: wrong length assertion; consider using `Expect("").Should(BeEmpty())` instead
./testdata/src/a/a.go:22:5: ginkgo-linter: wrong length assertion; consider using `Expect("").Should(BeEmpty())` instead
```

### Wrong `nil` Assertion
The linter finds assertion of the comparison to nil, with all kind of matchers, instead of using the existing `BeNil()` matcher; We want to assert the item, rather than a comparison result.

There are several wrong patterns:

```go
Expect(x == nil).To(Equal(true)) // should be: Expect(x).To(BeNil())
Expect(nil == x).To(Equal(true)) // should be: Expect(x).To(BeNil())
Expect(x != nil).To(Equal(true)) // should be: Expect(x).ToNot(BeNil())
Expect(nil != nil).To(Equal(true)) // should be: Expect(x).ToNot(BeNil())

Expect(x == nil).To(BeTrue()) // should be: Expect(x).To(BeNil())
Expect(x == nil).To(BeFalse()) // should be: Expect(x).ToNot(BeNil())
```
It also supports the embedded `Not()` matcher; e.g.

`Ω(x == nil).Should(Not(BeTrue()))` => `Ω(x).ShouldNot(BeNil())`

Or even (double negative):

`Ω(x != nil).Should(Not(BeTrue()))` => `Ω(x).Should(BeNil())`

## Suppress the linter
### Suppress warning from command line
* Use the `suppress-len-assertion=true` flag to suppress the wrong length assertion warning
* Use the `suppress-nil-assertion=true` flag to suppress the wrong nil assertion warning

### Suppress warning from the code
To suppress the wrong length assertion warning, add a comment with (only)

`ginkgo-linter:ignore-len-assert-warning`. 

To suppress the wrong nil assertion warning, add a comment with (only)

`ginkgo-linter:ignore-nil-assert-warning`. 

There are two options to use these comments:
1. If the comment is at the top of the file, supress the warning for the whole file; e.g.:
   ```go
   package mypackage
   
   // ginkgo-linter:ignore-len-assert-warning
   
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
   ```golang
   It("should test something", func() {
       // ginkgo-linter:ignore-nil-assert-warning
       Expect(x == nil).Should(BeTrue()) // this line will not trigger the warning
       Expect(x == nil).Should(BeTrue()) // this line will trigger the warning
   }
   ```
