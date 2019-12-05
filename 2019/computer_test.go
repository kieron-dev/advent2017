package advent2019_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

var _ = Describe("Computer", func() {

	var (
		c     *advent2019.Computer
		input *gbytes.Buffer
	)

	BeforeEach(func() {
		input = gbytes.NewBuffer()
		c = advent2019.NewComputer(input)
	})

	It("calculates simple inputs", func() {
		c.SetInput("1,0,0,0,99")
		Expect(c.Calculate()).To(Equal(2))
	})

	It("can panic with slice indexes then return -1", func() {
		c.SetInput("1,0,0,0,99")
		c.Prime(12, 02)
		Expect(c.TryCalculate()).To(Equal(-1))
	})

	It("works with new modes", func() {
		c.SetInput("3,0,4,0,99")
		input.Write([]byte("44\n"))
		res := c.Calculate()
		Expect(res).To(Equal(44))
	})

})
