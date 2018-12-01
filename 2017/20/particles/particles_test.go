package particles_test

import (
	"github.com/kieron-pivotal/advent2017/20/particles"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Particles", func() {
	Context("vectors", func() {
		It("can be added to a vector", func() {
			Expect(particles.NewVector(1, 2, 3).Add(particles.NewVector(4, 2, 1))).To(Equal(particles.NewVector(5, 4, 4)))
		})
	})

	It("calcs distance", func() {
		p0 := particles.New(particles.Vector{}, particles.Vector{}, particles.Vector{}, 0)
		Expect(p0.Distance()).To(Equal(0))
		p1 := particles.New(particles.NewVector(1, 2, 3), particles.Vector{}, particles.Vector{}, 0)
		Expect(p1.Distance()).To(Equal(6))
		p2 := particles.New(particles.NewVector(-1, 2, 3), particles.Vector{}, particles.Vector{}, 0)
		Expect(p2.Distance()).To(Equal(6))
	})

	It("starts at t 0", func() {
		p := particles.New(particles.Vector{}, particles.Vector{}, particles.Vector{}, 0)
		Expect(p.Time()).To(Equal(0))
	})

	It("can advance one time period", func() {
		p := particles.New(
			particles.NewVector(1, 2, 3),
			particles.NewVector(2, 3, 4),
			particles.NewVector(3, 4, 5),
			0,
		)
		p.Advance(1)
		Expect(p.Position()).To(Equal(particles.NewVector(6, 9, 12)))
		Expect(p.Time()).To(Equal(1))
	})

})

var _ = Describe("distribution", func() {
	It("can determine steady ordering", func() {
		d1 := particles.Distribution{
			Particles: []particles.Separation{
				{Id: 1, DistFromLast: 1},
				{Id: 2, DistFromLast: 2},
			},
		}
		d2 := particles.Distribution{
			Particles: []particles.Separation{
				{Id: 1, DistFromLast: 2},
				{Id: 2, DistFromLast: 3},
			},
		}
		Expect(d2.HasSteadyOrder(&d1)).To(BeTrue())
		d3 := particles.Distribution{
			Particles: []particles.Separation{
				{Id: 1, DistFromLast: 2},
				{Id: 2, DistFromLast: 1},
			},
		}
		Expect(d3.HasSteadyOrder(&d1)).To(BeFalse())
	})

	It("can generate distribution from slice of particles", func() {
		ps := []*particles.Particle{
			particles.New(
				particles.NewVector(1, 2, 3),
				particles.NewVector(2, 3, 4),
				particles.NewVector(3, 4, 5),
				1,
			),
			particles.New(
				particles.NewVector(2, 3, 4),
				particles.NewVector(2, 3, 4),
				particles.NewVector(3, 4, 5),
				2,
			),
		}
		distr := particles.DistrFromParticles(ps)
		ExpectedDistr := particles.Distribution{
			Particles: []particles.Separation{
				{Id: 1, DistFromLast: 6},
				{Id: 2, DistFromLast: 3},
			},
		}
		Expect(distr).To(Equal(&ExpectedDistr))
	})
})

var _ = Describe("expansion", func() {
	It("can determine eventual nearest to origin", func() {
		p0 := particles.New(
			particles.NewVector(3, 0, 0),
			particles.NewVector(2, 0, 0),
			particles.NewVector(-1, 0, 0),
			0,
		)
		p1 := particles.New(
			particles.NewVector(4, 0, 0),
			particles.NewVector(0, 0, 0),
			particles.NewVector(-2, 0, 0),
			1,
		)
		ps := []*particles.Particle{p0, p1}
		pClosest := particles.GetEventualClosest(ps)
		Expect(pClosest.Id()).To(Equal(0))
	})
})

var _ = Describe("slice of particles", func() {
	It("can remove duplicate positions", func() {

		p0 := particles.New(
			particles.NewVector(3, 0, 0),
			particles.NewVector(2, 0, 0),
			particles.NewVector(-1, 0, 0),
			0,
		)
		p1 := particles.New(
			particles.NewVector(3, 0, 0),
			particles.NewVector(2, 0, 0),
			particles.NewVector(-1, 0, 0),
			0,
		)
		p2 := particles.New(
			particles.NewVector(3, 1, 0),
			particles.NewVector(2, 0, 0),
			particles.NewVector(-1, 0, 0),
			0,
		)
		p3 := particles.New(
			particles.NewVector(3, 2, 0),
			particles.NewVector(2, 0, 0),
			particles.NewVector(-1, 0, 0),
			0,
		)
		p4 := particles.New(
			particles.NewVector(3, 1, 0),
			particles.NewVector(2, 0, 0),
			particles.NewVector(-1, 0, 0),
			0,
		)

		ps := []*particles.Particle{p0, p1, p2, p3, p4}
		ps = particles.RemoveCollisions(ps)
		Expect(len(ps)).To(Equal(1))
	})
})
