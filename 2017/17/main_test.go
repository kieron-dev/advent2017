package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Dance", func() {
	It("runs successfully with a valid arg", func() {
		command := exec.Command(cmd, "./input.txt")
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session, 600).Should(gexec.Exit(0))
		Expect(session).To(gbytes.Say("An incorrect value, so I can see the real result"))
	})

	It("errors when no arg passed", func() {
		command := exec.Command(cmd)
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session).Should(gexec.Exit(1))
	})
})
