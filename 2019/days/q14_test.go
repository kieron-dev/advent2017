package days_test

import (
	"io"
	"os"

	"github.com/kieron-pivotal/advent2017/advent2019/fuel"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q14", func() {

	var (
		equations io.Reader
		calc      *fuel.Calculator
	)

	BeforeEach(func() {
		var err error
		equations, err = os.Open("./input14")
		if err != nil {
			panic(err)
		}
		calc = fuel.NewCalculator()
	})

	JustBeforeEach(func() {
		calc.SetProgram(equations)
	})

	It("does part A", func() {
		ore, _ := calc.OreForFuel(map[string]int{})
		Expect(ore).To(Equal(201324))
	})

	FIt("does part B", func() {
		Expect(calc.FuelForOre(1000000000000)).To(Equal(6326857))
	})
})
