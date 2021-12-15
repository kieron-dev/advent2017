package days_test

import (
	"bufio"
	"container/heap"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("15", func() {
	It("does part A", func() {
		input, err := os.Open("input15")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		grid := [][]int{}
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			row := make([]int, len(line))
			for i, r := range line {
				row[i] = AToI(string(r))
			}
			grid = append(grid, row)
		}

		Expect(shortestPath(grid, NewCoord(0, 0), NewCoord(len(grid)-1, len(grid)-1))).To(Equal(390))
	})

	It("does part B", func() {
		input, err := os.Open("input15")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		grid := [][]int{}
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			row := make([]int, 5*len(line))
			for i, r := range line {
				row[i] = AToI(string(r))
			}
			for c := 1; c < 5; c++ {
				for i := 0; i < len(line); i++ {
					n := row[i] + c
					if n > 9 {
						n -= 9
					}
					row[c*len(line)+i] = n
				}
			}
			grid = append(grid, row)
		}

		l := len(grid)
		for a := 1; a < 5; a++ {
			for r := 0; r < l; r++ {
				row := make([]int, len(grid[0]))
				for c := 0; c < len(grid[0]); c++ {
					n := grid[r][c] + a
					if n > 9 {
						n -= 9
					}
					row[c] = n
				}
				grid = append(grid, row)
			}
		}

		Expect(shortestPath(grid, NewCoord(0, 0), NewCoord(len(grid)-1, len(grid)-1))).To(Equal(2814))
	})
})

type heapItem struct {
	distance int
	coord    Coord
	index    int
}

type priorityQueue []*heapItem

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*heapItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *priorityQueue) Update(item *heapItem, distance int) {
	item.distance = distance
	heap.Fix(pq, item.index)
}

func shortestPath(grid [][]int, start, end Coord) int {
	items := map[Coord]*heapItem{}
	pq := priorityQueue{}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid); c++ {
			distance := 1000000000
			coord := NewCoord(r, c)
			if coord == start {
				distance = 0
			}
			item := &heapItem{distance: distance, coord: coord}
			items[coord] = item
			heap.Push(&pq, item)
		}
	}

	visited := map[*heapItem]bool{}

	for {
		cur := heap.Pop(&pq).(*heapItem)
		if visited[cur] {
			continue
		}

		if cur.coord == end {
			break
		}

		for _, neighbour := range squareNeighbours(grid, cur.coord.R, cur.coord.C) {
			nv := grid[neighbour.R][neighbour.C]
			nd := cur.distance + nv
			currentNeighbourDistance := items[neighbour].distance
			if nd < currentNeighbourDistance {
				pq.Update(items[neighbour], nd)
			}
		}
		visited[cur] = true
	}

	return items[end].distance
}
