package advent2019_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

var _ = Describe("Computer", func() {

	var (
		c   *advent2019.Computer
		in  chan int
		out chan int
	)

	BeforeEach(func() {
		in = make(chan int, 2)
		out = make(chan int, 2)
		c = advent2019.NewComputer(in, out)
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
		in <- 44
		res := c.Calculate()
		Expect(res).To(Equal(44))
		output := <-out
		Expect(output).To(Equal(44))
	})

})

var _ = Describe("Computer Array", func() {
	var (
		arr   *advent2019.ComputerArray
		size  int
		prog  string
		phase []int
		out   int
	)

	BeforeEach(func() {
		size = 5
		arr = advent2019.NewArray(size)
		prog = "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
		phase = []int{4, 3, 2, 1, 0}
	})

	JustBeforeEach(func() {
		arr.SetProgram(prog)
		go func() { arr.SetPhase(phase) }()
		go func() { arr.WriteInitialInput(0) }()
		out = arr.Run()
	})

	It("runs a pipeline", func() {
		Expect(out).To(Equal(43210))
	})
})
