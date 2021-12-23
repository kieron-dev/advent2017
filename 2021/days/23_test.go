package days_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type burrow struct {
	rooms   [][]byte
	hallway []byte
	cost    int
	last    *burrow
}

func inputBurrow() *burrow {
	return &burrow{
		rooms: [][]byte{
			{'b', 'c'},
			{'b', 'a'},
			{'d', 'd'},
			{'a', 'c'},
		},
		hallway: make([]byte, 7),
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
		hallway: make([]byte, 7),
	}
}

func (b burrow) Key() [15]byte {
	key := [15]byte{}
	for i := 0; i < 7; i++ {
		key[i] = b.hallway[i]
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			key[7+2*i+j] = b.rooms[i][j]
		}
	}

	return key
}

func (b burrow) String() string {
	return fmt.Sprintf(`
#############
#%c%c.%c.%c.%c.%c%c#
###%c#%c#%c#%c###
  #%c#%c#%c#%c#
  #########
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
	return b.rooms[0][0] == 'a' &&
		b.rooms[0][1] == 'a' &&
		b.rooms[1][0] == 'b' &&
		b.rooms[1][1] == 'b' &&
		b.rooms[2][0] == 'c' &&
		b.rooms[2][1] == 'c' &&
		b.rooms[3][0] == 'd' &&
		b.rooms[3][1] == 'd'
}

func (b *burrow) clone() *burrow {
	n := &burrow{
		cost:    b.cost,
		rooms:   make([][]byte, 4),
		hallway: make([]byte, 7),
	}

	for i := 0; i < 4; i++ {
		n.rooms[i] = make([]byte, 2)
		for j := 0; j < 2; j++ {
			n.rooms[i][j] = b.rooms[i][j]
		}
	}
	for i := 0; i < 7; i++ {
		n.hallway[i] = b.hallway[i]
	}
	n.last = b

	return n
}

var costs = map[byte]int{
	'a': 1,
	'b': 10,
	'c': 100,
	'd': 1000,
}

func (b *burrow) moveToHallway(from, to int) *burrow {
	if b.rooms[from][0] == 0 && b.rooms[from][1] == 0 {
		return nil
	}

	if b.rooms[from][0] == b.rooms[from][1] && b.rooms[from][0] == 'a'+byte(from) {
		return nil
	}

	if b.hallway[to] != 0 {
		return nil
	}

	for i := from + 1; i > to; i-- {
		if b.hallway[i] != 0 {
			return nil
		}
	}
	for i := from + 2; i < to; i++ {
		if b.hallway[i] != 0 {
			return nil
		}
	}

	clone := b.clone()
	var letter byte
	if clone.rooms[from][0] != 0 {
		letter = clone.rooms[from][0]
		clone.rooms[from][0] = 0
	} else {
		letter = clone.rooms[from][1]
		clone.rooms[from][1] = 0
		clone.cost += costs[letter]
	}
	clone.hallway[to] = letter

	hPos := to
	if hPos > 1 && hPos < 5 {
		hPos = (hPos-1)*2 + 1
	} else if hPos > 4 {
		hPos += 4
	}
	rPos := from*2 + 2
	clone.cost += costs[letter] * (abs(rPos-hPos) + 1)

	return clone
}

func (b *burrow) moveToRoom(from int) *burrow {
	if b.hallway[from] == 0 {
		return nil
	}

	letter := b.hallway[from]
	roomIdx := int(letter - 'a')

	if b.rooms[roomIdx][0] != 0 {
		return nil
	}
	if b.rooms[roomIdx][1] != 0 && b.rooms[roomIdx][1] != letter {
		return nil
	}

	for i := roomIdx + 1; i > from; i-- {
		if b.hallway[i] != 0 {
			return nil
		}
	}
	for i := roomIdx + 2; i < from; i++ {
		if b.hallway[i] != 0 {
			return nil
		}
	}

	c := b.clone()
	c.hallway[from] = 0
	if c.rooms[roomIdx][1] == 0 {
		c.rooms[roomIdx][1] = letter
		c.cost += costs[letter]
	} else {
		c.rooms[roomIdx][0] = letter
	}

	hPos := from
	if hPos > 1 && hPos < 5 {
		hPos = (hPos-1)*2 + 1
	} else if hPos > 4 {
		hPos += 4
	}
	rPos := roomIdx*2 + 2
	c.cost += costs[letter] * (abs(rPos-hPos) + 1)

	return c
}

func (b *burrow) possibleMoves() []*burrow {
	res := []*burrow{}
	for _, to := range []int{0, 1, 2, 3, 4, 5, 6} {
		c := b.moveToRoom(to)
		if c != nil {
			res = append(res, c)
		}
		for _, from := range []int{0, 1, 2, 3} {
			c = b.moveToHallway(from, to)
			if c != nil {
				res = append(res, c)
			}
		}
	}

	return res
}

var _ = Describe("23", func() {
	DescribeTable("can move from room top to hallway", func(from, to, expCost int) {
		b := inputBurrow()
		c := b.moveToHallway(from, to)
		Expect(c.cost).To(Equal(expCost))
		Expect(c.hallway[to]).To(Equal(b.rooms[from][0]))
		Expect(c.rooms[from][0]).To(Equal(byte(0)))
	},
		Entry("1", 0, 0, 30),
		Entry("2", 1, 0, 50),
		Entry("3", 2, 0, 7000),
		Entry("4", 3, 0, 9),
		Entry("5", 0, 1, 20),
		Entry("6", 0, 2, 20),
		Entry("7", 0, 3, 40),
		Entry("8", 0, 4, 60),
		Entry("9", 0, 5, 80),
		Entry("10", 0, 6, 90),
		Entry("11", 1, 1, 40),
		Entry("12", 1, 2, 20),
		Entry("13", 1, 3, 20),
	)

	DescribeTable("can move from room bottom to hallway", func(from, to, expCost int) {
		b := inputBurrow()
		b.rooms[from][0] = 0
		c := b.moveToHallway(from, to)
		Expect(c.cost).To(Equal(expCost))
		Expect(c.hallway[to]).To(Equal(b.rooms[from][1]))
		Expect(c.rooms[from][1]).To(Equal(byte(0)))
	},
		Entry("1", 0, 0, 400),
		Entry("2", 1, 0, 6),
		Entry("3", 2, 0, 8000),
		Entry("4", 3, 0, 1000),
		Entry("5", 0, 1, 300),
		Entry("6", 0, 2, 300),
		Entry("7", 0, 3, 500),
		Entry("8", 0, 4, 700),
		Entry("9", 0, 5, 900),
		Entry("10", 0, 6, 1000),
		Entry("11", 1, 1, 5),
		Entry("12", 1, 2, 3),
		Entry("13", 1, 3, 3),
	)

	It("can't move through others", func() {
		b := inputBurrow()
		c := b.moveToHallway(1, 3)
		Expect(c).ToNot(BeNil())
		d := c.moveToHallway(0, 5)
		Expect(d).To(BeNil())

		d = c.moveToHallway(0, 3)
		Expect(d).To(BeNil())
	})

	DescribeTable("can move from hallway to room", func(letter byte, from, expCost int) {
		b := inputBurrow()
		b.hallway[from] = letter
		b.rooms[letter-'a'][0] = 0
		b.rooms[letter-'a'][1] = letter

		c := b.moveToRoom(from)
		Expect(c).ToNot(BeNil())
		Expect(c.cost).To(Equal(expCost))
		Expect(c.hallway[from]).To(Equal(byte(0)))
		Expect(c.rooms[letter-'a'][0]).To(Equal(letter))
	},
		Entry("1", byte('a'), 0, 3),
		Entry("2", byte('b'), 0, 50),
		Entry("3", byte('c'), 0, 700),
		Entry("4", byte('d'), 0, 9000),
		Entry("5", byte('a'), 6, 9),
	)
})

var _ = FDescribe("Day 23", func() {
	It("agrees with the example cost", func() {
		b := exampleBurrow()
		b = b.moveToHallway(2, 2)
		b = b.moveToHallway(1, 3)
		b = b.moveToRoom(3)
		b = b.moveToHallway(1, 3)
		b = b.moveToRoom(2)
		b = b.moveToHallway(0, 2)
		b = b.moveToRoom(2)
		b = b.moveToHallway(3, 4)
		b = b.moveToHallway(3, 5)
		b = b.moveToRoom(4)
		b = b.moveToRoom(3)
		b = b.moveToRoom(5)
		Expect(b.cost).To(Equal(12521))
	})

	It("does part A", func() {
		b := inputBurrow()
		next := []*burrow{b}
		minCost := 1000000000
		var minBurrow *burrow
		visited := map[[15]byte]int{}

		for len(next) > 0 {
			cur := next[0]
			// fmt.Println(cur)
			next = next[1:]

			if v, ok := visited[cur.Key()]; ok && v < cur.cost {
				continue
			}

			if cur.isDone() {
				if cur.cost < minCost {
					minCost = cur.cost
					fmt.Printf("mincost = %+v\n", minCost)
					minBurrow = cur
					continue
				}
			}

			for _, n := range cur.possibleMoves() {
				if v, ok := visited[n.Key()]; !ok || v > n.cost {
					next = append(next, n)
				}
			}
			visited[cur.Key()] = cur.cost
		}

		x := minBurrow
		for x != nil {
			fmt.Println(x)
			x = x.last
		}
		Expect(minCost).To(Equal(11608))
	})
})
