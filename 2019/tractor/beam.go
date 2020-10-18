package tractor

import (
	"github.com/kieron-dev/advent2017/advent2019/grid"
	"github.com/kieron-dev/advent2017/advent2019/intcode"
)

type Beam struct {
	prog string
}

func NewBeam() *Beam {
	b := Beam{}
	return &b
}

func (b *Beam) SetProg(prg string) {
	b.prog = prg
}

func (b *Beam) IsInBeamRange(c grid.Coord) bool {
	if c.X() < 0 || c.Y() < 0 {
		return false
	}
	in := make(chan int)
	out := make(chan int)
	computer := intcode.NewComputer(in, out)
	computer.SetInput(b.prog)
	go computer.Calculate()
	in <- c.X()
	in <- c.Y()
	resp := <-out
	return resp == 1
}

func (b *Beam) FirstSquare(size int) grid.Coord {
	left := grid.NewCoord(-size+1, 0)
	down := grid.NewCoord(-size+1, size-1)

	right1 := grid.NewCoord(1, 0)
	down1 := grid.NewCoord(0, 1)

	c := grid.NewCoord(size, 0)

	for !(b.IsInBeamRange(c) && b.IsInBeamRange(c.Add(left)) && b.IsInBeamRange(c.Add(down))) {
		c = c.Add(right1)
		for !b.IsInBeamRange(c) {
			c = c.Add(down1)
		}
	}
	return c.Add(left)
}
