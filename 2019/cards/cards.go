// Package cards deals with permutation of a prime-sized deck of cards
package cards

import (
	"bufio"
	"io"
	"log"
	"math/big"
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

func (d *Deck) equivalentTransform() (mult, offset int) {
	offset = d.CardAt(0)
	mult = (d.CardAt(1) - d.CardAt(0)) % d.size
	if mult < 1 {
		mult += d.size
	}
	return
}

func (d *Deck) EquivalentCardAt(pos, iterations int) int {
	// ax + b
	a, b := d.equivalentTransform()

	A := big.NewInt(int64(a))
	B := big.NewInt(int64(b))
	I := big.NewInt(int64(iterations))
	M := big.NewInt(int64(d.size))

	aToI := new(big.Int).Exp(A, I, M)

	mult := aToI
	var offset *big.Int

	if a == 1 {
		offset = new(big.Int).Mul(B, I)
	} else {
		AMinusOneInv := new(big.Int).ModInverse(big.NewInt(int64(a-1)), M)
		if AMinusOneInv == nil {
			log.Fatalf("no inverse of %d in ring size %d", a-1, d.size)
		}
		top := new(big.Int).Mul(B, new(big.Int).Sub(aToI, big.NewInt(1)))
		offset = new(big.Int).Mul(top, AMinusOneInv)
	}

	P := big.NewInt(int64(pos))
	prod := new(big.Int).Mul(P, mult)
	sum := new(big.Int).Add(prod, offset)
	mod := new(big.Int).Mod(sum, M)

	return int(mod.Int64())
}

func (d *Deck) Cards() []int {
	res := make([]int, d.size)

	for i := 0; i < d.size; i++ {
		res[d.mapping(i)%d.size] = i
	}

	return res
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
	bigMod := big.NewInt(int64(d.size))
	bigInc := big.NewInt(int64(inc))

	bigInv := new(big.Int).ModInverse(bigInc, bigMod)

	return func(i int) int {
		bigI := big.NewInt(int64(i))
		prod := new(big.Int).Mul(bigI, bigInv)
		mod := new(big.Int).Mod(prod, bigMod)

		return fn(int(mod.Int64()))
	}
}
