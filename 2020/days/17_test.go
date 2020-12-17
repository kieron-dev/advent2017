package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/gameoflife"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("17", func() {
	var (
		data *os.File
		cube gameoflife.Cube
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input17")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		data.Close()
	})

	It("does part A", func() {
		cube = gameoflife.NewCube(3)
		cube.Load(data)

		for i := 0; i < 6; i++ {
			cube.Evolve()
		}

		Expect(cube.ActiveCount()).To(Equal(255))
	})

	It("does part B", func() {
		cube = gameoflife.NewCube(4)
		cube.Load(data)

		for i := 0; i < 6; i++ {
			cube.Evolve()
		}

		Expect(cube.ActiveCount()).To(Equal(-1))
	})
})
