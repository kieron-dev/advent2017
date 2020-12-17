package gameoflife

import (
	"bufio"
	"io"
	"strings"
)

type Cube struct {
	dimension int

	mins []int
	maxs []int

	cells map[int]bool
}

func NewCube(dimension int) Cube {
	return Cube{
		dimension: dimension,
		cells:     map[int]bool{},
		mins:      make([]int, dimension),
		maxs:      make([]int, dimension),
	}
}

func (c *Cube) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	coord := make([]int, c.dimension)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		c.maxs[c.dimension-1] = len(line) - 1
		c.maxs[c.dimension-2] = coord[c.dimension-2]

		for x := 0; x < len(line); x++ {
			if line[x] == '#' {
				coord[c.dimension-1] = x
				c.cells[toKey(coord)] = true
			}
		}

		coord[c.dimension-2]++
	}
}

func toKey(coord []int) int {
	n := 0

	for _, i := range coord {
		n += 100*n + i
	}

	return n
}

func (c *Cube) Evolve() {
	sizes := make([]int, c.dimension)
	prod := 1

	for i := 0; i < c.dimension; i++ {
		// accomodate 1 each side of existing range
		sizes[i] = c.maxs[i] - c.mins[i] + 3
		prod *= sizes[i]
	}

	newState := map[int]bool{}
	set := [][]int{}

	for i := 0; i < prod; i++ {
		coord := make([]int, c.dimension)
		work := i
		for j := c.dimension - 1; j >= 0; j-- {
			coord[j] = c.mins[j] - 1 + (work % sizes[j])
			work /= sizes[j]
		}

		key := toKey(coord)
		oldVal := c.cells[key]
		neighbourCount := c.neighbours(coord)

		if oldVal && (neighbourCount == 2 || neighbourCount == 3) {
			newState[key] = true
			set = append(set, coord[:])
		}

		if !oldVal && neighbourCount == 3 {
			newState[key] = true
			set = append(set, coord[:])
		}
	}

	c.cells = newState
	for _, coord := range set {
		c.setMinMaxes(coord)
	}
}

func (c *Cube) setMinMaxes(coord []int) {
	for i, n := range coord {
		if n < c.mins[i] {
			c.mins[i] = n
		}
		if n > c.maxs[i] {
			c.maxs[i] = n
		}
	}
}

func (c Cube) neighbours(around []int) int {
	n := 0
	aroundKey := toKey(around)

	max := 1
	for i := 0; i < c.dimension; i++ {
		max *= 3
	}

	coord := make([]int, c.dimension)

	for i := 0; i < max; i++ {
		work := i
		for j := 0; j < c.dimension; j++ {
			coord[j] = around[j] - 1 + (work % 3)
			work /= 3
		}

		key := toKey(coord)
		if key == aroundKey {
			continue
		}

		if c.cells[key] {
			n++
		}
	}

	return n
}

func (c Cube) ActiveCount() int {
	return len(c.cells)
}
