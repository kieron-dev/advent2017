package advent2019_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

var _ = Describe("Computer", func() {

	var (
		c *advent2019.Computer
	)

	BeforeEach(func() {
		c = advent2019.NewComputer()
	})

	It("calculates simple inputs", func() {
		c.SetInput("1,0,0,0,99")
		Expect(c.Calculate()).To(Equal(2))
	})

	It("can prime then calculate", func() {
		c.SetInput("1,0,0,0,99")
		c.Prime(12, 02)
		Expect(c.Calculate()).To(Equal(2))
	})

})
