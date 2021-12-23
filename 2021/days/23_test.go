package days_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type burrow struct {
	rooms    [][]byte
	hallway  []byte
	cost     int
	roomSize int
}

func inputBurrowB() *burrow {
	return &burrow{
		rooms: [][]byte{
			{'b', 'd', 'd', 'c'},
			{'b', 'c', 'b', 'a'},
			{'d', 'b', 'a', 'd'},
			{'a', 'a', 'c', 'c'},
		},
		hallway:  make([]byte, 7),
		roomSize: 4,
	}
}

func exampleBurrowB() *burrow {
	return &burrow{
		rooms: [][]byte{
			{'b', 'd', 'd', 'a'},
			{'c', 'c', 'b', 'd'},
			{'b', 'b', 'a', 'c'},
			{'d', 'a', 'c', 'a'},
		},
		hallway:  make([]byte, 7),
		roomSize: 4,
	}
}

func inputBurrow() *burrow {
	return &burrow{
		rooms: [][]byte{
			{'b', 'c'},
			{'b', 'a'},
			{'d', 'd'},
			{'a', 'c'},
		},
		hallway:  make([]byte, 7),
		roomSize: 2,
	}
}

func exampleBurrow() *burrow {
	return &burrow{
		rooms: [][]byte{
			{'b', 'a'},
			{'c', 'd'},
			{'b', 'c'},
			{'d', 'a'},
		},
		hallway:  make([]byte, 7),
		roomSize: 2,
	}
}

func (b burrow) Key() [23]byte {
	key := [23]byte{}
	for i := 0; i < 7; i++ {
		key[i] = b.hallway[i]
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < b.roomSize; j++ {
			key[7+4*i+j] = b.rooms[i][j]
		}
	}
	return key
}

func (b *burrow) minCostSolution(stackDepth int, visited map[[23]byte]int, maxDepth int) int {
	minCost := 1000000
	visited[b.Key()] = b.cost

	if stackDepth > maxDepth {
		return minCost
	}

	if b.isDone() {
		return b.cost
	}

	for hall := 0; hall < 7; hall++ {
		cost, roomIdx, depth := b.moveToRoom(hall)
		if cost > 0 {
			b.cost += cost
			b.rooms[roomIdx][depth] = b.hallway[hall]
			b.hallway[hall] = 0
			key := b.Key()
			if v, ok := visited[key]; !ok || v > b.cost {
				mcost := b.minCostSolution(stackDepth+1, visited, maxDepth)
				if mcost < minCost {
					minCost = mcost
				}
			}
			b.hallway[hall] = b.rooms[roomIdx][depth]
			b.rooms[roomIdx][depth] = 0
			b.cost -= cost
			continue
		}
		for room := 0; room < 4; room++ {
			cost, depth := b.moveToHallway(room, hall)
			if cost > 0 {
				b.cost += cost
				b.hallway[hall] = b.rooms[room][depth]
				b.rooms[room][depth] = 0
				key := b.Key()
				if v, ok := visited[key]; !ok || v > b.cost {
					mcost := b.minCostSolution(stackDepth+1, visited, maxDepth)
					if mcost < minCost {
						minCost = mcost
					}
				}
				b.rooms[room][depth] = b.hallway[hall]
				b.hallway[hall] = 0
				b.cost -= cost
			}
		}
	}

	return minCost
}

func (b burrow) String() string {
	extra := ""
	if b.roomSize == 4 {
		extra = fmt.Sprintf(`#%c#%c#%c#%c#
  #%c#%c#%c#%c#`,
			orDot(b.rooms[0][2]),
			orDot(b.rooms[1][2]),
			orDot(b.rooms[2][2]),
			orDot(b.rooms[3][2]),
			orDot(b.rooms[0][3]),
			orDot(b.rooms[1][3]),
			orDot(b.rooms[2][3]),
			orDot(b.rooms[3][3]),
		)
	}
	return fmt.Sprintf(`
#############
#%c%c.%c.%c.%c.%c%c#
###%c#%c#%c#%c###
  #%c#%c#%c#%c#
  %s#########
                     %d
`,
		orDot(b.hallway[0]),
		orDot(b.hallway[1]),
		orDot(b.hallway[2]),
		orDot(b.hallway[3]),
		orDot(b.hallway[4]),
		orDot(b.hallway[5]),
		orDot(b.hallway[6]),
		orDot(b.rooms[0][0]),
		orDot(b.rooms[1][0]),
		orDot(b.rooms[2][0]),
		orDot(b.rooms[3][0]),
		orDot(b.rooms[0][1]),
		orDot(b.rooms[1][1]),
		orDot(b.rooms[2][1]),
		orDot(b.rooms[3][1]),
		extra,
		b.cost,
	)
}

func orDot(c byte) byte {
	if c == 0 {
		return '.'
	}
	return c
}

func (b *burrow) isDone() bool {
	for i := 0; i < 4; i++ {
		for _, v := range b.rooms[i] {
			if v != 'a'+byte(i) {
				return false
			}
		}
	}
	return true
}

var costs = map[byte]int{
	'a': 1,
	'b': 10,
	'c': 100,
	'd': 1000,
}

func (b burrow) moveToHallway(from, to int) (int, int) {
	if b.hallway[to] != 0 {
		return 0, 0
	}

	any := false
	for i := 0; i < b.roomSize; i++ {
		if b.rooms[from][i] != 0 {
			any = true
			break
		}
	}
	if !any {
		return 0, 0
	}

	complete := 0
	for i := 0; i < b.roomSize; i++ {
		if b.rooms[from][i] == 'a'+byte(from) {
			complete++
		}
	}
	if complete == b.roomSize {
		return 0, 0
	}

	for i := from + 1; i > to; i-- {
		if b.hallway[i] != 0 {
			return 0, 0
		}
	}
	for i := from + 2; i < to; i++ {
		if b.hallway[i] != 0 {
			return 0, 0
		}
	}

	var letter byte
	cost := 0
	var depth int
	for i := 0; i < b.roomSize; i++ {
		if b.rooms[from][i] != 0 {
			depth = i
			letter = b.rooms[from][i]
			break
		}
	}
	cost += costs[letter] * depth

	hPos := to
	if hPos > 1 && hPos < 5 {
		hPos = (hPos-1)*2 + 1
	} else if hPos > 4 {
		hPos += 4
	}
	rPos := from*2 + 2
	cost += costs[letter] * (abs(rPos-hPos) + 1)

	return cost, depth
}

func (b burrow) moveToRoom(from int) (int, int, int) {
	if b.hallway[from] == 0 {
		return 0, 0, 0
	}

	letter := b.hallway[from]
	roomIdx := int(letter - 'a')

	if b.rooms[roomIdx][0] != 0 {
		return 0, 0, 0
	}

	for i := 0; i < b.roomSize; i++ {
		if b.rooms[roomIdx][i] != 0 && b.rooms[roomIdx][i] != letter {
			return 0, 0, 0
		}
	}

	for i := roomIdx + 1; i > from; i-- {
		if b.hallway[i] != 0 {
			return 0, 0, 0
		}
	}
	for i := roomIdx + 2; i < from; i++ {
		if b.hallway[i] != 0 {
			return 0, 0, 0
		}
	}

	cost := 0
	depth := 0

	for i := b.roomSize - 1; i >= 0; i-- {
		if b.rooms[roomIdx][i] == 0 {
			depth = i
			break
		}
	}
	cost += costs[letter] * depth

	hPos := from
	if hPos > 1 && hPos < 5 {
		hPos = (hPos-1)*2 + 1
	} else if hPos > 4 {
		hPos += 4
	}
	rPos := roomIdx*2 + 2
	cost += costs[letter] * (abs(rPos-hPos) + 1)

	return cost, roomIdx, depth
}

var _ = Describe("23", func() {
	It("does part A", func() {
		b := inputBurrow()
		minCost := b.minCostSolution(0, map[[23]byte]int{}, 16)
		Expect(minCost).To(Equal(11608))
	})

	It("does part B", func() {
		b := inputBurrowB()
		minCost := b.minCostSolution(0, map[[23]byte]int{}, 32)
		Expect(minCost).To(Equal(46754))
	})
})
