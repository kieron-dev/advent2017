package advent2019

import (
	"errors"
)

const len = 6

type PWGen struct {
	min      int
	max      int
	exactly2 bool
	cur      []int
}

func NewPWGen(min, max int, exactly2 bool) *PWGen {
	g := &PWGen{
		min:      min,
		max:      max,
		exactly2: exactly2,
	}
	g.cur = make([]int, 6)
	for i := 0; i < len; i++ {
		g.cur[i] = 1
	}
	for g.Current() < g.min {
		_, err := g.Next()
		if err != nil {
			panic(err)
		}
	}
	return g
}

func (g *PWGen) Current() int {
	r := 0
	for i := 0; i < len; i++ {
		r = 10*r + g.cur[i]
	}
	return r
}

func (g *PWGen) Next() (int, error) {
	if err := g.increment(); err != nil {
		return 0, err
	}
	for (!g.exactly2 && !g.hasConsecSameDigits()) || (g.exactly2 && !g.HasExactly2ConsecSameDigits()) {
		if err := g.increment(); err != nil {
			return 0, err
		}
	}
	if g.Current() > g.max {
		return 0, errors.New("end")
	}
	return g.Current(), nil
}

func (g *PWGen) hasConsecSameDigits() bool {
	last := g.cur[0]

	for i := 1; i < len; i++ {
		if g.cur[i] == last {
			return true
		}
		last = g.cur[i]
	}
	return false
}

func (g *PWGen) HasExactly2ConsecSameDigits() bool {
	last := -1
	l := 0
	for i := 0; i < len; i++ {
		if g.cur[i] == last {
			l++
			continue
		}
		if l == 2 {
			return true
		}

		last = g.cur[i]
		l = 1
	}
	return l == 2
}

func (g *PWGen) increment() error {
	for i := len - 1; i >= 0; i-- {
		if g.cur[i] < 9 {
			g.cur[i]++
			return nil
		}
		if i == 0 {
			return errors.New("eol")
		}
		val := g.cur[i-1] + 1
		for j := i; j < len; j++ {
			g.cur[j] = val
		}
	}
	return errors.New("oops")
}
