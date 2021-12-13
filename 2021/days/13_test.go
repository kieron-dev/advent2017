package days_test

import (
	"bufio"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("13", func() {
	It("does part A", func() {
		grid, folds := load13()

		newGrid := foldGrid(grid, folds[0])
		Expect(len(newGrid)).To(Equal(631))
	})

	It("does part B", func() {
		grid, folds := load13()

		for _, f := range folds {
			grid = foldGrid(grid, f)
		}
		lines := printGrid(grid)
		Expect(lines).To(Equal([]string{
			"#### #### #    ####   ##  ##  ###  ####",
			"#    #    #    #       # #  # #  # #   ",
			"###  ###  #    ###     # #    #  # ### ",
			"#    #    #    #       # # ## ###  #   ",
			"#    #    #    #    #  # #  # # #  #   ",
			"#### #    #### #     ##   ### #  # #   ",
		}))
	})
})

func printGrid(grid map[Coord]bool) []string {
	maxX := 0
	maxY := 0
	for c := range grid {
		if c.R > maxX {
			maxX = c.R
		}
		if c.C > maxY {
			maxY = c.C
		}
	}

	res := []string{}
	for y := 0; y <= maxY; y++ {
		line := ""
		for x := 0; x <= maxX; x++ {
			c := NewCoord(x, y)
			if grid[c] {
				line += "#"
			} else {
				line += " "
			}
		}
		res = append(res, line)
	}

	return res
}

func foldGrid(grid map[Coord]bool, f fold) map[Coord]bool {
	newGrid := map[Coord]bool{}

	for coord := range grid {
		if f.direction == "x" {
			if coord.R < f.position {
				newGrid[coord] = true
			} else {
				newGrid[NewCoord(f.position-(coord.R-f.position), coord.C)] = true
			}
		}
		if f.direction == "y" {
			if coord.C < f.position {
				newGrid[coord] = true
			} else {
				newGrid[NewCoord(coord.R, f.position-(coord.C-f.position))] = true
			}
		}
	}

	return newGrid
}

type fold struct {
	direction string
	position  int
}

func load13() (map[Coord]bool, []fold) {
	input, err := os.Open("input13")
	Expect(err).NotTo(HaveOccurred())
	defer input.Close()

	coords := map[Coord]bool{}
	folds := []fold{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ",") {
			nums := parseNumList(line, ",")
			Expect(nums).To(HaveLen(2))
			coords[NewCoord(nums[0], nums[1])] = true
			continue
		}

		if strings.Contains(line, "fold") {
			splits := strings.Split(line, "=")
			Expect(splits).To(HaveLen(2))
			pos := AToI(splits[1])
			if strings.Contains(line, "x") {
				folds = append(folds, fold{direction: "x", position: pos})
			} else {
				folds = append(folds, fold{direction: "y", position: pos})
			}
		}
	}

	return coords, folds
}
