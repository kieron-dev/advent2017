package gameoflife_test

import (
	"io"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kieron-dev/adventofcode/2020/gameoflife"
)

var _ = Describe("Cube", func() {
	var (
		data io.Reader
		cube gameoflife.Cube
	)

	JustBeforeEach(func() {
		data = strings.NewReader(`
.#.
..#
###
`)
		cube.Load(data)
	})

	Context("3D", func() {
		BeforeEach(func() {
			cube = gameoflife.NewCube(3)
		})

		It("has 11 active after 1 cycle", func() {
			cube.Evolve()
			Expect(cube.ActiveCount()).To(Equal(11))
		})

		It("has 21 active after 2 cycles", func() {
			for i := 0; i < 2; i++ {
				cube.Evolve()
			}
			Expect(cube.ActiveCount()).To(Equal(21))
		})

		It("has 112 active after 6 cycles", func() {
			for i := 0; i < 6; i++ {
				cube.Evolve()
			}
			Expect(cube.ActiveCount()).To(Equal(112))
		})
	})

	Context("4D", func() {
		BeforeEach(func() {
			cube = gameoflife.NewCube(4)
		})

		FIt("has 112 active after 6 cycles", func() {
			for i := 0; i < 6; i++ {
				cube.Evolve()
			}
			Expect(cube.ActiveCount()).To(Equal(848))
		})
	})
})
