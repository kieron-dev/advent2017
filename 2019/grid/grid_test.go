package grid_test

import (
	"io"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019/grid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grid", func() {

	var (
		g        *grid.Grid
		contents io.Reader
	)

	BeforeEach(func() {
		g = grid.New()
	})

	It("can load a file", func() {
		contents = strings.NewReader(`##.#
.#.#`)
		g.Load(contents)
		Expect(g.Height()).To(Equal(2))
		Expect(g.Width()).To(Equal(4))
		Expect(g.AsteroidCount()).To(Equal(5))
	})

	It("can count visible asteroids from a point", func() {
		contents = strings.NewReader(`.#..#
.....
#####
....#
...##
`)
		g.Load(contents)
		Expect(g.VisibleFrom(grid.NewCoord(3, 4))).To(HaveLen(8))
		Expect(g.VisibleFrom(grid.NewCoord(4, 2))).To(HaveLen(5))
		Expect(g.VisibleFrom(grid.NewCoord(4, 0))).To(HaveLen(7))
		Expect(g.VisibleFrom(grid.NewCoord(1, 0))).To(HaveLen(7))
		Expect(g.BestAsteroid()).To(Equal(grid.NewCoord(3, 4)))
	})

	It("can order visible asteroids by angle", func() {

		contents = strings.NewReader(`.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....X...###..
..#.#.....#....##
`)
		g.Load(contents)
		best := grid.NewCoord(8, 3)
		foo := g.VisibleFrom(best)
		g.Sort(best, foo)

		Expect(foo[0]).To(Equal(grid.NewCoord(8, 1)))
		Expect(foo[1]).To(Equal(grid.NewCoord(9, 0)))
		Expect(foo[2]).To(Equal(grid.NewCoord(9, 1)))
		Expect(foo[3]).To(Equal(grid.NewCoord(10, 0)))
		Expect(foo[4]).To(Equal(grid.NewCoord(9, 2)))
		Expect(foo[5]).To(Equal(grid.NewCoord(11, 1)))
		Expect(foo[6]).To(Equal(grid.NewCoord(12, 1)))
		Expect(foo[7]).To(Equal(grid.NewCoord(11, 2)))
		Expect(foo[8]).To(Equal(grid.NewCoord(15, 1)))
		Expect(foo[9]).To(Equal(grid.NewCoord(12, 2)))
		Expect(foo[10]).To(Equal(grid.NewCoord(13, 2)))
		Expect(foo[11]).To(Equal(grid.NewCoord(14, 2)))
		Expect(foo[12]).To(Equal(grid.NewCoord(15, 2)))
		Expect(foo[13]).To(Equal(grid.NewCoord(12, 3)))
		Expect(foo[14]).To(Equal(grid.NewCoord(16, 4)))
		Expect(foo[15]).To(Equal(grid.NewCoord(15, 4)))
		Expect(foo[16]).To(Equal(grid.NewCoord(10, 4)))
		Expect(foo[17]).To(Equal(grid.NewCoord(4, 4)))
		Expect(foo[18]).To(Equal(grid.NewCoord(2, 4)))
		Expect(foo[19]).To(Equal(grid.NewCoord(2, 3)))
		Expect(foo[20]).To(Equal(grid.NewCoord(0, 2)))
		Expect(foo[21]).To(Equal(grid.NewCoord(1, 2)))
		Expect(foo[22]).To(Equal(grid.NewCoord(0, 1)))
		Expect(foo[23]).To(Equal(grid.NewCoord(1, 1)))
		Expect(foo[24]).To(Equal(grid.NewCoord(5, 2)))
		Expect(foo[25]).To(Equal(grid.NewCoord(1, 0)))
		Expect(foo[26]).To(Equal(grid.NewCoord(5, 1)))
		Expect(foo[27]).To(Equal(grid.NewCoord(6, 1)))
		Expect(foo[28]).To(Equal(grid.NewCoord(6, 0)))
		Expect(foo[29]).To(Equal(grid.NewCoord(7, 0)))
	})

	It("can find the 200th asteroid to be lasered", func() {
		contents = strings.NewReader(`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##
`)
		g.Load(contents)
		best := g.BestAsteroid()
		Expect(best).To(Equal(grid.NewCoord(11, 13)))
		Expect(g.LaserN(1, best)).To(Equal(grid.NewCoord(11, 12)))
		Expect(g.LaserN(2, best)).To(Equal(grid.NewCoord(12, 1)))
		Expect(g.LaserN(3, best)).To(Equal(grid.NewCoord(12, 2)))
		Expect(g.LaserN(10, best)).To(Equal(grid.NewCoord(12, 8)))

		Expect(g.LaserN(100, best)).To(Equal(grid.NewCoord(10, 16)))
		Expect(g.LaserN(200, best)).To(Equal(grid.NewCoord(8, 2)))
	})

})
