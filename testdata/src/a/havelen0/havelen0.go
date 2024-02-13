package havelen0

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const EMPTY = 0

var _ = Describe("test HaveLen(0)", func() {
	It("should replace HaveLen(0) with BeEmpty()", func() {
		x := make([]int, 0)
		Expect(x).To(HaveLen(0))               // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\(x\)\.To\(BeEmpty\(\)\). instead`
		Expect(x).WithOffset(1).To(HaveLen(0)) // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\(x\)\.WithOffset\(1\)\.To\(BeEmpty\(\)\). instead`
		Expect(x).To(HaveLen(EMPTY))           // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\(x\)\.To\(BeEmpty\(\)\). instead`

		x = append(x, 1)
		Expect(x).To(Not(HaveLen(0)))     // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\(x\)\.ToNot\(BeEmpty\(\)\). instead`
		Expect(x).To(Not(HaveLen(EMPTY))) // want `ginkgo-linter: wrong length assertion\. Consider using .Expect\(x\)\.ToNot\(BeEmpty\(\)\). instead`
	})
})
