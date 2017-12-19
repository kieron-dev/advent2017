package assembly

import (
	"strconv"
	"strings"
)

type Machine struct {
	registers    map[rune]int
	instructions []string
	ptr          int
	lastVal      int
}

func NewMachine() *Machine {
	m := Machine{}
	m.initialiseRegisters()
	m.instructions = []string{}
	return &m
}

func (m *Machine) initialiseRegisters() {
	m.registers = map[rune]int{}
	for i := 0; i < 26; i++ {
		r := rune('a' + i)
		m.registers[r] = 0
	}
}

func (m *Machine) AppendInstruction(instr string) {
	m.instructions = append(m.instructions, instr)
}

func (m *Machine) Run() {
	for m.ptr < len(m.instructions) {
		exit := m.Execute(m.instructions[m.ptr])
		if exit {
			break
		}
	}
}

func (m *Machine) RecoverVal() int {
	return m.lastVal
}

func (m *Machine) Execute(instr string) bool {
	words := strings.Split(instr, " ")
	switch words[0] {
	case "set":
		reg := getReg(words[1])
		val := m.getNum(words[2])
		m.registers[reg] = val
		m.ptr++
	case "add":
		reg := getReg(words[1])
		addend := m.getNum(words[2])
		m.registers[reg] += addend
		m.ptr++
	case "mul":
		reg := getReg(words[1])
		factor := m.getNum(words[2])
		m.registers[reg] *= factor
		m.ptr++
	case "mod":
		reg := getReg(words[1])
		mod := m.getNum(words[2])
		m.registers[reg] %= mod
		m.ptr++
	case "snd":
		num := m.getNum(words[1])
		m.lastVal = num
		m.ptr++
	case "rcv":
		num := m.getNum(words[1])
		if num != 0 {
			return true
		}
		m.ptr++
	case "jgz":
		cond := m.getNum(words[1])
		num := m.getNum(words[2])
		if cond > 0 {
			m.ptr += num
		} else {
			m.ptr++
		}
	default:
		panic("unknown instruction: " + instr)
	}
	return false
}

func (m *Machine) GetRegister(r rune) int {
	return m.registers[r]
}

func getReg(reg string) rune {
	if len(reg) > 1 {
		panic("Not a register: " + reg)
	}
	return rune(reg[0])
}

func (m *Machine) getNum(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		if len(s) == 1 && s[0] >= 'a' && s[0] <= 'z' {
			return m.registers[rune(s[0])]
		}
		panic(err)
	}
	return val
}
