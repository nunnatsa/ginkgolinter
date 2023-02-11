package havelen0

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestHaveLen0(t *testing.T) {
	g := NewWithT(t)

	x := make([]int, 0)
	g.Expect(x).Should(HaveLen(0))

	x = append(x, 1)
	g.Expect(x).Should(Not(HaveLen(0)))
}
