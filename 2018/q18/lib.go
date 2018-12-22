package q18

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type Coord struct {
	X int
	Y int
}

func C(x, y int) Coord {
	return Coord{X: x, Y: y}
}

type Area struct {
	Grid [][]rune
}

func NewArea(in io.Reader) *Area {
	a := Area{}
	a.Grid = [][]rune{}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, "\n")
		a.Grid = append(a.Grid, []rune(line))
	}

	return &a
}

func (a *Area) Print() {
	for _, row := range a.Grid {
		for _, c := range row {
			fmt.Printf("%s", string(c))
		}
		fmt.Println("")
	}
}

func (a *Area) Step() {
	newGrid := [][]rune{}

	for y, row := range a.Grid {
		newRow := []rune{}
		for x := range row {
			c := C(x, y)
			newRow = append(newRow, a.NextState(c))
		}
		newGrid = append(newGrid, newRow)
	}
	a.Grid = newGrid
}

func (a *Area) At(c Coord) rune {
	if c.X < 0 || c.X >= len(a.Grid[0]) || c.Y < 0 || c.Y >= len(a.Grid) {
		return '.'
	}
	return a.Grid[c.Y][c.X]
}

func (a *Area) Score() int {
	yards := 0
	trees := 0
	for _, row := range a.Grid {
		for _, r := range row {
			if r == '#' {
				yards++
			} else if r == '|' {
				trees++
			}
		}
	}
	return trees * yards
}

func (a *Area) NextState(c Coord) rune {
	switch a.At(c) {
	case '.':
		treeCount := 0
		for _, r := range a.Surrounding(c) {
			if r == '|' {
				treeCount++
			}
		}

		if treeCount > 2 {
			return '|'
		}
		return '.'
	case '|':
		yardCount := 0
		for _, r := range a.Surrounding(c) {
			if r == '#' {
				yardCount++
			}
		}

		if yardCount > 2 {
			return '#'
		}
		return '|'
	case '#':
		yardCount := 0
		treeCount := 0
		for _, r := range a.Surrounding(c) {
			switch r {
			case '#':
				yardCount++
			case '|':
				treeCount++
			}
		}
		if yardCount > 0 && treeCount > 0 {
			return '#'
		}
		return '.'
	default:
		log.Fatal("unknown state")
	}
	return '.'
}

func (a *Area) Surrounding(c Coord) []rune {
	res := []rune{}
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			t := Coord{X: c.X + x, Y: c.Y + y}
			if t == c {
				continue
			}
			res = append(res, a.At(t))
		}
	}
	return res
}
