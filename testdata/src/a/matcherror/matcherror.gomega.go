package matcherror

import (
	"errors"
	"fmt"
	"github.com/onsi/gomega/types"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("Check MatchError", func() {
	e1 := e("example")
	e2 := e("example")

	It("one argument - valid", func() {
		gomega.Expect(e1).To(gomega.MatchError(e2))
		gomega.Expect(errors.New("example")).To(gomega.MatchError("example"))
		gomega.Expect(e1).To(gomega.MatchError("example"))
		gomega.Expect(e1).To(gomega.MatchError(gomega.ContainSubstring("amp")))
		gomega.Expect(e1).To(gomega.MatchError(gomega.Not(gomega.ContainSubstring("abc"))))
		gomega.Expect(e1).To(gomega.MatchError(CustomMatcher()))
		gomega.Expect(e1).To(gomega.MatchError(CustomMatcherP()))
		gomega.Expect(e1).To(gomega.MatchError(customMatcher{}))
		gomega.Expect(e1).To(gomega.MatchError(&customMatcher{}))

		var gm types.GomegaMatcher = customMatcher{}
		gomega.Expect(e1).To(gomega.MatchError(gm))
	})

	It("one argument - error", func() {
		gomega.Expect("example").To(gomega.MatchError("example")) // want `ginkgo-linter: the MatchError matcher used to assert a non error type \("example"\)`
		one := 1
		gomega.Expect(one).To(gomega.MatchError("example")) // want `ginkgo-linter: the MatchError matcher used to assert a non error type \(one\)`
		gomega.Expect(e1).To(gomega.MatchError(isExample))  // want `ginkgo-linter: missing function description as second parameter of MatchError`
		gomega.Expect(e1).To(gomega.MatchError(f1))         // want `ginkgo-linter: multiple issues: MatchError first parameter \(f1\) must be error, string, GomegaMatcher or func\(error\)bool are allowed; redundant MatchError arguments; consider removing them`
	})

	It("two arguments - valid", func() {
		gomega.Expect(e1).To(gomega.MatchError(func(_ error) bool {
			return true
		}, "a description"))
		gomega.Expect(e1).ToNot(gomega.MatchError(func(_ error) bool {
			return false
		}, "a description"))
		gomega.Expect(e1).To(gomega.MatchError(func(err error) bool {
			err.Error()
			return true
		}, "a description"))
		gomega.Expect(e1).ToNot(gomega.MatchError(func(err error) bool {
			err.Error()
			return false
		}, "a description"))
		gomega.Expect(e1).To(gomega.MatchError(isExample, "is it the example error"))
		gomega.Expect(e1).To(gomega.MatchError(isExample, fmt.Sprint("is it the example error")))
	})

	It("two arguments - wrong usage", func() {
		gomega.Expect(e1).To(gomega.MatchError(e2, "not used"), "this text is ok")                             // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		gomega.Expect(e1).To(gomega.MatchError("example", "not used"), "this text is ok")                      // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		gomega.Expect(e1).To(gomega.MatchError(gomega.ContainSubstring("amp"), "not used"), "this text is ok") // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
	})

	It("two arguments - invalid", func() {
		gomega.Expect(e1).To(gomega.MatchError(1, "not used"), "wrong type - int")         // want `ginkgo-linter: multiple issues: MatchError first parameter \(1\) must be error, string, GomegaMatcher or func\(error\)bool are allowed; redundant MatchError arguments; consider removing them`
		gomega.Expect(e1).To(gomega.MatchError(compareString, "wrong function signature")) // want `ginkgo-linter: multiple issues: MatchError first parameter \(compareString\) must be error, string, GomegaMatcher or func\(error\)bool are allowed; redundant MatchError arguments; consider removing them`
	})

	It("two arguments. Second is not a string", func() {
		gomega.Expect(e1).To(gomega.MatchError(isExample, 1))                         // want `ginkgo-linter: The second parameter of MatchError must be the function description \(string\)`
		gomega.Expect(e1).To(gomega.MatchError(isExample, func() int { return 1 }))   // want `ginkgo-linter: The second parameter of MatchError must be the function description \(string\)`
		gomega.Expect(e1).To(gomega.MatchError(isExample, func() int { return 1 }())) // want `ginkgo-linter: The second parameter of MatchError must be the function description \(string\)`
	})

	It("multiple arguments", func() {
		gomega.Expect(e1).To(gomega.MatchError(isExample, "is it the example error", "not used"))             // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		gomega.Expect(e1).To(gomega.MatchError(isExample, "is it the example error", "not used", "not used")) // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		gomega.Expect(e1).To(gomega.MatchError(e2, "not used", "not used"))                                   // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		gomega.Expect(e1).To(gomega.MatchError("example", "not used", "not used"))                            // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		gomega.Expect(e1).To(gomega.MatchError(gomega.ContainSubstring("amp"), "not used", "not used"))       // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
	})

	It("funcs - valid", func() {
		gomega.Expect(f1()).To(gomega.MatchError(errors.New("example")))
		gomega.Expect(f1()).To(gomega.MatchError("example"))
		gomega.Expect(f1()).To(gomega.MatchError(gomega.ContainSubstring("amp")))
	})

	It("funcs - invalid", func() {
		gomega.Expect(f2()).To(gomega.MatchError(errors.New("example")))          // want `ginkgo-linter: the MatchError matcher used to assert a non error type \(f2\(\)\)`
		gomega.Expect(f2()).To(gomega.MatchError("example"))                      // want `ginkgo-linter: the MatchError matcher used to assert a non error type \(f2\(\)\)`
		gomega.Expect(f2()).To(gomega.MatchError(gomega.ContainSubstring("amp"))) // want `ginkgo-linter: the MatchError matcher used to assert a non error type \(f2\(\)\)`
		gomega.Expect(f3()).To(gomega.MatchError(errors.New("example")))          // want `ginkgo-linter: the MatchError matcher used to assert a non error type \(f3\(\)\)`
		gomega.Expect(f3()).To(gomega.MatchError("example"))                      // want `ginkgo-linter: the MatchError matcher used to assert a non error type \(f3\(\)\)`
		gomega.Expect(f3()).To(gomega.MatchError(gomega.ContainSubstring("amp"))) // want `ginkgo-linter: the MatchError matcher used to assert a non error type \(f3\(\)\)`
	})

	It("and - invalid", func() {
		var err error = e1
		gomega.Expect(err).To(gomega.And(gomega.Not(gomega.BeNil()), gomega.MatchError("expect", "not used")))        // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		gomega.Expect(err).To(gomega.And(gomega.MatchError("expect"), gomega.MatchError(e2)))                         // valid
		gomega.Expect(err).To(gomega.And(gomega.MatchError("expect", "not used"), gomega.MatchError(e2, "not used"))) // want `ginkgo-linter: multiple issues: redundant MatchError arguments; consider removing them; redundant MatchError arguments; consider removing them`
		gomega.Expect(err).To(gomega.And(gomega.MatchError("expect", "not used"), gomega.MatchError(e2)))             // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		gomega.Expect(err).To(gomega.And(gomega.MatchError("expect"), gomega.MatchError(e2, "not used")))             // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
	})

	It("or - invalid", func() {
		var err error = e1
		gomega.Expect(err).To(gomega.Or(gomega.Not(gomega.BeNil()), gomega.MatchError("expect", "not used")))        // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		gomega.Expect(err).To(gomega.Or(gomega.MatchError("expect"), gomega.MatchError(e2)))                         // valid
		gomega.Expect(err).To(gomega.Or(gomega.MatchError("expect", "not used"), gomega.MatchError(e2, "not used"))) // want `ginkgo-linter: multiple issues: redundant MatchError arguments; consider removing them; redundant MatchError arguments; consider removing them`
		gomega.Expect(err).To(gomega.Or(gomega.MatchError("expect", "not used"), gomega.MatchError(e2)))             // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
		gomega.Expect(err).To(gomega.Or(gomega.MatchError("expect"), gomega.MatchError(e2, "not used")))             // want `ginkgo-linter: redundant MatchError arguments; consider removing them`
	})

})
