package tiled

import (
	"bufio"
	"io"
	"strings"
)

type Floor struct {
	instructions []string
	blackTiles   map[[2]int]bool
}

func NewFloor() Floor {
	return Floor{
		blackTiles: map[[2]int]bool{},
	}
}

func (f *Floor) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		f.instructions = append(f.instructions, line)
	}

	f.followInstructions()
}

func (f *Floor) BlackCount() int {
	return len(f.blackTiles)
}

func getSurrounding(coord [2]int) [][2]int {
	return [][2]int{
		{coord[0] - 2, coord[1]},
		{coord[0] + 2, coord[1]},
		{coord[0] - 1, coord[1] + 1},
		{coord[0] + 1, coord[1] + 1},
		{coord[0] - 1, coord[1] - 1},
		{coord[0] + 1, coord[1] - 1},
	}
}

func (f *Floor) Evolve() {
	newMap := map[[2]int]bool{}
	visited := map[[2]int]bool{}

	for coord := range f.blackTiles {
		for _, c := range append(getSurrounding(coord), coord) {
			if visited[c] {
				continue
			}
			visited[c] = true

			neighbours := f.neighbours(c)
			if f.blackTiles[c] && (neighbours == 1 || neighbours == 2) {
				newMap[c] = true
			}
			if !f.blackTiles[c] && neighbours == 2 {
				newMap[c] = true
			}
		}
	}

	f.blackTiles = newMap
}

func (f Floor) neighbours(coord [2]int) int {
	n := 0

	for _, c := range getSurrounding(coord) {
		if f.blackTiles[c] {
			n++
		}
	}

	return n
}

func (f *Floor) followInstructions() {
	for _, line := range f.instructions {
		coord := [2]int{0, 0}

		for i := 0; i < len(line); i++ {
			if line[i] == 'e' {
				coord[0] += 2
				continue
			} else if line[i] == 'w' {
				coord[0] -= 2
				continue
			} else if line[i] == 's' {
				coord[1]++
			} else if line[i] == 'n' {
				coord[1]--
			}

			if line[i+1] == 'e' {
				coord[0]++
			} else if line[i+1] == 'w' {
				coord[0]--
			}

			i++
		}
		if f.blackTiles[coord] {
			delete(f.blackTiles, coord)
		} else {
			f.blackTiles[coord] = true
		}
	}
}
