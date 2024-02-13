package asyncconfig

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func slowInt() int {
	time.Sleep(time.Second)
	return 42
}

func slowBool() bool {
	time.Sleep(time.Second)
	return true
}

var _ = Describe("test async suppress configuration", func() {
	It("should not trigger warning", func() {
		// this is a wrong assertion. It will take 1 second to return from slowInt(), and only then, we'll start polling -
		// this will give a false positive result. But here the linter is configured to ignore this error.
		Eventually(slowInt()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(42))
	})
	It("should trigger comparison waring", func() {
		Eventually(slowBool()).WithPolling(time.Millisecond * 100).WithTimeout(time.Second * 2).Should(Equal(true)) // want `ginkgo-linter: wrong boolean assertion\. Consider using .Eventually\(slowBool\(\)\)\.WithPolling\(time\.Millisecond \* 100\)\.WithTimeout\(time\.Second \* 2\)\.Should\(BeTrue\(\)\). instead`
	})
})
