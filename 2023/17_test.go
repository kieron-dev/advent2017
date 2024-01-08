package two023_test

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type shortestPath struct {
	plan []string
}

type item struct {
	location coord
	lostHeat int
	dir      direction
	index    int
}

func (i item) key() string {
	return fmt.Sprintf("%v %v", i.location, i.dir)
}

type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].lostHeat < pq[j].lostHeat
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x any) {
	n := len(*pq)
	it := x.(*item)
	it.index = n
	*pq = append(*pq, it)
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	it := old[n-1]
	old[n-1] = nil
	it.index = -1
	*pq = old[0 : n-1]
	return it
}

func (pq *priorityQueue) update(it *item, lostHeat int) {
	it.lostHeat = lostHeat
	heap.Fix(pq, it.index)
}

func newShortestPath(filename string) shortestPath {
	f, err := os.Open(filename)
	Expect(err).NotTo(HaveOccurred())
	defer f.Close()

	var s shortestPath
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		s.plan = append(s.plan, line)
	}

	return s
}

func (s shortestPath) heatLoss() int {
	pq := priorityQueue{}
	heap.Init(&pq)

	start1 := &item{
		location: coord{0, 0},
		lostHeat: 0,
		dir:      right,
	}

	start2 := &item{
		location: coord{0, 0},
		lostHeat: 0,
		dir:      down,
	}

	visited := map[string]bool{}
	items := map[string]*item{start1.key(): start1, start2.key(): start2}
	heap.Push(&pq, start1)
	heap.Push(&pq, start2)

	res := -1

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*item)
		if visited[cur.key()] {
			continue
		}
		visited[cur.key()] = true

		if (cur.location == coord{len(s.plan) - 1, len(s.plan[0]) - 1}) {
			res = cur.lostHeat
			break
		}

		for _, d := range []direction{up, down, left, right} {
			if d == cur.dir {
				continue
			}
			if (coord(cur.dir).add(coord(d)) == coord{}) {
				continue
			}
			newHeat := 0
			for i := 1; i < 4; i++ {
				n := cur.location.add(coord(d).mult(i))
				if n[0] < 0 || n[0] >= len(s.plan[0]) || n[1] < 0 || n[1] >= len(s.plan) {
					continue
				}
				nItem := item{location: n, dir: d}
				existingItem, ok := items[nItem.key()]
				newHeat += int(s.plan[n[0]][n[1]] - '0')
				if !ok {
					nItem.lostHeat = cur.lostHeat + newHeat
					heap.Push(&pq, &nItem)
					items[nItem.key()] = &nItem
				} else if cur.lostHeat+newHeat < existingItem.lostHeat {
					pq.update(existingItem, cur.lostHeat+newHeat)
				}
			}
		}
	}

	return res
}

func (s shortestPath) heatLossB() int {
	pq := priorityQueue{}
	heap.Init(&pq)

	start1 := &item{
		location: coord{0, 0},
		lostHeat: 0,
		dir:      right,
	}

	start2 := &item{
		location: coord{0, 0},
		lostHeat: 0,
		dir:      down,
	}

	visited := map[string]bool{}
	items := map[string]*item{start1.key(): start1, start2.key(): start2}
	heap.Push(&pq, start1)
	heap.Push(&pq, start2)

	res := -1

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*item)
		if visited[cur.key()] {
			continue
		}
		visited[cur.key()] = true

		if (cur.location == coord{len(s.plan) - 1, len(s.plan[0]) - 1}) {
			res = cur.lostHeat
			break
		}

		for _, d := range []direction{up, down, left, right} {
			if d == cur.dir {
				continue
			}
			if (coord(cur.dir).add(coord(d)) == coord{}) {
				continue
			}
			newHeat := 0
			for i := 1; i < 11; i++ {
				n := cur.location.add(coord(d).mult(i))
				if n[0] < 0 || n[0] >= len(s.plan[0]) || n[1] < 0 || n[1] >= len(s.plan) {
					break
				}
				newHeat += int(s.plan[n[0]][n[1]] - '0')
				if i < 4 {
					continue
				}
				nItem := item{location: n, dir: d}
				existingItem, ok := items[nItem.key()]
				if !ok {
					nItem.lostHeat = cur.lostHeat + newHeat
					heap.Push(&pq, &nItem)
					items[nItem.key()] = &nItem
				} else if cur.lostHeat+newHeat < existingItem.lostHeat {
					pq.update(existingItem, cur.lostHeat+newHeat)
				}
			}
		}
	}

	return res
}

var _ = Describe("17", func() {
	It("does part A", func() {
		s := newShortestPath("input17")
		Expect(s.heatLoss()).To(Equal(1076))
	})

	It("does part B", func() {
		s := newShortestPath("input17")
		Expect(s.heatLossB()).To(Equal(1219))
	})
})
