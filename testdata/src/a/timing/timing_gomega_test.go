package timing

import (
	tm "time"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("check async intervals", func() {
	const (
		timeout = tm.Second * 10
		factor  = 5
	)

	var interval = 60 * tm.Second
	gomega.Eventually(func() bool { return true }, 5, factor).Should(gomega.BeTrue())                                      // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\).*tm\.Second\*5, tm\.Second\*factor`
	gomega.Eventually(func() bool { return true }, 5, 1).Should(gomega.BeTrue())                                           // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\).*, tm\.Second\*5, tm\.Second`
	gomega.Eventually(func() bool { return true }, tm.Second, interval).Should(gomega.BeTrue())                            // can't check variables
	gomega.Eventually(func() bool { return true }, tm.Second, tm.Millisecond*10).Should(gomega.BeTrue())                   // valid
	gomega.Eventually(func() bool { return true }, tm.Millisecond*10, tm.Second).Should(gomega.BeTrue())                   // want `timeout must not be shorter than the polling interval`
	gomega.Eventually(func() bool { return true }, tm.Second*(10+factor), tm.Second*60*60*24*356).Should(gomega.BeTrue())  // want `timeout must not be shorter than the polling interval`
	gomega.Eventually(func() bool { return true }, tm.Second*60*60*24*356, tm.Second*(10+factor)).Should(gomega.BeTrue())  // valid, many multiplication
	gomega.Eventually(func() bool { return true }, timeout, tm.Millisecond*9).Should(gomega.BeTrue())                      // const
	gomega.Eventually(func() bool { return true }, tm.Millisecond*9, timeout).Should(gomega.BeTrue())                      // want `timeout must not be shorter than the polling interval`
	gomega.Eventually(func() bool { return true }, timeout, tm.Millisecond*9).WithPolling(timeout).Should(gomega.BeTrue()) // want `polling defined more than once`

	gomega.Eventually(func() bool { return true }, 1+0, uint(2)).Should(gomega.BeTrue())    // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
	gomega.Eventually(func() bool { return true }, 1.1, float64(2)).Should(gomega.BeTrue()) // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
	gomega.Eventually(func() bool { return true }, "3s", int32(2)).Should(gomega.BeTrue())  // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
})
