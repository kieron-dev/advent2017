package q19

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type Computer struct {
	Registers    [6]int
	Instructions []string
	IPReg        int
	IP           int
}

func NewComputer(in io.Reader) *Computer {
	c := Computer{}
	scanner := bufio.NewScanner(in)
	first := true
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, "\n")
		if first {
			first = false
			n, err := fmt.Sscanf(line, "#ip %d", &c.IPReg)
			if err != nil {
				log.Fatal("ip load: ", err)
			}
			if n != 1 {
				log.Fatal("eh?")
			}
			continue
		}
		c.Instructions = append(c.Instructions, line)
	}
	return &c
}

func (c *Computer) Execute() int {
	i := 0
	for {
		if !c.ExecuteNext() {
			break
		}
		if i%100000 == 0 {
			fmt.Printf("i = %+v\n", i)
			fmt.Printf("c.Registers = %+v\n", c.Registers)
		}
		i++
	}
	return c.Registers[0]
}

func (c *Computer) ExecuteNext() bool {
	if c.IP < 0 || c.IP >= len(c.Instructions) {
		return false
	}
	c.Registers[c.IPReg] = c.IP
	line := c.Instructions[c.IP]
	c.ExecuteLine(line)
	c.IP = c.Registers[c.IPReg] + 1
	return true
}

func (c *Computer) ExecuteLine(line string) {
	var inst string
	var x, y, z int
	n, err := fmt.Sscanf(line, "%s %d %d %d", &inst, &x, &y, &z)
	if err != nil {
		fmt.Printf("line = %+v\n", line)
		panic("oops")
	}
	if n != 4 {
		log.Fatal("can't execute", line)
	}
	c.Ops()[inst](x, y, z)
}

func (c *Computer) Equals(other *Computer) bool {
	for i, r := range c.Registers {
		if r != other.Registers[i] {
			return false
		}
	}
	return true
}

func (comp *Computer) SetRegisters(a, b, c, d, e, f int) {
	comp.Registers[0] = a
	comp.Registers[1] = b
	comp.Registers[2] = c
	comp.Registers[3] = d
	comp.Registers[4] = e
	comp.Registers[5] = f
}

type Op func(a, b, c int)

func (c *Computer) Ops() map[string]Op {
	return map[string]Op{
		"addr": c.Addr,
		"addi": c.Addi,
		"mulr": c.Mulr,
		"muli": c.Muli,
		"banr": c.Banr,
		"bani": c.Bani,
		"borr": c.Borr,
		"bori": c.Bori,
		"setr": c.Setr,
		"seti": c.Seti,
		"gtir": c.Gtir,
		"gtri": c.Gtri,
		"gtrr": c.Gtrr,
		"eqir": c.Eqir,
		"eqri": c.Eqri,
		"eqrr": c.Eqrr,
	}
}

func (c *Computer) Addr(R1, R2, R3 int) {
	c.Registers[R3] = c.Registers[R1] + c.Registers[R2]
}

func (c *Computer) Addi(R1, V2, R3 int) {
	c.Registers[R3] = c.Registers[R1] + V2
}

func (c *Computer) Mulr(R1, R2, R3 int) {
	c.Registers[R3] = c.Registers[R1] * c.Registers[R2]
}

func (c *Computer) Muli(R1, V2, R3 int) {
	c.Registers[R3] = c.Registers[R1] * V2
}

func (c *Computer) Banr(R1, R2, R3 int) {
	c.Registers[R3] = c.Registers[R1] & c.Registers[R2]
}

func (c *Computer) Bani(R1, V2, R3 int) {
	c.Registers[R3] = c.Registers[R1] & V2
}

func (c *Computer) Borr(R1, R2, R3 int) {
	c.Registers[R3] = c.Registers[R1] | c.Registers[R2]
}

func (c *Computer) Bori(R1, V2, R3 int) {
	c.Registers[R3] = c.Registers[R1] | V2
}

func (c *Computer) Setr(R1, _, R3 int) {
	c.Registers[R3] = c.Registers[R1]
}

func (c *Computer) Seti(V1, _, R3 int) {
	c.Registers[R3] = V1
}

func (c *Computer) Gtir(V1, R2, R3 int) {
	if V1 > c.Registers[R2] {
		c.Registers[R3] = 1
	} else {
		c.Registers[R3] = 0
	}
}

func (c *Computer) Gtri(R1, V2, R3 int) {
	if c.Registers[R1] > V2 {
		c.Registers[R3] = 1
	} else {
		c.Registers[R3] = 0
	}
}

func (c *Computer) Gtrr(R1, R2, R3 int) {
	if c.Registers[R1] > c.Registers[R2] {
		c.Registers[R3] = 1
	} else {
		c.Registers[R3] = 0
	}
}

func (c *Computer) Eqir(V1, R2, R3 int) {
	if V1 == c.Registers[R2] {
		c.Registers[R3] = 1
	} else {
		c.Registers[R3] = 0
	}
}

func (c *Computer) Eqri(R1, V2, R3 int) {
	if c.Registers[R1] == V2 {
		c.Registers[R3] = 1
	} else {
		c.Registers[R3] = 0
	}
}

func (c *Computer) Eqrr(R1, R2, R3 int) {
	if c.Registers[R1] == c.Registers[R2] {
		c.Registers[R3] = 1
	} else {
		c.Registers[R3] = 0
	}
}
