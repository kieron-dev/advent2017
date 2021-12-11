package days_test

import (
	"bufio"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("11", func() {
	It("does part A", func() {
		grid := load11()

		flashes := 0
		for i := 0; i < 100; i++ {
			var newFlashes int
			grid, newFlashes = advanceOctopuses(grid)
			flashes += newFlashes
		}

		Expect(flashes).To(Equal(1717))
	})

	It("does part B", func() {
		grid := load11()

		count := 0
		for {
			count++
			var flashes int
			grid, flashes = advanceOctopuses(grid)
			if flashes == 100 {
				break
			}
		}

		Expect(count).To(Equal(476))
	})
})

func load11() [][]int {
	input, err := os.Open("input11")
	Expect(err).NotTo(HaveOccurred())
	defer input.Close()

	grid := [][]int{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, r := range line {
			row = append(row, AToI(string(r)))
		}
		grid = append(grid, row)
	}

	return grid
}

func advanceOctopuses(grid [][]int) ([][]int, int) {
	nines := []Coord{}
	newGrid := make([][]int, len(grid))
	for r, row := range grid {
		newRow := make([]int, len(row))
		for c, n := range row {
			newRow[c] = n + 1
			if n == 9 {
				nines = append(nines, NewCoord(r, c))
			}
		}
		newGrid[r] = newRow
	}

	visited := map[Coord]bool{}
	for len(nines) > 0 {
		cur := nines[0]
		nines = nines[1:]
		if visited[cur] {
			continue
		}

		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if i == 0 && j == 0 {
					continue
				}
				if cur.R+i < 0 || cur.R+i >= len(grid) || cur.C+j < 0 || cur.C+j >= len(grid[0]) {
					continue
				}
				newGrid[cur.R+i][cur.C+j]++
				if newGrid[cur.R+i][cur.C+j] > 9 {
					nines = append(nines, NewCoord(cur.R+i, cur.C+j))
				}
			}
		}

		visited[cur] = true
	}

	flashes := 0
	for r, row := range newGrid {
		for c, n := range row {
			if n > 9 {
				newGrid[r][c] = 0
				flashes++
			}
		}
	}

	return newGrid, flashes
}
