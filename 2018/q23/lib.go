package q23

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type Teleport struct {
	Nanobots []*Nanobot
}

type Coord struct {
	X int
	Y int
	Z int
}

func (c Coord) Dist(d Coord) int {
	x := c.X - d.X
	y := c.Y - d.Y
	z := c.Z - d.Z
	return Abs(x) + Abs(y) + Abs(z)
}

type Nanobot struct {
	Coord        Coord
	SignalRadius int
}

func (n *Nanobot) Dist(m *Nanobot) int {
	return n.Coord.Dist(m.Coord)
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func NewTeleport(in io.Reader) *Teleport {
	t := Teleport{}
	t.Nanobots = []*Nanobot{}
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, "\n")
		nano := Nanobot{}
		n, err := fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &nano.Coord.X, &nano.Coord.Y, &nano.Coord.Z, &nano.SignalRadius)
		if err != nil {
			log.Fatal("scanf ", err)
		}
		if n < 4 {
			log.Fatal("scanf n ")
		}
		t.Nanobots = append(t.Nanobots, &nano)
	}
	return &t
}

func (t *Teleport) Strongest() *Nanobot {
	maxSignal := 0
	var maxNanobot *Nanobot

	for _, n := range t.Nanobots {
		if n.SignalRadius > maxSignal {
			maxSignal = n.SignalRadius
			maxNanobot = n
		}
	}

	return maxNanobot
}

func (t *Teleport) InRange(n *Nanobot) int {
	count := 0
	for _, m := range t.Nanobots {
		if n.Dist(m) <= n.SignalRadius {
			count++
		}
	}
	return count
}
