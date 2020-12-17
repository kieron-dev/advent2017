package gameoflife

import (
	"bufio"
	"io"
	"strings"
)

type Cube struct {
	minX, maxX int
	minY, maxY int
	minZ, maxZ int

	cells map[[3]int]bool
}

func NewCube() Cube {
	return Cube{
		cells: map[[3]int]bool{},
	}
}

func (c *Cube) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	z := 0
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		c.maxX = len(line) - 1
		c.maxY = y

		for x := 0; x < len(line); x++ {
			if line[x] == '#' {
				c.cells[[3]int{x, y, z}] = true
			}
		}

		y++
	}
}

func (c *Cube) Evolve() {
	newState := map[[3]int]bool{}

	for x := c.minX - 1; x < c.maxX+2; x++ {
		for y := c.minY - 1; y < c.maxY+2; y++ {
			for z := c.minZ - 1; z < c.maxZ+2; z++ {
				oldVal := c.cells[[3]int{x, y, z}]
				neighbourCount := c.neighbours(x, y, z)

				if oldVal && (neighbourCount == 2 || neighbourCount == 3) {
					newState[[3]int{x, y, z}] = true
					c.setMinMaxes(x, y, z)
				}

				if !oldVal && neighbourCount == 3 {
					newState[[3]int{x, y, z}] = true
					c.setMinMaxes(x, y, z)
				}
			}
		}
	}

	c.cells = newState
}

func (c *Cube) setMinMaxes(x, y, z int) {
	if x < c.minX {
		c.minX = x
	}
	if x > c.maxX {
		c.maxX = x
	}
	if y < c.minY {
		c.minY = y
	}
	if y > c.maxY {
		c.maxY = y
	}
	if z < c.minZ {
		c.minZ = z
	}
	if z > c.maxZ {
		c.maxZ = z
	}
}

func (c Cube) neighbours(x, y, z int) int {
	n := 0

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}

				if c.cells[[3]int{x + i, y + j, z + k}] {
					n++
				}
			}
		}
	}

	return n
}

func (c Cube) ActiveCount() int {
	return len(c.cells)
}
