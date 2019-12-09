package advent2019_test

import (
	"math/big"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

var _ = Describe("Computer", func() {

	var (
		c   *advent2019.Computer
		in  chan big.Int
		out chan big.Int
	)

	BeforeEach(func() {
		in = make(chan big.Int, 20)
		out = make(chan big.Int, 200)
		c = advent2019.NewComputer(in, out)
	})

	It("calculates simple inputs", func() {
		c.SetInput("1,0,0,0,99")
		Expect(c.Calculate().Int64()).To(Equal(int64(2)))
	})

	It("can panic with slice indexes then return -1", func() {
		c.SetInput("1,0,0,0,99")
		c.Prime(12, -2)
		Expect(c.TryCalculate().Int64()).To(Equal(int64(-1)))
	})

	It("works with new modes", func() {
		c.SetInput("3,0,4,0,99")
		in <- *big.NewInt(44)
		res := c.Calculate()
		Expect(res.Int64()).To(Equal(int64(44)))
		output := <-out
		Expect((&output).Cmp(big.NewInt(44))).To(Equal(0))
	})

	It("works with relative base mode", func() {
		input := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"

		c.SetInput(input)
		c.Calculate()
		close(out)
		output := []string{}
		for n := range out {
			output = append(output, n.String())
		}
		Expect(strings.Join(output, ",")).To(Equal(input))
	})

	It("works with big ints", func() {
		input := "1102,34915192,34915192,7,4,7,99,0"
		c.SetInput(input)
		c.Calculate()
		n := <-out
		Expect(n.String()).To(HaveLen(16))
	})

	It("works with another big input", func() {
		input := "104,1125899906842624,99"
		c.SetInput(input)
		c.Calculate()
		n := <-out
		Expect(n.String()).To(Equal("1125899906842624"))
	})

})

var _ = Describe("Computer Array", func() {
	var (
		arr   *advent2019.ComputerArray
		size  int
		prog  string
		phase []int64
		out   *big.Int
	)

	BeforeEach(func() {
		size = 5
		arr = advent2019.NewArray(size)
		prog = "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
		phase = []int64{4, 3, 2, 1, 0}
	})

	JustBeforeEach(func() {
		arr.SetProgram(prog)
		arr.SetPhase(phase)
		arr.WriteInitialInput(0)
		arr.Run()
		out = arr.GetResult()
	})

	It("runs a pipeline", func() {
		Expect(out.Cmp(big.NewInt(43210))).To(Equal(0))
	})
})
