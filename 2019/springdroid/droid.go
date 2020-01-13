package springdroid

import (
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019/intcode"
)

type Droid struct {
	computer *intcode.Computer
	in, out  chan int
}

func NewDroid() *Droid {
	d := Droid{}
	d.in = make(chan int)
	d.out = make(chan int)
	d.computer = intcode.NewComputer(d.in, d.out)

	return &d
}

func (d *Droid) LoadProgram(p string) {
	d.computer.SetInput(p)
}

func (d *Droid) RunProgram() {
	go d.computer.Calculate()
}

func (d *Droid) Output() int {
	return <-d.out
}

func (d *Droid) Input(s string) {
	s = strings.TrimSpace(s)
	for _, b := range s {
		d.in <- int(b)
	}
	d.in <- 10
}
