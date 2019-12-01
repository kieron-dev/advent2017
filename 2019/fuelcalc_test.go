package advent2019_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"advent2019"
)

var _ = Describe("FuelCalc", func() {
	var (
		calculator advent2019.FuelCalc
	)

	BeforeEach(func() {
		calculator = advent2019.FuelCalc{}
	})

	DescribeTable("FuelForMass", func(mass, expectedFuel int) {
		fuel := calculator.FuelForMass(mass)
		Expect(fuel).To(Equal(expectedFuel))
	},
		Entry("12", 12, 2),
		Entry("14", 14, 2),
		Entry("1969", 1969, 654),
		Entry("100756", 100756, 33583),
	)

	DescribeTable("FuelForFuel", func(fuel, expectedFuel int) {
		extraFuel := calculator.FuelForFuel(fuel)
		Expect(extraFuel).To(Equal(expectedFuel))
	},
		Entry("2", 2, 0),
		Entry("654", 654, 312),
		Entry("33583", 33583, 50346-33583),
	)
})
