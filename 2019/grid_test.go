package advent2019_test

import (
	"github.com/kieron-dev/advent2017/advent2019"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grid", func() {

	var (
		grid *advent2019.Grid
	)

	BeforeEach(func() {
		grid = advent2019.NewGrid()
	})

	It("accepts movements", func() {
		grid.Move("U10")
		Expect(true).To(BeTrue())
	})

	It("finds closest common point", func() {
		grid.Move("R8,U5,L5,D3")

		grid2 := advent2019.NewGrid()
		grid2.Move("U7,R6,D4,L4")

		Expect(grid.ClosestCommonDist(grid2)).To(Equal(6))
	})

})
