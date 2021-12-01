package days_test

import (
	"bufio"
	"os"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("01", func() {
	It("does part A", func() {
		nums := getInput01()

		last := nums[0]
		increases := 0
		for i := 1; i < len(nums); i++ {
			if nums[i] > last {
				increases++
			}
			last = nums[i]
		}

		Expect(increases).To(Equal(1387))
	})

	It("does part B", func() {
		nums := getInput01()

		last := nums[0] + nums[1] + nums[2]
		increases := 0
		for i := 1; i < len(nums)-2; i++ {
			sum := last - nums[i-1] + nums[i+2]
			if sum > last {
				increases++
			}
			last = sum
		}

		Expect(increases).To(Equal(1362))
	})
})

func AToI(a string) int {
	n, err := strconv.Atoi(a)
	Expect(err).NotTo(HaveOccurred())

	return n
}

func getInput01() []int {
	input, err := os.Open("input01")
	Expect(err).NotTo(HaveOccurred())
	defer input.Close()

	scanner := bufio.NewScanner(input)
	nums := []int{}
	for scanner.Scan() {
		numStr := scanner.Text()
		nums = append(nums, AToI(numStr))
	}

	return nums
}
