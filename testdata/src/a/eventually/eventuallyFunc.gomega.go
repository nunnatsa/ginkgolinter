package eventually

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("Eventually with function", func() {
	Context("Eventually", func() {
		ch := make(chan int)
		gomega.Eventually(ch).Should(gomega.BeClosed())                                                                                        // valid
		gomega.Eventually(slowInt).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42))                   // valid
		gomega.Eventually(slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42))                 // want `ginkgo-linter: use a function call in Eventually. This actually checks nothing, because Eventually receives the function returned value, instead of function itself, and this value is never changed`
		gomega.Eventually(context.TODO(), slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42)) // want `ginkgo-linter: use a function call in Eventually. This actually checks nothing, because Eventually receives the function returned value, instead of function itself, and this value is never changed`
		ctx := context.TODO()
		gomega.Eventually(ctx, slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42)) // want `ginkgo-linter: use a function call in Eventually. This actually checks nothing, because Eventually receives the function returned value, instead of function itself, and this value is never changed`
		gomega.Eventually(ctx, slowInt).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42))   // valid
		gomega.Eventually(ctx, func(g gomega.Gomega) {                                                                              // valid
			// make sure the existing rules are still applied within "Eventually"
			g.Expect(len("a")).Should(gomega.Equal(1)) // want `ginkgo-linter: wrong length assertion; consider using .g.Expect\("a"\)\.Should\(gomega\.HaveLen\(1\)\). instead`
		}).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42)) // valid
	})

	Context("EventuallyWithOffset", func() {
		ch := make(chan int)
		gomega.EventuallyWithOffset(1, ch).Should(gomega.BeClosed())                                                                                        // valid
		gomega.EventuallyWithOffset(1, slowInt).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42))                   // valid
		gomega.EventuallyWithOffset(1, slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42))                 // want `ginkgo-linter: use a function call in EventuallyWithOffset. This actually checks nothing, because EventuallyWithOffset receives the function returned value, instead of function itself, and this value is never changed`
		gomega.EventuallyWithOffset(1, context.TODO(), slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42)) // want `ginkgo-linter: use a function call in EventuallyWithOffset. This actually checks nothing, because EventuallyWithOffset receives the function returned value, instead of function itself, and this value is never changed`
		ctx := context.TODO()
		gomega.EventuallyWithOffset(1, ctx, slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42)) // want `ginkgo-linter: use a function call in EventuallyWithOffset. This actually checks nothing, because EventuallyWithOffset receives the function returned value, instead of function itself, and this value is never changed`
		gomega.EventuallyWithOffset(1, ctx, slowInt).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42))   // valid
		gomega.EventuallyWithOffset(1, ctx, func(g gomega.Gomega) {                                                                              // valid
			// make sure the existing rules are still applied within "EventuallyWithOffset"
			g.Expect(len("a")).Should(gomega.Equal(1)) // want `ginkgo-linter: wrong length assertion; consider using .g.Expect\("a"\)\.Should\(gomega\.HaveLen\(1\)\). instead`
		}).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42)) // valid
	})

	Context("Consistently", func() {
		ch := make(chan int)
		gomega.Consistently(ch).Should(gomega.BeClosed())                                                                                        // valid
		gomega.Consistently(slowInt).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42))                   // valid
		gomega.Consistently(slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42))                 // want `ginkgo-linter: use a function call in Consistently. This actually checks nothing, because Consistently receives the function returned value, instead of function itself, and this value is never changed`
		gomega.Consistently(context.TODO(), slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42)) // want `ginkgo-linter: use a function call in Consistently. This actually checks nothing, because Consistently receives the function returned value, instead of function itself, and this value is never changed`
		ctx := context.TODO()
		gomega.Consistently(ctx, slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42)) // want `ginkgo-linter: use a function call in Consistently. This actually checks nothing, because Consistently receives the function returned value, instead of function itself, and this value is never changed`
		gomega.Consistently(ctx, slowInt).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42))   // valid
		gomega.Consistently(ctx, func(g gomega.Gomega) {                                                                              // valid
			// make sure the existing rules are still applied within "Consistently"
			g.Expect(len("a")).Should(gomega.Equal(1)) // want `ginkgo-linter: wrong length assertion; consider using .g.Expect\("a"\)\.Should\(gomega\.HaveLen\(1\)\). instead`
		}).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42)) // valid
	})

	Context("ConsistentlyWithOffset", func() {
		ch := make(chan int)
		gomega.ConsistentlyWithOffset(1, ch).Should(gomega.BeClosed())                                                                                        // valid
		gomega.ConsistentlyWithOffset(1, slowInt).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42))                   // valid
		gomega.ConsistentlyWithOffset(1, slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42))                 // want `ginkgo-linter: use a function call in ConsistentlyWithOffset. This actually checks nothing, because ConsistentlyWithOffset receives the function returned value, instead of function itself, and this value is never changed`
		gomega.ConsistentlyWithOffset(1, context.TODO(), slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42)) // want `ginkgo-linter: use a function call in ConsistentlyWithOffset. This actually checks nothing, because ConsistentlyWithOffset receives the function returned value, instead of function itself, and this value is never changed`
		ctx := context.TODO()
		gomega.ConsistentlyWithOffset(1, ctx, slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42)) // want `ginkgo-linter: use a function call in ConsistentlyWithOffset. This actually checks nothing, because ConsistentlyWithOffset receives the function returned value, instead of function itself, and this value is never changed`
		gomega.ConsistentlyWithOffset(1, ctx, slowInt).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42))   // valid
		gomega.ConsistentlyWithOffset(1, ctx, func(g gomega.Gomega) {                                                                              // valid
			// make sure the existing rules are still applied within "ConsistentlyWithOffset"
			g.Expect(len("a")).Should(gomega.Equal(1)) // want `ginkgo-linter: wrong length assertion; consider using .g.Expect\("a"\)\.Should\(gomega\.HaveLen\(1\)\). instead`
		}).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(gomega.Equal(42)) // valid
	})
})
