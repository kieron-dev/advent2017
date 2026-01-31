package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/masked"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("14", func() {
	var (
		data   *os.File
		memory masked.Memory
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input14")
		Expect(err).NotTo(HaveOccurred())

		memory = masked.NewMemory()
		memory.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("does part A", func() {
		memory.ProcessMaskingVals()
		Expect(memory.GetSum()).To(Equal(9879607673316))
	})

	It("does part B", func() {
		memory.ProcessMaskingAddrs()
		Expect(memory.GetSum()).To(Equal(3435342392262))
	})
})
