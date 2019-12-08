package advent2019

import (
	"strconv"
	"strings"
	"sync"
)

type Computer struct {
	registers []int
	in        chan int
	out       chan int
}

func NewComputer(in, out chan int) *Computer {
	c := Computer{
		in:        in,
		out:       out,
		registers: []int{},
	}
	return &c
}

type ComputerArray struct {
	size       int
	isFeedback bool
	computers  []*Computer
	inputs     []chan int
}

func NewArray(size int) *ComputerArray {
	arr := ComputerArray{size: size}
	for i := 0; i < size; i++ {
		arr.inputs = append(arr.inputs, make(chan int, 100))
	}
	arr.inputs = append(arr.inputs, make(chan int, 100))
	for i := 0; i < size; i++ {
		comp := NewComputer(arr.inputs[i], arr.inputs[i+1])
		arr.computers = append(arr.computers, comp)
	}
	return &arr
}

func NewFeedbackArray(size int) *ComputerArray {
	arr := ComputerArray{size: size, isFeedback: true}
	for i := 0; i < size; i++ {
		arr.inputs = append(arr.inputs, make(chan int, 100))
	}
	for i := 0; i < size; i++ {
		comp := NewComputer(arr.inputs[i], arr.inputs[(i+1)%size])
		arr.computers = append(arr.computers, comp)
	}
	return &arr
}

func (a *ComputerArray) SetPhase(phases []int) {
	for i := 0; i < a.size; i++ {
		a.inputs[i] <- phases[i]
	}
}

func (a *ComputerArray) WriteInitialInput(n int) {
	a.inputs[0] <- n
}

func (a *ComputerArray) SetProgram(prog string) {
	for i := 0; i < a.size; i++ {
		a.computers[i].SetInput(prog)
	}
}

func (a *ComputerArray) Run() {
	var wg sync.WaitGroup

	wg.Add(a.size)
	for i := 0; i < a.size; i++ {
		go func(n int) {
			defer wg.Done()
			a.computers[n].Calculate()
		}(i)
	}

	wg.Wait()
}

func (a *ComputerArray) GetResult() int {
	var out int
	if a.isFeedback {
		out = <-a.inputs[0]
	} else {
		out = <-a.inputs[a.size]
	}
	return out
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
			c.registers[c.registers[ip+1]] = <-c.in
			jump = 2
		case 4:
			c.out <- c.ValueAt(ip, 1)
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
