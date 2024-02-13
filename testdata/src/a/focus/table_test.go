package focus

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("test focused tables", func() {
	FDescribeTable("focused table", func(s []int, l int) {
		Expect(len(s)).Should(Equal(l)) // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\(s\)\.Should\(HaveLen\(l\)\). instead`
	},
		Entry([]int{1, 2, 3, 4}, 4),
		FEntry([]int{1, 2, 3, 4, 5}, 5),
	)

	DescribeTable("non-focused table", func(s []int, l int) {
		Expect(s).Should(HaveLen(l))
	},
		Entry([]int{1, 2, 3, 4}, 4),
		FEntry([]int{1, 2, 3, 4, 5}, 5),
	)
})
