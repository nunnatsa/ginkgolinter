package matcherror

import (
	"errors"
	"fmt"

	"github.com/onsi/gomega/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type e string

func (err e) Error() string {
	return string(err)
}

func isExample(err error) bool {
	return err.Error() == "example"
}

func compareString(err error, s string) bool {
	return err.Error() == s
}

func f1() error {
	return errors.New("example")
}

func f2() string {
	return "example"
}

func f3() (string, error) {
	return "example", errors.New("example")
}

type customMatcher struct{}

func (customMatcher) Match(actual interface{}) (success bool, err error) {
	return true, nil
}
func (customMatcher) FailureMessage(actual interface{}) (message string) {
	return "FailureMessage"
}
func (customMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return "NegatedFailureMessage"
}

func CustomMatcher() customMatcher {
	return customMatcher{}
}

func CustomMatcherP() *customMatcher {
	return &customMatcher{}
}

var _ = Describe("Check MatchError", func() {
	e1 := e("example")
	e2 := e("example")

	It("one argument - valid", func() {
		Expect(e1).To(MatchError(e2))
		Expect(errors.New("example")).To(MatchError("example"))
		Expect(e1).To(MatchError("example"))
		Expect(e1).To(MatchError(ContainSubstring("amp")))
		Expect(e1).To(MatchError(Not(ContainSubstring("abc"))))
		Expect(e1).To(MatchError(CustomMatcher()))
		Expect(e1).To(MatchError(CustomMatcherP()))
		Expect(e1).To(MatchError(customMatcher{}))
		Expect(e1).To(MatchError(&customMatcher{}))

		var gm types.GomegaMatcher = customMatcher{}
		Expect(e1).To(MatchError(gm))
	})

	It("one argument - error", func() {
		Expect("example").To(MatchError("example")) // want `ginkgo-linter: the MatchError matcher used to assert a non error type \("example"\)`
		one := 1
		Expect(one).To(MatchError("example")) // want `ginkgo-linter: the MatchError matcher used to assert a non error type \(one\)`
		Expect(e1).To(MatchError(isExample))  // want `ginkgo-linter: missing function description as second parameter of MatchError`
		Expect(e1).To(MatchError(f1))         // want `ginkgo-linter: MatchError first parameter \(f1\) must be error, string, GomegaMatcher or func\(error\)bool are allowed`
	})

	It("two arguments - valid", func() {
		Expect(e1).To(MatchError(func(_ error) bool {
			return true
		}, "a description"))
		Expect(e1).ToNot(MatchError(func(_ error) bool {
			return false
		}, "a description"))
		Expect(e1).To(MatchError(func(err error) bool {
			err.Error()
			return true
		}, "a description"))
		Expect(e1).ToNot(MatchError(func(err error) bool {
			err.Error()
			return false
		}, "a description"))
		Expect(e1).To(MatchError(isExample, "is it the example error"))
		Expect(e1).To(MatchError(isExample, fmt.Sprint("is it the example error")))
	})

	It("two arguments - wrong usage", func() {
		Expect(e1).To(MatchError(e2, "not used"), "this text is ok")                      // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		Expect(e1).To(MatchError("example", "not used"), "this text is ok")               // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		Expect(e1).To(MatchError(ContainSubstring("amp"), "not used"), "this text is ok") // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
	})

	It("two arguments - invalid", func() {
		Expect(e1).To(MatchError(1, "not used"), "wrong type - int")         // want `ginkgo-linter: multiple issues: MatchError first parameter \(1\) must be error, string, GomegaMatcher or func\(error\)bool are allowed; redundant MatchError arguments; consider removing them`
		Expect(e1).To(MatchError(compareString, "wrong function signature")) // want `ginkgo-linter: multiple issues: MatchError first parameter \(compareString\) must be error, string, GomegaMatcher or func\(error\)bool are allowed; redundant MatchError arguments; consider removing them`
	})

	It("two arguments. Second is not a string", func() {
		Expect(e1).To(MatchError(isExample, 1))                         // want `ginkgo-linter: The second parameter of MatchError must be the function description \(string\)`
		Expect(e1).To(MatchError(isExample, func() int { return 1 }))   // want `ginkgo-linter: The second parameter of MatchError must be the function description \(string\)`
		Expect(e1).To(MatchError(isExample, func() int { return 1 }())) // want `ginkgo-linter: The second parameter of MatchError must be the function description \(string\)`
	})

	It("multiple arguments", func() {
		Expect(e1).To(MatchError(isExample, "is it the example error", "not used"))             // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		Expect(e1).To(MatchError(isExample, "is it the example error", "not used", "not used")) // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		Expect(e1).To(MatchError(e2, "not used", "not used"))                                   // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		Expect(e1).To(MatchError("example", "not used", "not used"))                            // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		Expect(e1).To(MatchError(ContainSubstring("amp"), "not used", "not used"))              // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
	})

	It("funcs - valid", func() {
		Expect(f1()).To(MatchError(errors.New("example")))
		Expect(f1()).To(MatchError("example"))
		Expect(f1()).To(MatchError(ContainSubstring("amp")))
	})

	It("funcs - invalid", func() {
		Expect(f2()).To(MatchError(errors.New("example")))   // want `ginkgo-linter: the MatchError matcher used to assert a non error type \(f2\(\)\)`
		Expect(f2()).To(MatchError("example"))               // want `ginkgo-linter: the MatchError matcher used to assert a non error type \(f2\(\)\)`
		Expect(f2()).To(MatchError(ContainSubstring("amp"))) // want `ginkgo-linter: the MatchError matcher used to assert a non error type \(f2\(\)\)`
		Expect(f3()).To(MatchError(errors.New("example")))   // want `ginkgo-linter: the MatchError matcher used to assert a non error type \(f3\(\)\)`
		Expect(f3()).To(MatchError("example"))               // want `ginkgo-linter: the MatchError matcher used to assert a non error type \(f3\(\)\)`
		Expect(f3()).To(MatchError(ContainSubstring("amp"))) // want `ginkgo-linter: the MatchError matcher used to assert a non error type \(f3\(\)\)`
	})

	It("and - invalid", func() {
		var err error = e1
		Expect(err).To(And(Not(BeNil()), MatchError("expect", "not used")))               // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		Expect(err).To(And(MatchError("expect"), MatchError(e2)))                         // valid
		Expect(err).To(And(MatchError("expect", "not used"), MatchError(e2, "not used"))) // want `ginkgo-linter: multiple issues: redundant MatchError arguments; consider removing them; redundant MatchError arguments; consider removing them`
		Expect(err).To(And(MatchError("expect", "not used"), MatchError(e2)))             // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		Expect(err).To(And(MatchError("expect"), MatchError(e2, "not used")))             // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
	})

	It("or - invalid", func() {
		var err error = e1
		Expect(err).To(Or(Not(BeNil()), MatchError("expect", "not used")))               // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		Expect(err).To(Or(MatchError("expect"), MatchError(e2)))                         // valid
		Expect(err).To(Or(MatchError("expect", "not used"), MatchError(e2, "not used"))) // want `ginkgo-linter: multiple issues: redundant MatchError arguments; consider removing them; redundant MatchError arguments; consider removing them`
		Expect(err).To(Or(MatchError("expect", "not used"), MatchError(e2)))             // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		Expect(err).To(Or(MatchError("expect"), MatchError(e2, "not used")))             // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
	})
})
