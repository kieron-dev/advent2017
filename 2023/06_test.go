package two023_test

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("06", func() {
	getMinTime := func(t, d int) int {
		for i := 1; i < t; i++ {
			dist := i * (t - i)
			if dist > d {
				return i
			}
		}
		return 0
	}

	getMaxTime := func(t, d int) int {
		for i := t - 1; i > 0; i-- {
			dist := i * (t - i)
			if dist > d {
				return i
			}
		}
		return 0
	}

	It("does part A", func() {
		f, err := os.Open("input06")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		scanner := bufio.NewScanner(f)
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, ":")
		times := alisttoi(parts[1])

		scanner.Scan()
		line = scanner.Text()
		parts = strings.Split(line, ":")
		distances := alisttoi(parts[1])

		prod := 1
		for i := range times {
			t := times[i]
			d := distances[i]

			minTime := getMinTime(t, d)
			maxTime := getMaxTime(t, d)

			prod *= maxTime - minTime + 1
		}

		Expect(prod).To(Equal(625968))
	})

	It("does part B", func() {
		f, err := os.Open("input06")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		scanner := bufio.NewScanner(f)
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, ":")
		ns := strings.ReplaceAll(parts[1], " ", "")
		t, err := strconv.Atoi(ns)
		Expect(err).NotTo(HaveOccurred())

		scanner.Scan()
		line = scanner.Text()
		parts = strings.Split(line, ":")
		ns = strings.ReplaceAll(parts[1], " ", "")
		d, err := strconv.Atoi(ns)
		Expect(err).NotTo(HaveOccurred())

		minTime := getMinTime(t, d)
		maxTime := getMaxTime(t, d)

		Expect(maxTime - minTime + 1).To(Equal(71503))
	})
})
