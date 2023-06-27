package eventually

import (
	"context"
	"os"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type GetFunc[T any] func(ctx context.Context) (T, error)

func retFuncType() GetFunc[int] {
	return func(_ context.Context) (int, error) {
		return 1, nil
	}
}

func slowInt() int {
	time.Sleep(time.Second)
	return 42
}

func withArguments(a, b int) int {
	return 42
}

type test struct{}

func (test) slowInt() int {
	time.Sleep(time.Second)
	return 42
}

func (test) withArguments(a, b int) int {
	time.Sleep(time.Second)
	return 42
}

func retFunc(a int) func() int {
	return func() int {
		return a
	}
}

func retMethod(tst test) func() int {
	return tst.slowInt
}

func retChan() chan int {
	return make(chan int)
}

var _ = Describe("Eventually with function", func() {
	Context("Eventually", func() {
		It("should not trigger warning", func() {
			Eventually(retFuncType()).Should(Equal(1))
		})

		Eventually(retFunc(5)).Should(Equal(5)) // valid. retFunc returns a function
		ch := make(chan int)
		Eventually(ch).Should(BeClosed())                                                                                        // valid
		Eventually(slowInt).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42))                   // valid
		Eventually(slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42))                 // want `ginkgo-linter: use a function call in Eventually. This actually checks nothing, because Eventually receives the function returned value, instead of function itself, and this value is never changed`
		Eventually(context.TODO(), slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42)) // want `ginkgo-linter: use a function call in Eventually. This actually checks nothing, because Eventually receives the function returned value, instead of function itself, and this value is never changed`
		ctx := context.TODO()
		Eventually(ctx, slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42)) // want `ginkgo-linter: use a function call in Eventually. This actually checks nothing, because Eventually receives the function returned value, instead of function itself, and this value is never changed`
		Eventually(ctx, slowInt).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42))   // valid
		Eventually(ctx, func(g Gomega) {
			// make sure the existing rules are still applied within "Eventually"
			g.Expect(len("a")).Should(Equal(1)) // want `ginkgo-linter: wrong length assertion; consider using .g.Expect\("a"\)\.Should\(HaveLen\(1\)\). instead`
		}).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42)) // valid
		Eventually(os.Getwd).Should(Equal("something"))   // valid
		Eventually(os.Getwd()).Should(Equal("something")) // want `ginkgo-linter: use a function call in Eventually. This actually checks nothing, because Eventually receives the function returned value, instead of function itself, and this value is never changed`

		tst := test{}
		Eventually(tst.slowInt).Should(Equal(42))                           // valid
		Eventually(retMethod(tst)).Should(Equal(42))                        // valid. The function returns a method
		Eventually(tst.slowInt()).Should(Equal(42))                         // want `ginkgo-linter: use a function call in Eventually. This actually checks nothing, because Eventually receives the function returned value, instead of function itself, and this value is never changed`
		Eventually(withArguments(4, 5)).Should(Equal(42))                   // want `ginkgo-linter: use a function call in Eventually. This actually checks nothing, because Eventually receives the function returned value, instead of function itself, and this value is never changed`
		Eventually(withArguments).WithArguments(4, 5).Should(Equal(42))     // valid
		Eventually(tst.withArguments).WithArguments(4, 5).Should(Equal(42)) // valid
		Eventually(tst.withArguments(4, 5)).Should(Equal(42))               // want `ginkgo-linter: use a function call in Eventually. This actually checks nothing, because Eventually receives the function returned value, instead of function itself, and this value is never changed; consider using .Eventually\(tst\.withArguments\)\.WithArguments\(4, 5\)\.Should\(Equal\(42\)\). instead`
		Eventually(func(a, b int) int {
			return withArguments(a, b)
		}).WithArguments(4, 5).Should(Equal(42)) // valid
		Eventually(retChan()).Should(BeClosed()) // valid. retChan returns a channel
	})
	Context("Two errors in one assertion", func() {
		Eventually(func() *test { return nil }()).Should(BeNil())
		Eventually(func() []int { return nil }()).Should(HaveLen(0))        // want `ginkgo-linter: use a function call in Eventually. This actually checks nothing, because Eventually receives the function returned value, instead of function itself, and this value is never changed; consider using .Eventually\(func\(\) \[\]int { return nil }\)\.Should\(BeEmpty\(\)\). instead`
		Eventually(func() []int { return nil }).Should(HaveLen(0))          // want `ginkgo-linter: wrong length assertion; consider using .Eventually\(func\(\) \[\]int { return nil }\).Should\(BeEmpty\(\)\). instead`
		Eventually(func() bool { return true }()).Should(Equal(true))       // want `ginkgo-linter: use a function call in Eventually. This actually checks nothing, because Eventually receives the function returned value, instead of function itself, and this value is never changed; consider using .Eventually\(func\(\) bool { return true }\).Should\(BeTrue\(\)\). instead`
		Eventually(func() bool { return true }).Should(Equal(true))         // want `ginkgo-linter: wrong boolean assertion; consider using .Eventually\(func\(\) bool { return true }\)\.Should\(BeTrue\(\)\). instead`
		Eventually(func() bool { return true }).Should(Not(Equal(false)))   // want `ginkgo-linter: wrong boolean assertion; consider using .Eventually\(func\(\) bool { return true }\)\.Should\(BeTrue\(\)\). instead`
		Eventually(func() bool { return true }()).Should(Not(Equal(false))) // want `ginkgo-linter: use a function call in Eventually. This actually checks nothing, because Eventually receives the function returned value, instead of function itself, and this value is never changed; consider using .Eventually\(func\(\) bool { return true }\)\.Should\(BeTrue\(\)\). instead`
	})

	Context("EventuallyWithOffset", func() {
		ch := make(chan int)
		EventuallyWithOffset(1, ch).Should(BeClosed())                                                                                        // valid
		EventuallyWithOffset(1, slowInt).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42))                   // valid
		EventuallyWithOffset(1, slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42))                 // want `ginkgo-linter: use a function call in EventuallyWithOffset. This actually checks nothing, because EventuallyWithOffset receives the function returned value, instead of function itself, and this value is never changed`
		EventuallyWithOffset(1, context.TODO(), slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42)) // want `ginkgo-linter: use a function call in EventuallyWithOffset. This actually checks nothing, because EventuallyWithOffset receives the function returned value, instead of function itself, and this value is never changed`
		ctx := context.TODO()
		EventuallyWithOffset(1, ctx, slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42)) // want `ginkgo-linter: use a function call in EventuallyWithOffset. This actually checks nothing, because EventuallyWithOffset receives the function returned value, instead of function itself, and this value is never changed`
		EventuallyWithOffset(1, ctx, slowInt).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42))   // valid
		EventuallyWithOffset(1, ctx, func(g Gomega) {
			// make sure the existing rules are still applied within "EventuallyWithOffset"
			g.Expect(len("a")).Should(Equal(1)) // want `ginkgo-linter: wrong length assertion; consider using .g.Expect\("a"\)\.Should\(HaveLen\(1\)\). instead`
		}).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42)) // valid
	})

	Context("Consistently", func() {
		ch := make(chan int)
		Consistently(ch).Should(BeClosed())                                                                                        // valid
		Consistently(slowInt).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42))                   // valid
		Consistently(slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42))                 // want `ginkgo-linter: use a function call in Consistently. This actually checks nothing, because Consistently receives the function returned value, instead of function itself, and this value is never changed`
		Consistently(context.TODO(), slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42)) // want `ginkgo-linter: use a function call in Consistently. This actually checks nothing, because Consistently receives the function returned value, instead of function itself, and this value is never changed`
		ctx := context.TODO()
		Consistently(ctx, slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42)) // want `ginkgo-linter: use a function call in Consistently. This actually checks nothing, because Consistently receives the function returned value, instead of function itself, and this value is never changed`
		Consistently(ctx, slowInt).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42))   // valid
		Consistently(ctx, func(g Gomega) {
			// make sure the existing rules are still applied within "Consistently"
			g.Expect(len("a")).Should(Equal(1)) // want `ginkgo-linter: wrong length assertion; consider using .g.Expect\("a"\)\.Should\(HaveLen\(1\)\). instead`
		}).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42)) // valid
	})

	Context("ConsistentlyWithOffset", func() {
		ch := make(chan int)
		ConsistentlyWithOffset(1, ch).Should(BeClosed())                                                                                        // valid
		ConsistentlyWithOffset(1, slowInt).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42))                   // valid
		ConsistentlyWithOffset(1, slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42))                 // want `ginkgo-linter: use a function call in ConsistentlyWithOffset. This actually checks nothing, because ConsistentlyWithOffset receives the function returned value, instead of function itself, and this value is never changed`
		ConsistentlyWithOffset(1, context.TODO(), slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42)) // want `ginkgo-linter: use a function call in ConsistentlyWithOffset. This actually checks nothing, because ConsistentlyWithOffset receives the function returned value, instead of function itself, and this value is never changed`
		ctx := context.TODO()
		ConsistentlyWithOffset(1, ctx, slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42)) // want `ginkgo-linter: use a function call in ConsistentlyWithOffset. This actually checks nothing, because ConsistentlyWithOffset receives the function returned value, instead of function itself, and this value is never changed`
		ConsistentlyWithOffset(1, ctx, slowInt).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42))   // valid
		ConsistentlyWithOffset(1, ctx, func(g Gomega) {
			// make sure the existing rules are still applied within "ConsistentlyWithOffset"
			g.Expect(len("a")).Should(Equal(1)) // want `ginkgo-linter: wrong length assertion; consider using .g.Expect\("a"\)\.Should\(HaveLen\(1\)\). instead`
		}).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42)) // valid
	})
})
