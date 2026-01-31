package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/maps"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("03", func() {
	var (
		grid  maps.Cylinder
		input *os.File
	)

	BeforeEach(func() {
		var err error
		input, err = os.Open("./input03")
		Expect(err).NotTo(HaveOccurred())

		grid = maps.NewCylinder()
		grid.Load(input)
	})

	AfterEach(func() {
		input.Close()
	})

	It("does part A", func() {
		Expect(grid.CountChars(maps.NewCoord(0, 0), maps.NewVector(3, 1), '#')).To(Equal(176))
	})

	It("does part B", func() {
		dirs := []maps.Vector{
			{X: 1, Y: 1},
			{X: 3, Y: 1},
			{X: 5, Y: 1},
			{X: 7, Y: 1},
			{X: 1, Y: 2},
		}

		prod := 1
		for _, dir := range dirs {
			prod *= grid.CountChars(maps.NewCoord(0, 0), dir, '#')
		}

		Expect(prod).To(Equal(5872458240))
	})
})
