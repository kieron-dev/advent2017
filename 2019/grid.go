package advent2019

import (
	"math"
	"strconv"
	"strings"
)

type Grid struct {
	cells map[[2]int]int
	x     int
	y     int
}

func NewGrid() *Grid {
	g := Grid{}
	g.cells = map[[2]int]int{}
	return &g
}

func (g *Grid) Move(moves string) {
	t := 0
	for _, move := range strings.Split(moves, ",") {
		n, err := strconv.Atoi(move[1:])
		if err != nil {
			panic(err)
		}
		delta := 1
		var attr *int

		switch move[0] {
		case 'U':
			attr = &g.y
		case 'D':
			delta = -1
			attr = &g.y
		case 'L':
			delta = -1
			attr = &g.x
		case 'R':
			attr = &g.x
		}

		for i := 0; i < n; i++ {
			*attr += delta
			t++
			g.recordCell(t)
		}
	}
}

func (g *Grid) recordCell(n int) {
	coord := [2]int{g.x, g.y}
	if g.cells[coord] == 0 {
		g.cells[coord] = n
	}
}

func (g *Grid) ClosestCommonDist(h *Grid) int {
	minDist := math.MaxInt32

	for k, _ := range g.cells {
		if h.cells[k] > 0 {
			d := mod(k[0]) + mod(k[1])
			if d < minDist {
				minDist = d
			}
		}
	}
	return minDist
}

func (g *Grid) QuickestCommonCell(h *Grid) int {
	minTime := math.MaxInt32

	for k, c := range g.cells {
		if h.cells[k] > 0 {
			t := c + h.cells[k]
			if t < minTime {
				minTime = t
			}
		}
	}
	return minTime
}

func mod(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
