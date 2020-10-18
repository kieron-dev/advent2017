package days_test

import (
	"io/ioutil"
	"strings"

	"github.com/kieron-dev/advent2017/advent2019/intcode"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q09", func() {
	var (
		progBytes []byte
		in, out   chan int
		c         *intcode.Computer
	)

	BeforeEach(func() {
		var err error
		progBytes, err = ioutil.ReadFile("./input09")
		if err != nil {
			panic(err)
		}

		in = make(chan int, 1)
		out = make(chan int, 20)

		c = intcode.NewComputer(in, out)
		c.SetInput(strings.TrimSpace(string(progBytes)))
	})

	It("does part A", func() {
		in <- 1
		c.Calculate()

		close(out)
		var last int
		for n := range out {
			last = n
		}
		Expect(last).To(Equal(3454977209))
	})

	It("does part B", func() {
		in <- 2
		c.Calculate()

		close(out)
		res := <-out
		Expect(res).To(Equal(50120))
	})
})
