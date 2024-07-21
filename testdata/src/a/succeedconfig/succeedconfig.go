package succeedconfig

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func returnsFloatAndErr(arg int) (float64, error) {
	if arg == 0 {
		return 0, errors.New("an error")
	}
	return 1.0 / float64(arg), nil
}

func returnsErr(arg bool) error {
	if !arg {
		return nil
	}
	return errors.New("an error")
}

var _ = Describe("Ensure succeed assertion is used correctly", Label("succeed"), func() {
	Expect(returnsFloatAndErr(0)).To(Not(Succeed())) // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`
	Expect(returnsFloatAndErr(0)).NotTo(Succeed())   // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`
	Expect(returnsFloatAndErr(1)).To(Succeed())      // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`

	value0, err0 := returnsFloatAndErr(0)
	value1, _ := returnsFloatAndErr(1)

	Expect(value0, err0).ToNot(Succeed()) // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`
	Expect(value1).To(Succeed())          // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`

	resp := struct {
		response any
		err      error
	}{
		response: nil,
		err:      err0,
	}

	Expect(err0).ToNot(Succeed())
	ExpectWithOffset(1, err0).NotTo(Succeed())
	Expect(err0).NotTo(Not(Succeed()))
	Expect(resp.err).To(Succeed())
	// ginkgo-linter:ignore-succeed-warning
	Expect(err0).ToNot(Succeed())

	Expect(returnsErr(false)).To(Succeed())
	Expect(returnsErr(true)).To(Not(Succeed()))
	Expect(returnsErr(false)).NotTo(Succeed())

	Eventually(func() error {
		return returnsErr(false)
	}).Should(Succeed())

	ConsistentlyWithOffset(2, func() error {
		return returnsErr(false)
	}).Should(Succeed())

	Eventually(func() (float64, error) { // want `ginkgo-linter: Succeed matcher must be asserted against a function that returns exactly one error or a function which takes in a single gomega.Gomega argument and returns nothing`
		return returnsFloatAndErr(1)
	}).Should(Succeed())

	Consistently(func() (float64, error) { // want `ginkgo-linter: Succeed matcher must be asserted against a function that returns exactly one error or a function which takes in a single gomega.Gomega argument and returns nothing`
		return returnsFloatAndErr(0)
	}).ShouldNot(Succeed())

	Consistently(func() (float64, error) { // want `ginkgo-linter: Succeed matcher must be asserted against a function that returns exactly one error or a function which takes in a single gomega.Gomega argument and returns nothing`
		return returnsFloatAndErr(0)
	}).ShouldNot(Not(Succeed()))

	Eventually(func(g Gomega) {
		err := returnsErr(false)
		g.Expect(err).ToNot(HaveOccurred())
	}).Should(Not(Succeed()))

	Consistently(func(g Gomega, ctx context.Context) {
		g.Expect(returnsFloatAndErr(1)).ToNot(Succeed()) // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`
	}).Should(Succeed())
})
