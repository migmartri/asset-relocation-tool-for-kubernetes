// +build feature

package features

import (
	"os/exec"

	. "github.com/bunniesandbeatings/goerkin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Report version", func() {
	steps := NewSteps()

	Scenario("version command reports version", func() {
		steps.When("running relok8s version")
		steps.Then("the command exits without error")
		steps.And("the version is printed")
	})

	steps.Define(func(define Definitions) {
		DefineCommonSteps(define)

		define.When(`^running relok8s version$`, func() {
			command := exec.Command(ChartMoverBinaryPath, "version")
			var err error
			CommandSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
		})

		define.Then(`^the version is printed$`, func() {
			Eventually(CommandSession.Out).Should(Say("relok8s version: 1.2.3"))
		})
	})
})
