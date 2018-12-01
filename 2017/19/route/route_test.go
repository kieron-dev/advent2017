package route_test

import (
	"github.com/kieron-pivotal/advent2017/19/route"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Route", func() {
	Context("baby steps", func() {
		var r *route.Map

		BeforeEach(func() {
			r = route.New([]string{
				"       |  ",
				"       +- ",
			})
		})

		It("can be created with a slice of strings", func() {
			route.New([]string{"", ""})
			Expect(true).To(BeTrue())
		})

		It("can identify the starting point", func() {
			Expect(r.CurPosition()).To(Equal(route.NewPoint(0, 7)))
			Expect(r.CurDirection()).To(Equal(route.Down))
		})

		It("can move one step", func() {
			r.Step()
			Expect(r.CurPosition()).To(Equal(route.NewPoint(1, 7)))
			Expect(r.CurDirection()).To(Equal(route.Down))
		})

		It("can change direction", func() {
			r.Step()
			r.Step()
			Expect(r.CurPosition()).To(Equal(route.NewPoint(1, 8)))
			Expect(r.CurDirection()).To(Equal(route.Right))
		})

		It("can terminate", func() {
			r.Step()
			r.Step()
			Expect(r.Step()).To(BeFalse())
		})
	})

	Context("let them run", func() {
		var r *route.Map
		BeforeEach(func() {
			r = route.New([]string{
				"     |          ",
				"     |  +--+    ",
				"     A  |  C    ",
				" F---|----E|--+ ",
				"     |  |  |  D ",
				"     +B-+  +--+ ",
			})
		})

		It("can follow the route", func() {
			r.Walk()
			Expect(r.CurPosition()).To(Equal(route.NewPoint(3, 1)))
		})

		It("records letters", func() {
			r.Walk()
			Expect(r.GetLetters()).To(Equal("ABCDEF"))
		})

		It("counts steps", func() {
			r.Walk()
			Expect(r.GetStepCount()).To(Equal(38))
		})
	})
})
