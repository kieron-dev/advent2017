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
	t     int
}

func NewGrid() *Grid {
	g := Grid{}
	g.cells = map[[2]int]int{}
	return &g
}

func (g *Grid) Move(moves string) {
	for _, move := range strings.Split(moves, ",") {
		g.advance(move)
	}
}

func (g *Grid) advance(move string) {
	n := getLength(move)
	moveFn := getMoveFn(move)

	for i := 0; i < n; i++ {
		moveFn(g)
		g.t++
		g.recordCell()
	}
}

func getLength(move string) int {
	n, err := strconv.Atoi(move[1:])
	if err != nil {
		panic(err)
	}
	return n
}

func getMoveFn(move string) func(g *Grid) {
	switch move[0] {
	case 'U':
		return func(g *Grid) { g.y++ }
	case 'D':
		return func(g *Grid) { g.y-- }
	case 'L':
		return func(g *Grid) { g.x-- }
	case 'R':
		return func(g *Grid) { g.x++ }
	}
	panic("eh?")
}

func (g *Grid) recordCell() {
	coord := [2]int{g.x, g.y}
	if g.cells[coord] == 0 {
		g.cells[coord] = g.t
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
