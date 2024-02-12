package noassersion

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGomega(t *testing.T) {
	g := NewWithT(t)

	g.Expect("aaa")                       // want `ginkgo-linter: "Expect": missing assertion method\. Expected "To\(\)", "ToNot\(\)" or "NotTo\(\)"`
	g.Expect("aaa").WithOffset(1)         // want `ginkgo-linter: "Expect": missing assertion method\. Expected "To\(\)", "ToNot\(\)" or "NotTo\(\)"`
	g.Expect("aaa").WithOffset(1).Error() // want `ginkgo-linter: "Expect": missing assertion method\. Expected "To\(\)", "ToNot\(\)" or "NotTo\(\)"`

	g.Eventually(func(gg Gomega) { gg.Expect(true) }) // want `ginkgo-linter: "Eventually": missing assertion method\. Expected "Should\(\)" or "ShouldNot\(\)"` `ginkgo-linter: "Expect": missing assertion method\. Expected "To\(\)", "ToNot\(\)" or "NotTo\(\)"`
}
