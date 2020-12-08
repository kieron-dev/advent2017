package vm_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/vm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Console", func() {
	var (
		data    io.Reader
		console vm.Console
	)

	BeforeEach(func() {
		data = strings.NewReader(`
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`)
		console = vm.NewConsole()
		console.Load(data)
	})

	It("starts infinite loop with correct acc value", func() {
		console.RunTillLoop()
		Expect(console.Acc()).To(Equal(5))
	})

	It("terminates with an inst toggle", func() {
		console.FixInstructionTillTerm()
		Expect(console.Acc()).To(Equal(8))
	})
})
