package days_test

import (
	"io/ioutil"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019/intcode"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q02", func() {
	var (
		all []byte
	)

	BeforeEach(func() {
		var err error
		all, err = ioutil.ReadFile("./input02")
		if err != nil {
			panic(err)
		}
	})

	It("does part A", func() {
		c := intcode.NewComputer(nil, nil)
		c.SetInput(strings.TrimSpace(string(all)))
		c.Prime(12, 02)
		out := c.Calculate()

		Expect(out).To(Equal(int64(5866714)))
	})

	It("does part B", func() {
		target := int64(19690720)
		var noun, verb int64
		soln := false

	out:
		for noun = 0; noun < 100; noun++ {
			for verb = 0; verb < 100; verb++ {
				c := intcode.NewComputer(nil, nil)
				c.SetInput(strings.TrimSpace(string(all)))
				c.Prime(noun, verb)

				if c.TryCalculate() == target {
					soln = true
					break out
				}
			}
		}
		Expect(soln).To(BeTrue())
		Expect(100*noun + verb).To(Equal(int64(5208)))
	})

})
