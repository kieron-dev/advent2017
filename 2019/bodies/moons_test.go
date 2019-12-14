package bodies_test

import (
	"io"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019/bodies"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Moons", func() {

	var (
		s            *bodies.System
		initialState io.Reader
	)

	BeforeEach(func() {
		s = bodies.NewSystem()
	})

	JustBeforeEach(func() {
		s.Load(initialState)
	})

	Context("example steps", func() {
		BeforeEach(func() {
			initialState = strings.NewReader(`<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>
`)
		})

		It("can load initial state", func() {
			Expect(s.Moons()).To(HaveLen(4))
			Expect(s.Moons()[0].Pos()).To(Equal(bodies.NewCoord3(-1, 0, 2)))
			Expect(s.Moons()[1].Pos()).To(Equal(bodies.NewCoord3(2, -10, -7)))
		})

		It("can do the first tick", func() {
			s.Tick()
			Expect(s.Moons()[0].Pos()).To(Equal(bodies.NewCoord3(2, -1, 1)))
			Expect(s.Moons()[1].Pos()).To(Equal(bodies.NewCoord3(3, -7, -4)))
			Expect(s.Moons()[2].Pos()).To(Equal(bodies.NewCoord3(1, -7, 5)))
			Expect(s.Moons()[3].Pos()).To(Equal(bodies.NewCoord3(2, 2, 0)))

			Expect(s.Moons()[0].Vel()).To(Equal(bodies.NewCoord3(3, -1, -1)))
			Expect(s.Moons()[1].Vel()).To(Equal(bodies.NewCoord3(1, 3, 3)))
			Expect(s.Moons()[2].Vel()).To(Equal(bodies.NewCoord3(-3, 1, -3)))
			Expect(s.Moons()[3].Vel()).To(Equal(bodies.NewCoord3(-1, -3, 1)))
		})

		It("can do the first 10 ticks", func() {
			for i := 0; i < 10; i++ {
				s.Tick()
			}
			Expect(s.Moons()[0].Pos()).To(Equal(bodies.NewCoord3(2, 1, -3)))
			Expect(s.Moons()[1].Pos()).To(Equal(bodies.NewCoord3(1, -8, 0)))
			Expect(s.Moons()[2].Pos()).To(Equal(bodies.NewCoord3(3, -6, 1)))
			Expect(s.Moons()[3].Pos()).To(Equal(bodies.NewCoord3(2, 0, 4)))

			Expect(s.Moons()[0].Vel()).To(Equal(bodies.NewCoord3(-3, -2, 1)))
			Expect(s.Moons()[1].Vel()).To(Equal(bodies.NewCoord3(-1, 1, 3)))
			Expect(s.Moons()[2].Vel()).To(Equal(bodies.NewCoord3(3, 2, -3)))
			Expect(s.Moons()[3].Vel()).To(Equal(bodies.NewCoord3(1, -1, -1)))
		})

		It("can calculate the energy in the system", func() {
			for i := 0; i < 10; i++ {
				s.Tick()
			}
			Expect(s.TotalEnergy()).To(Equal(179))
		})

		Context("getting back to the same state", func() {

			BeforeEach(func() {
				initialState = strings.NewReader(`<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>
`)
			})

			It("finds the first X initial state repeat", func() {
				Expect(s.FirstXRepeat()).To(Equal(2028))
			})

			It("finds the first Y initial state repeat", func() {
				Expect(s.FirstYRepeat()).To(Equal(5898))
			})

			It("finds the first Z initial state repeat", func() {
				Expect(s.FirstZRepeat()).To(Equal(4702))
			})

			It("finds the first whole repeat", func() {
				Expect(s.FirstRepeat()).To(Equal(int64(4686774924)))
			})
		})
	})

})
