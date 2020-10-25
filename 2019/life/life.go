package life

import (
	"fmt"
	"strings"
)

type Life int

func New(chart string) Life {
	var l Life

	for i, c := range strings.ReplaceAll(chart, "\n", "") {
		if c == '#' {
			l |= 1 << i
		}
	}

	return l
}

func (l Life) Chart() string {
	out := ""

	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if l.Infested(r, c) {
				out += "#"
			} else {
				out += "."
			}
		}
		if r < 4 {
			out += "\n"
		}
	}

	return out
}

func (l Life) Neighbours(row, col int) int {
	count := 0

	neighbours := [][2]int{}
	if row > 0 {
		neighbours = append(neighbours, [2]int{row - 1, col})
	}
	if col > 0 {
		neighbours = append(neighbours, [2]int{row, col - 1})
	}
	if col < 4 {
		neighbours = append(neighbours, [2]int{row, col + 1})
	}
	if row < 4 {
		neighbours = append(neighbours, [2]int{row + 1, col})
	}

	for _, coord := range neighbours {
		if l.Infested(coord[0], coord[1]) {
			count++
		}
	}

	return count
}

func (l Life) Evolve() Life {
	var next Life

	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			neighbours := l.Neighbours(r, c)

			if l.Infested(r, c) && neighbours == 1 {
				next = next.Infest(r, c)
				continue
			}

			if !l.Infested(r, c) && (neighbours == 1 || neighbours == 2) {
				next = next.Infest(r, c)
			}
		}
	}

	return next
}

func (l Life) Infested(r, c int) bool {
	pos := 5*r + c
	mask := 1 << pos

	return int(l)&mask > 0
}

func (l Life) Infest(r, c int) Life {
	mask := 1 << (5*r + c)
	i := int(l)
	i |= mask

	return Life(i)
}

func (l Life) CountBugs() int {
	count := 0
	n := int(l)

	for n > 0 {
		if n&1 > 0 {
			count++
		}
		n >>= 1
	}

	return count
}

func (l Life) FirstRepeat() Life {
	visited := map[Life]bool{l: true}

	next := l
	for {
		next = next.Evolve()
		if visited[next] {
			return next
		}
		visited[next] = true
	}
}

func (l Life) PlutoniumNeighbours(row, col int, outer, inner Life) int {
	count := 0
	coords := [][2]int{{row - 1, col}, {row + 1, col}, {row, col - 1}, {row, col + 1}}

	for _, coord := range coords {
		r, c := coord[0], coord[1]

		if r == 2 && c == 2 {
			if col == 1 {
				for i := 0; i < 5; i++ {
					if inner.Infested(i, 0) {
						count++
					}
				}
			}
			if col == 3 {
				for i := 0; i < 5; i++ {
					if inner.Infested(i, 4) {
						count++
					}
				}
			}
			if row == 1 {
				for i := 0; i < 5; i++ {
					if inner.Infested(0, i) {
						count++
					}
				}
			}
			if row == 3 {
				for i := 0; i < 5; i++ {
					if inner.Infested(4, i) {
						count++
					}
				}
			}
			continue
		}

		if r >= 0 && r < 5 && c >= 0 && c < 5 && l.Infested(r, c) {
			count++
			continue
		}

		if r < 0 && outer.Infested(1, 2) {
			count++
			continue
		}

		if r > 4 && outer.Infested(3, 2) {
			count++
			continue
		}

		if c < 0 && outer.Infested(2, 1) {
			count++
			continue
		}

		if c > 4 && outer.Infested(2, 3) {
			count++
			continue
		}
	}

	return count
}

type Line struct {
	line map[int]Life
	min  int
	max  int
}

func NewLine(chart string) Line {
	l := Line{
		line: map[int]Life{0: New(chart)},
		min:  0,
		max:  0,
	}

	return l
}

func (l Line) Evolve() Line {
	next := Line{
		line: map[int]Life{},
		min:  l.min,
		max:  l.max,
	}

	for i := l.min - 1; i <= l.max+1; i++ {
		outer := l.line[i-1]
		cur := l.line[i]
		inner := l.line[i+1]

		nextLife := calcNext(outer, cur, inner)
		if int(nextLife) == 0 {
			continue
		}

		next.line[i] = nextLife

		if i < l.min {
			next.min = i
		}
		if i > l.max {
			next.max = i
		}
	}

	return next
}

func calcNext(outer, cur, inner Life) Life {
	var next Life

	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if r == 2 && c == 2 {
				continue
			}

			neighbours := cur.PlutoniumNeighbours(r, c, outer, inner)

			if cur.Infested(r, c) && neighbours == 1 {
				next = next.Infest(r, c)
			}

			if !cur.Infested(r, c) && (neighbours == 1 || neighbours == 2) {
				next = next.Infest(r, c)
			}
		}
	}

	return next
}

func (l Line) Print() {
	for i := l.min; i <= l.max; i++ {
		fmt.Printf("%d\n%s\n\n", i, l.line[i].Chart())
	}
}

func (l Line) CountBugs() int {
	count := 0

	for _, life := range l.line {
		count += life.CountBugs()
	}

	return count
}

func (l Line) Len() int {
	return l.max - l.min + 1
}
