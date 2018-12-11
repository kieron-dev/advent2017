package q11_test

import (
	"github.com/kieron-pivotal/advent2017/2018/q11"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q11", func() {

	DescribeTable("cell power", func(serial, row, col int, power int) {
		grid := q11.NewGrid(serial, 300, 300)
		Expect(grid.CellPower(row, col)).To(Equal(power))
	},

		Entry("easy", 8, 4, 2, 4),
		Entry("ex1", 57, 78, 121, -5),
		Entry("ex2", 39, 195, 216, 0),
		Entry("ex3", 71, 152, 100, 4),
	)

	DescribeTable("max power fuel cell", func(serial, expectedRow, expectedCol int) {
		grid := q11.NewGrid(serial, 300, 300)
		r, c := grid.Largest3x3Cell()
		Expect(r).To(Equal(expectedRow))
		Expect(c).To(Equal(expectedCol))
	},
		Entry("ex1", 18, 45, 33),
		Entry("ex2", 42, 61, 21),
	)

	DescribeTable("max power n x n fuel cell", func(serial, expectedRow, expectedCol, expectedSize, expectedPower int) {
		grid := q11.NewGrid(serial, 300, 300)
		p, r, c, n := grid.LargestCell()
		Expect(p).To(Equal(expectedPower), "power")
		Expect(r).To(Equal(expectedRow), "row")
		Expect(c).To(Equal(expectedCol), "col")
		Expect(n).To(Equal(expectedSize), "size")
	},
		Entry("ex1", 18, 269, 90, 16, 113),
		Entry("ex2", 42, 251, 232, 12, 119),
	)
})
