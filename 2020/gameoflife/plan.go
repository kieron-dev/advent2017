// Package gameoflife does Conway's game of life
package gameoflife

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/maps"
)

type SeatingPlan struct {
	state [][]byte
	rows  int
	cols  int
}

func NewSeatingPlan() SeatingPlan {
	return SeatingPlan{
		state: [][]byte{},
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

	p.rows = len(p.state)
}

func (p *SeatingPlan) AddLine(line string) {
	p.state = append(p.state, []byte(line))
	p.cols = len(line)
}

func (p SeatingPlan) State() string {
	res := ""

	for _, line := range p.state {
		res += string(line) + "\n"
	}

	return res
}

func (p SeatingPlan) OccupiedAround(coord maps.Coord) int {
	n := 0

	for r := -1; r < 2; r++ {
		for c := -1; c < 2; c++ {
			if (r == 0 && c == 0) || coord.Y+r < 0 || coord.X+c < 0 || coord.Y+r >= p.rows || coord.X+c >= p.cols {
				continue
			}

			if p.state[coord.Y+r][coord.X+c] == '#' {
				n++
			}
		}
	}

	return n
}

var (
	n  = maps.NewVector(0, -1)
	s  = maps.NewVector(0, 1)
	w  = maps.NewVector(-1, 0)
	e  = maps.NewVector(1, 0)
	nw = maps.NewVector(-1, -1)
	ne = maps.NewVector(1, -1)
	se = maps.NewVector(1, 1)
	sw = maps.NewVector(-1, 1)

	all = []maps.Vector{n, s, e, w, nw, ne, se, sw}
)

func (p SeatingPlan) VisiblyOccupiedAround(coord maps.Coord) int {
	res := 0

	for _, dir := range all {
		r := coord.Y
		c := coord.X

		for {
			r += dir.Y
			c += dir.X

			if r < 0 || c < 0 || r >= p.rows || c >= p.cols {
				break
			}

			if p.state[r][c] == '.' {
				continue
			}
			if p.state[r][c] == '#' {
				res++
			}

			// if p.state[r][c] == 'L'
			break
		}
	}

	return res
}

func (p SeatingPlan) IsInPlan(c maps.Coord) bool {
	return c.X >= 0 && c.X < p.cols && c.Y >= 0 && c.Y < p.rows
}

// Evolve returns true if a change has occurred
func (p *SeatingPlan) Evolve(partB bool) bool {
	newState := make([][]byte, p.rows)
	change := false
	lim := 4
	if partB {
		lim = 5
	}

	for r := 0; r < p.rows; r++ {
		newState[r] = make([]byte, p.cols)
		for c := 0; c < p.cols; c++ {
			coord := maps.NewCoord(c, r)

			var occupiedAround int
			if partB {
				occupiedAround = p.VisiblyOccupiedAround(coord)
			} else {
				occupiedAround = p.OccupiedAround(coord)
			}

			if p.state[r][c] == 'L' && occupiedAround == 0 {
				newState[r][c] = '#'
				change = true
				continue
			}

			if p.state[r][c] == '#' && occupiedAround >= lim {
				newState[r][c] = 'L'
				change = true
				continue
			}

			newState[r][c] = p.state[r][c]
		}
	}

	p.state = newState

	return change
}

func (p *SeatingPlan) Stabilise(partB bool) {
	for p.Evolve(partB) {
	}
}

func (p SeatingPlan) OccupiedSeats() int {
	n := 0

	for _, line := range p.state {
		n += bytes.Count(line, []byte("#"))
	}

	return n
}
