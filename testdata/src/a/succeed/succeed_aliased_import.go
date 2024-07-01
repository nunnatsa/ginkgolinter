package succeed

import (
	. "github.com/onsi/ginkgo/v2"
	gom "github.com/onsi/gomega"
)

var _ = Describe("Ensure succeed assertion is used correctly", Label("succeed"), func() {
	gom.Expect(returnsFloatAndErr(0)).To(gom.Not(gom.Succeed())) // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`
	gom.Expect(returnsFloatAndErr(0)).NotTo(gom.Succeed())       // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`
	gom.Expect(returnsFloatAndErr(1)).To(gom.Succeed())          // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`

	value0, err0 := returnsFloatAndErr(0)
	value1, _ := returnsFloatAndErr(1)

	gom.Expect(value0, err0).ToNot(gom.Succeed()) // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`
	gom.Expect(value1).To(gom.Succeed())          // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`

	resp := struct {
		response any
		err      error
	}{
		response: nil,
		err:      err0,
	}

	gom.Expect(err0).ToNot(gom.Succeed())              // want `ginkgo-linter: Succeed matcher should be asserted against a function call; consider replacing with gom.Expect\(err0\).To\(gom.HaveOccurred\(\)\)`
	gom.ExpectWithOffset(1, err0).NotTo(gom.Succeed()) // want `ginkgo-linter: Succeed matcher should be asserted against a function call; consider replacing with gom.ExpectWithOffset\(1, err0\).To\(gom.HaveOccurred\(\)\)`
	gom.Expect(err0).NotTo(gom.Not(gom.Succeed()))     // want `ginkgo-linter: Succeed matcher should be asserted against a function call; consider replacing with gom.Expect\(err0\).To\(gom.Not\(gom.HaveOccurred\(\)\)\)`
	gom.Expect(resp.err).To(gom.Succeed())             // want `ginkgo-linter: Succeed matcher should be asserted against a function call; consider replacing with gom.Expect\(resp.err\).ToNot\(gom.HaveOccurred\(\)\)`
	// ginkgo-linter:ignore-succeed-warning
	gom.Expect(err0).ToNot(gom.Succeed())

	gom.Expect(returnsErr(false)).To(gom.Succeed())
	gom.Expect(returnsErr(true)).To(gom.Not(gom.Succeed()))
	gom.Expect(returnsErr(false)).NotTo(gom.Succeed())

	gom.Eventually(func() error {
		return returnsErr(false)
	}).Should(gom.Succeed())

	gom.ConsistentlyWithOffset(2, func() error {
		return returnsErr(false)
	}).Should(gom.Succeed())

	gom.Eventually(func() (float64, error) { // want `ginkgo-linter: Succeed matcher must be asserted against a function that returns exactly one error or a function which takes in a single gomega.Gomega argument and returns nothing`
		return returnsFloatAndErr(1)
	}).Should(gom.Succeed())

	gom.Consistently(func() (float64, error) { // want `ginkgo-linter: Succeed matcher must be asserted against a function that returns exactly one error or a function which takes in a single gomega.Gomega argument and returns nothing`
		return returnsFloatAndErr(0)
	}).ShouldNot(gom.Succeed())

	gom.Consistently(func() (float64, error) { // want `ginkgo-linter: Succeed matcher must be asserted against a function that returns exactly one error or a function which takes in a single gomega.Gomega argument and returns nothing`
		return returnsFloatAndErr(0)
	}).ShouldNot(gom.Not(gom.Succeed()))

	gom.Eventually(func(g gom.Gomega) {
		err := returnsErr(false)
		g.Expect(err).ToNot(gom.HaveOccurred())
	}).Should(gom.Not(gom.Succeed()))

	gom.Consistently(func(g gom.Gomega) {
		g.Expect(returnsFloatAndErr(1)).ToNot(gom.Succeed()) // want `ginkgo-linter: Succeed matcher must be asserted against exactly one error value or a function call returning the same, or it will always fail`
	}).Should(gom.Succeed())
})
