package acceptance

import (
	"fmt"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

func verifyUsageText(session *gexec.Session) {

	By("Then I should see help message printed with the product name", func() {
		Eventually(session).Should(gbytes.Say("MCCP CLI"))
	})

	By("And Usage category", func() {
		Eventually(session).Should(gbytes.Say("Usage:"))
		Eventually(string(session.Wait().Out.Contents())).Should(ContainSubstring("mccp [command]"))
	})

	By("And Available-Commands category", func() {
		Eventually(session).Should(gbytes.Say("Available Commands:"))
		Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`help[\s]+Help about any command`))
		Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`templates[\s]+Interact with CAPI templates`))
	})

	By("And Flags category", func() {
		Eventually(session).Should(gbytes.Say("Flags:"))
		Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`-h, --help[\s]+help for mccp`))
	})

	By("And command help usage", func() {
		Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`Use "mccp \[command\] --help" for more information about a command`))
	})

}

func DescribeMccpCliHelp() {
	var _ = Describe("MCCP Help Tests", func() {

		MCCP_BIN_PATH := GetMCCBinPath()
		CAPI_ENDPOINT_URL := GetCapiEndpointUrl()

		var session *gexec.Session
		var err error

		BeforeEach(func() {

			By("Given I have a mccp binary installed on my local machine", func() {
				Expect(FileExists(MCCP_BIN_PATH)).To(BeTrue(), fmt.Sprintf("%s can not be found.", MCCP_BIN_PATH))
			})
		})

		It("Verify that mccp displays error message when provided with the wrong flag", func() {

			By("When I run 'mccp foo'", func() {
				command := exec.Command(MCCP_BIN_PATH, "foo")
				session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
				Expect(err).ShouldNot(HaveOccurred())
			})

			By("Then I should see mccp error message", func() {
				Eventually(session.Err).Should(gbytes.Say("Error: unknown command \"foo\" for \"mccp\""))
				Eventually(session.Err).Should(gbytes.Say("Run 'mccp --help' for usage."))
			})
		})

		It("Verify that mccp help flag prints the help text", func() {

			By("When I run the command 'mccp --help' ", func() {
				command := exec.Command(MCCP_BIN_PATH, "--help")
				session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
				Expect(err).ShouldNot(HaveOccurred())
			})

			verifyUsageText(session)

		})

		It("Verify that mccp command prints the help text", func() {

			By("When I run the command 'mccp'", func() {
				command := exec.Command(MCCP_BIN_PATH)
				session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
				Expect(err).ShouldNot(HaveOccurred())
			})

			verifyUsageText(session)

		})

		It("Verify that mccp command prints the sub help text for the list command", func() {

			By("When I run the command 'mccp templates list help'", func() {
				command := exec.Command(MCCP_BIN_PATH, "templates", "list", "--help", "--endpoint", CAPI_ENDPOINT_URL)
				session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
				Expect(err).ShouldNot(HaveOccurred())
			})

			By("Then I should see help message printed with the product name", func() {
				Eventually(session).Should(gbytes.Say("List CAPI templates"))
			})

			By("And Usage category", func() {
				Eventually(session).Should(gbytes.Say("Usage:"))
				Eventually(string(session.Wait().Out.Contents())).Should(ContainSubstring("mccp templates list [flags]"))
			})

			By("And Examples category", func() {
				Eventually(session).Should(gbytes.Say("Examples:"))
				Eventually(session).Should(gbytes.Say("mccp templates list"))
			})

			By("And Flags category", func() {
				Eventually(session).Should(gbytes.Say("Flags:"))
				Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`-h, --help[\s]+help for list`))
			})

			By("And  Global Flags category", func() {
				Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`-e, --endpoint string\s+The CAPI templates HTTP API endpoint`))
			})

		})

		It("Verify that mccp command prints the sub help text for the render command", func() {

			By("When I run the command 'mccp templates render help'", func() {
				command := exec.Command(MCCP_BIN_PATH, "templates", "render", "--help", "--endpoint", CAPI_ENDPOINT_URL)
				session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
				Expect(err).ShouldNot(HaveOccurred())
			})

			By("Then I should see help message printed with the product name", func() {
				Eventually(session).Should(gbytes.Say("Render CAPI template"))
			})

			By("And Usage category", func() {
				Eventually(session).Should(gbytes.Say("Usage:"))
				Eventually(string(session.Wait().Out.Contents())).Should(ContainSubstring("mccp templates render [flags]"))
			})

			By("And Examples category", func() {
				Eventually(session).Should(gbytes.Say("Examples:"))
				Eventually(session).Should(gbytes.Say("mccp templates render <template-name>"))
			})

			By("And Flags category", func() {
				Eventually(session).Should(gbytes.Say("Flags:"))
				Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`-h, --help[\s]+help for render`))
				Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`--create-pr[\s]+Indicates whether to create a pull request for the CAPI template`))
				Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`--list-parameters[\s]+The CAPI templates HTTP API endpoint`))
				Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`--pr-base string[\s]+The base branch to open the pull request against`))
				Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`--pr-branch string[\s]+The branch to create the pull request from`))
				Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`--pr-commit-message string[\s]+The commit message to use when adding the CAPI template`))
				Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`--pr-description string[\s]+The description of the pull request`))
				Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`--pr-repo string[\s]+The repository to open a pull request against`))
				Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`--pr-title string[\s]+The title of the pull request`))
				Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`--set stringArray[\s]+Set parameter values on the command line \(can specify multiple or separate values with commas: key1=val1,key2=val2\)`))
			})

			By("And  Global Flags category", func() {
				Eventually(string(session.Wait().Out.Contents())).Should(MatchRegexp(`-e, --endpoint string\s+The CAPI templates HTTP API endpoint`))
			})

		})
	})
}
