package tractor

import (
	"github.com/kieron-pivotal/advent2017/advent2019/grid"
	"github.com/kieron-pivotal/advent2017/advent2019/intcode"
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
