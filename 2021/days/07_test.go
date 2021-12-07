package days_test

import (
	"io/ioutil"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("07", func() {
	It("does part A", func() {
		nums := load07()
		min := 10000000
		for i := 0; i < 2000; i++ {
			cost := fuelCost(i, nums)
			if cost < min {
				min = cost
			}
		}

		Expect(min).To(Equal(356958))
	})

	It("does part B", func() {
		nums := load07()
		min := 1000000000000
		for i := 0; i < 2000; i++ {
			cost := fuelCost2(i, nums)
			if cost < min {
				min = cost
			}
		}

		Expect(min).To(Equal(105461913))
	})
})

func fuelCost2(from int, nums []int) int {
	c := 0
	for _, n := range nums {
		d := abs(from - n)
		c += d * (1 + d) / 2
	}

	return c
}

func fuelCost(from int, nums []int) int {
	c := 0
	for _, n := range nums {
		c += abs(from - n)
	}

	return c
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func load07() []int {
	bytes, err := ioutil.ReadFile("input07")
	Expect(err).NotTo(HaveOccurred())
	line := strings.TrimSpace(string(bytes))
	return parseNumList(line, ",")
}
