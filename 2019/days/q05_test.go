package days_test

import (
	"io/ioutil"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019/intcode"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q05", func() {
	var (
		all []byte
		in  chan int64
		out chan int64
		c   *intcode.Computer
	)

	BeforeEach(func() {
		var err error
		all, err = ioutil.ReadFile("./input05")
		if err != nil {
			panic(err)
		}
		in = make(chan int64, 1)
		out = make(chan int64, 10)
		c = intcode.NewComputer(in, out)
	})

	It("does part A", func() {
		in <- 1

		c.SetInput(strings.TrimSpace(string(all)))
		c.Calculate()
		close(out)

		var last int64
		for res := range out {
			last = res
		}
		Expect(last).To(Equal(int64(13210611)))
	})

	It("does part B", func() {
		in <- 5
		c.SetInput(strings.TrimSpace(string(all)))
		c.Calculate()

		res := <-out
		Expect(res).To(Equal(int64(584126)))
	})
})
