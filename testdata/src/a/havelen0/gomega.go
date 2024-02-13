package havelen0

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestHaveLen0(t *testing.T) {
	g := NewWithT(t)

	x := make([]int, 0)
	g.Expect(x).Should(HaveLen(0))               // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(x\)\.Should\(BeEmpty\(\)\). instead`
	g.Expect(x).WithOffset(1).Should(HaveLen(0)) // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(x\)\.WithOffset\(1\)\.Should\(BeEmpty\(\)\). instead`

	x = append(x, 1)
	g.Expect(x).Should(Not(HaveLen(0))) // want `ginkgo-linter: wrong length assertion\. Consider using .g\.Expect\(x\)\.ShouldNot\(BeEmpty\(\)\). instead`
}
