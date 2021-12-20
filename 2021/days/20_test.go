package days_test

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("20", func() {
	It("does part A", func() {
		input, err := os.Open("input20")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		instr := ""
		image := map[Coord]bool{}
		row := 0
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) > 500 {
				instr = line
				continue
			}

			if line == "" {
				continue
			}

			for col := 0; col < len(line); col++ {
				if line[col] == '#' {
					image[NewCoord(row, col)] = true
				}
			}

			row++
		}

		for i := 0; i < 2; i++ {
			image = enhance(image, instr, i%2 == 1)
		}

		Expect(len(image)).To(Equal(5680))
	})

	It("does part B", func() {
		input, err := os.Open("input20")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		instr := ""
		image := map[Coord]bool{}
		row := 0
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) > 500 {
				instr = line
				continue
			}

			if line == "" {
				continue
			}

			for col := 0; col < len(line); col++ {
				if line[col] == '#' {
					image[NewCoord(row, col)] = true
				}
			}

			row++
		}

		for i := 0; i < 50; i++ {
			image = enhance(image, instr, i%2 == 1)
		}

		Expect(len(image)).To(Equal(19766))
	})
})

func drawImage(image map[Coord]bool) {
	minR, maxR, minC, maxC := getLimits(image)

	fmt.Println()
	for r := minR; r <= maxR; r++ {
		for c := minC; c <= maxC; c++ {
			if image[NewCoord(r, c)] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func getLimits(image map[Coord]bool) (int, int, int, int) {
	big := 10000000000
	minR, maxR, minC, maxC := big, -big, big, -big
	for c := range image {
		if c.R < minR {
			minR = c.R
		}
		if c.R > maxR {
			maxR = c.R
		}
		if c.C < minC {
			minC = c.C
		}
		if c.C > maxC {
			maxC = c.C
		}
	}

	return minR, maxR, minC, maxC
}

func enhance(image map[Coord]bool, instr string, odd bool) map[Coord]bool {
	res := map[Coord]bool{}

	minR, maxR, minC, maxC := getLimits(image)

	if instr[0] == '#' {
		if !odd {
			minR -= 4
			maxR += 4
			minC -= 4
			maxC += 4
		} else {
			minR += 2
			maxR -= 2
			minC += 2
			maxC -= 2
		}
	} else {
		minR -= 1
		maxR += 1
		minC -= 1
		maxC += 1
	}

	for r := minR; r <= maxR; r++ {
		for c := minC; c <= maxC; c++ {
			idx := 0
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					idx <<= 1
					if image[NewCoord(r+i, c+j)] {
						idx++
					}
				}
			}

			if instr[idx] == '#' {
				res[NewCoord(r, c)] = true
			}
		}
	}

	return res
}
