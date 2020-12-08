package vm

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type Instruction struct {
	opCode string
	param  int
}

type Console struct {
	acc          int
	instructions []Instruction
	pos          int
	visited      map[int]bool
}

func NewConsole() Console {
	return Console{
		visited: map[int]bool{},
	}
}

func (c *Console) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		c.AddInstruction(line)
	}
}

func (c *Console) AddInstruction(line string) {
	c.instructions = append(c.instructions, ParseLine(line))
}

func ParseLine(line string) Instruction {
	items := strings.Split(line, " ")
	opCode := items[0]

	param, err := strconv.Atoi(items[1])
	if err != nil {
		log.Fatalf("failed to convert %q to int", items[1])
	}

	return Instruction{
		opCode: opCode,
		param:  param,
	}
}

// RunTillLoop returns true on successful termination, i.e. it didn't loop
func (c *Console) RunTillLoop() bool {
	for !c.visited[c.pos] {
		if c.pos >= len(c.instructions) {
			return true
		}

		c.visited[c.pos] = true

		inst := c.instructions[c.pos]
		switch inst.opCode {
		case "nop":
			c.pos++
		case "jmp":
			c.pos += inst.param
		case "acc":
			c.acc += inst.param
			c.pos++
		default:
			log.Fatalf("unknown instruction %+v", inst)
		}
	}

	return false
}

func (c *Console) FixInstructionTillTerm() {
	for i, inst := range c.instructions {
		if inst.opCode == "acc" {
			continue
		}

		newOp := "jmp"
		if inst.opCode == "jmp" {
			newOp = "nop"
		}

		c.instructions[i] = Instruction{
			opCode: newOp,
			param:  inst.param,
		}

		if c.RunTillLoop() {
			return
		}

		c.instructions[i] = inst
		c.Reset()
	}

	log.Fatal("failed to fix instructions")
}

func (c *Console) Reset() {
	c.visited = map[int]bool{}
	c.pos = 0
	c.acc = 0
}

func (c Console) Acc() int {
	return c.acc
}
