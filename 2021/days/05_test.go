package days_test

import (
	"bufio"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("05", func() {
	It("does part A", func() {
		input, err := os.Open("input05")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		grid := map[[2]int]int{}
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.TrimSpace(line)
			sides := strings.Split(line, " -> ")
			Expect(sides).To(HaveLen(2))

			from := parseNumList(sides[0], ",")
			to := parseNumList(sides[1], ",")

			fromX, fromY := from[0], from[1]
			toX, toY := to[0], to[1]

			dirX := 1
			if fromX > toX {
				dirX = -1
			}
			dirY := 1
			if fromY > toY {
				dirY = -1
			}

			if fromX == toX {
				y := fromY
				for {
					grid[[2]int{fromX, y}]++
					if y == toY {
						break
					}
					y += dirY
				}
				continue
			}

			if fromY == toY {
				x := fromX
				for {
					grid[[2]int{x, fromY}]++
					if x == toX {
						break
					}
					x += dirX
				}
				continue
			}

		}

		count := 0
		for _, n := range grid {
			if n > 1 {
				count++
			}
		}

		Expect(count).To(Equal(5167))
	})

	It("does part B", func() {
		input, err := os.Open("input05")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		grid := map[[2]int]int{}
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.TrimSpace(line)
			sides := strings.Split(line, " -> ")
			Expect(sides).To(HaveLen(2))

			from := parseNumList(sides[0], ",")
			to := parseNumList(sides[1], ",")

			fromX, fromY := from[0], from[1]
			toX, toY := to[0], to[1]

			dirX := 1
			if fromX > toX {
				dirX = -1
			}
			dirY := 1
			if fromY > toY {
				dirY = -1
			}

			if fromX == toX {
				y := fromY
				for {
					grid[[2]int{fromX, y}]++
					if y == toY {
						break
					}
					y += dirY
				}
				continue
			}

			if fromY == toY {
				x := fromX
				for {
					grid[[2]int{x, fromY}]++
					if x == toX {
						break
					}
					x += dirX
				}
				continue
			}

			x, y := fromX, fromY
			for {
				grid[[2]int{x, y}]++

				if x == toX {
					break
				}

				x += dirX
				y += dirY
			}
		}

		count := 0
		for _, n := range grid {
			if n > 1 {
				count++
			}
		}

		Expect(count).To(Equal(17604))
	})
})
