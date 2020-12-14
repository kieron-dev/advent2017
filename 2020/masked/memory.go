package masked

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Memory struct {
	valMask0       int
	valMask1       int
	addrMask0      int
	addrMask1      int
	addrXPositions []int
	memory         map[int]int
	lines          []string
}

func NewMemory() Memory {
	return Memory{
		memory:    map[int]int{},
		valMask0:  (1 << 36) - 1,
		valMask1:  0,
		addrMask0: (1 << 36) - 1,
		addrMask1: 0,
	}
}

var re = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

func (m *Memory) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		m.lines = append(m.lines, line)
	}
}

func (m *Memory) ProcessMaskingVals() {
	for _, line := range m.lines {
		if strings.HasPrefix(line, "mask = ") {
			m.valMask0, m.valMask1 = m.getMasks(line[7:])
			continue
		}

		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			addr, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatalf("str to int failed: %v", err)
			}

			val, err := strconv.Atoi(matches[2])
			if err != nil {
				log.Fatalf("str to int failed: %v", err)
			}

			m.setMem(addr, val)
			continue
		}

		log.Fatal("didn't expect to get here!")
	}
}

func (m *Memory) ProcessMaskingAddrs() {
	for _, line := range m.lines {
		if strings.HasPrefix(line, "mask = ") {
			m.addrMask0, m.addrMask1 = m.getMasks(line[7:])
			m.addrXPositions = m.getXPositions(line[7:])
			continue
		}

		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			addr, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatalf("str to int failed: %v", err)
			}

			val, err := strconv.Atoi(matches[2])
			if err != nil {
				log.Fatalf("str to int failed: %v", err)
			}

			m.setVarMem(addr, val)
			continue
		}

		log.Fatal("didn't expect to get here!")
	}
}

func (m Memory) getXPositions(mask string) []int {
	positions := []int{}

	for i, b := range mask {
		if b == 'X' {
			positions = append(positions, len(mask)-i-1)
		}
	}

	return positions
}

func (m *Memory) getMasks(mask string) (int, int) {
	mask0 := strings.ReplaceAll(mask, "X", "1")
	mask1 := strings.ReplaceAll(mask, "X", "0")

	mask064, err := strconv.ParseInt(mask0, 2, 64)
	if err != nil {
		log.Fatalf("mask0 getting failed: %v", err)
	}

	mask164, err := strconv.ParseInt(mask1, 2, 64)
	if err != nil {
		log.Fatalf("mask1 getting failed: %v", err)
	}

	return int(mask064), int(mask164)
}

func (m *Memory) setMem(addr, val int) {
	val &= m.valMask0
	val |= m.valMask1
	m.memory[addr] = val
}

func (m *Memory) setVarMem(addr, val int) {
	addr |= m.addrMask1

	for i := 0; i < (1 << len(m.addrXPositions)); i++ {
		for j := 0; j < len(m.addrXPositions); j++ {
			if i&(1<<j) > 0 {
				addr |= (1 << m.addrXPositions[j])
			} else {
				addr &^= (1 << m.addrXPositions[j])
			}
		}
		m.setMem(addr, val)
	}
}

func (m Memory) Get(addr int) int {
	return m.memory[addr]
}

func (m Memory) GetSum() int {
	sum := 0

	for _, v := range m.memory {
		sum += v
	}

	return sum
}
