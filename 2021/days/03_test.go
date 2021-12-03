package days_test

import (
	"bufio"
	"os"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("03", func() {
	getInput03 := func() []int {
		input, err := os.Open("input03")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		nums := []int{}
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			n, err := strconv.ParseInt(line, 2, 64)
			Expect(err).NotTo(HaveOccurred())
			nums = append(nums, int(n))
		}

		return nums
	}

	countOnesInPos := func(nums []int, bit int) int {
		count1 := 0
		for _, n := range nums {
			if n&bit == bit {
				count1++
			}
		}

		return count1
	}

	It("does part A", func() {
		nums := getInput03()

		var gamma, epsilon int

		bit := 1 << 11
		for bit > 0 {
			ones := countOnesInPos(nums, bit)
			if ones > len(nums)/2 {
				gamma |= bit
			} else {
				epsilon |= bit
			}

			bit >>= 1
		}

		Expect(gamma * epsilon).To(BeEquivalentTo(3959450))
	})

	filter := func(nums []int, rev bool) int {
		list := make([]int, len(nums))
		copy(list, nums)

		testbit := 1 << 11
		for testbit > 0 {
			countOn := 0
			for _, n := range list {
				if n&testbit > 0 {
					countOn++
				}
			}

			discardValue := testbit
			if rev {
				discardValue = 0
			}
			if countOn >= len(list)/2 {
				discardValue = 0
				if rev {
					discardValue = testbit
				}
			}

			newNums := []int{}
			for _, n := range list {
				if n&testbit != discardValue {
					newNums = append(newNums, n)
				}
			}

			list = newNums
			if len(list) == 1 {
				break
			}
			testbit >>= 1
		}

		Expect(list).To(HaveLen(1))
		return list[0]
	}

	It("does part B", func() {
		nums := getInput03()

		oxygen := filter(nums, false)
		co2 := filter(nums, true)

		Expect(oxygen * co2).To(Equal(7440311))
	})
})
