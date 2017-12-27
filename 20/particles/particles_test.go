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
		Expect(p0.ManhattanDistance()).To(Equal(float64(0)))
		p1 := particles.New(particles.NewVector(1, 2, 3), particles.Vector{}, particles.Vector{}, 0)
		Expect(p1.ManhattanDistance()).To(Equal(float64(6)))
	})

	It("calcs pos after time t", func() {
		p := particles.New(
			particles.NewVector(1, 2, 3),
			particles.NewVector(2, 3, 4),
			particles.NewVector(3, 4, 5),
			0,
		)
		Expect(p.Position(1)).To(Equal(particles.NewVector(6, 9, 12)))
	})

})
