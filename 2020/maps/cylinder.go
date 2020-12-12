// Package maps does grids and stuff
package maps

import (
	"bufio"
	"io"
	"log"
	"strings"
)

type Cylinder struct {
	width   int
	height  int
	pattern []string
}

type Tuple struct {
	X, Y int
}

type (
	Coord  Tuple
	Vector Tuple
)

func NewCoord(x, y int) Coord {
	return Coord{X: x, Y: y}
}

func (c Coord) Plus(v Vector) Coord {
	return Coord{
		X: c.X + v.X,
		Y: c.Y + v.Y,
	}
}

func (v Vector) Plus(w Vector) Vector {
	return Vector{
		X: v.X + w.X,
		Y: v.Y + w.Y,
	}
}

func (v Vector) Times(n int) Vector {
	return Vector{
		X: v.X * n,
		Y: v.Y * n,
	}
}

func (v Vector) RotateLeft() Vector {
	return Vector{
		X: v.Y,
		Y: -v.X,
	}
}

func (v Vector) RotateRight() Vector {
	return Vector{
		X: -v.Y,
		Y: v.X,
	}
}

func (c Coord) Neighbours(rows, cols int) []Coord {
	res := []Coord{}

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}

			row := c.Y + i
			if row < 0 || row >= rows {
				continue
			}

			col := c.X + j
			if col < 0 || col >= cols {
				continue
			}

			res = append(res, NewCoord(col, row))
		}
	}

	return res
}

func NewVector(x, y int) Vector {
	return Vector{X: x, Y: y}
}

func NewCylinder() Cylinder {
	return Cylinder{}
}

func (c *Cylinder) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		c.pattern = append(c.pattern, line)
		lineLen := len(line)
		if c.width != 0 && lineLen != c.width {
			log.Fatalf("unexpected line length: %s, cylinder: %+v", line, c)
		}
		c.width = lineLen
	}

	c.height = len(c.pattern)
}

func (c Cylinder) CountChars(start Coord, dir Vector, char byte) int {
	n := 0

	for pos := start; pos.Y < c.height; pos = pos.Plus(dir) {
		if c.At(pos) == char {
			n++
		}
	}

	return n
}

func (c Cylinder) At(coord Coord) byte {
	return c.pattern[coord.Y][coord.X%c.width]
}
