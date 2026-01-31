package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/vm"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("08", func() {
	var (
		data    *os.File
		console vm.Console
	)

	BeforeEach(func() {
		var err error
		data, err := os.Open("./input08")
		Expect(err).NotTo(HaveOccurred())

		console = vm.NewConsole()
		console.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("part A", func() {
		console.RunTillLoop()
		Expect(console.Acc()).To(Equal(1928))
	})

	It("part B", func() {
		console.FixInstructionTillTerm()
		Expect(console.Acc()).To(Equal(1319))
	})
})
