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

The linter supports the `Expect`, `ExpectWithOffset` and the `Ω` function, and the `Should`, `ShouldNot`, `To`, `ToNot` and `NotTo` assertion functions.

It also supports the embedded `Not()` function; e.g.

`Ω(len(x)).Should(Not(Equal(4)))` => `Ω(x).ShouldNot(HaveLen(4))`

Or even (double negative):

`Ω(len(x)).To(Not(BeNumeric(">", 0)))` => `Ω(x).To(BeEmpty())`
