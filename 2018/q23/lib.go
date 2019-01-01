package q23

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"log"
	"strings"
)

type Teleport struct {
	Nanobots []*Nanobot
}

type Coord struct {
	X int
	Y int
	Z int
}

func C(x, y, z int) Coord {
	return Coord{X: x, Y: y, Z: z}
}

func (c Coord) Dist(d Coord) int {
	x := c.X - d.X
	y := c.Y - d.Y
	z := c.Z - d.Z
	return Abs(x) + Abs(y) + Abs(z)
}

type Nanobot struct {
	Coord        Coord
	SignalRadius int
}

func (n *Nanobot) Dist(m *Nanobot) int {
	return n.Coord.Dist(m.Coord)
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func NewTeleport(in io.Reader) *Teleport {
	t := Teleport{}
	t.Nanobots = []*Nanobot{}
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, "\n")
		nano := Nanobot{}
		n, err := fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &nano.Coord.X, &nano.Coord.Y, &nano.Coord.Z, &nano.SignalRadius)
		if err != nil {
			log.Fatal("scanf ", err)
		}
		if n < 4 {
			log.Fatal("scanf n ")
		}
		t.Nanobots = append(t.Nanobots, &nano)
	}
	return &t
}

func (t *Teleport) Strongest() *Nanobot {
	maxSignal := 0
	var maxNanobot *Nanobot

	for _, n := range t.Nanobots {
		if n.SignalRadius > maxSignal {
			maxSignal = n.SignalRadius
			maxNanobot = n
		}
	}

	return maxNanobot
}

func (t *Teleport) InRange(n *Nanobot) int {
	count := 0
	for _, m := range t.Nanobots {
		if n.Dist(m) <= n.SignalRadius {
			count++
		}
	}
	return count
}

func (t *Teleport) GetLimits() (min, max Coord) {
	min.X = t.Nanobots[0].Coord.X
	max.X = min.X
	min.Y = t.Nanobots[0].Coord.Y
	max.Y = min.Y
	min.Z = t.Nanobots[0].Coord.Z
	max.Z = min.Z

	for _, n := range t.Nanobots {
		lowX := n.Coord.X - n.SignalRadius
		highX := n.Coord.X + n.SignalRadius
		lowY := n.Coord.Y - n.SignalRadius
		highY := n.Coord.Y + n.SignalRadius
		lowZ := n.Coord.Z - n.SignalRadius
		highZ := n.Coord.Z + n.SignalRadius

		if lowX < min.X {
			min.X = lowX
		}
		if highX > max.X {
			max.X = highX
		}
		if lowY < min.Y {
			min.Y = lowY
		}
		if highY > max.Y {
			max.Y = highY
		}
		if lowZ < min.Z {
			min.Z = lowZ
		}
		if highZ > max.Z {
			max.Z = highZ
		}
	}
	return
}

func (t *Teleport) Sample(min, max, step Coord) int {
	inRangeNanobots := map[*Nanobot]bool{}

	for x := min.X; x < max.X; x += step.X {
		for y := min.Y; y < max.Y; y += step.Y {
			for z := min.Z; z < max.Z; z += step.Z {

				for _, n := range t.Nanobots {
					if inRangeNanobots[n] {
						continue
					}
					if n.Coord.Dist(C(x, y, z)) <= n.SignalRadius {
						inRangeNanobots[n] = true
					}
				}
			}
		}
	}
	return len(inRangeNanobots)
}

func GetStepForSample(min, max Coord, num int) Coord {
	return C(
		maxi((max.X-min.X)/num, 1),
		maxi((max.Y-min.Y)/num, 1),
		maxi((max.Z-min.Z)/num, 1),
	)
}

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SplitCube(min, max Coord) [][]Coord {
	mid := C(min.X+(max.X-min.X)/2, min.Y+(max.Y-min.Y)/2, min.Z+(max.Z-min.Z)/2)

	return [][]Coord{
		{C(min.X, min.Y, min.Z), C(mid.X, mid.Y, mid.Z)},
		{C(min.X, min.Y, mid.Z), C(mid.X, mid.Y, max.Z)},
		{C(min.X, mid.Y, min.Z), C(mid.X, max.Y, mid.Z)},
		{C(min.X, mid.Y, mid.Z), C(mid.X, max.Y, max.Z)},
		{C(mid.X, min.Y, min.Z), C(max.X, mid.Y, mid.Z)},
		{C(mid.X, min.Y, mid.Z), C(max.X, mid.Y, max.Z)},
		{C(mid.X, mid.Y, min.Z), C(max.X, max.Y, mid.Z)},
		{C(mid.X, mid.Y, mid.Z), C(max.X, max.Y, max.Z)},
	}
}

func (t *Teleport) FindBestCoord(sampleSize int) Coord {
	min, max := t.GetLimits()

	start := Cube{
		Min: min,
		Max: max,
	}

	pq := &PriorityQueue{&start}
	heap.Init(pq)

	for len(*pq) > 0 {
		cur := heap.Pop(pq).(*Cube)
		if cur.UnitVolume() {
			return cur.Min
		}
		fmt.Printf("cur.Max.X - cur.Min.X = %+v\n", cur.Max.X-cur.Min.X)
		for _, subcube := range SplitCube(cur.Min, cur.Max) {
			step := GetStepForSample(subcube[0], subcube[1], sampleSize)
			inRange := t.Sample(subcube[0], subcube[1], step)
			c := Cube{
				Min:         subcube[0],
				Max:         subcube[1],
				InRangeBots: inRange,
			}
			heap.Push(pq, &c)
		}
	}
	return Coord{}
}
