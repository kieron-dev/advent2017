package two022_test

import (
	"bufio"
	"os"
	"sort"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("01", func() {
	It("does part A", func() {
		input, err := os.Open("input01")
		Expect(err).NotTo(HaveOccurred())

		scanner := bufio.NewScanner(input)
		i := 1
		sum := 0
		var maxSum int

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				if sum > maxSum {
					maxSum = sum
				}
				i++
				sum = 0
				continue
			}

			num, err := strconv.Atoi(line)
			Expect(err).NotTo(HaveOccurred())
			sum += num
		}

		if sum > maxSum {
			maxSum = sum
		}

		Expect(maxSum).To(Equal(71934))
	})

	It("does part B", func() {
		input, err := os.Open("input01")
		Expect(err).NotTo(HaveOccurred())
		var cals []int

		scanner := bufio.NewScanner(input)
		var sum int

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				cals = append(cals, sum)
				sum = 0
				continue
			}

			num, err := strconv.Atoi(line)
			Expect(err).NotTo(HaveOccurred())
			sum += num
		}

		cals = append(cals, sum)

		sort.Ints(cals)

		l := len(cals)
		res := cals[l-1] + cals[l-2] + cals[l-3]

		Expect(res).To(Equal(211447))
	})
})
