package two023_test

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type cosmos struct {
	data            [][]byte
	galaxies        []coord
	expansionRows   map[int]bool
	expansionCols   map[int]bool
	expansionFactor int
}

func newCosmos(fileName string, expansionFactor int) cosmos {
	c := cosmos{}
	c.expansionFactor = expansionFactor
	c.load(fileName)
	c.getExpansions()
	c.getGalaxies()

	return c
}

func (c *cosmos) load(fileName string) {
	f, err := os.Open(fileName)
	Expect(err).NotTo(HaveOccurred())
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		c.data = append(c.data, []byte(line))
	}
}

func (c *cosmos) getExpansions() {
	c.expansionRows = map[int]bool{}
	c.expansionCols = map[int]bool{}

	for r, line := range c.data {
		if !bytes.Contains(line, []byte("#")) {
			c.expansionRows[r] = true
		}
	}
	for col := range c.data[0] {
		empty := true
		for r := range c.data {
			if c.data[r][col] == '#' {
				empty = false
				break
			}
		}
		if empty {
			c.expansionCols[col] = true
		}
	}
}

func (c *cosmos) getGalaxies() {
	for r := range c.data {
		for col := range c.data[r] {
			if c.data[r][col] == '#' {
				c.galaxies = append(c.galaxies, coord{r, col})
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (c cosmos) galaxiesDists() int {
	d := 0
	for i := 0; i < len(c.galaxies)-1; i++ {
		for j := i + 1; j < len(c.galaxies); j++ {
			galA := c.galaxies[i]
			galB := c.galaxies[j]
			d += galA.dist(galB)

			for i := min(galA[0], galB[0]) + 1; i < max(galA[0], galB[0]); i++ {
				if c.expansionRows[i] {
					d += c.expansionFactor - 1
				}
			}
			for i := min(galA[1], galB[1]) + 1; i < max(galA[1], galB[1]); i++ {
				if c.expansionCols[i] {
					d += c.expansionFactor - 1
				}
			}
		}
	}

	return d
}

func (c cosmos) print() {
	fmt.Println()
	for _, l := range c.data {
		fmt.Printf("%s\n", l)
	}
	fmt.Println()
}

var _ = Describe("11", func() {
	It("does part A", func() {
		c := newCosmos("input11", 2)
		Expect(c.galaxiesDists()).To(Equal(9693756))
	})

	It("does part B", func() {
		c := newCosmos("input11", 1_000_000)
		Expect(c.galaxiesDists()).To(Equal(717878258016))
	})
})
