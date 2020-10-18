package grid

import (
	"io"
	"sort"

	"github.com/kieron-dev/advent2017/advent2019"
)

type Grid struct {
	height    int
	width     int
	asteroids map[Coord]asteroidType
}

type asteroidType int

const (
	_ asteroidType = iota
	unknown
	visible
	hidden
)

func New() *Grid {
	g := Grid{
		asteroids: map[Coord]asteroidType{},
	}
	return &g
}

func (g *Grid) Load(contents io.Reader) {
	r := advent2019.FileReader{}
	lines := 0
	r.Each(contents, func(line string) {
		g.width = len(line)
		for x, c := range line {
			if c == '#' {
				g.asteroids[NewCoord(x, lines)] = unknown
			}
		}
		lines++
	})
	g.height = lines
}

func (g *Grid) AsteroidCount() int {
	return len(g.asteroids)
}

func (g *Grid) Height() int {
	return g.height
}

func (g *Grid) Width() int {
	return g.width
}

func (g *Grid) testAndHide(from, coord Coord) bool {
	asteroid, present := g.asteroids[coord]
	if !present || asteroid == hidden {
		return false
	}

	g.asteroids[coord] = visible

	vector := coord.Minus(from)

	for k, v := range g.asteroids {
		if k == from || v != unknown {
			continue
		}

		hideVec := k.Minus(from)

		if hideVec.x*vector.y == hideVec.y*vector.x &&
			hideVec.Mag2() > vector.Mag2() &&
			hideVec.quadrant() == vector.quadrant() {
			g.asteroids[k] = hidden
		}
	}
	return true
}

func (g *Grid) Sort(from Coord, coords []Coord) {
	sort.Slice(coords, func(i, j int) bool {
		v1 := coords[i].Minus(from)
		v2 := coords[j].Minus(from)

		if v1.quadrant() == v2.quadrant() {
			return v1.x*v2.y > v1.y*v2.x
		}
		return v1.quadrant() < v2.quadrant()
	})
}

func (g *Grid) VisibleFrom(c Coord) []Coord {
	g.reset()
	visible := []Coord{}
	for row := c.y; row >= 0; row-- {
		for col := c.x; col >= 0; col-- {
			coord := NewCoord(col, row)
			if coord == c {
				continue
			}
			if g.testAndHide(c, coord) {
				visible = append(visible, coord)
			}
		}
		for col := c.x + 1; col < g.width; col++ {
			coord := NewCoord(col, row)
			if g.testAndHide(c, coord) {
				visible = append(visible, coord)
			}
		}
	}
	for row := c.y + 1; row < g.height; row++ {
		for col := c.x; col >= 0; col-- {
			coord := NewCoord(col, row)
			if g.testAndHide(c, coord) {
				visible = append(visible, coord)
			}
		}
		for col := c.x + 1; col < g.width; col++ {
			coord := NewCoord(col, row)
			if g.testAndHide(c, coord) {
				visible = append(visible, coord)
			}
		}
	}
	return visible
}

func (g *Grid) reset() {
	for k := range g.asteroids {
		g.asteroids[k] = unknown
	}
}

func (g *Grid) BestAsteroid() Coord {
	max := 0
	var maxCoord Coord

	for k := range g.asteroids {
		count := len(g.VisibleFrom(k))
		if count > max {
			max = count
			maxCoord = k
		}
	}
	return maxCoord
}

func (g *Grid) LaserN(n int, pos Coord) Coord {
	need := n
	for need > 0 {
		visible := g.VisibleFrom(pos)
		if len(visible) >= need {
			g.Sort(pos, visible)
			return visible[need-1]
		}
		need -= len(visible)
		for _, c := range visible {
			delete(g.asteroids, c)
		}
	}
	panic("whoops")
}
