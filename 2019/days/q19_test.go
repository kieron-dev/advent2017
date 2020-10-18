package days_test

import (
	"io/ioutil"
	"strings"

	"github.com/kieron-dev/advent2017/advent2019/grid"
	"github.com/kieron-dev/advent2017/advent2019/tractor"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q19", func() {
	var (
		b    *tractor.Beam
		prog string
	)

	BeforeEach(func() {
		input, err := ioutil.ReadFile("../days/input19")
		if err != nil {
			panic(err)
		}

		prog = strings.TrimSpace(string(input))

		b = tractor.NewBeam()
		b.SetProg(prog)
	})

	It("does part A", func() {
		count := 0
		for i := 0; i < 50; i++ {
			for j := 0; j < 50; j++ {
				if b.IsInBeamRange(grid.NewCoord(i, j)) {
					count++
				}
			}
		}
		Expect(count).To(Equal(215))
	})

	It("does part B", func() {
		coords := b.FirstSquare(100)
		Expect(10000*coords.X() + coords.Y()).To(Equal(7720975))
	})
})
