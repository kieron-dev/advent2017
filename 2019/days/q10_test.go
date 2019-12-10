package days_test

import (
	"os"

	"github.com/kieron-pivotal/advent2017/advent2019/grid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q10", func() {
	var (
		file *os.File
		g    *grid.Grid
	)

	BeforeEach(func() {
		var err error
		file, err = os.Open("./input10")
		if err != nil {
			panic(err)
		}
		g = grid.New()
		g.Load(file)
	})

	It("does part A", func() {
		coord := g.BestAsteroid()
		Expect(g.VisibleFrom(coord)).To(HaveLen(247))
	})

	It("does part B", func() {
		coord := g.BestAsteroid()
		res := g.LaserN(200, coord)
		Expect(res).To(Equal(grid.NewCoord(19, 19)))
	})

})
