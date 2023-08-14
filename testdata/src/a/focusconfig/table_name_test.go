package focus

import (
	gnk "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = gnk.Describe("test focused tables", func() {
	gnk.FDescribeTable("focused table", func(s []int, l int) { // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code, consider to replace with "DescribeTable"`
		Expect(len(s)).Should(Equal(l)) // want `ginkgo-linter: wrong length assertion; consider using .Expect\(s\)\.Should\(HaveLen\(l\)\). instead`
	},
		gnk.Entry("check slice not focused", []int{1, 2, 3, 4}, 4),
		gnk.Entry("check slice focus spec", gnk.Focus, []int{1, 2, 3, 4}, 4), // want `ginkgo-linter: Focus spec found. This is used only for local debug and should not be part of the actual source code, consider to remove it`
		gnk.FEntry("check slice focused", []int{1, 2, 3, 4, 5}, 5),           // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code, consider to replace with "Entry"`
	)

	gnk.DescribeTable("non-focused table", func(s []int, l int) {
		Expect(s).Should(HaveLen(l))
	},
		gnk.Entry("check slice not focused", []int{1, 2, 3, 4}, 4),
		gnk.Entry("check slice focus spec", gnk.Focus, []int{1, 2, 3, 4}, 4), // want `ginkgo-linter: Focus spec found. This is used only for local debug and should not be part of the actual source code, consider to remove it`
		gnk.FEntry("check slice focused", []int{1, 2, 3, 4, 5}, 5),           // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code, consider to replace with "Entry"`
	)

	gnk.DescribeTable("spec focused table", gnk.Focus, func(s []int, l int) { // want `ginkgo-linter: Focus spec found. This is used only for local debug and should not be part of the actual source code, consider to remove it`
		Expect(s).Should(HaveLen(l))
	},
		gnk.Entry("check slice not focused", []int{1, 2, 3, 4}, 4),
		gnk.Entry("check slice focus spec", gnk.Focus, []int{1, 2, 3, 4}, 4), // want `ginkgo-linter: Focus spec found. This is used only for local debug and should not be part of the actual source code, consider to remove it`
		gnk.FEntry("check slice focused", []int{1, 2, 3, 4, 5}, 5),           // want `ginkgo-linter: Focus container found. This is used only for local debug and should not be part of the actual source code, consider to replace with "Entry"`
	)
})
