package days_test

import (
	"os"

	"github.com/kieron-dev/advent2017/advent2019"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q03", func() {
	var (
		file       *os.File
		fileReader advent2019.FileReader
		grids      []*advent2019.Grid
	)

	BeforeEach(func() {
		var err error
		file, err = os.Open("./input03")
		if err != nil {
			panic(err)
		}
		fileReader = advent2019.FileReader{}
		grids = []*advent2019.Grid{}
	})

	It("does part A", func() {
		fileReader.Each(file, func(line string) {
			grid := advent2019.NewGrid()
			grid.Move(line)
			grids = append(grids, grid)
		})

		res := grids[0].ClosestCommonDist(grids[1])
		Expect(res).To(Equal(1626))
	})

	It("does part B", func() {
		fileReader.Each(file, func(line string) {
			grid := advent2019.NewGrid()
			grid.Move(line)
			grids = append(grids, grid)
		})

		res := grids[0].QuickestCommonCell(grids[1])
		Expect(res).To(Equal(27330))
	})
})
