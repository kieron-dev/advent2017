package q17

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type Coord struct {
	X int
	Y int
}

func NewCoord(x, y int) Coord {
	return Coord{
		X: x,
		Y: y,
	}
}

type Slice struct {
	Grid map[Coord]rune
	MinX int
	MaxX int
	MinY int
	MaxY int
}

func NewSlice(in io.Reader) *Slice {
	bn := 99999999
	s := Slice{MinX: bn, MaxX: -bn, MinY: bn, MaxY: -bn}
	s.Grid = map[Coord]rune{}
	s.Set(NewCoord(500, 0), '+')

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, "\n")
		s.SetVein(line)
	}
	return &s
}

func (s *Slice) SetVein(vein string) {
	var a, b, c int
	if n, err := fmt.Sscanf(vein, "x=%d, y=%d..%d", &a, &b, &c); err == nil && n > 0 {
		s.AddVertical(a, b, c)
	} else if n, err := fmt.Sscanf(vein, "y=%d, x=%d..%d", &a, &b, &c); err == nil && n > 0 {
		s.AddHorizontal(a, b, c)
	} else {
		log.Fatal(vein)
	}
}

func (s *Slice) AddVertical(x, y1, y2 int) {
	for y := y1; y <= y2; y++ {
		c := NewCoord(x, y)
		s.Set(c, '#')
	}
}

func (s *Slice) AddHorizontal(y, x1, x2 int) {
	for x := x1; x <= x2; x++ {
		c := NewCoord(x, y)
		s.Set(c, '#')
	}
}

func (s *Slice) Set(c Coord, v rune) {
	s.Grid[c] = v
	if c.X < s.MinX {
		s.MinX = c.X
	}
	if c.X > s.MaxX {
		s.MaxX = c.X
	}
	if c.Y < s.MinY {
		s.MinY = c.Y
	}
	if c.Y > s.MaxY {
		s.MaxY = c.Y
	}
}

func (s *Slice) Print() {
	for y := s.MinY; y <= s.MaxY; y++ {
		for x := s.MinX; x <= s.MaxX; x++ {
			c := NewCoord(x, y)
			if v, ok := s.Grid[c]; ok {
				fmt.Printf("%s", string(v))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (s *Slice) DripTillOverFlow(from Coord) []Coord {
	return nil
}
