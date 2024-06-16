package succeedconfig

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
	Expect(safeDivide(0, 0)).To(Not(Succeed())) // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`
	Expect(safeDivide(0, 0)).NotTo(Succeed())   // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`
	Expect(safeDivide(0, 1)).To(Succeed())      // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`

	value0, err0 := safeDivide(0, 0)
	value1, _ := safeDivide(0, 1)

	Expect(value0, err0).ToNot(Succeed()) // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`
	Expect(value1).To(Succeed())          // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`

	wrap := struct {
		err error
	}{
		err: err0,
	}

	Expect(err0).ToNot(Succeed())
	ExpectWithOffset(1, err0).NotTo(Succeed())
	Expect(err0).NotTo(Not(Succeed()))
	Expect(wrap.err).To(Succeed())
	// ginkgo-linter:ignore-succeed-warning
	Expect(err0).ToNot(Succeed())

	Expect(attemptToCallNetwork(false)).To(Succeed())
	Expect(attemptToCallNetwork(true)).To(Not(Succeed()))
	Expect(attemptToCallNetwork(false)).NotTo(Succeed())

	Eventually(func() error {
		return attemptToCallNetwork(false)
	}).Should(Succeed())

	ConsistentlyWithOffset(1, func() error {
		return attemptToCallNetwork(false)
	}).Should(Succeed())

	Eventually(func() (int, error) { // want `ginkgo-linter: Succeed matcher must be asserted against a function that returns exactly one error`
		return safeDivide(0, 1)
	}).Should(Succeed())

	Consistently(func() (int, error) { // want `ginkgo-linter: Succeed matcher must be asserted against a function that returns exactly one error`
		return safeDivide(0, 0)
	}).ShouldNot(Succeed())

	Consistently(func() (int, error) { // want `ginkgo-linter: Succeed matcher must be asserted against a function that returns exactly one error`
		return safeDivide(0, 0)
	}).ShouldNot(Not(Succeed()))
})
