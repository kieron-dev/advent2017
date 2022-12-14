package two022_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("14", func() {
	It("does part A", func() {
		f, err := os.Open("input14")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		scanner := bufio.NewScanner(f)
		grid := map[Coord]byte{}
		minX := 99999999
		minY := 0
		maxX := -99999999
		maxY := -99999999

		setMinMax := func(c Coord) {
			if c.X < minX {
				minX = c.X
			}
			if c.X > maxX {
				maxX = c.X
			}
			if c.Y < minY {
				minY = c.Y
			}
			if c.Y > maxY {
				maxY = c.Y
			}
		}

		addSegment := func(from, to Coord) {
			var dir string
			if from.X == to.X {
				if from.Y < to.Y {
					dir = "U"
				} else {
					dir = "D"
				}
			} else {
				if from.X < to.X {
					dir = "R"
				} else {
					dir = "L"
				}
			}

			for c := from; c != to; c = c.Move(dir, 1) {
				grid[c] = '#'
			}
			grid[to] = '#'
		}

		for scanner.Scan() {
			line := scanner.Text()
			var last *Coord
			for _, segment := range strings.Split(line, " -> ") {
				var x, y int
				_, err := fmt.Sscanf(segment, "%d,%d", &x, &y)
				Expect(err).NotTo(HaveOccurred())
				c := NewCoord(x, y)
				setMinMax(c)
				if last != nil {
					addSegment(*last, c)
				}
				last = &c
			}
		}

		var count int
		for {
			s := NewCoord(500, 0)
			settled := false
			for s.X >= minX && s.X <= maxX {
				down := s.Move("U", 1)
				if _, filled := grid[down]; !filled {
					s = down
					continue
				}
				left := down.Move("L", 1)
				if _, filled := grid[left]; !filled {
					s = left
					continue
				}
				right := down.Move("R", 1)
				if _, filled := grid[right]; !filled {
					s = right
					continue
				}
				grid[s] = 'o'
				settled = true
				count++
				break
			}
			if !settled {
				break
			}
		}

		Expect(count).To(Equal(696), func() string {
			for y := minY; y <= maxY; y++ {
				for x := minX; x <= maxX; x++ {
					if b, ok := grid[NewCoord(x, y)]; ok {
						fmt.Printf("%c", b)
					} else {
						fmt.Print(".")
					}
				}
				fmt.Println()
			}
			return ""
		})
	})

	It("does part B", func() {
		f, err := os.Open("input14")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		scanner := bufio.NewScanner(f)
		grid := map[Coord]byte{}
		minX := 99999999
		minY := 0
		maxX := -99999999
		maxY := -99999999

		setMinMax := func(c Coord) {
			if c.X < minX {
				minX = c.X
			}
			if c.X > maxX {
				maxX = c.X
			}
			if c.Y < minY {
				minY = c.Y
			}
			if c.Y > maxY {
				maxY = c.Y
			}
		}

		addSegment := func(from, to Coord) {
			var dir string
			if from.X == to.X {
				if from.Y < to.Y {
					dir = "U"
				} else {
					dir = "D"
				}
			} else {
				if from.X < to.X {
					dir = "R"
				} else {
					dir = "L"
				}
			}

			for c := from; c != to; c = c.Move(dir, 1) {
				grid[c] = '#'
			}
			grid[to] = '#'
		}

		for scanner.Scan() {
			line := scanner.Text()
			var last *Coord
			for _, segment := range strings.Split(line, " -> ") {
				var x, y int
				_, err := fmt.Sscanf(segment, "%d,%d", &x, &y)
				Expect(err).NotTo(HaveOccurred())
				c := NewCoord(x, y)
				setMinMax(c)
				if last != nil {
					addSegment(*last, c)
				}
				last = &c
			}
		}

		var count int
		for {
			s := NewCoord(500, 0)
			if _, ok := grid[s]; ok {
				break
			}
			for {
				if s.Y == maxY+1 {
					grid[s] = 'o'
					count++
					break
				}
				down := s.Move("U", 1)
				if _, filled := grid[down]; !filled {
					s = down
					continue
				}
				left := down.Move("L", 1)
				if _, filled := grid[left]; !filled {
					s = left
					continue
				}
				right := down.Move("R", 1)
				if _, filled := grid[right]; !filled {
					s = right
					continue
				}
				grid[s] = 'o'
				count++
				break
			}
		}

		Expect(count).To(Equal(23610))
	})
})
