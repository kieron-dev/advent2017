package days_test

import (
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q05", func() {
	var (
		all []byte
		in  chan big.Int
		out chan big.Int
		c   *advent2019.Computer
	)

	BeforeEach(func() {
		var err error
		all, err = ioutil.ReadFile("./input05")
		if err != nil {
			panic(err)
		}
		in = make(chan big.Int, 1)
		out = make(chan big.Int, 10)
		c = advent2019.NewComputer(in, out)
	})

	It("does part A", func() {
		in <- *big.NewInt(1)

		c.SetInput(strings.TrimSpace(string(all)))
		c.Calculate()
		close(out)

		var last string
		for res := range out {
			last = res.String()
		}
		Expect(last).To(Equal("13210611"))
	})

	It("does part B", func() {
		in <- *big.NewInt(5)
		c.SetInput(strings.TrimSpace(string(all)))
		c.Calculate()

		res := <-out
		Expect(res.String()).To(Equal("584126"))
	})
})
