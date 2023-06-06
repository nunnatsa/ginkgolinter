package eventually

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
	"os"
	"os/exec"
	"testing"
	"time"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "test session")
}

var _ = Describe("", func() {
	createSession := func() *Session {
		s, err := Start(exec.Command(os.Args[0]), GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		return s
	}

	It("", func() {
		Eventually(createSession()).WithTimeout(time.Second * 10).Should(Exit(0))
	})
})
