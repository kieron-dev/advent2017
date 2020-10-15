package cards

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type (
	mapping func(curPos int) (newPos int)
)

type Deck struct {
	size        int
	mapping     mapping
	positioning mapping
}

func NewDeck(size int) *Deck {
	deck := &Deck{
		size:        size,
		mapping:     identity,
		positioning: identity,
	}
	return deck
}

func (d *Deck) SetShuffle(steps io.Reader) {
	bufReader := bufio.NewReader(steps)

	var err error
	var line string

	for err != io.EOF {
		line, err = bufReader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		d.AddTransform(strings.TrimSpace(line))
	}
}

func (d *Deck) PosOf(i int) int {
	return d.mapping(i) % d.size
}

func (d *Deck) CardAt(pos int) int {
	return d.positioning(pos)
}

func (d *Deck) Cards() []int {
	res := make([]int, d.size)

	for i := 0; i < d.size; i++ {
		res[d.mapping(i)%d.size] = i
	}

	return res
}

func (d *Deck) Period(pos int) int {
	cur := pos

	i := 0
	for {
		cur = d.CardAt(cur)
		i++
		if cur == pos {
			break
		}
		if i%1000000 == 0 {
			fmt.Printf("i = %+v, pos = %+v\n", i, cur)
		}
	}

	return i
}

func (d *Deck) AddTransform(line string) {
	if line == "deal into new stack" {
		d.mapping = d.reverse(d.mapping)
		d.positioning = d.invReverse(d.positioning)
		return
	}

	if strings.HasPrefix(line, "cut ") {
		cutNum, err := strconv.Atoi(line[4:])
		if err != nil {
			panic(err)
		}
		d.mapping = d.cut(d.mapping, cutNum)
		d.positioning = d.invCut(d.positioning, cutNum)
		return
	}

	if strings.HasPrefix(line, "deal with increment ") {
		inc, err := strconv.Atoi(line[20:])
		if err != nil {
			panic(err)
		}
		d.mapping = d.dealWithInc(d.mapping, inc)
		d.positioning = d.invDealWithInc(d.positioning, inc)
		return
	}
}

func identity(i int) int {
	return i
}

func (d *Deck) reverse(fn mapping) mapping {
	return func(i int) int {
		return d.size - fn(i) - 1
	}
}

func (d *Deck) invReverse(fn mapping) mapping {
	return func(i int) int {
		return fn(d.size - i - 1)
	}
}

func (d *Deck) cut(fn mapping, cutNum int) mapping {
	return func(i int) int {
		return (fn(i) - cutNum + d.size) % d.size
	}
}

func (d *Deck) invCut(fn mapping, cutNum int) mapping {
	return func(i int) int {
		return fn((i + cutNum + d.size) % d.size)
	}
}

func (d *Deck) dealWithInc(fn mapping, inc int) mapping {
	return func(i int) int {
		return (fn(i) * inc) % d.size
	}
}

func (d *Deck) invDealWithInc(fn mapping, inc int) mapping {
	_, inv, _ := xgcd(inc, d.size)
	if inv < 0 {
		inv += d.size
	}

	return func(i int) int {
		return fn((i * inv) % d.size)
	}
}

func xgcd(a, b int) (int, int, int) {
	a0, a1, b0, b1 := 1, 0, 0, 1

	for {
		q := a / b
		a = a % b
		a0 = a0 - q*a1
		b0 = b0 - q*b1

		if a == 0 {
			return b, a1, b1
		}

		q = b / a
		b = b % a
		a1 = a1 - q*a0
		b1 = b1 - q*b0

		if b == 0 {
			return a, a0, b0
		}
	}
}
