package gameoflife_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/gameoflife"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Seating Plan", func() {
	var (
		data io.Reader
		plan gameoflife.SeatingPlan
	)

	BeforeEach(func() {
		data = strings.NewReader(`
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`)
		plan = gameoflife.NewSeatingPlan()
		plan.Load(data)
	})

	It("can load the state", func() {
		Expect(plan.State()).To(Equal(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`))
	})

	Describe("part A", func() {
		It("can do a single iteration", func() {
			plan.Evolve(false)
			Expect(plan.State()).To(Equal(`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
`))
		})

		It("can iterate until stabilisation", func() {
			plan.Stabilise(false)
			Expect(plan.State()).To(Equal(`#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##
`))
			Expect(plan.OccupiedSeats()).To(Equal(37))
		})
	})

	Describe("part B", func() {
		It("can do a single iteration", func() {
			plan.Evolve(true)
			Expect(plan.State()).To(Equal(`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
`))
		})

		It("can do a two iterations", func() {
			plan.Evolve(true)
			plan.Evolve(true)
			Expect(plan.State()).To(Equal(`#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#
`))
		})

		It("can iterate until stabilisation", func() {
			plan.Stabilise(true)
			Expect(plan.State()).To(Equal(`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#
`))
			Expect(plan.OccupiedSeats()).To(Equal(26))
		})
	})
})
