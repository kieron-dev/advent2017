package q23

type Cube struct {
	Min         Coord
	Max         Coord
	InRangeBots int
	dist        int
	index       int
}

func (c Cube) UnitVolume() bool {
	return c.Max.X-c.Min.X == 1 && c.Max.Y-c.Min.Y == 1 && c.Max.Z-c.Min.Z == 1
}

type PriorityQueue []*Cube

func (c Cube) DistFromO() int {
	if c.dist == 0 {
		var x, y, z int
		if c.Min.X*c.Max.X < 0 {
			x = 0
		} else {
			x = Abs(c.Min.X)
			if Abs(c.Max.X) < x {
				x = Abs(c.Max.X)
			}
		}
		if c.Min.Y*c.Max.Y < 0 {
			y = 0
		} else {
			y = Abs(c.Min.Y)
			if Abs(c.Max.Y) < y {
				y = Abs(c.Max.Y)
			}
		}
		if c.Min.Z*c.Max.Z < 0 {
			z = 0
		} else {
			z = Abs(c.Min.Z)
			if Abs(c.Max.Z) < z {
				z = Abs(c.Max.Z)
			}
		}
		c.dist = x + y + z
	}
	return c.dist
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	a := pq[i]
	b := pq[j]
	if a.InRangeBots == b.InRangeBots {
		return a.DistFromO() < b.DistFromO()
	}
	return a.InRangeBots > b.InRangeBots
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Cube)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
