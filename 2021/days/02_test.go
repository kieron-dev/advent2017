package days_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("02", func() {
	It("does part A", func() {
		input, err := os.Open("input02")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		depth := 0
		horiz := 0

		move := func(dir string, n int) {
			switch dir {
			case "up":
				depth -= n
				if depth < 0 {
					fmt.Printf("breach: %d\n", depth)
				}
			case "down":
				depth += n
			case "forward":
				horiz += n
			}
		}

		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, " ")
			Expect(parts).To(HaveLen(2))
			move(parts[0], AToI(parts[1]))
		}

		Expect(depth * horiz).To(Equal(1693300))
	})

	It("does part B", func() {
		input, err := os.Open("input02")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		depth := 0
		horiz := 0
		aim := 0

		move := func(dir string, n int) {
			switch dir {
			case "up":
				aim -= n
			case "down":
				aim += n
			case "forward":
				horiz += n
				depth += n * aim
			}
		}

		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, " ")
			Expect(parts).To(HaveLen(2))
			move(parts[0], AToI(parts[1]))
		}

		Expect(depth * horiz).To(Equal(1857958050))
	})
})
