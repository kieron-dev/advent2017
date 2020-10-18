package vacuumbot

import (
	"github.com/kieron-dev/advent2017/advent2019/grid"
	"github.com/kieron-dev/advent2017/advent2019/intcode"
)

type System struct {
	in, out  chan int
	computer *intcode.Computer
	grid     map[grid.Coord]byte
}

func NewSystem() *System {
	s := System{}
	s.in = make(chan int, 21)
	s.out = make(chan int, 10000)
	s.computer = intcode.NewComputer(s.in, s.out)
	s.grid = map[grid.Coord]byte{}

	return &s
}

func (s *System) SetProg(prog string) {
	s.computer.SetInput(prog)
}

func (s *System) Poke(addr, val int) {
	s.computer.SetAddr(addr, val)
}

func (s *System) Run() {
	go s.computer.Calculate()
}

func (s *System) Input(in string) {
	for _, b := range in {
		s.in <- int(b)
	}
	s.in <- 10
}

func (s *System) AcquireGrid() {
	row := 0
	col := 0
	for _, c := range s.GetOutput() {
		if c == 10 {
			row++
			col = 0
			continue
		}
		char := byte(c)
		if char != '.' {
			s.grid[grid.NewCoord(col, row)] = char
		}
		col++
	}
}

func (s *System) GetCrossovers() []grid.Coord {
	var out []grid.Coord

	for coord, char := range s.grid {
		if char != '#' {
			continue
		}
		north := coord.Move(grid.North)
		south := coord.Move(grid.South)
		east := coord.Move(grid.East)
		west := coord.Move(grid.West)
		if s.grid[north] == '#' &&
			s.grid[south] == '#' &&
			s.grid[west] == '#' &&
			s.grid[east] == '#' {
			out = append(out, coord)
		}
	}
	return out
}

func (s *System) GetOutputChan() chan int {
	return s.out
}

func (s *System) GetOutput() []int {
	var out []int
	var last int
	for n := range s.out {
		out = append(out, n)
		if n == 10 && last == 10 {
			break
		}
		last = n
	}
	return out
}
