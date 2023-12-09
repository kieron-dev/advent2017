package two023_test

import (
	"bufio"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func reduce(nums []int) []int {
	res := make([]int, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		res[i] = nums[i+1] - nums[i]
	}
	return res
}

func allZero(nums []int) bool {
	s := map[int]bool{}
	for _, n := range nums {
		s[n] = true
	}
	return len(s) == 1 && s[0]
}

func extrapolateString(line string) int {
	nums := alisttoi(line)
	return extrapolate(nums)
}

func extrapolate(nums []int) int {
	numsLs := [][]int{nums}

	for !allZero(numsLs[len(numsLs)-1]) {
		numsLs = append(numsLs, reduce(numsLs[len(numsLs)-1]))
	}

	add := 0
	for i := len(numsLs) - 2; i >= 0; i-- {
		add += numsLs[i+1][len(numsLs[i+1])-1]
	}

	return nums[len(nums)-1] + add
}

func extrapolateBack(line string) int {
	nums := alisttoi(line)
	l := len(nums)
	copy := make([]int, l)
	for i, n := range nums {
		copy[l-1-i] = n
	}
	return extrapolate(copy)
}

var _ = DescribeTable("extrapolate", func(in string, expected int) {
	Expect(extrapolateString(in)).To(Equal(expected))
},
	Entry("0 3 6 9 12 15", "0 3 6 9 12 15", 18),
	Entry("1 3 6 10 15 21", "1 3 6 10 15 21", 28),
	Entry("10 13 16 21 30 45", "10 13 16 21 30 45", 68),
)

var _ = DescribeTable("extrapolateBack", func(in string, expected int) {
	Expect(extrapolateBack(in)).To(Equal(expected))
},
	Entry("0 3 6 9 12 15", "0 3 6 9 12 15", -3),
	Entry("1 3 6 10 15 21", "1 3 6 10 15 21", 0),
	Entry("10 13 16 21 30 45", "10 13 16 21 30 45", 5),
)

var _ = Describe("day 09", func() {
	It("does part A", func() {
		f, err := os.Open("input09")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		sum := 0
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}

			sum += extrapolateString(line)
		}

		Expect(sum).To(Equal(1938800261))
	})

	It("does part B", func() {
		f, err := os.Open("input09")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		sum := 0
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}

			sum += extrapolateBack(line)
		}

		Expect(sum).To(Equal(1112))
	})
})
