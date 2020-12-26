// Package cups is for crab cup games
package cups

import (
	"strconv"
)

type Cup struct {
	num  int
	next *Cup
}

type Game struct {
	cups    map[int]*Cup
	current int
	max     int
}

func NewGame() Game {
	return Game{
		cups: map[int]*Cup{},
	}
}

func (g *Game) Load(data string, to1000000 bool) {
	var prev *Cup
	var first *Cup

	for i := 0; i < len(data); i++ {
		n := int(data[i] - '0')
		cup := Cup{
			num: n,
		}

		g.cups[n] = &cup

		if first == nil {
			first = &cup
			g.current = n
		}

		if prev != nil {
			prev.next = &cup
		}
		prev = &cup

		if n > g.max {
			g.max = n
		}
	}

	if !to1000000 {
		prev.next = first
		return
	}

	for n := g.max + 1; n <= 1000000; n++ {
		cup := Cup{
			num: n,
		}

		g.cups[n] = &cup

		prev.next = &cup
		prev = &cup
	}

	prev.next = first
	g.max = 1000000
}

func (g Game) Cups() string {
	s := ""

	cur := g.cups[1].next

	for cur.num != 1 {
		s += strconv.Itoa(cur.num)
		cur = cur.next
	}

	return s
}

func (g *Game) Play() {
	cur := g.cups[g.current]
	r1 := cur.next
	r2 := r1.next
	r3 := r2.next

	cur.next = r3.next

	next := g.current - 1
	if next == 0 {
		next = g.max
	}

	for next == r1.num || next == r2.num || next == r3.num {
		next--
		if next == 0 {
			next = g.max
		}
	}

	nextCup := g.cups[next]

	oldRight := nextCup.next
	nextCup.next = r1
	r3.next = oldRight

	g.current = cur.next.num
}

func (g Game) After1Prod() int {
	next1 := g.cups[1].next
	next2 := next1.next

	return next1.num * next2.num
}
