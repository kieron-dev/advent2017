package life

import "strings"

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

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			pos := i*5 + j
			if l&(1<<pos) > 0 {
				out += "#"
			} else {
				out += "."
			}
		}
		if i < 4 {
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
		pos := 5*coord[0] + coord[1]
		if int(l)&(1<<pos) > 0 {
			count++
		}
	}

	return count
}

func (l Life) Evolve() Life {
	var next int

	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			pos := 5*r + c
			neighbours := l.Neighbours(r, c)
			mask := 1 << pos

			if int(l)&mask > 0 && neighbours == 1 {
				next |= mask
				continue
			}

			if int(l)&mask == 0 && (neighbours == 1 || neighbours == 2) {
				next |= mask
			}
		}
	}

	return Life(next)
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
