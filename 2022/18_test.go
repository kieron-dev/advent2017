package two022_test

import (
	"bytes"
	"math"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type cubes struct {
	minX, minY, minZ int
	maxX, maxY, maxZ int
	topEdges         map[[3]int]int
	rightEdges       map[[3]int]int
	frontEdges       map[[3]int]int
	cubes            map[[3]int]bool
}

func newCubes() *cubes {
	return &cubes{
		topEdges:   map[[3]int]int{},
		rightEdges: map[[3]int]int{},
		frontEdges: map[[3]int]int{},
		cubes:      map[[3]int]bool{},
		minX:       math.MaxInt,
		minY:       math.MaxInt,
		minZ:       math.MaxInt,
		maxX:       -math.MaxInt,
		maxY:       -math.MaxInt,
		maxZ:       -math.MaxInt,
	}
}

func (c *cubes) add(x, y, z int) {
	if c.cubes[[3]int{x, y, z}] {
		return
	}

	c.cubes[[3]int{x, y, z}] = true

	c.topEdges[[3]int{x, y, z}]++
	c.rightEdges[[3]int{x, y, z}]++
	c.frontEdges[[3]int{x, y, z}]++
	c.topEdges[[3]int{x, y, z - 1}]++
	c.rightEdges[[3]int{x - 1, y, z}]++
	c.frontEdges[[3]int{x, y - 1, z}]++

	if x < c.minX {
		c.minX = x
	}
	if y < c.minY {
		c.minY = y
	}
	if z < c.minZ {
		c.minZ = z
	}
	if x > c.maxX {
		c.maxX = x
	}
	if y > c.maxY {
		c.maxY = y
	}
	if z > c.maxZ {
		c.maxZ = z
	}
}

func (c *cubes) surfaceArea() int {
	res := 0
	for _, n := range c.topEdges {
		if n == 1 {
			res++
		}
	}
	for _, n := range c.rightEdges {
		if n == 1 {
			res++
		}
	}
	for _, n := range c.frontEdges {
		if n == 1 {
			res++
		}
	}

	return res
}

func (c *cubes) loadCubes(bs []byte) {
	for _, l := range bytes.Fields(bs) {
		if len(l) == 0 {
			continue
		}
		nums := bytes.Split(l, []byte(","))
		Expect(nums).To(HaveLen(3))
		x := stoi(string(nums[0]))
		y := stoi(string(nums[1]))
		z := stoi(string(nums[2]))
		c.add(x, y, z)
	}
}

func (c *cubes) getNeighbours(p [3]int) [][3]int {
	x, y, z := p[0], p[1], p[2]
	moves := [][3]int{
		{-1, 0, 0},
		{1, 0, 0},
		{0, -1, 0},
		{0, 1, 0},
		{0, 0, -1},
		{0, 0, 1},
	}

	var res [][3]int

	for _, mv := range moves {
		x1, y1, z1 := x+mv[0], y+mv[1], z+mv[2]
		if x1 < c.minX-1 || x1 > c.maxX+1 {
			continue
		}
		if y1 < c.minY-1 || y1 > c.maxY+1 {
			continue
		}
		if z1 < c.minZ-1 || z1 > c.maxZ+1 {
			continue
		}
		if c.cubes[[3]int{x1, y1, z1}] {
			continue
		}
		res = append(res, [3]int{x1, y1, z1})
	}

	return res
}

var _ = Describe("18", func() {
	var c *cubes

	BeforeEach(func() {
		c = newCubes()
	})

	It("does 1 cube", func() {
		c.add(1, 2, 3)
		Expect(c.surfaceArea()).To(Equal(6))
	})

	It("does 2 cubes", func() {
		c.add(1, 1, 1)
		c.add(2, 1, 1)
		Expect(c.surfaceArea()).To(Equal(10))
	})

	It("does the example", func() {
		bs := []byte(`2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`)
		c.loadCubes(bs)
		Expect(c.surfaceArea()).To(Equal(64))
	})

	It("does part A", func() {
		bs, err := os.ReadFile("input18")
		Expect(err).NotTo(HaveOccurred())
		c.loadCubes(bs)
		Expect(c.surfaceArea()).To(Equal(3542))
	})

	It("does part B", func() {
		bs, err := os.ReadFile("input18")
		Expect(err).NotTo(HaveOccurred())
		c.loadCubes(bs)

		queue := [][3]int{{c.minX - 1, c.minY - 1, c.minZ - 1}}

		d := newCubes()
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			if d.cubes[cur] {
				continue
			}

			d.add(cur[0], cur[1], cur[2])

			queue = append(queue, c.getNeighbours(cur)...)
		}

		area := d.surfaceArea()
		area -= 2 * (d.maxX - d.minX + 1) * (d.maxY - d.minY + 1)
		area -= 2 * (d.maxX - d.minX + 1) * (d.maxZ - d.minZ + 1)
		area -= 2 * (d.maxY - d.minY + 1) * (d.maxZ - d.minZ + 1)

		Expect(area).To(Equal(2080))
	})
})
