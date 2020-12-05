// Package seating parses boarding passes for aeroplane seating
package seating

import (
	"bufio"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Plan struct {
	seats []int
}

func NewPlan() Plan {
	return Plan{}
}

func (p *Plan) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		p.AddSeat(line)
	}
}

func (p *Plan) AddSeat(code string) {
	code = strings.ReplaceAll(code, "B", "1")
	code = strings.ReplaceAll(code, "F", "0")
	code = strings.ReplaceAll(code, "R", "1")
	code = strings.ReplaceAll(code, "L", "0")

	id, err := strconv.ParseInt(code, 2, 32)
	if err != nil {
		log.Fatalf("failed to parse %q as binary int32", code)
	}

	p.seats = append(p.seats, int(id))
}

func (p Plan) MaxSeatID() int {
	sort.Ints(p.seats)

	return p.seats[len(p.seats)-1]
}

func (p Plan) MissingSeat() int {
	sort.Ints(p.seats)

	for i, s := range p.seats {
		if p.seats[i+1] != s+1 {
			return s + 1
		}
	}

	log.Fatal("failed to find missing seat")
	return -1
}
