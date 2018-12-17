package q15_test

import (
	"strings"

	"github.com/kieron-pivotal/advent2017/2018/q15"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q15", func() {

	var (
		input01 *strings.Reader
		input02 *strings.Reader
		input03 *strings.Reader
		input04 *strings.Reader
		// input05 *strings.Reader
		fight *q15.Fight
	)

	BeforeEach(func() {
		input01 = strings.NewReader(`#######
#E..G.#
#...#.#
#.G.#G#
#######
`)

		input02 = strings.NewReader(`#######
#.E...#
#.....#
#...G.#
#######
`)

		input03 = strings.NewReader(`#########
#G..G..G#
#.......#
#.......#
#G..E..G#
#.......#
#.......#
#G..G..G#
#########
`)

		input04 = strings.NewReader(`#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######
`)

		// 		input05 = strings.NewReader(`#######
		// #...G.#
		// #..G.G#
		// #.#.#G#
		// #...#E#
		// #.....#
		// #######
		// `)
	})

	Context("initialising", func() {

		It("can read the input", func() {
			fight = q15.NewFight(input01)
			Expect(fight.At(q15.Coord{Row: 0, Col: 0})).To(Equal('#'))
			Expect(fight.At(q15.Coord{Row: 1, Col: 4})).To(Equal('G'))
		})

	})

	Context("identifying squares", func() {
		It("lists units in the correct order", func() {
			fight = q15.NewFight(input01)
			list := fight.GetActorCoords()
			Expect(list).To(HaveLen(4))
			Expect(list[0]).To(Equal(q15.Coord{Row: 1, Col: 1}))
			Expect(list[1]).To(Equal(q15.Coord{Row: 1, Col: 4}))
		})

		It("identifies attack squares", func() {
			fight = q15.NewFight(input01)
			list := fight.GetAttackSquares(q15.Coord{Row: 1, Col: 1})
			Expect(list).To(HaveLen(6))
			Expect(list[0]).To(Equal(q15.Coord{Row: 1, Col: 3}))
			Expect(list[1]).To(Equal(q15.Coord{Row: 1, Col: 5}))
			Expect(list[2]).To(Equal(q15.Coord{Row: 2, Col: 2}))
		})

		It("can find the nearest attack square", func() {
			fight = q15.NewFight(input01)
			Expect(fight.NearestAttack(q15.Coord{Row: 1, Col: 1})).To(Equal(q15.Coord{Row: 1, Col: 3}))
		})

		It("can give the correct next move", func() {
			fight = q15.NewFight(input02)
			attacker := q15.Coord{Row: 1, Col: 2}
			target := fight.NearestAttack(attacker)
			Expect(fight.NextSquare(attacker, target)).To(Equal(q15.Coord{Row: 1, Col: 3}))
		})
	})

	Context("performing a set of moves", func() {
		It("puts things in the right place", func() {
			fight = q15.NewFight(input03)
			fight.Print()
			fight.Step()
			fight.Print()
			Expect(fight.At(q15.Coord{Row: 1, Col: 1})).To(Equal('.'))
			Expect(fight.At(q15.Coord{Row: 1, Col: 2})).To(Equal('G'))
			Expect(fight.At(q15.Coord{Row: 2, Col: 4})).To(Equal('G'))
			Expect(fight.At(q15.Coord{Row: 1, Col: 6})).To(Equal('G'))
			fight.Step()
			fight.Print()
			fight.Step()
			fight.Print()
			Expect(fight.At(q15.Coord{Row: 3, Col: 3})).To(Equal('G'))
			Expect(fight.At(q15.Coord{Row: 3, Col: 4})).To(Equal('E'))
			Expect(fight.At(q15.Coord{Row: 3, Col: 5})).To(Equal('G'))
		})

	})

	Context("Full battle", func() {
		It("gets the right score", func() {
			f := q15.NewFight(input04)
			Expect(f.Run()).To(Equal(27730))
		})
	})

})
