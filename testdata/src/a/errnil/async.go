package errnil

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type myError string

func (e myError) Error() string {
	return string(e)
}

func f1() error {
	return errors.New("example")
}

var _ = Describe("async and error", func() {
	It("should report wrong error assertion in async functions", func() {
		ctx := context.Background()
		Eventually(func() error { return nil }).Should(BeNil())                                            // want `wrong error assertion. Consider using .Eventually\(func\(\) error { return nil }\)\.Should\(Succeed\(\)\). instead`
		Eventually(errFunc).Should(BeNil())                                                                // want `wrong error assertion. Consider using .Eventually\(errFunc\)\.Should\(Succeed\(\)\). instead`
		Eventually(func() error { return nil }).Should(Equal(nil))                                         // want `wrong error assertion. Consider using .Eventually\(func\(\) error { return nil }\)\.Should\(Succeed\(\)\). instead`
		Eventually(errFunc).Should(Equal(nil))                                                             // want `wrong error assertion. Consider using .Eventually\(errFunc\)\.Should\(Succeed\(\)\). instead`
		Eventually(func() error { return myError("fake error") }).ShouldNot(BeNil())                       // want `wrong error assertion. Consider using .Eventually\(func\(\) error { return myError\("fake error"\) }\)\.ShouldNot\(Succeed\(\)\). instead`
		Eventually(ctx, func() error { return myError("fake error") }).ShouldNot(Equal(nil))               // want `wrong error assertion. Consider using .Eventually\(ctx, func\(\) error { return myError\("fake error"\) }\)\.ShouldNot\(Succeed\(\)\). instead`
		Eventually(ctx, func() error { return myError("fake error") }).ShouldNot(MatchError("fake error")) // valid
		Eventually(func() error { return myError("fake error") }).ShouldNot(MatchError(f1))                // want `MatchError first parameter \(f1\) must be error, string, GomegaMatcher or func\(error\)bool are allowed`
		Eventually(func() string { return "fake error" }).ShouldNot(MatchError(Equal("fake error")))       // want `the MatchError matcher used to assert a non error type`
	})
	It("should report wrong error assertion in async functions", func() {
		Consistently(func() error { return nil }).Should(BeNil())   // want `wrong error assertion. Consider using .Consistently\(func\(\) error { return nil }\)\.Should\(Succeed\(\)\). instead`
		EventuallyWithOffset(1, errFunc).Should(BeNil())            // want `wrong error assertion. Consider using .EventuallyWithOffset\(1, errFunc\)\.Should\(Succeed\(\)\). instead`
		ConsistentlyWithOffset(1, errFunc).Should(BeNil())          // want `wrong error assertion. Consider using .ConsistentlyWithOffset\(1, errFunc\)\.Should\(Succeed\(\)\). instead`
		Consistently(context.Background(), errFunc).Should(BeNil()) // want `wrong error assertion. Consider using .Consistently\(context.Background\(\), errFunc\)\.Should\(Succeed\(\)\). instead`
		ctx := context.Background()
		ConsistentlyWithOffset(1, ctx, errFunc).Should(BeNil()) // want `wrong error assertion. Consider using .ConsistentlyWithOffset\(1, ctx, errFunc\)\.Should\(Succeed\(\)\). instead`
	})
})
