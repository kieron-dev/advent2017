package robot

import (
	"fmt"

	"github.com/kieron-dev/advent2017/advent2019/grid"
	"github.com/kieron-dev/advent2017/advent2019/intcode"
)

type Color int

const (
	Black Color = 0
	White Color = 1
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Robot struct {
	computer   *intcode.Computer
	pos        grid.Coord
	in, out    chan int
	direction  Direction
	grid       map[grid.Coord]Color
	minX, maxX int
	minY, maxY int
}

func New() *Robot {
	r := Robot{}
	r.in = make(chan int, 10)
	r.out = make(chan int, 10)
	r.computer = intcode.NewComputer(r.in, r.out)
	r.grid = map[grid.Coord]Color{}
	r.maxX = -1000
	r.maxY = -1000
	r.minX = 1000
	r.minY = 1000
	return &r
}

func (r *Robot) RunProg(prog string) {
	r.computer.SetInput(prog)
	r.computer.Calculate()
}

func (r *Robot) Visited() int {
	return len(r.grid)
}

func (r *Robot) Move() {
	color := r.grid[r.pos]

	if r.pos.X() > r.maxX {
		r.maxX = r.pos.X()
	}
	if r.pos.X() < r.minX {
		r.minX = r.pos.X()
	}
	if r.pos.Y() > r.maxY {
		r.maxY = r.pos.Y()
	}
	if r.pos.Y() < r.minY {
		r.minY = r.pos.Y()
	}

	r.in <- int(color)

	newColor := <-r.out
	r.grid[r.pos] = Color(newColor)
	if r.grid[r.pos] < 0 || r.grid[r.pos] > 1 {
		panic(fmt.Sprintf("invalid return color: %d", r.grid[r.pos]))
	}

	turn := <-r.out
	if turn == 0 {
		r.direction = (r.direction + 3) % 4
	} else {
		r.direction = (r.direction + 1) % 4
	}

	switch r.direction {
	case Up:
		r.pos = r.pos.Add(grid.NewCoord(0, -1))
	case Down:
		r.pos = r.pos.Add(grid.NewCoord(0, 1))
	case Left:
		r.pos = r.pos.Add(grid.NewCoord(-1, 0))
	case Right:
		r.pos = r.pos.Add(grid.NewCoord(1, 0))
	}
}

func (r *Robot) GridToString() string {
	out := ""
	for row := r.minY; row <= r.maxY; row++ {
		for col := r.minX; col <= r.maxX; col++ {
			if r.grid[grid.NewCoord(col, row)] == White {
				out += "#"
			} else {
				out += " "
			}
		}
		out += "\n"
	}

	return out
}

func (r *Robot) Set(c grid.Coord, color Color) {
	r.grid[c] = color
}
