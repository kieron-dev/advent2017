package days_test

import (
	"io/ioutil"
	"strings"

	"github.com/kieron-dev/advent2017/advent2019/intcode"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q05", func() {
	var (
		all     []byte
		in, out chan int
		c       *intcode.Computer
	)

	BeforeEach(func() {
		var err error
		all, err = ioutil.ReadFile("./input05")
		if err != nil {
			panic(err)
		}
		in = make(chan int, 1)
		out = make(chan int, 10)
		c = intcode.NewComputer(in, out)
	})

	It("does part A", func() {
		in <- 1

		c.SetInput(strings.TrimSpace(string(all)))
		c.Calculate()
		close(out)

		var last int
		for res := range out {
			last = res
		}
		Expect(last).To(Equal(13210611))
	})

	It("does part B", func() {
		in <- 5
		c.SetInput(strings.TrimSpace(string(all)))
		c.Calculate()

		res := <-out
		Expect(res).To(Equal(584126))
	})
})
