package timing

import (
	tm "time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func helperFunc(timeout int) {
	Eventually(func() bool { return true }, timeout, 5).Should(BeTrue()) // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
}

var _ = Describe("check async intervals", func() {
	const (
		timeout = tm.Second * 10
		factor  = 5
	)

	var interval = 60 * tm.Second

	Eventually(func() bool { return true }, 5, factor).Should(BeTrue())                                     // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
	Eventually(func() bool { return true }, 5, 1).Should(BeTrue())                                          // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
	Eventually(func() bool { return true }, tm.Second, interval).Should(BeTrue())                           // can't check variables
	Eventually(func() bool { return true }, tm.Second, tm.Millisecond*10).Should(BeTrue())                  // valid
	Eventually(func() bool { return true }, tm.Millisecond*10, tm.Second).Should(BeTrue())                  // want `timeout must not be shorter than the polling interval`
	Eventually(func() bool { return true }, tm.Second*(10+factor), tm.Second*60*60*24*356).Should(BeTrue()) // want `timeout must not be shorter than the polling interval`
	Eventually(func() bool { return true }, tm.Second*60*60*24*356, tm.Second*(10+factor)).Should(BeTrue()) // valid, many multiplication
	Eventually(func() bool { return true }, timeout, tm.Millisecond*9).Should(BeTrue())                     // const
	Eventually(func() bool { return true }, tm.Millisecond*9, timeout).Should(BeTrue())                     // want `timeout must not be shorter than the polling interval`
	Eventually(func() bool { return true }, 1+0, uint(2)).Should(BeTrue())                                  // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
	Eventually(func() bool { return true }, 1.1, float64(2)).Should(BeTrue())                               // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
	Eventually(func() bool { return true }, "3s", int32(2)).Should(BeTrue())                                // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
})
