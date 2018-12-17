package q16_test

import (
	"github.com/kieron-pivotal/advent2017/2018/q16"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q16", func() {

	DescribeTable("Ops Table", func(op,
		in0, in1, in2, in3,
		arg1, arg2, arg3,
		out0, out1, out2, out3 int) {
		c := q16.NewComputer()
		d := q16.NewComputer()
		c.SetRegisters(in0, in1, in2, in3)
		d.SetRegisters(out0, out1, out2, out3)
		c.Ops()[op](arg1, arg2, arg3)
		Expect(c.Equals(d)).To(BeTrue())
	},
		Entry("addr", 0,
			1, 2, 3, 4,
			3, 2, 1,
			1, 7, 3, 4),
		Entry("addi", 1,
			2, 3, 5, 7,
			1, 3, 0,
			6, 3, 5, 7),
		Entry("mulr", 2,
			2, 3, 5, 7,
			0, 2, 1,
			2, 10, 5, 7),
		Entry("muli", 3,
			2, 3, 5, 7,
			0, 2, 1,
			2, 4, 5, 7),
		Entry("banr", 4,
			2, 3, 5, 7,
			0, 1, 2,
			2, 3, 2, 7),
		Entry("bani", 5,
			6, 3, 5, 7,
			0, 1, 3,
			6, 3, 5, 0),
		Entry("borr", 6,
			2, 3, 5, 7,
			0, 1, 2,
			2, 3, 3, 7),
		Entry("bori", 7,
			6, 3, 5, 7,
			0, 1, 3,
			6, 3, 5, 7),
		Entry("setr", 8,
			6, 3, 5, 7,
			0, 1, 3,
			6, 3, 5, 6),
		Entry("seti", 9,
			6, 3, 5, 7,
			9, 1, 3,
			6, 3, 5, 9),
		Entry("gtir", 10,
			2, 3, 5, 7,
			1, 1, 3,
			2, 3, 5, 0),
		Entry("gtri", 11,
			2, 3, 5, 7,
			1, 1, 3,
			2, 3, 5, 1),
		Entry("gtrr", 12,
			2, 3, 5, 7,
			1, 2, 3,
			2, 3, 5, 0),
		Entry("eqir", 13,
			2, 3, 5, 7,
			5, 2, 3,
			2, 3, 5, 1),
		Entry("eqri", 14,
			2, 3, 5, 7,
			2, 5, 3,
			2, 3, 5, 1),
		Entry("eqrr", 15,
			7, 3, 5, 7,
			0, 3, 3,
			7, 3, 5, 1),
	)

	It("can get the write number of matches", func() {
		c := q16.NewComputer()
		c.SetRegisters(3, 2, 1, 1)
		d := q16.NewComputer()
		d.SetRegisters(3, 2, 2, 1)
		Expect(c.NumOps(d, 2, 1, 2)).To(Equal(3))
	})

})
