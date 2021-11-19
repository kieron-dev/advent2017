package days_test

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type action string

const (
	turnOn  action = "turnOn"
	turnOff action = "turnOff"
	toggle  action = "toggle"
)

type instruction struct {
	action       action
	fromX, fromY int
	toX, toY     int
}

func newInstruction(line string) instruction {
	var a action
	switch {
	case strings.HasPrefix(line, "turn on"):
		a = turnOn
	case strings.HasPrefix(line, "turn off"):
		a = turnOff
	case strings.HasPrefix(line, "toggle"):
		a = toggle
	default:
		Fail("unexpected action: " + line)
	}

	re := regexp.MustCompile(`(\d+),(\d+) through (\d+),(\d+)`)
	matches := re.FindStringSubmatch(line)
	Expect(matches).ToNot(BeNil())

	fromX, err := strconv.Atoi(matches[1])
	Expect(err).NotTo(HaveOccurred())
	fromY, err := strconv.Atoi(matches[2])
	Expect(err).NotTo(HaveOccurred())
	toX, err := strconv.Atoi(matches[3])
	Expect(err).NotTo(HaveOccurred())
	toY, err := strconv.Atoi(matches[4])
	Expect(err).NotTo(HaveOccurred())

	return instruction{
		action: a,
		fromX:  fromX,
		fromY:  fromY,
		toX:    toX,
		toY:    toY,
	}
}

func (instr instruction) DoA(grid *[1_000_000]bool) {
	for x := instr.fromX; x <= instr.toX; x++ {
		for y := instr.fromY; y <= instr.toY; y++ {
			el := 1000*x + y
			switch instr.action {
			case turnOn:
				grid[el] = true
			case turnOff:
				grid[el] = false
			case toggle:
				grid[el] = !grid[el]
			}
		}
	}
}

func (instr instruction) DoB(grid *[1_000_000]int) {
	for x := instr.fromX; x <= instr.toX; x++ {
		for y := instr.fromY; y <= instr.toY; y++ {
			el := 1000*x + y
			switch instr.action {
			case turnOn:
				grid[el]++
			case turnOff:
				if grid[el] > 0 {
					grid[el]--
				}
			case toggle:
				grid[el] += 2
			}
		}
	}
}

var _ = Describe("06", func() {
	It("does part A", func() {
		input, err := os.Open("input06")
		Expect(err).NotTo(HaveOccurred())

		defer input.Close()

		var grid [1_000_000]bool
		scanner := bufio.NewScanner(input)

		for scanner.Scan() {
			line := scanner.Text()
			instr := newInstruction(line)
			instr.DoA(&grid)
		}

		on := 0
		for i := 0; i < 1_000_000; i++ {
			if grid[i] {
				on++
			}
		}

		Expect(on).To(Equal(569999))
	})

	It("does part B", func() {
		input, err := os.Open("input06")
		Expect(err).NotTo(HaveOccurred())

		defer input.Close()

		var grid [1_000_000]int
		scanner := bufio.NewScanner(input)

		for scanner.Scan() {
			line := scanner.Text()
			instr := newInstruction(line)
			instr.DoB(&grid)
		}

		brightness := 0
		for i := 0; i < 1_000_000; i++ {
			brightness += grid[i]
		}

		Expect(brightness).To(Equal(17836115))
	})
})
