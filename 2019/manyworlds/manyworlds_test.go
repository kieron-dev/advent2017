package manyworlds_test

import (
	"io"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019/grid"
	"github.com/kieron-pivotal/advent2017/advent2019/manyworlds"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Manyworlds", func() {

	var (
		w    *manyworlds.World
		maze io.Reader
	)

	BeforeEach(func() {
		w = manyworlds.NewWorld()
	})

	JustBeforeEach(func() {
		w.LoadMap(maze)
	})

	Context("a simple map", func() {
		BeforeEach(func() {
			maze = strings.NewReader(`#########
#b.A.@.a#
#########
`)
		})

		It("can read it", func() {
			Expect(w.KeysCount()).To(Equal(2))
		})

		It("records the start position", func() {
			Expect(w.StartPos()).To(Equal(grid.NewCoord(5, 1)))
		})

		It("records the start position as a dot", func() {
			Expect(w.CharAt(w.StartPos())).To(Equal(rune('.')))
		})

		It("takes 8 steps to collect all keys", func() {
			Expect(w.MinStepsToCollectKeys()).To(Equal(8))
		})
	})

	Context("slightly harder", func() {
		BeforeEach(func() {
			maze = strings.NewReader(`########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################
`)
		})

		It("takes 86 steps to collect all keys", func() {
			Expect(w.MinStepsToCollectKeys()).To(Equal(86))
		})
	})

	Context("another slightly harder", func() {
		BeforeEach(func() {
			maze = strings.NewReader(`########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################
`)
		})

		It("takes 132 steps to collect all keys", func() {
			Expect(w.MinStepsToCollectKeys()).To(Equal(132))
		})
	})

	Context("one more slightly harder", func() {
		BeforeEach(func() {
			maze = strings.NewReader(`#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################
`)
		})

		It("takes 136 steps to collect all keys", func() {
			Expect(w.MinStepsToCollectKeys()).To(Equal(136))
		})
	})
})
