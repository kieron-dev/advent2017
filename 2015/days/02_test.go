package days_test

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("02", func() {
	It("does part A", func() {
		file, err := os.Open("input02")
		Expect(err).NotTo(HaveOccurred())
		defer file.Close()

		scanner := bufio.NewScanner(file)
		area := 0

		for scanner.Scan() {
			line := scanner.Text()
			line = strings.TrimSpace(line)

			nums := strings.Split(line, "x")
			Expect(nums).To(HaveLen(3))

			a, err := strconv.Atoi(nums[0])
			Expect(err).NotTo(HaveOccurred())
			b, err := strconv.Atoi(nums[1])
			Expect(err).NotTo(HaveOccurred())
			c, err := strconv.Atoi(nums[2])
			Expect(err).NotTo(HaveOccurred())

			dims := []int{a, b, c}
			sort.Ints(dims)

			area += 2*(dims[0]*dims[1]+dims[0]*dims[2]+dims[1]*dims[2]) + dims[0]*dims[1]
		}

		Expect(area).To(Equal(1586300))
	})

	It("does part B", func() {
		file, err := os.Open("input02")
		Expect(err).NotTo(HaveOccurred())
		defer file.Close()

		scanner := bufio.NewScanner(file)
		length := 0

		for scanner.Scan() {
			line := scanner.Text()
			line = strings.TrimSpace(line)

			nums := strings.Split(line, "x")
			Expect(nums).To(HaveLen(3))

			a, err := strconv.Atoi(nums[0])
			Expect(err).NotTo(HaveOccurred())
			b, err := strconv.Atoi(nums[1])
			Expect(err).NotTo(HaveOccurred())
			c, err := strconv.Atoi(nums[2])
			Expect(err).NotTo(HaveOccurred())

			dims := []int{a, b, c}
			sort.Ints(dims)

			length += 2*(dims[0]+dims[1]) + dims[0]*dims[1]*dims[2]
		}

		Expect(length).To(Equal(3737498))
	})
})
