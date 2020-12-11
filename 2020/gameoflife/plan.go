// Package gameoflife does Conway's game of life
package gameoflife

import (
	"bufio"
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/maps"
)

type SeatingPlan struct {
	state map[maps.Coord]byte
	rows  int
	cols  int
}

func NewSeatingPlan() SeatingPlan {
	return SeatingPlan{
		state: map[maps.Coord]byte{},
	}
}

func (p *SeatingPlan) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		p.AddLine(line)
	}
}

func (p *SeatingPlan) AddLine(line string) {
	for col := 0; col < len(line); col++ {
		coord := maps.NewCoord(col, p.rows)
		p.state[coord] = line[col]
	}

	p.rows++
	p.cols = len(line)
}

func (p SeatingPlan) State() string {
	res := ""

	for r := 0; r < p.rows; r++ {
		line := ""
		for c := 0; c < p.cols; c++ {
			coord := maps.NewCoord(c, r)
			line += string(p.state[coord])
		}
		res += line + "\n"
	}

	return res
}

func (p SeatingPlan) OccupiedAround(coord maps.Coord) int {
	n := 0

	for _, c := range coord.Neighbours(p.rows, p.cols) {
		if p.state[c] == '#' {
			n++
		}
	}

	return n
}

// Evolve returns true if a change has occurred
func (p *SeatingPlan) Evolve() bool {
	newState := map[maps.Coord]byte{}
	change := false

	for r := 0; r < p.rows; r++ {
		for c := 0; c < p.cols; c++ {
			coord := maps.NewCoord(c, r)
			occupiedAround := p.OccupiedAround(coord)

			if p.state[coord] == 'L' && occupiedAround == 0 {
				newState[coord] = '#'
				change = true
				continue
			}

			if p.state[coord] == '#' && occupiedAround >= 4 {
				newState[coord] = 'L'
				change = true
				continue
			}

			newState[coord] = p.state[coord]
		}
	}

	p.state = newState

	return change
}

func (p *SeatingPlan) Stabilise() {
	for p.Evolve() {
	}
}

func (p SeatingPlan) OccupiedSeats() int {
	n := 0

	for _, s := range p.state {
		if s == '#' {
			n++
		}
	}

	return n
}
