package focus

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("test focused tables", func() {
	FDescribeTable("focused table", func(s []int, l int) { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code, consider to replace with "DescribeTable"`
		Expect(len(s)).Should(Equal(l)) // want `ginkgo-linter: wrong length assertion; consider using .Expect\(s\)\.Should\(HaveLen\(l\)\). instead`
	},
		Entry("check slice not focused", []int{1, 2, 3, 4}, 4),
		Entry("check slice focus spec", Focus, []int{1, 2, 3, 4}, 4), // want `ginkgo-linter: Focus spec found. This is used only for local debug and should not be part of the actual source code, consider to remove it`
		FEntry("check slice focused", []int{1, 2, 3, 4, 5}, 5),       // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code, consider to replace with "Entry"`
	)

	DescribeTable("non-focused table", func(s []int, l int) {
		Expect(s).Should(HaveLen(l))
	},
		Entry("check slice not focused", []int{1, 2, 3, 4}, 4),
		Entry("check slice focus spec", Focus, []int{1, 2, 3, 4}, 4), // want `ginkgo-linter: Focus spec found. This is used only for local debug and should not be part of the actual source code, consider to remove it`
		FEntry("check slice focused", []int{1, 2, 3, 4, 5}, 5),       // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code, consider to replace with "Entry"`
	)

	DescribeTable("spec focused table", Focus, func(s []int, l int) { // want `ginkgo-linter: Focus spec found. This is used only for local debug and should not be part of the actual source code, consider to remove it`
		Expect(s).Should(HaveLen(l))
	},
		Entry("check slice not focused", []int{1, 2, 3, 4}, 4),
		Entry("check slice focus spec", Focus, []int{1, 2, 3, 4}, 4), // want `ginkgo-linter: Focus spec found. This is used only for local debug and should not be part of the actual source code, consider to remove it`
		FEntry("check slice focused", []int{1, 2, 3, 4, 5}, 5),       // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code, consider to replace with "Entry"`
	)
})
