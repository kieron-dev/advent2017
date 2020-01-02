package manyworlds_test

import (
	"io"
	"strings"

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

	Context("multiple start points: 1", func() {
		BeforeEach(func() {
			maze = strings.NewReader(`###############
#d.ABC.#.....a#
######@#@######
###############
######@#@######
#b.....#.....c#
###############
`)
		})

		It("takes 24 steps", func() {
			Expect(w.MinStepsToCollectKeys()).To(Equal(24))
		})
	})

	Context("multiple start points: 2", func() {
		BeforeEach(func() {
			maze = strings.NewReader(`#############
#DcBa.#.GhKl#
#.###@#@#I###
#e#d#####j#k#
###C#@#@###J#
#fEbA.#.FgHi#
#############`)
		})

		It("takes 32 steps", func() {
			Expect(w.MinStepsToCollectKeys()).To(Equal(32))
		})
	})

	Context("multiple start points: 3", func() {
		BeforeEach(func() {
			maze = strings.NewReader(`#############
#g#f.D#..h#l#
#F###e#E###.#
#dCba@#@BcIJ#
#############
#nK.L@#@G...#
#M###N#H###.#
#o#m..#i#jk.#
#############`)
		})

		It("takes 72 steps", func() {
			Expect(w.MinStepsToCollectKeys()).To(Equal(72))
		})
	})
})
