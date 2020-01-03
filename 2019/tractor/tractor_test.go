package tractor_test

import (
	"io/ioutil"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019/grid"
	"github.com/kieron-pivotal/advent2017/advent2019/tractor"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tractor", func() {

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

	It("responds with 1 to (0,0) coord", func() {
		Expect(b.IsInBeamRange(grid.NewCoord(0, 0))).To(BeTrue())
	})

	It("responds with 0 to (1,0) coord", func() {
		Expect(b.IsInBeamRange(grid.NewCoord(1, 0))).To(BeFalse())
	})
})
