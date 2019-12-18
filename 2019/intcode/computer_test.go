package intcode_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kieron-pivotal/advent2017/advent2019/intcode"
)

var _ = Describe("Intcode Computer", func() {

	var (
		c       *intcode.Computer
		in, out chan int
	)

	BeforeEach(func() {
		in = make(chan int, 20)
		out = make(chan int, 200)
		c = intcode.NewComputer(in, out)
	})

	It("calculates simple inputs", func() {
		c.SetInput("1,0,0,0,99")
		Expect(c.Calculate()).To(Equal(2))
	})

	It("can panic with slice indexes then return -1", func() {
		c.SetInput("1,0,0,0,99")
		c.Prime(12, -2)
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

	It("works with relative base mode", func() {
		input := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"

		c.SetInput(input)
		c.Calculate()
		close(out)
		output := []string{}
		for n := range out {
			output = append(output, fmt.Sprintf("%d", n))
		}
		Expect(strings.Join(output, ",")).To(Equal(input))
	})

	It("works with big ints", func() {
		input := "1102,34915192,34915192,7,4,7,99,0"
		c.SetInput(input)
		c.Calculate()
		n := <-out
		Expect(fmt.Sprintf("%d", n)).To(HaveLen(16))
	})

	It("works with another big input", func() {
		input := "104,1125899906842624,99"
		c.SetInput(input)
		c.Calculate()
		n := <-out
		Expect(n).To(Equal(1125899906842624))
	})

	Context("operations", func() {
		It("does addition", func() {
			input := "1,7,8,0,4,0,99,100,201"
			c.SetInput(input)
			c.Calculate()
			n := <-out
			Expect(n).To(Equal(301))
		})

		It("does addition with immediate numbers to add", func() {
			input := "1101,42,-1,0,4,0,99"
			c.SetInput(input)
			c.Calculate()
			n := <-out
			Expect(n).To(Equal(41))
		})

		It("does addition storing in a relative position", func() {
			input := "109,3,21001,9,-1,10,4,13,99,42"
			c.SetInput(input)
			c.Calculate()
			n := <-out
			Expect(n).To(Equal(41))
		})
	})
})

var _ = Describe("Computer Array", func() {
	var (
		arr   *intcode.ComputerArray
		size  int
		prog  string
		phase []int
		out   int
	)

	BeforeEach(func() {
		size = 5
		arr = intcode.NewArray(size)
		prog = "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
		phase = []int{4, 3, 2, 1, 0}
	})

	JustBeforeEach(func() {
		arr.SetProgram(prog)
		arr.SetPhase(phase)
		arr.WriteInitialInput(0)
		arr.Run()
		out = arr.GetResult()
	})

	It("runs a pipeline", func() {
		Expect(out).To(Equal(43210))
	})
})
