package days_test

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("16", func() {
	props := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	It("does part A", func() {
		input, err := os.Open("input16")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		re := regexp.MustCompile(`Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)`)
		scanner := bufio.NewScanner(input)
		sue := ""
		for scanner.Scan() {
			line := scanner.Text()
			matches := re.FindStringSubmatch(line)
			Expect(matches).ToNot(BeNil())

			match := true
			for i := 2; i < 7; i += 2 {
				num, err := strconv.Atoi(matches[i+1])
				Expect(err).NotTo(HaveOccurred())
				if props[matches[i]] != num {
					match = false
					continue
				}
			}
			if match {
				sue = matches[1]
			}
		}

		Expect(sue).To(Equal("103"))
	})

	It("does part B", func() {
		input, err := os.Open("input16")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		re := regexp.MustCompile(`Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)`)
		scanner := bufio.NewScanner(input)
		sue := ""
		for scanner.Scan() {
			line := scanner.Text()
			matches := re.FindStringSubmatch(line)
			Expect(matches).ToNot(BeNil())

			match := true
			for i := 2; i < 7; i += 2 {
				num, err := strconv.Atoi(matches[i+1])
				Expect(err).NotTo(HaveOccurred())
				switch matches[i] {
				case "cats":
					fallthrough
				case "trees":
					if props[matches[i]] >= num {
						match = false
					}
				case "pomeranians":
					fallthrough
				case "goldfish":
					if props[matches[i]] <= num {
						match = false
					}
				default:
					if props[matches[i]] != num {
						match = false
					}
				}
			}
			if match {
				sue = matches[1]
			}
		}

		Expect(sue).To(Equal("405"))
	})
})
