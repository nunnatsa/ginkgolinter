package timing

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"a/timing/pkg"
)

var _ = Describe("check async intervals", func() {

	const (
		factor  = 5
		timeout = time.Second * 10
		polling = time.Millisecond * 500
	)

	It("timeout shorter than polling", func() {
		Eventually(func() bool { return true }, timeout, pkg.Timeout).Should(BeTrue())                                                  // valid
		Eventually(func() bool { return true }).WithTimeout(time.Second * 10).WithPolling(time.Second * (10 + factor)).Should(BeTrue()) // want `timeout must not be shorter than the polling interval`
		Eventually(func() bool { return true }).Within(time.Second * 10).WithPolling(time.Second * (10 + factor)).Should(BeTrue())      // want `timeout must not be shorter than the polling interval`
		Eventually(func() bool { return true }).Within(time.Second * 10).ProbeEvery(time.Second * (10 + factor)).Should(BeTrue())       // want `timeout must not be shorter than the polling interval`
		Eventually(func() bool { return true }).WithTimeout(time.Second * 10).ProbeEvery(time.Second * (10 + factor)).Should(BeTrue())  // want `timeout must not be shorter than the polling interval`
		Eventually(func() bool { return true }, time.Second*60*60*24*365, time.Second*(10+factor)).Should(BeTrue())                     // valid
		Eventually(func() bool { return true }, time.Second, time.Millisecond*10).Should(BeTrue())                                      //valid
		Eventually(func() bool { return true }, time.Millisecond*10, time.Second).Should(BeTrue())                                      // want `timeout must not be shorter than the polling interval`
		Eventually(func() bool { return true }, timeout, polling).Should(BeTrue())
		Eventually(func() bool { return true }, polling, timeout).Should(BeTrue()) // want `timeout must not be shorter than the polling interval`
		Eventually(func() bool { return true }, polling*1000, timeout).Should(BeTrue())
		Eventually(func() bool { return true }, polling*1000, timeout+10000000000000).Should(BeTrue())               // want `timeout must not be shorter than the polling interval`
		Eventually(func() bool { return true }, polling*1000).WithPolling(timeout + 10000000000000).Should(BeTrue()) // want `timeout must not be shorter than the polling interval`
		Eventually(func() bool { return true }, polling*1000).ProbeEvery(timeout + 10000000000000).Should(BeTrue())  // want `timeout must not be shorter than the polling interval`
	})

	It("Consistently timeout shorter than polling", func() {
		Consistently(func() bool { return true }, timeout, pkg.Timeout).Should(BeTrue())                                                  // valid
		Consistently(func() bool { return true }).WithTimeout(time.Second * 10).WithPolling(time.Second * (10 + factor)).Should(BeTrue()) // want `timeout must not be shorter than the polling interval`
		Consistently(func() bool { return true }).Within(time.Second * 10).WithPolling(time.Second * (10 + factor)).Should(BeTrue())      // want `timeout must not be shorter than the polling interval`
		Consistently(func() bool { return true }).Within(time.Second * 10).ProbeEvery(time.Second * (10 + factor)).Should(BeTrue())       // want `timeout must not be shorter than the polling interval`
		Consistently(func() bool { return true }).WithTimeout(time.Second * 10).ProbeEvery(time.Second * (10 + factor)).Should(BeTrue())  // want `timeout must not be shorter than the polling interval`
		Consistently(func() bool { return true }, time.Second*60*60*24*365, time.Second*(10+factor)).Should(BeTrue())                     // valid
		Consistently(func() bool { return true }, time.Second, time.Millisecond*10).Should(BeTrue())                                      //valid
		Consistently(func() bool { return true }, time.Millisecond*10, time.Second).Should(BeTrue())                                      // want `timeout must not be shorter than the polling interval`
		Consistently(func() bool { return true }, timeout, polling).Should(BeTrue())
		Consistently(func() bool { return true }, polling, timeout).Should(BeTrue()) // want `timeout must not be shorter than the polling interval`
		Consistently(func() bool { return true }, polling*1000, timeout).Should(BeTrue())
		Consistently(func() bool { return true }, polling*1000, timeout+10000000000000).Should(BeTrue())               // want `timeout must not be shorter than the polling interval`
		Consistently(func() bool { return true }, polling*1000).WithPolling(timeout + 10000000000000).Should(BeTrue()) // want `timeout must not be shorter than the polling interval`
		Consistently(func() bool { return true }, polling*1000).ProbeEvery(timeout + 10000000000000).Should(BeTrue())  // want `timeout must not be shorter than the polling interval`
	})

	It("timeout shorter than polling + context", func() {
		Eventually(context.Background(), func() bool { return true }, timeout, pkg.Timeout).Should(BeTrue())                                                  // valid
		Eventually(context.Background(), func() bool { return true }).WithTimeout(time.Second * 10).WithPolling(time.Second * (10 + factor)).Should(BeTrue()) // want `timeout must not be shorter than the polling interval`
		Eventually(context.Background(), func() bool { return true }).Within(time.Second * 10).WithPolling(time.Second * (10 + factor)).Should(BeTrue())      // want `timeout must not be shorter than the polling interval`
		Eventually(context.Background(), func() bool { return true }).Within(time.Second * 10).ProbeEvery(time.Second * (10 + factor)).Should(BeTrue())       // want `timeout must not be shorter than the polling interval`
		Eventually(context.Background(), func() bool { return true }).WithTimeout(time.Second * 10).ProbeEvery(time.Second * (10 + factor)).Should(BeTrue())  // want `timeout must not be shorter than the polling interval`
		Eventually(context.Background(), func() bool { return true }, time.Second*60*60*24*365, time.Second*(10+factor)).Should(BeTrue())                     // valid
		Eventually(context.Background(), func() bool { return true }, time.Second, time.Millisecond*10).Should(BeTrue())                                      //valid
		Eventually(context.Background(), func() bool { return true }, time.Millisecond*10, time.Second).Should(BeTrue())                                      // want `timeout must not be shorter than the polling interval`
		Eventually(context.Background(), func() bool { return true }, timeout, polling).Should(BeTrue())
		Eventually(context.Background(), func() bool { return true }, polling, timeout).Should(BeTrue()) // want `timeout must not be shorter than the polling interval`
		Eventually(context.Background(), func() bool { return true }, polling*1000, timeout).Should(BeTrue())
		Eventually(context.Background(), func() bool { return true }, polling*1000, timeout+10000000000000).Should(BeTrue())               // want `timeout must not be shorter than the polling interval`
		Eventually(context.Background(), func() bool { return true }, polling*1000).WithPolling(timeout + 10000000000000).Should(BeTrue()) // want `timeout must not be shorter than the polling interval`
		Eventually(context.Background(), func() bool { return true }, polling*1000).ProbeEvery(timeout + 10000000000000).Should(BeTrue())  // want `timeout must not be shorter than the polling interval`
	})

	It("non-duration values", func() {
		Eventually(func() bool { return true }, 1+0).Should(BeTrue())                   // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
		Eventually(func() bool { return true }, uint(2)).Should(BeTrue())               // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
		Eventually(func() bool { return true }, 1.1).Should(BeTrue())                   // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
		Eventually(func() bool { return true }, float64(2)).Should(BeTrue())            // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
		Eventually(func() bool { return true }, "3s").Should(BeTrue())                  // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
		Eventually(func() bool { return true }, int32(2)).Should(BeTrue())              // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
		Eventually(func() bool { return true }, int64(2)).Should(BeTrue())              // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
		Eventually(func() bool { return true }, time.Second, int64(2)).Should(BeTrue()) // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
		Eventually(func() bool { return true }, time.Second, "2s").Should(BeTrue())     // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
	})
	It("non-duration values + context", func() {
		Eventually(context.Background(), func() bool { return true }, 1+0).Should(BeTrue())                   // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
		Eventually(context.Background(), func() bool { return true }, uint(2)).Should(BeTrue())               // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
		Eventually(context.Background(), func() bool { return true }, 1.1).Should(BeTrue())                   // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
		Eventually(context.Background(), func() bool { return true }, float64(2)).Should(BeTrue())            // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
		Eventually(context.Background(), func() bool { return true }, "3s").Should(BeTrue())                  // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
		Eventually(context.Background(), func() bool { return true }, int32(2)).Should(BeTrue())              // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
		Eventually(context.Background(), func() bool { return true }, int64(2)).Should(BeTrue())              // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
		Eventually(context.Background(), func() bool { return true }, time.Second, int64(2)).Should(BeTrue()) // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
		Eventually(context.Background(), func() bool { return true }, time.Second, "2s").Should(BeTrue())     // want `only use time.Duration for timeout and polling in Eventually\(\) or Consistently\(\)`
	})

	It("multiple timeout", func() {
		Eventually(func() bool { return true }).WithTimeout(time.Second * 10).WithTimeout(time.Second * 10).WithPolling(time.Millisecond * 500).Should(BeTrue()) // want `timeout defined more than once`
		Eventually(func() bool { return true }).WithTimeout(time.Second * 10).Within(time.Second * 10).WithPolling(time.Millisecond * 500).Should(BeTrue())      // want `timeout defined more than once`
		Eventually(func() bool { return true }).Within(time.Second * 10).Within(time.Second * 10).WithPolling(time.Millisecond * 500).Should(BeTrue())           // want `timeout defined more than once`
		Eventually(func() bool { return true }).Within(time.Second * 10).WithTimeout(time.Second * 10).WithPolling(time.Millisecond * 500).Should(BeTrue())      // want `timeout defined more than once`
		Eventually(func() bool { return true }, time.Second*10).WithTimeout(time.Second * 10).WithPolling(time.Millisecond * 500).Should(BeTrue())               //want `timeout defined more than once`
		Eventually(func() bool { return true }, time.Second*10).Within(time.Second * 10).WithPolling(time.Millisecond * 500).Should(BeTrue())                    //want `timeout defined more than once`
	})

	It("multiple timeout + context", func() {
		Eventually(context.Background(), func() bool { return true }).WithTimeout(time.Second * 10).WithTimeout(time.Second * 10).WithPolling(time.Millisecond * 500).Should(BeTrue()) // want `timeout defined more than once`
		Eventually(context.Background(), func() bool { return true }).WithTimeout(time.Second * 10).Within(time.Second * 10).WithPolling(time.Millisecond * 500).Should(BeTrue())      // want `timeout defined more than once`
		Eventually(context.Background(), func() bool { return true }).Within(time.Second * 10).Within(time.Second * 10).WithPolling(time.Millisecond * 500).Should(BeTrue())           // want `timeout defined more than once`
		Eventually(context.Background(), func() bool { return true }).Within(time.Second * 10).WithTimeout(time.Second * 10).WithPolling(time.Millisecond * 500).Should(BeTrue())      // want `timeout defined more than once`
		Eventually(context.Background(), func() bool { return true }, time.Second*10).WithTimeout(time.Second * 10).WithPolling(time.Millisecond * 500).Should(BeTrue())               //want `timeout defined more than once`
		Eventually(context.Background(), func() bool { return true }, time.Second*10).Within(time.Second * 10).WithPolling(time.Millisecond * 500).Should(BeTrue())                    //want `timeout defined more than once`
	})

	It("multiple polling", func() {
		Eventually(func() bool { return true }).WithTimeout(time.Second * 10).WithPolling(500 * time.Millisecond).WithPolling(time.Millisecond * 500).Should(BeTrue()) //want `polling defined more than once`
		Eventually(func() bool { return true }).WithTimeout(time.Second * 10).WithPolling(500 * time.Millisecond).ProbeEvery(time.Millisecond * 500).Should(BeTrue())  //want `polling defined more than once`
		Eventually(func() bool { return true }).WithTimeout(time.Second * 10).ProbeEvery(500 * time.Millisecond).ProbeEvery(time.Millisecond * 500).Should(BeTrue())   //want `polling defined more than once`
		Eventually(func() bool { return true }).WithTimeout(time.Second * 10).ProbeEvery(500 * time.Millisecond).WithPolling(time.Millisecond * 500).Should(BeTrue())  //want `polling defined more than once`
		Eventually(func() bool { return true }, time.Second*10, 500*time.Millisecond).WithPolling(time.Millisecond * 500).Should(BeTrue())                             // want `polling defined more than once`
		Eventually(func() bool { return true }, time.Second*10, 500*time.Millisecond).ProbeEvery(time.Millisecond * 500).Should(BeTrue())                              // want `polling defined more than once`
	})

	It("multiple polling", func() {
		Eventually(context.Background(), func() bool { return true }).WithTimeout(time.Second * 10).WithPolling(500 * time.Millisecond).WithPolling(time.Millisecond * 500).Should(BeTrue()) //want `polling defined more than once`
		Eventually(context.Background(), func() bool { return true }).WithTimeout(time.Second * 10).WithPolling(500 * time.Millisecond).ProbeEvery(time.Millisecond * 500).Should(BeTrue())  //want `polling defined more than once`
		Eventually(context.Background(), func() bool { return true }).WithTimeout(time.Second * 10).ProbeEvery(500 * time.Millisecond).ProbeEvery(time.Millisecond * 500).Should(BeTrue())   //want `polling defined more than once`
		Eventually(context.Background(), func() bool { return true }).WithTimeout(time.Second * 10).ProbeEvery(500 * time.Millisecond).WithPolling(time.Millisecond * 500).Should(BeTrue())  //want `polling defined more than once`
		Eventually(context.Background(), func() bool { return true }, time.Second*10, 500*time.Millisecond).WithPolling(time.Millisecond * 500).Should(BeTrue())                             // want `polling defined more than once`
		Eventually(context.Background(), func() bool { return true }, time.Second*10, 500*time.Millisecond).ProbeEvery(time.Millisecond * 500).Should(BeTrue())                              // want `polling defined more than once`
	})

	It("mix multiple timeout &polling", func() {
		Eventually(func() bool { return true }).WithTimeout(time.Second * 10).WithPolling(500 * time.Millisecond).WithTimeout(time.Second * 10).WithPolling(500 * time.Millisecond).Should(BeTrue()) // want `multiple issues: timeout defined more than once; polling defined more than once`
		Eventually(func() bool { return true }).WithTimeout(time.Second * 10).WithPolling(500 * time.Millisecond).Within(time.Second * 10).ProbeEvery(500 * time.Millisecond).Should(BeTrue())       // want `multiple issues: timeout defined more than once; polling defined more than once`
		Eventually(func() bool { return true }).Within(time.Second * 10).ProbeEvery(500 * time.Millisecond).Within(time.Second * 10).ProbeEvery(500 * time.Millisecond).Should(BeTrue())             // want `multiple issues: timeout defined more than once; polling defined more than once`
		Eventually(func() bool { return true }, time.Second*10, time.Millisecond*500).WithTimeout(time.Second * 60 * 60 * 24 * 365).WithPolling(time.Second * (10 + factor)).Should(BeTrue())        //want `multiple issues: timeout defined more than once; polling defined more than once`
		Eventually(func() bool { return true }, time.Second*10, time.Millisecond*500).Within(time.Second * 60 * 60 * 24 * 365).ProbeEvery(time.Second * (10 + factor)).Should(BeTrue())              //want `multiple issues: timeout defined more than once; polling defined more than once`
		Eventually(func() bool { return true }, time.Second*10, time.Millisecond*500).WithTimeout(time.Second * 60 * 60 * 24 * 365).ProbeEvery(time.Second * (10 + factor)).Should(BeTrue())         //want `multiple issues: timeout defined more than once; polling defined more than once`
		Eventually(func() bool { return true }, time.Second*10, time.Millisecond*500).Within(time.Second * 60 * 60 * 24 * 365).WithPolling(time.Second * (10 + factor)).Should(BeTrue())             //want `multiple issues: timeout defined more than once; polling defined more than once`
	})
})
