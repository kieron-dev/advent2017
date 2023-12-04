package two023_test

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("04", func() {
	countWins := func(line string) int {
		parts := strings.Split(line, ":")
		Expect(parts).To(HaveLen(2))

		parts2 := strings.Split(parts[1], "|")
		Expect(parts2).To(HaveLen(2))

		winners := []int{}
		for _, ns := range strings.Split(strings.TrimSpace(parts2[0]), " ") {
			if ns == "" {
				continue
			}
			n, err := strconv.Atoi(ns)
			Expect(err).NotTo(HaveOccurred())
			winners = append(winners, n)
		}

		yours := []int{}
		for _, ns := range strings.Split(strings.TrimSpace(parts2[1]), " ") {
			if ns == "" {
				continue
			}
			n, err := strconv.Atoi(ns)
			Expect(err).NotTo(HaveOccurred())
			yours = append(yours, n)
		}

		winSet := map[int]bool{}
		for _, n := range winners {
			winSet[n] = true
		}

		count := 0
		for _, n := range yours {
			if winSet[n] {
				count++
			}
		}

		return count
	}

	It("does part A", func() {
		f, err := os.Open("input04")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		sum := 0

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}

			winCount := countWins(line)
			score := int(math.Pow(2.0, float64(winCount-1)))

			sum += score
		}

		Expect(sum).To(Equal(18653))
	})

	It("does part B", func() {
		f, err := os.Open("input04")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		wins := []int{}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}

			wins = append(wins, countWins(line))
		}

		counts := make([]int, len(wins))
		for i := range counts {
			counts[i] = 1
		}
		for i, n := range wins {
			for j := 0; j < n; j++ {
				counts[i+1+j] += counts[i]
			}
		}

		sum := 0
		for _, n := range counts {
			sum += n
		}

		Expect(sum).To(Equal(5921508))
	})
})
