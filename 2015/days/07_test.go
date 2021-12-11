package days_test

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type input struct {
	value    int
	variable string
}

func newInput(in string) input {
	num, err := strconv.Atoi(in)
	if err != nil {
		return input{variable: in}
	}

	return input{value: num}
}

func (i input) Value(vars map[string]int) (int, error) {
	if i.variable != "" {
		val, ok := vars[i.variable]
		if !ok {
			return 0, fmt.Errorf("%q not assigned yet", i.variable)
		}

		return val, nil
	}

	return i.value, nil
}

type op string

const (
	AND    op = "AND"
	OR     op = "OR"
	LSHIFT op = "LSHIFT"
	RSHIFT op = "RSHIFT"
	NOT    op = "NOT"
	ASSIGN op = "ASSIGN"
)

type bitOp struct {
	op             op
	input1, input2 input
	output         string
}

var (
	assignRE = regexp.MustCompile(`^(\w+) -> (\w+)$`)
	notRE    = regexp.MustCompile(`^NOT (\w+) -> (\w+)$`)
	binOpRE  = regexp.MustCompile(`^(\w+) (\w+) (\w+) -> (\w+)$`)
)

func newBitOp(line string) bitOp {
	matches := assignRE.FindStringSubmatch(line)
	if matches != nil {
		return bitOp{
			op:     ASSIGN,
			input1: newInput(matches[1]),
			output: matches[2],
		}
	}

	matches = notRE.FindStringSubmatch(line)
	if matches != nil {
		return bitOp{
			op:     NOT,
			input1: newInput(matches[1]),
			output: matches[2],
		}
	}

	matches = binOpRE.FindStringSubmatch(line)
	if matches != nil {
		return bitOp{
			op:     op(matches[2]),
			input1: newInput(matches[1]),
			input2: newInput(matches[3]),
			output: matches[4],
		}
	}

	Fail("unrecognized op: " + line)
	return bitOp{}
}

func (b bitOp) Eval(vars map[string]int) error {
	i1, err := b.input1.Value(vars)
	if err != nil {
		return err
	}

	switch b.op {
	case ASSIGN:
		vars[b.output] = i1
		return nil
	case NOT:
		vars[b.output] = 1 ^ i1
		return nil
	}

	i2, err := b.input2.Value(vars)
	if err != nil {
		return err
	}

	switch b.op {
	case AND:
		vars[b.output] = i1 & i2
		return nil
	case OR:
		vars[b.output] = i1 | i2
		return nil
	case LSHIFT:
		vars[b.output] = i1 << i2
		return nil
	case RSHIFT:
		vars[b.output] = i1 >> i2
		return nil
	}

	Fail("unknown op: " + string(b.op))
	return nil
}

var _ = Describe("07", func() {
	It("does part A", func() {
		input, err := os.Open("input07")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		vars := map[string]int{}
		instructions := []bitOp{}

		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			instructions = append(instructions, newBitOp(line))
		}

		for len(instructions) > 0 {
			newInstructions := []bitOp{}

			for _, instr := range instructions {
				if instr.Eval(vars) != nil {
					newInstructions = append(newInstructions, instr)
				}
			}
			instructions = newInstructions
		}

		Expect(vars["a"]).To(Equal(3176))
	})

	It("does part B", func() {
		file, err := os.Open("input07")
		Expect(err).NotTo(HaveOccurred())
		defer file.Close()

		vars := map[string]int{}
		instructions := []bitOp{}
		instructions2 := []bitOp{}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			instructions = append(instructions, newBitOp(line))
			instructions2 = append(instructions2, newBitOp(line))
		}

		for len(instructions) > 0 {
			newInstructions := []bitOp{}

			for _, instr := range instructions {
				if instr.Eval(vars) != nil {
					newInstructions = append(newInstructions, instr)
				}
			}
			instructions = newInstructions
		}

		aVal := vars["a"]

		vars = map[string]int{}

		for len(instructions2) > 0 {
			newInstructions := []bitOp{}

			for _, instr := range instructions2 {
				if instr.output == "b" {
					instr.input1 = input{value: aVal}
					instr.op = ASSIGN
				}
				if instr.Eval(vars) != nil {
					newInstructions = append(newInstructions, instr)
				}
			}
			instructions2 = newInstructions
		}

		Expect(vars["a"]).To(Equal(14710))
	})
})
