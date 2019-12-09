package days_test

import (
	"os"
	"strconv"

	"github.com/kieron-pivotal/advent2017/advent2019"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q01", func() {
	var (
		file       *os.File
		fileReader advent2019.FileReader
		fuelCalc   advent2019.FuelCalc
		fuel       int
	)

	BeforeEach(func() {
		var err error
		file, err = os.Open("./input01")
		if err != nil {
			panic(err)
		}
		fileReader = advent2019.FileReader{}
		fuelCalc = advent2019.FuelCalc{}
		fuel = 0
	})

	It("does part A", func() {
		fileReader.Each(file, func(line string) {
			n, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			fuel += fuelCalc.FuelForMass(n)
		})

		Expect(fuel).To(Equal(3423511))
	})

	It("does part B", func() {
		fileReader.Each(file, func(line string) {
			n, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			origFuel := fuelCalc.FuelForMass(n)
			extraFuel := fuelCalc.FuelForFuel(origFuel)
			fuel += origFuel + extraFuel
		})
		Expect(fuel).To(Equal(5132379))
	})
})
