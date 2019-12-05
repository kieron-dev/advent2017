package advent2019

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Computer struct {
	registers   []int
	inputReader io.Reader
}

func NewComputer(inputReader io.Reader) *Computer {
	c := Computer{
		inputReader: inputReader,
		registers:   []int{},
	}
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
	ip := 0
	var jump int
	for {
		switch c.registers[ip] % 100 {
		case 1:
			c.registers[c.registers[ip+3]] = c.ValueAt(ip, 1) + c.ValueAt(ip, 2)
			jump = 4
		case 2:
			c.registers[c.registers[ip+3]] = c.ValueAt(ip, 1) * c.ValueAt(ip, 2)
			jump = 4
		case 3:
			c.registers[c.registers[ip+1]] = c.readInput()
			jump = 2
		case 4:
			fmt.Printf("--- %d\n", c.ValueAt(ip, 1))
			jump = 2
		case 5:
			if c.ValueAt(ip, 1) != 0 {
				ip = c.ValueAt(ip, 2)
				jump = 0
			} else {
				jump = 3
			}
		case 6:
			if c.ValueAt(ip, 1) == 0 {
				ip = c.ValueAt(ip, 2)
				jump = 0
			} else {
				jump = 3
			}
		case 7:
			if c.ValueAt(ip, 1) < c.ValueAt(ip, 2) {
				c.registers[c.registers[ip+3]] = 1
			} else {
				c.registers[c.registers[ip+3]] = 0
			}
			jump = 4
		case 8:
			if c.ValueAt(ip, 1) == c.ValueAt(ip, 2) {
				c.registers[c.registers[ip+3]] = 1
			} else {
				c.registers[c.registers[ip+3]] = 0
			}
			jump = 4
		case 99:
			return c.registers[0]
		}
		ip += jump
	}
}

func (c *Computer) readInput() int {
	var input int
	fmt.Print("> ")
	n, err := fmt.Fscanf(c.inputReader, "%d", &input)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("expected to read an int")
	}
	return input
}

func (c *Computer) ValueAt(base, offset int) int {
	mask := 100
	for i := 1; i < offset; i++ {
		mask *= 10
	}
	if (c.registers[base]/mask)%10 == 1 {
		return c.registers[base+offset]
	}
	return c.registers[c.registers[base+offset]]
}

func (c *Computer) TryCalculate() (ret int) {

	defer func() {
		if r := recover(); r != nil {
			ret = -1
		}
	}()

	return c.Calculate()
}
