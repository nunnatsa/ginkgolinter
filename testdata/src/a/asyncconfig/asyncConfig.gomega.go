package asyncconfig

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	g "github.com/onsi/gomega"
)

var _ = Describe("test async suppress configuration", func() {
	It("should not trigger warning", func() {
		// this is a wrong assertion. It will take 1 second to return from slowInt(), and only then, we'll start polling -
		// this will give a false positive result. But here the linter is configured to ignore this error.
		g.Eventually(slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(g.Equal(42))
	})
	It("should trigger comparison waring", func() {
		g.Eventually(slowBool()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(g.Equal(true)) // want `ginkgo-linter: wrong boolean assertion\. Consider using .g\.Eventually\(slowBool\(\)\)\.WithPolling\(time\.Millisecond \* 100\)\.WithTimeout\(time\.Second \* 2\)\.Should\(g\.BeTrue\(\)\). instead`
	})
})
