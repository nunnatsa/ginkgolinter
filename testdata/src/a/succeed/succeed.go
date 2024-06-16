package succeed

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func safeDivide(num int, denom int) (int, error) {
	if denom == 0 {
		return 0, errors.New("divide by zero")
	}
	return num / denom, nil
}

func attemptToCallNetwork(shouldTry bool) error {
	if !shouldTry {
		return nil
	}
	return errors.New("don't know how to connect to network")
}

var _ = Describe("Ensure succeed assertion is used correctly", Label("succeed"), func() {
	Expect(safeDivide(0, 0)).To(Not(Succeed())) // want `ginkgo-linter: Succeed matcher should only be asserted against a function call, and the function must return exactly one error`
	Expect(safeDivide(0, 0)).NotTo(Succeed())   // want `ginkgo-linter: Succeed matcher should only be asserted against a function call, and the function must return exactly one error`
	Expect(safeDivide(0, 1)).To(Succeed())      // want `ginkgo-linter: Succeed matcher should only be asserted against a function call, and the function must return exactly one error`

	_, err0 := safeDivide(0, 0)
	_, err1 := safeDivide(0, 1)

	Expect(err0).To(Not(Succeed())) // want `ginkgo-linter: Succeed matcher should only be asserted against a function call, and the function must return exactly one error`
	Expect(err0).NotTo(Succeed())   // want `ginkgo-linter: Succeed matcher should only be asserted against a function call, and the function must return exactly one error`
	Expect(err1).To(Succeed())      // want `ginkgo-linter: Succeed matcher should only be asserted against a function call, and the function must return exactly one error`

	Expect(attemptToCallNetwork(false)).To(Succeed())
	Expect(attemptToCallNetwork(true)).To(Not(Succeed()))
	Expect(attemptToCallNetwork(false)).NotTo(Succeed())

	Eventually(func() error {
		return attemptToCallNetwork(false)
	}).Should(Succeed())

	Eventually(func() error {
		return attemptToCallNetwork(false)
	}).Should(Succeed())

	Eventually(func() (int, error) { // want `ginkgo-linter: Succeed matcher must only be asserted against a function that returns exactly one error`
		return safeDivide(0, 1)
	}).Should(Succeed())

	Consistently(func() (int, error) { // want `ginkgo-linter: Succeed matcher must only be asserted against a function that returns exactly one error`
		return safeDivide(0, 0)
	}).ShouldNot(Succeed())

	Consistently(func() (int, error) { // want `ginkgo-linter: Succeed matcher must only be asserted against a function that returns exactly one error`
		return safeDivide(0, 0)
	}).ShouldNot(Not(Succeed()))
})
