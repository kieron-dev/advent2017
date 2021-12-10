package days_test

import (
	"bufio"
	"os"
	"sort"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("09", func() {
	It("does part A", func() {
		input, err := os.Open("input09")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		grid := [][]int{}
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.TrimSpace(line)
			lineNums := []int{}
			for i := 0; i < len(line); i++ {
				lineNums = append(lineNums, AToI(string(line[i])))
			}
			grid = append(grid, lineNums)
		}

		sum := 0
		for _, cell := range minima(grid) {
			sum += grid[cell.R][cell.C] + 1
		}

		Expect(sum).To(Equal(498))
	})

	It("does part B", func() {
		input, err := os.Open("input09")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		grid := [][]int{}
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.TrimSpace(line)
			lineNums := []int{}
			for i := 0; i < len(line); i++ {
				lineNums = append(lineNums, AToI(string(line[i])))
			}
			grid = append(grid, lineNums)
		}

		basinSizes := []int{}

		for _, minimum := range minima(grid) {
			basinSizes = append(basinSizes, basinSize(grid, minimum))
		}

		sort.Ints(basinSizes)
		l := len(basinSizes)
		prod := basinSizes[l-1] * basinSizes[l-2] * basinSizes[l-3]
		Expect(prod).To(Equal(1071000))
	})
})

func basinSize(grid [][]int, sink Coord) int {
	next := []Coord{sink}
	size := 0
	visited := map[Coord]bool{}

	for len(next) > 0 {
		cur := next[0]
		next = next[1:]
		if visited[cur] {
			continue
		}
		size++

		for _, neighbour := range squareNeighbours(grid, cur.R, cur.C) {
			nVal := grid[neighbour.R][neighbour.C]
			if nVal < 9 && nVal > grid[cur.R][cur.C] {
				next = append(next, neighbour)
			}
		}

		visited[cur] = true
	}

	return size
}

type Coord struct {
	R, C int
}

func NewCoord(r, c int) Coord {
	return Coord{R: r, C: c}
}

func minima(grid [][]int) []Coord {
	res := []Coord{}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			smallest := true
			for _, neighbour := range squareNeighbours(grid, r, c) {
				if grid[r][c] >= grid[neighbour.R][neighbour.C] {
					smallest = false
					break
				}
			}
			if smallest {
				res = append(res, NewCoord(r, c))
			}
		}
	}

	return res
}

func squareNeighbours(grid [][]int, r, c int) []Coord {
	res := []Coord{}

	if r-1 >= 0 {
		res = append(res, NewCoord(r-1, c))
	}
	if r+1 < len(grid) {
		res = append(res, NewCoord(r+1, c))
	}
	if c-1 >= 0 {
		res = append(res, NewCoord(r, c-1))
	}
	if c+1 < len(grid[r]) {
		res = append(res, NewCoord(r, c+1))
	}

	return res
}
