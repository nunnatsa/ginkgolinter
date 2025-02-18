package issue_190

import (
	"context"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"

	"github.com/onsi/gomega"
)

var _ = Describe("", func() {
	It("should work", func(ctx SpecContext) {
		gomega.Eventually(ctx, func() error {
			return fmt.Errorf("foo")
		}).Should(gomega.MatchError("foo"))
	}, SpecTimeout(time.Second))
	It("should work", func(ctx context.Context) {
		gomega.Eventually(ctx, func() error {
			return fmt.Errorf("foo")
		}).Should(gomega.MatchError("foo"))
	}, SpecTimeout(time.Second))
})
