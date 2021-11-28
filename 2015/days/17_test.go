package days_test

import (
	"bufio"
	"os"
	"sort"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("17", func() {
	It("does part A", func() {
		target := 150

		input, err := os.Open("input17")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		scanner := bufio.NewScanner(input)
		sizes := []int{}
		for scanner.Scan() {
			numStr := scanner.Text()
			n, err := strconv.Atoi(numStr)
			Expect(err).NotTo(HaveOccurred())
			sizes = append(sizes, n)
		}

		sort.Ints(sizes)

		Expect(numSolns(sizes, target, 0)).To(Equal(654))
	})

	It("does part B", func() {
		target := 150

		input, err := os.Open("input17")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		scanner := bufio.NewScanner(input)
		sizes := []int{}
		for scanner.Scan() {
			numStr := scanner.Text()
			n, err := strconv.Atoi(numStr)
			Expect(err).NotTo(HaveOccurred())
			sizes = append(sizes, n)
		}

		sort.Ints(sizes)

		Expect(numSolnsOfLength(sizes, target, 3, 0)).To(Equal(57))
	})
})

func numSolns(options []int, target int, length int) int {
	if len(options) == 0 || target < 0 {
		return 0
	}

	first := options[0]

	if first == target {
		return 1 + numSolns(options[1:], target, length)
	}

	return numSolns(options[1:], target-first, length+1) + numSolns(options[1:], target, length)
}

func numSolnsOfLength(options []int, target int, requiredLength, length int) int {
	if len(options) == 0 || target < 0 || length > requiredLength {
		return 0
	}

	first := options[0]

	if first == target {
		return 1 + numSolnsOfLength(options[1:], target, requiredLength, length)
	}

	return numSolnsOfLength(options[1:], target-first, requiredLength, length+1) +
		numSolnsOfLength(options[1:], target, requiredLength, length)
}
