package issue_211

import (
	"context"

	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
)

var x = 8

var _ = Describe("make sure issue 211 is fixed", func() {

	_ = func(sz int) {
		var buf []int
		Expect(len(buf)).To(BeNumerically("==", 1<<(fastLog2(uint64(sz))))) // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\(buf\)\.To\(HaveLen\(1 << \(fastLog2\(uint64\(sz\)\)\)\)\). instead`
		Expect(cap(buf)).To(BeNumerically("==", 1<<(fastLog2(uint64(sz))))) // want `ginkgo-linter: wrong cap assertion\. Consider using .Expect\(buf\)\.To\(HaveCap\(1 << \(fastLog2\(uint64\(sz\)\)\)\)\). instead`
	}

	It("should not exit with nil reference error", func(ctx context.Context) {
		slice := make([]int, 9, 10)
		Expect(cap(slice)).To(BeNumerically(">=", 1))   // want `ginkgo-linter: wrong cap assertion\. Consider using .Expect\(slice\)\.ToNot\(HaveCap\(0\)\). instead`
		Expect(cap(slice)).To(BeNumerically(">", 0))    // want `ginkgo-linter: wrong cap assertion\. Consider using .Expect\(slice\)\.ToNot\(HaveCap\(0\)\). instead`
		Expect(cap(slice)).To(BeNumerically("==", x+2)) // want `ginkgo-linter: wrong cap assertion\. Consider using .Expect\(slice\)\.To\(HaveCap\(x \+ 2\)\). instead`
	})
})
