package q22

import (
	"container/heap"
	"fmt"
)

type CaveType int

const (
	Rocky CaveType = iota
	Wet
	Narrow
)

type Map struct {
	Depth           int
	MaxX            int
	MaxY            int
	Target          Coord
	GeologicIndices map[Coord]int
}

type Coord struct {
	X int
	Y int
}

type Cave struct {
	index    int
	Coord    Coord
	Time     int
	Carrying Items
	Previous *Cave
	Done     bool
}

type Items int

const (
	Torch Items = iota
	ClimbingGear
	Neither
)

var AllowableItems = map[CaveType][]Items{
	Rocky:  []Items{Torch, ClimbingGear},
	Wet:    []Items{ClimbingGear, Neither},
	Narrow: []Items{Torch, Neither},
}

func (t CaveType) OkFor(carrying Items) bool {
	for _, item := range AllowableItems[t] {
		if item == carrying {
			return true
		}
	}
	return false
}

func C(x, y int) Coord {
	return Coord{X: x, Y: y}
}

func NewMap(target Coord, depth int) *Map {
	m := Map{
		Target: target,
		Depth:  depth,
	}
	m.GeologicIndices = map[Coord]int{}

	return &m
}

func (m *Map) PopulateGeologicIndices(maxX, maxY int) {
	prevMaxX := m.MaxX
	prevMaxY := m.MaxY
	if maxX > m.MaxX {
		m.MaxX = maxX
	}
	if maxY > m.MaxY {
		m.MaxY = maxY
	}

	for x := prevMaxX; x <= m.MaxX; x++ {
		m.GeologicIndices[C(x, 0)] = 16807 * x
	}
	for y := prevMaxY; y <= m.MaxY; y++ {
		m.GeologicIndices[C(0, y)] = 48271 * y
	}
	for x := 1; x <= m.MaxX; x++ {
		for y := 1; y <= m.MaxY; y++ {
			if x == m.Target.X && y == m.Target.Y {
				m.GeologicIndices[C(x, y)] = 0
			} else {
				m.GeologicIndices[C(x, y)] = m.ErosionLevel(C(x-1, y)) * m.ErosionLevel(C(x, y-1))
			}
		}
	}
}

func (m *Map) PrintGeoIdx() {
	fmt.Println()
	for y := 0; y <= m.MaxY; y++ {
		for x := 0; x <= m.MaxX; x++ {
			fmt.Printf("%12d ", m.GeologicIndices[C(x, y)])
		}
		fmt.Println()
	}
}

func (m *Map) PrintTypes() {
	fmt.Println()
	for y := 0; y <= m.MaxY; y++ {
		for x := 0; x <= m.MaxX; x++ {
			switch m.Type(C(x, y)) {
			case Rocky:
				fmt.Printf(".")
			case Wet:
				fmt.Printf("=")
			case Narrow:
				fmt.Printf("|")
			}
		}
		fmt.Println()
	}
}

func (m *Map) GeologicIndex(c Coord) int {
	if c.X > m.MaxX || c.Y > m.MaxY {
		m.PopulateGeologicIndices(c.X, c.Y)
	}
	return m.GeologicIndices[c]
}

func (m *Map) ErosionLevel(c Coord) int {
	return (m.GeologicIndex(c) + m.Depth) % 20183
}

func (m *Map) Type(c Coord) CaveType {
	return CaveType(m.ErosionLevel(c) % 3)
}

func (m *Map) RiskLevel() int {
	s := 0
	for x := 0; x <= m.Target.X; x++ {
		for y := 0; y <= m.Target.Y; y++ {
			s += int(m.Type(C(x, y)))
		}
	}
	return s
}

func (c Coord) Neighbours() []Coord {
	res := []Coord{}
	if c.X > 0 {
		res = append(res, Coord{X: c.X - 1, Y: c.Y})
	}
	if c.Y > 0 {
		res = append(res, Coord{X: c.X, Y: c.Y - 1})
	}
	res = append(res, Coord{X: c.X + 1, Y: c.Y}, Coord{X: c.X, Y: c.Y + 1})
	return res
}

func (m *Map) ShortestToTarget() int {
	const inf = 9999999999

	caves := map[Coord]map[Items]*Cave{}
	start := Cave{
		Coord:    C(0, 0),
		Carrying: Torch,
	}
	startII := Cave{
		Coord:    C(0, 0),
		Carrying: ClimbingGear,
		Time:     7,
	}
	caves[start.Coord] = map[Items]*Cave{Torch: &start, ClimbingGear: &startII}

	pq := PriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, &start)
	heap.Push(&pq, &startII)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Cave)
		curType := m.Type(cur.Coord)

		for _, neighbourCoord := range cur.Coord.Neighbours() {
			_, ok := caves[neighbourCoord]
			if !ok {
				caves[neighbourCoord] = map[Items]*Cave{}
				for _, gear := range AllowableItems[m.Type(neighbourCoord)] {
					nextCave := &Cave{
						Coord:    neighbourCoord,
						Time:     inf,
						Carrying: gear,
					}
					caves[neighbourCoord][gear] = nextCave
					heap.Push(&pq, nextCave)
				}
			}
			for gear, nextCave := range caves[neighbourCoord] {
				if nextCave.Done {
					continue
				}
				var newTime int
				if gear == cur.Carrying {
					newTime = cur.Time + 1
				} else if curType.OkFor(gear) {
					newTime = cur.Time + 8
				} else {
					newTime = inf
				}
				if newTime < nextCave.Time {
					nextCave.Previous = cur
					pq.update(nextCave, newTime)
				}
			}
		}
		cur.Done = true
		if cur.Coord == m.Target && cur.Carrying == Torch {
			return cur.Time
		}
	}

	return 0
}
