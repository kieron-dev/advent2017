package days_test

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const gridSize = 100

var _ = Describe("18", func() {
	It("does part A", func() {
		input, err := os.Open("input18")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		grid := [gridSize * gridSize]bool{}
		scanner := bufio.NewScanner(input)
		r := 0
		for scanner.Scan() {
			line := scanner.Text()
			for i, c := range line {
				if c == '#' {
					grid[r*gridSize+i] = true
				}
			}
			r++
		}

		for i := 0; i < 100; i++ {
			grid = evolve(grid)
		}

		count := 0
		for r := 0; r < gridSize; r++ {
			for c := 0; c < gridSize; c++ {
				if grid[gridSize*r+c] {
					count++
				}
			}
		}

		Expect(count).To(Equal(814))
	})

	It("does part B", func() {
		input, err := os.Open("input18")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		grid := [gridSize * gridSize]bool{}
		scanner := bufio.NewScanner(input)
		r := 0
		for scanner.Scan() {
			line := scanner.Text()
			for i, c := range line {
				if c == '#' {
					grid[r*gridSize+i] = true
				}
			}
			r++
		}
		grid[0] = true
		grid[gridSize-1] = true
		grid[(gridSize-1)*gridSize] = true
		grid[gridSize*gridSize-1] = true

		for i := 0; i < 100; i++ {
			grid = evolve(grid)
			grid[0] = true
			grid[gridSize-1] = true
			grid[(gridSize-1)*gridSize] = true
			grid[gridSize*gridSize-1] = true
		}

		count := 0
		for r := 0; r < gridSize; r++ {
			for c := 0; c < gridSize; c++ {
				if grid[gridSize*r+c] {
					count++
				}
			}
		}

		Expect(count).To(Equal(924))
	})
})

func printGrid(grid [gridSize * gridSize]bool) {
	fmt.Println()
	for r := 0; r < gridSize; r++ {
		for c := 0; c < gridSize; c++ {
			if grid[r*gridSize+c] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func evolve(grid [gridSize * gridSize]bool) [gridSize * gridSize]bool {
	newGrid := [gridSize * gridSize]bool{}

	for r := 0; r < gridSize; r++ {
		for c := 0; c < gridSize; c++ {
			count := countNeighbours(r, c, grid)
			switch grid[r*gridSize+c] {
			case true:
				newGrid[gridSize*r+c] = (count == 2 || count == 3)
			case false:
				newGrid[gridSize*r+c] = count == 3
			}
		}
	}

	return newGrid
}

func countNeighbours(r, c int, grid [gridSize * gridSize]bool) int {
	count := 0

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if r+i < 0 || r+i >= gridSize {
				continue
			}

			if c+j < 0 || c+j >= gridSize {
				continue
			}

			if grid[gridSize*(r+i)+c+j] {
				count++
			}
		}
	}

	return count
}
