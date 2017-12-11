package registers_test

import (
	"github.com/kieron-pivotal/advent2017/08/registers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Registers", func() {
	It("parses an instruction", func() {
		Expect(registers.ParseInstruction("fi inc 284 if cp == 976")).To(Equal(registers.Instruction{
			Register:   "fi",
			Op:         registers.INC,
			Amt:        284,
			CondTarget: "cp",
			Comparison: registers.EQ,
			CompAmt:    976,
		}))

		Expect(registers.ParseInstruction("fi dec -284 if cp >= -976")).To(Equal(registers.Instruction{
			Register:   "fi",
			Op:         registers.DEC,
			Amt:        -284,
			CondTarget: "cp",
			Comparison: registers.GTE,
			CompAmt:    -976,
		}))
	})

})
