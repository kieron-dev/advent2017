package two022_test

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("10", func() {
	It("does part A", func() {
		f, err := os.Open("input10")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		clock := 0
		x := 1
		ans := 0
		incrClock := func() {
			clock++
			if clock <= 220 && ((clock-20)%40) == 0 {
				ans += clock * x
			}
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			incrClock()
			line := scanner.Text()

			if line == "noop" {
				continue
			}

			var n int
			_, err := fmt.Sscanf(line, "addx %d", &n)
			Expect(err).NotTo(HaveOccurred())

			incrClock()
			x += n
		}

		Expect(ans).To(Equal(12640))
	})

	It("does part B", func() {
		f, err := os.Open("input10")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		clock := 0
		x := 1
		var screen [6][40]byte

		incrClock := func() {
			clock++
			c := (clock - 1) % 40
			r := (clock - 1) / 40
			if x >= c-1 && x <= c+1 {
				screen[r][c] = '#'
			} else {
				screen[r][c] = ' '
			}
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			incrClock()
			line := scanner.Text()

			if line == "noop" {
				continue
			}

			var n int
			_, err := fmt.Sscanf(line, "addx %d", &n)
			Expect(err).NotTo(HaveOccurred())

			incrClock()
			x += n
		}
		fmt.Println()
		for _, r := range screen {
			fmt.Printf("%s\n", r)
		}
	})
})
