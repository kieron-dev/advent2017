package q25

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/kieron-pivotal/advent2017/2018/q23"
)

type Space struct {
	Points         []Point
	Constellations [][]Point
}

type Point struct {
	D1 int
	D2 int
	D3 int
	D4 int
}

func P(a, b, c, d int) Point {
	return Point{
		D1: a,
		D2: b,
		D3: c,
		D4: d,
	}
}

func (p Point) Dist(to Point) int {
	return q23.Abs(p.D1-to.D1) +
		q23.Abs(p.D2-to.D2) +
		q23.Abs(p.D3-to.D3) +
		q23.Abs(p.D4-to.D4)
}

func NewSpace(in io.Reader) *Space {
	s := Space{}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, "\n")
		coords := strings.Split(line, ",")
		p := P(atoi(coords[0]), atoi(coords[1]), atoi(coords[2]), atoi(coords[3]))
		s.Points = append(s.Points, p)
	}
	return &s
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func (s *Space) Partition() int {
	for _, p := range s.Points {
		in := map[int]bool{}
		for i, c := range s.Constellations {
			for _, s := range c {
				if s.Dist(p) <= 3 {
					in[i] = true
					break
				}
			}
		}

		newConstellations := [][]Point{}
		constel := []Point{p}
		for i, c := range s.Constellations {
			if in[i] {
				constel = append(constel, c...)
			} else {
				newConstellations = append(newConstellations, c)
			}
		}
		newConstellations = append(newConstellations, constel)
		s.Constellations = newConstellations
	}
	return len(s.Constellations)
}
