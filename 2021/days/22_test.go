package days_test

import (
	"bufio"
	"os"
	"regexp"
	"sort"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("22", func() {
	It("does part A", func() {
		input, err := os.Open("input22")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		scanner := bufio.NewScanner(input)

		grid := map[coord3d]bool{}
		re := regexp.MustCompile(`(o.+) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)`)

	outer:
		for scanner.Scan() {
			line := scanner.Text()
			matches := re.FindStringSubmatch(line)
			Expect(matches).To(HaveLen(8))

			toggle := matches[1] == "on"
			x1, x2 := AToI(matches[2]), AToI(matches[3])
			y1, y2 := AToI(matches[4]), AToI(matches[5])
			z1, z2 := AToI(matches[6]), AToI(matches[7])

			for _, n := range []int{x1, x2, y1, y2, z1, z2} {
				if n < -50 || n > 50 {
					continue outer
				}
			}

			for x := AToI(matches[2]); x <= AToI(matches[3]); x++ {
				for y := AToI(matches[4]); y <= AToI(matches[5]); y++ {
					for z := AToI(matches[6]); z <= AToI(matches[7]); z++ {
						if toggle {
							grid[newCoord3d(x, y, z)] = true
						} else {
							delete(grid, newCoord3d(x, y, z))
						}
					}
				}
			}
		}

		Expect(len(grid)).To(Equal(542711))
	})

	It("does part B", func() {
		input, err := os.Open("input22")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		scanner := bufio.NewScanner(input)

		xcoordmap := map[int]bool{}
		ycoordmap := map[int]bool{}
		zcoordmap := map[int]bool{}
		rects := []*rect{}
		re := regexp.MustCompile(`(o.+) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)`)

		for scanner.Scan() {
			line := scanner.Text()
			matches := re.FindStringSubmatch(line)
			Expect(matches).To(HaveLen(8))

			toggle := matches[1] == "on"
			_ = toggle
			x1, x2 := AToI(matches[2]), AToI(matches[3])
			y1, y2 := AToI(matches[4]), AToI(matches[5])
			z1, z2 := AToI(matches[6]), AToI(matches[7])

			xcoordmap[x1] = true
			xcoordmap[x2+1] = true
			ycoordmap[y1] = true
			ycoordmap[y2+1] = true
			zcoordmap[z1] = true
			zcoordmap[z2+1] = true

			rect := newRect(toggle, x1, x2, y1, y2, z1, z2)
			rects = append(rects, rect)
		}

		xcoords := []int{}
		for n := range xcoordmap {
			xcoords = append(xcoords, n)
		}
		ycoords := []int{}
		for n := range ycoordmap {
			ycoords = append(ycoords, n)
		}
		zcoords := []int{}
		for n := range zcoordmap {
			zcoords = append(zcoords, n)
		}

		sort.Ints(xcoords)
		sort.Ints(ycoords)
		sort.Ints(zcoords)

		xmap := map[int]int{}
		for i, x := range xcoords {
			xmap[x] = i
		}
		ymap := map[int]int{}
		for i, y := range ycoords {
			ymap[y] = i
		}
		zmap := map[int]int{}
		for i, z := range zcoords {
			zmap[z] = i
		}

		grid := newBitmap(1000000000)

		for _, r := range rects {
			xIdx1 := xmap[r.x1]
			xIdx2 := xmap[r.x2+1]
			yIdx1 := ymap[r.y1]
			yIdx2 := ymap[r.y2+1]
			zIdx1 := zmap[r.z1]
			zIdx2 := zmap[r.z2+1]

			for x := xIdx1; x < xIdx2; x++ {
				for y := yIdx1; y < yIdx2; y++ {
					for z := zIdx1; z < zIdx2; z++ {
						grid.set(x*1000000+y*1000+z, r.state)
					}
				}
			}
		}

		sum := 0
		for i := 0; i < 1000000000; i++ {
			if !grid.get(i) {
				continue
			}
			c := i
			cz := c % 1000
			c /= 1000
			cy := c % 1000
			c /= 1000
			cx := c
			x := xcoords[cx+1] - xcoords[cx]
			y := ycoords[cy+1] - ycoords[cy]
			z := zcoords[cz+1] - zcoords[cz]
			sum += x * y * z
		}

		Expect(sum).To(Equal(1160303042684776))
	})

	Describe("bitmap", func() {
		It("does something", func() {
			bm := newBitmap(8)
			Expect(bm.get(0)).To(BeFalse())
			Expect(bm.get(7)).To(BeFalse())
			bm.set(5, true)
			Expect(bm.get(5)).To(BeTrue())
			bm.set(5, false)
			Expect(bm.get(5)).To(BeFalse())
		})
		It("does something", func() {
			bm := newBitmap(9)
			bm.set(9, true)
			Expect(bm.get(9)).To(BeTrue())
		})
	})
})

type bitmap struct {
	bits []uint8
	size int
}

func newBitmap(size int) *bitmap {
	return &bitmap{
		bits: make([]byte, size/8+1),
		size: size,
	}
}

func (b bitmap) get(n int) bool {
	b8 := b.bits[n/8]
	offset := n % 8
	mask := uint8(1) << offset
	return b8&mask > 0
}

func (b *bitmap) set(n int, val bool) {
	b8 := b.bits[n/8]
	offset := n % 8
	mask := uint8(1) << offset
	if val {
		b8 |= mask
	} else {
		b8 &= ^mask
	}

	b.bits[n/8] = b8
}

type rect struct {
	state  bool
	x1, x2 int
	y1, y2 int
	z1, z2 int
}

func newRect(state bool, x1, x2, y1, y2, z1, z2 int) *rect {
	return &rect{
		state: state, x1: x1, x2: x2, y1: y1, y2: y2, z1: z1, z2: z2,
	}
}
