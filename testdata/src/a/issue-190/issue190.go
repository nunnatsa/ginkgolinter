package issue_190

import (
	"context"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
)

var _ = Describe("make sure issue 190 is fixed", func() {
	It("should work", func(ctx context.Context) {
		Eventually(ctx, func() error {
			return fmt.Errorf("foo")
		}).Should(MatchError("foo"))
	}, SpecTimeout(time.Second))

	It("should work", func(ctx SpecContext) {
		Eventually(ctx, func() error {
			return fmt.Errorf("foo")
		}).Should(MatchError("foo"))
	}, SpecTimeout(time.Second))
})
