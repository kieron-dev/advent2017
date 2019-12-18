package intcode

import (
	"strconv"
	"strings"
)

type Computer struct {
	registers    map[int]int
	relativeBase int
	in           chan int
	out          chan int
}

func NewComputer(in, out chan int) *Computer {
	c := Computer{
		in:           in,
		out:          out,
		registers:    map[int]int{},
		relativeBase: 0,
	}
	return &c
}

func (c *Computer) SetAddr(addr, val int) {
	c.registers[addr] = val
}

func (c *Computer) SetInput(in string) {
	for i, nstr := range strings.Split(in, ",") {
		n, err := strconv.Atoi(nstr)
		if err != nil {
			panic(err)
		}
		c.registers[i] = n
	}
}

func (c *Computer) Prime(noun, verb int) {
	c.registers[1] = noun
	c.registers[2] = verb
}

func (c *Computer) Calculate() int {
	ip := 0
	for {
		opCode := c.registers[ip] % 100
		switch opCode {

		case 1:
			idx := c.IndexFor(ip, 3)
			c.registers[idx] = c.ValueAt(ip, 1) + c.ValueAt(ip, 2)
			ip += 4

		case 2:
			idx := c.IndexFor(ip, 3)
			c.registers[idx] = c.ValueAt(ip, 1) * c.ValueAt(ip, 2)
			ip += 4

		case 3:
			idx := c.IndexFor(ip, 1)
			c.registers[idx] = <-c.in
			ip += 2

		case 4:
			c.out <- c.ValueAt(ip, 1)
			ip += 2

		case 5:
			if c.ValueAt(ip, 1) != 0 {
				ip = c.ValueAt(ip, 2)
			} else {
				ip += 3
			}

		case 6:
			if c.ValueAt(ip, 1) == 0 {
				ip = c.ValueAt(ip, 2)
			} else {
				ip += 3
			}

		case 7:
			idx := c.IndexFor(ip, 3)
			if c.ValueAt(ip, 1) < c.ValueAt(ip, 2) {
				c.registers[idx] = 1
			} else {
				c.registers[idx] = 0
			}
			ip += 4

		case 8:
			idx := c.IndexFor(ip, 3)
			if c.ValueAt(ip, 1) == c.ValueAt(ip, 2) {
				c.registers[idx] = 1
			} else {
				c.registers[idx] = 0
			}
			ip += 4

		case 9:
			c.relativeBase += c.ValueAt(ip, 1)
			ip += 2

		case 99:
			return c.registers[0]
		}
	}
}

func (c *Computer) ValueAt(base, offset int) int {
	idx := c.IndexFor(base, offset)
	if idx < 0 {
		panic("idx out of range")
	}
	return c.registers[idx]
}

func (c *Computer) IndexFor(base, offset int) int {
	mask := 100
	for i := 1; i < offset; i++ {
		mask *= 10
	}

	switch (c.registers[base] / mask) % 10 {

	case 0:
		return c.registers[base+offset]

	case 1:
		return base + offset

	case 2:
		return c.relativeBase + c.registers[base+offset]

	default:
		panic("eh?")
	}
}

func (c *Computer) TryCalculate() (ret int) {

	defer func() {
		if r := recover(); r != nil {
			ret = -1
		}
	}()

	return c.Calculate()
}
