package masked_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/masked"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Memory", func() {
	var (
		data   io.Reader
		memory masked.Memory
	)

	JustBeforeEach(func() {
		memory = masked.NewMemory()
		memory.Load(data)
	})

	Context("masking vals", func() {
		BeforeEach(func() {
			data = strings.NewReader(`
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
`)
		})

		It("can process the data", func() {
			memory.ProcessMaskingVals()

			Expect(memory.Get(8)).To(Equal(64))
			Expect(memory.Get(7)).To(Equal(101))
		})
	})

	Context("masking addrs", func() {
		BeforeEach(func() {
			data = strings.NewReader(`
mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1
`)
		})

		It("can do the variable memory addr thing", func() {
			memory.ProcessMaskingAddrs()

			Expect(memory.GetSum()).To(Equal(208))
		})
	})
})
