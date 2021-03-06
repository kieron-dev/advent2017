package assembly

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Machine struct {
	registers    map[rune]int
	instructions []string
	ptr          int
	rcvQueue     []int
	sndCount     int
	idx          int
	partner      *Machine
	waiting      bool
}

func NewMachine(idx int) *Machine {
	m := Machine{}
	m.initialiseRegisters()
	m.registers['p'] = idx
	m.idx = idx
	m.instructions = []string{}
	m.rcvQueue = []int{}
	return &m
}

func (m *Machine) send(num int) {
	m.partner.rcvQueue = append(m.partner.rcvQueue, num)
	m.partner.waiting = false
}

func (m *Machine) rcvNum() (int, error) {
	if len(m.rcvQueue) == 0 {
		return 0, errors.New("waiting")
	}
	num := m.rcvQueue[0]
	m.rcvQueue = m.rcvQueue[1:]
	return num, nil
}

func (m *Machine) Duet(other *Machine) {
	m.partner = other
	other.partner = m
}

func (m *Machine) GetCount() int {
	return m.sndCount
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
	for m.isActive() {
		m.Execute(m.instructions[m.ptr])
	}
}

func (m *Machine) isActive() bool {
	return !m.waiting && !m.isFinished()
}

func (m *Machine) isFinished() bool {
	return m.ptr >= len(m.instructions)
}

func (m *Machine) Execute(instr string) {
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
		m.send(num)
		m.sndCount++
		m.ptr++
	case "rcv":
		reg := getReg(words[1])
		num, err := m.rcvNum()
		if err != nil {
			m.waiting = true
		} else {
			m.waiting = false
			m.registers[reg] = num
			m.ptr++
		}
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

func RunMachines(machines []*Machine) {
	m1 := machines[0]
	m2 := machines[1]
	for !m1.waiting || !m2.waiting {
		m1.Run()
		m2.Run()
		if m1.isFinished() && m2.isFinished() {
			fmt.Println("Normal exit")
			break
		}
	}

	fmt.Println("Sends from 2nd machine:", machines[1].sndCount)
}
