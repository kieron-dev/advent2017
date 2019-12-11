package advent2019

import (
	"math/big"
	"strings"
)

type Computer struct {
	registers    map[int64]*big.Int
	relativeBase int64
	in           chan string
	out          chan string
}

func NewComputer(in, out chan string) *Computer {
	c := Computer{
		in:           in,
		out:          out,
		registers:    map[int64]*big.Int{},
		relativeBase: 0,
	}
	return &c
}

func (c *Computer) SetInput(in string) {
	for i, nstr := range strings.Split(in, ",") {
		c.registers[int64(i)] = new(big.Int)
		c.registers[int64(i)].SetString(nstr, 10)
	}
}

func (c *Computer) Prime(noun, verb int64) {
	c.registers[1].SetInt64(noun)
	c.registers[2].SetInt64(verb)
}

func (c *Computer) Calculate() *big.Int {
	ip := int64(0)
	for {
		inst := new(big.Int)
		opCode := int(inst.Mod(c.registers[ip], big.NewInt(100)).Int64())
		switch opCode {

		case 1:
			idx := c.IndexFor(ip, 3)
			_, ok := c.registers[idx]
			if !ok {
				c.registers[idx] = new(big.Int)
			}
			c.registers[idx].Add(c.ValueAt(ip, 1), c.ValueAt(ip, 2))
			ip += 4

		case 2:
			idx := c.IndexFor(ip, 3)
			_, ok := c.registers[idx]
			if !ok {
				c.registers[idx] = new(big.Int)
			}
			c.registers[idx].Mul(c.ValueAt(ip, 1), c.ValueAt(ip, 2))
			ip += 4

		case 3:
			idx := c.IndexFor(ip, 1)
			_, ok := c.registers[idx]
			if !ok {
				c.registers[idx] = new(big.Int)
			}
			input := <-c.in
			c.registers[idx].SetString(input, 10)
			ip += 2

		case 4:
			c.out <- c.ValueAt(ip, 1).String()
			ip += 2

		case 5:
			if c.ValueAt(ip, 1).Cmp(big.NewInt(0)) != 0 {
				ip = c.ValueAt(ip, 2).Int64()
			} else {
				ip += 3
			}

		case 6:
			if c.ValueAt(ip, 1).Cmp(big.NewInt(0)) == 0 {
				ip = c.ValueAt(ip, 2).Int64()
			} else {
				ip += 3
			}

		case 7:
			idx := c.IndexFor(ip, 3)
			_, ok := c.registers[idx]
			if !ok {
				c.registers[idx] = new(big.Int)
			}
			if c.ValueAt(ip, 1).Cmp(c.ValueAt(ip, 2)) < 0 {
				c.registers[idx].SetInt64(1)
			} else {
				c.registers[idx].SetInt64(0)
			}
			ip += 4

		case 8:
			idx := c.IndexFor(ip, 3)
			_, ok := c.registers[idx]
			if !ok {
				c.registers[idx] = new(big.Int)
			}
			if c.ValueAt(ip, 1).Cmp(c.ValueAt(ip, 2)) == 0 {
				c.registers[idx].SetInt64(1)
			} else {
				c.registers[idx].SetInt64(0)
			}
			ip += 4

		case 9:
			c.relativeBase += c.ValueAt(ip, 1).Int64()
			ip += 2

		case 99:
			return c.registers[0]
		}
	}
}

func (c *Computer) ValueAt(base int64, offset int64) *big.Int {
	idx := c.IndexFor(base, offset)
	if idx < 0 {
		panic("idx out of range")
	}
	n, ok := c.registers[idx]
	if !ok {
		n = big.NewInt(0)
	}
	return n
}

func (c *Computer) IndexFor(base int64, offset int64) int64 {
	mask := int64(100)
	for i := int64(1); i < offset; i++ {
		mask *= 10
	}

	switch (c.registers[base].Int64() / mask) % 10 {

	case 0:
		return c.registers[base+offset].Int64()
	case 1:
		return base + offset

	case 2:
		return c.relativeBase + c.registers[base+offset].Int64()

	default:
		panic("eh?")
	}
}

func (c *Computer) TryCalculate() (ret *big.Int) {

	defer func() {
		if r := recover(); r != nil {
			ret = big.NewInt(-1)
		}
	}()

	return c.Calculate()
}
