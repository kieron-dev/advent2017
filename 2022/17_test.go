package two022_test

import (
	"bytes"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type rock struct {
	coords []*Coord
}

func newRock(t, bottom int) *rock {
	switch t {
	case 0:
		return newRockA(bottom)
	case 1:
		return newRockB(bottom)
	case 2:
		return newRockC(bottom)
	case 3:
		return newRockD(bottom)
	case 4:
		return newRockE(bottom)
	}
	panic("invalid rock type")
}

func newRockA(bottom int) *rock {
	return &rock{
		coords: []*Coord{
			{X: 2, Y: bottom},
			{X: 3, Y: bottom},
			{X: 4, Y: bottom},
			{X: 5, Y: bottom},
		},
	}
}

func newRockB(bottom int) *rock {
	return &rock{
		coords: []*Coord{
			{X: 3, Y: bottom + 2},
			{X: 2, Y: bottom + 1},
			{X: 3, Y: bottom + 1},
			{X: 4, Y: bottom + 1},
			{X: 3, Y: bottom},
		},
	}
}

func newRockC(bottom int) *rock {
	return &rock{
		coords: []*Coord{
			{X: 4, Y: bottom + 2},
			{X: 4, Y: bottom + 1},
			{X: 2, Y: bottom},
			{X: 3, Y: bottom},
			{X: 4, Y: bottom},
		},
	}
}

func newRockD(bottom int) *rock {
	return &rock{
		coords: []*Coord{
			{X: 2, Y: bottom + 3},
			{X: 2, Y: bottom + 2},
			{X: 2, Y: bottom + 1},
			{X: 2, Y: bottom},
		},
	}
}

func newRockE(bottom int) *rock {
	return &rock{
		coords: []*Coord{
			{X: 2, Y: bottom + 1},
			{X: 3, Y: bottom + 1},
			{X: 2, Y: bottom},
			{X: 3, Y: bottom},
		},
	}
}

func (r *rock) print() {
	fmt.Printf("coords: ")
	for _, c := range r.coords {
		fmt.Printf("%v ", *c)
	}
	fmt.Println()
}

func (r *rock) shift(dir byte, g rockGrid) {
	minX := r.coords[0].X
	maxX := minX

	for _, c := range r.coords {
		if c.X < minX {
			minX = c.X
		}
		if c.X > maxX {
			maxX = c.X
		}
	}

	switch dir {
	case '<':
		if minX == 0 {
			return
		}
		for _, c := range r.coords {
			if g.filled[Coord{X: c.X - 1, Y: c.Y}] {
				return
			}
		}
		for _, c := range r.coords {
			c.X--
		}
	case '>':
		if maxX == 6 {
			return
		}
		for _, c := range r.coords {
			if g.filled[Coord{X: c.X + 1, Y: c.Y}] {
				return
			}
		}
		for _, c := range r.coords {
			c.X++
		}
	default:
		panic("invalid direction " + string(dir))
	}
}

func (r *rock) drop(g rockGrid) {
	for _, c := range r.coords {
		c.Y--
	}
}

func (r *rock) canDrop(g rockGrid) bool {
	for _, c := range r.coords {
		if g.filled[Coord{X: c.X, Y: c.Y - 1}] {
			return false
		}
	}

	return true
}

type rockGrid struct {
	filled     map[Coord]bool
	highestRow int
}

func newRockGrid() *rockGrid {
	r := &rockGrid{
		filled: map[Coord]bool{},
	}
	for i := 0; i < 7; i++ {
		r.filled[Coord{X: i, Y: 0}] = true
	}

	return r
}

func (g *rockGrid) markFilled(r rock) {
	for _, c := range r.coords {
		g.filled[*c] = true
		if c.Y > g.highestRow {
			g.highestRow = c.Y
		}
	}
}

func (g *rockGrid) print() {
	for y := g.highestRow; y > 0; y-- {
		for x := 0; x < 7; x++ {
			if g.filled[Coord{X: x, Y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
		// fmt.Printf(" %5d\n", y)
	}
}

func (g *rockGrid) printN() {
	for y := 0; y <= g.highestRow; y++ {
		n := 0
		for x := 0; x < 7; x++ {
			n *= 2
			if g.filled[Coord{X: x, Y: y}] {
				n += 1
			}
		}
		fmt.Printf("%d ", n)
	}
}

func (g rockGrid) getRowBytes() []byte {
	nos := make([]byte, g.highestRow+1)
	for y := 0; y <= g.highestRow; y++ {
		var n byte
		for x := 0; x < 7; x++ {
			n *= 2
			if g.filled[Coord{X: x, Y: y}] {
				n += 1
			}
		}
		nos[y] = n
	}

	return nos
}

var _ = Describe("17", func() {
	It("does part A", func() {
		bs, err := os.ReadFile("input17")
		Expect(err).NotTo(HaveOccurred())
		lenbs := len(bs)

		g := newRockGrid()
		bpos := 0

		for i := 0; i < 2022; i++ {
			r := newRock(i%5, g.highestRow+4)
			for {
				// r.print()
				r.shift(bs[bpos], *g)
				bpos++
				if bpos > lenbs-1 {
					bpos -= lenbs
				}
				if !r.canDrop(*g) {
					g.markFilled(*r)
					break
				}
				r.drop(*g)
			}
		}

		Expect(g.highestRow).To(Equal(3179))
	})

	It("does part B", func() {
		bs, err := os.ReadFile("input17")
		Expect(err).NotTo(HaveOccurred())
		// bs := []byte(">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>")
		lenbs := len(bs)

		g := newRockGrid()
		bpos := 0

		var startHeight int
		var end int
		var endHeight int
		var suf []byte
		const sampleLen = 10
		const samplePoint = 100
		const target = 1000000000000

		for i := 0; i < 20000; i++ {
			r := newRock(i%5, g.highestRow+4)
			for {
				r.shift(bs[bpos], *g)
				bpos++
				if bpos > lenbs-1 {
					bpos -= lenbs
				}
				if !r.canDrop(*g) {
					g.markFilled(*r)
					break
				}
				r.drop(*g)
			}

			nos := g.getRowBytes()
			if i == samplePoint {
				suf = nos[len(nos)-sampleLen:]
				startHeight = g.highestRow
				continue
			}

			if suf != nil && bytes.HasSuffix(nos, suf) {
				end = i
				endHeight = g.highestRow
				break
			}
		}

		more := target - end
		interval := end - samplePoint
		for i := end + 1; i < end+(more%interval); i++ {
			r := newRock(i%5, g.highestRow+4)
			for {
				r.shift(bs[bpos], *g)
				bpos++
				if bpos > lenbs-1 {
					bpos -= lenbs
				}
				if !r.canDrop(*g) {
					g.markFilled(*r)
					break
				}
				r.drop(*g)
			}
		}
		times := more / interval
		res := g.highestRow + (endHeight-startHeight)*times

		Expect(res).To(Equal(1567723342929))
	})
})
