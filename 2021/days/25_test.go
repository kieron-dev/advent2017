package days_test

import (
	"bufio"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("25", func() {
	It("does part A", func() {
		input, err := os.Open("input25")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		grid := [][]byte{}
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			row := []byte(line)
			grid = append(grid, row)
		}

		i := 0
		for {
			i++
			if !advance(grid) {
				break
			}
		}

		Expect(i).To(Equal(492))
	})
})

func advance(grid [][]byte) bool {
	rows := len(grid)
	cols := len(grid[0])

	any := false

	easts := []Coord{}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			t := grid[r][c]
			if t == '>' && grid[r][(c+1)%cols] == '.' {
				easts = append(easts, NewCoord(r, c))
				any = true
			}
		}
	}

	for _, e := range easts {
		grid[e.R][e.C] = '.'
		grid[e.R][(e.C+1)%cols] = '>'
	}

	souths := []Coord{}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			t := grid[r][c]
			if t == 'v' && grid[(r+1)%rows][c] == '.' {
				souths = append(souths, NewCoord(r, c))
				any = true
			}
		}
	}

	for _, s := range souths {
		grid[s.R][s.C] = '.'
		grid[(s.R+1)%rows][s.C] = 'v'
	}

	return any
}
