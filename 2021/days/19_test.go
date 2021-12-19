package days_test

import (
	"bufio"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type coord3d struct {
	X, Y, Z int
}

func newCoord3d(x, y, z int) coord3d {
	return coord3d{
		X: x,
		Y: y,
		Z: z,
	}
}

func (c coord3d) distanceFrom(d coord3d) int {
	return (c.X-d.X)*(c.X-d.X) +
		(c.Y-d.Y)*(c.Y-d.Y) +
		(c.Z-d.Z)*(c.Z-d.Z)
}

type beacon struct {
	coord     coord3d
	distances []int
	sameAs    []*beacon
	scanner   *scann3r
}

func (b *beacon) isSameAs(c *beacon) bool {
	distances := map[int]int{}
	count := 0

	for _, d := range b.distances {
		distances[d]++
	}

	for _, d := range c.distances {
		if _, ok := distances[d]; ok {
			count++
		}
	}

	if count >= 11 {
		b.sameAs = append(b.sameAs, c)
		c.sameAs = append(c.sameAs, b)
		return true
	}

	return false
}

type scann3r struct {
	idx     int
	beacons []*beacon
	nextTo  []*scann3r
	pos     coord3d
}

func (s *scann3r) setDistances() {
	for i, b := range s.beacons {
		for j := 0; j < len(s.beacons); j++ {
			if j == i {
				continue
			}
			b.distances = append(b.distances, b.coord.distanceFrom(s.beacons[j].coord))
		}
	}
}

func (s *scann3r) compareTo(t *scann3r) bool {
	same := false
	for i := 0; i < len(s.beacons); i++ {
		for j := 0; j < len(t.beacons); j++ {
			if s.beacons[i].isSameAs(t.beacons[j]) {
				same = true
				break
			}
		}
	}

	if same {
		s.nextTo = append(s.nextTo, t)
		t.nextTo = append(t.nextTo, s)
	}

	return same
}

func (s *scann3r) alignFrom(t *scann3r) {
	sBeacons := map[*beacon]bool{}
	for _, b := range s.beacons {
		sBeacons[b] = true
	}

	sPair := []*beacon{}
	tPair := []*beacon{}
outer:
	for _, b := range t.beacons {
		for _, c := range b.sameAs {
			if sBeacons[c] {
				sPair = append(sPair, c)
				tPair = append(tPair, b)
				if len(sPair) > 1 {
					break outer
				}
			}
		}
	}
	Expect(sPair).To(HaveLen(2))

	realX := tPair[0].coord.X - tPair[1].coord.X
	realY := tPair[0].coord.Y - tPair[1].coord.Y
	realZ := tPair[0].coord.Z - tPair[1].coord.Z

	otherX := sPair[0].coord.X - sPair[1].coord.X
	otherY := sPair[0].coord.Y - sPair[1].coord.Y
	otherZ := sPair[0].coord.Z - sPair[1].coord.Z

	fn := transformIdentity

	switch abs(realX) {
	case abs(otherX):
		if realX == -otherX {
			fn = transformXNegX(fn)
		}
	case abs(otherY):
		fn = transformXY(fn)
		if realX == -otherY {
			fn = transformXNegX(fn)
		}
	case abs(otherZ):
		fn = transformXZ(fn)
		if realX == -otherZ {
			fn = transformXNegX(fn)
		}
	default:
		Fail("eh?")
	}

	switch abs(realY) {
	case abs(otherX):
		fn = transformYX(fn)
		if realY == -otherX {
			fn = transformYNegY(fn)
		}
	case abs(otherY):
		if realY == -otherY {
			fn = transformYNegY(fn)
		}
	case abs(otherZ):
		fn = transformYZ(fn)
		if realY == -otherZ {
			fn = transformYNegY(fn)
		}
	default:
		Fail("eh?")
	}

	switch abs(realZ) {
	case abs(otherX):
		fn = transformZX(fn)
		if realZ == -otherX {
			fn = transformZNegZ(fn)
		}
	case abs(otherY):
		fn = transformZY(fn)
		if realZ == -otherY {
			fn = transformZNegZ(fn)
		}
	case abs(otherZ):
		if realZ == -otherZ {
			fn = transformZNegZ(fn)
		}
	default:
		Fail("eh?")
	}

	d := fn(sPair[0].coord)
	transX := tPair[0].coord.X - d.X
	transY := tPair[0].coord.Y - d.Y
	transZ := tPair[0].coord.Z - d.Z

	for _, b := range s.beacons {
		b.coord = fn(b.coord)
		b.coord.X += transX
		b.coord.Y += transY
		b.coord.Z += transZ
	}

	s.pos.X = transX
	s.pos.Y = transY
	s.pos.Z = transZ
}

func transformIdentity(c coord3d) coord3d {
	return c
}

func transformXNegX(fn func(coord3d) coord3d) func(coord3d) coord3d {
	return func(c coord3d) coord3d {
		n := fn(c)
		n.X = -n.X
		return n
	}
}

func transformXY(fn func(coord3d) coord3d) func(coord3d) coord3d {
	return func(c coord3d) coord3d {
		n := fn(c)
		n.X = c.Y
		return n
	}
}

func transformXZ(fn func(coord3d) coord3d) func(coord3d) coord3d {
	return func(c coord3d) coord3d {
		n := fn(c)
		n.X = c.Z
		return n
	}
}

func transformYNegY(fn func(coord3d) coord3d) func(coord3d) coord3d {
	return func(c coord3d) coord3d {
		n := fn(c)
		n.Y = -n.Y
		return n
	}
}

func transformYX(fn func(coord3d) coord3d) func(coord3d) coord3d {
	return func(c coord3d) coord3d {
		n := fn(c)
		n.Y = c.X
		return n
	}
}

func transformYZ(fn func(coord3d) coord3d) func(coord3d) coord3d {
	return func(c coord3d) coord3d {
		n := fn(c)
		n.Y = c.Z
		return n
	}
}

func transformZNegZ(fn func(coord3d) coord3d) func(coord3d) coord3d {
	return func(c coord3d) coord3d {
		n := fn(c)
		n.Z = -n.Z
		return n
	}
}

func transformZX(fn func(coord3d) coord3d) func(coord3d) coord3d {
	return func(c coord3d) coord3d {
		n := fn(c)
		n.Z = c.X
		return n
	}
}

func transformZY(fn func(coord3d) coord3d) func(coord3d) coord3d {
	return func(c coord3d) coord3d {
		n := fn(c)
		n.Z = c.Y
		return n
	}
}

var _ = FDescribe("19", func() {
	It("does part A", func() {
		input, err := os.Open("input19")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		scanners := []*scann3r{}
		scanner := bufio.NewScanner(input)
		i := 0
		var cur *scann3r
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "---") {
				cur = &scann3r{idx: i}
				scanners = append(scanners, cur)
				i++
				continue
			}
			if line == "" {
				continue
			}
			nums := parseNumList(line, ",")
			Expect(nums).To(HaveLen(3))
			coord := newCoord3d(nums[0], nums[1], nums[2])
			cur.beacons = append(cur.beacons, &beacon{coord: coord, scanner: cur})
		}

		for _, s := range scanners {
			s.setDistances()
		}

		for i := 0; i < len(scanners); i++ {
			for j := i + 1; j < len(scanners); j++ {
				scanners[i].compareTo(scanners[j])
			}
		}

		next := []*scann3r{scanners[0]}
		visited := map[*scann3r]bool{}
		for len(next) > 0 {
			cur := next[0]
			next = next[1:]
			if visited[cur] {
				continue
			}

			for _, neighbour := range cur.nextTo {
				neighbour.alignFrom(cur)
				next = append(next, neighbour)
			}

			visited[cur] = true
		}
		Expect(len(visited)).To(Equal(len(scanners)))

		beacons := map[coord3d]bool{}
		for _, s := range scanners {
			for _, b := range s.beacons {
				beacons[b.coord] = true
			}
		}

		Expect(len(beacons)).To(Equal(430))

		max := 0
		for i := 0; i < len(scanners); i++ {
			for j := i + 1; j < len(scanners); j++ {
				dist := abs(scanners[i].pos.X-scanners[j].pos.X) +
					abs(scanners[i].pos.Y-scanners[j].pos.Y) +
					abs(scanners[i].pos.Z-scanners[j].pos.Z)
				if dist > max {
					max = dist
				}
			}
		}

		Expect(max).To(Equal(11860))
	})
})
