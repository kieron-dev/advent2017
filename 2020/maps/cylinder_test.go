package maps_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/maps"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cylinder", func() {
	var (
		input io.Reader
		grid  maps.Cylinder
	)

	BeforeEach(func() {
		grid = maps.NewCylinder()

		input = strings.NewReader(`
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`)
		grid.Load(input)
	})

	It("counts the #s in a 3,1 downwards direction from 0,0", func() {
		Expect(grid.CountChars(maps.NewCoord(0, 0), maps.NewVector(3, 1), '#')).To(Equal(7))
	})
})
