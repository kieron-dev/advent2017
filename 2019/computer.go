package advent2019

import (
	"strconv"
	"strings"
)

type Computer struct {
	registers []int
}

func NewComputer() *Computer {
	c := Computer{}
	c.registers = []int{}
	return &c
}

func (c *Computer) SetInput(in string) {
	for _, nstr := range strings.Split(in, ",") {
		n, err := strconv.Atoi(nstr)
		if err != nil {
			panic(err)
		}
		c.registers = append(c.registers, n)
	}
}

func (c *Computer) Prime(noun, verb int) {
	c.registers[1] = noun
	c.registers[2] = verb
}

func (c *Computer) Calculate() int {
	pos := 0
	for {
		switch c.registers[pos] {
		case 1:
			c.registers[c.registers[pos+3]] = c.registers[c.registers[pos+1]] + c.registers[c.registers[pos+2]]
		case 2:
			c.registers[c.registers[pos+3]] = c.registers[c.registers[pos+1]] * c.registers[c.registers[pos+2]]
		case 99:
			return c.registers[0]
		}
		pos += 4
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
