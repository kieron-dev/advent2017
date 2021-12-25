package days_test

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type instruction struct {
	op   string
	arg1 string
	arg2 string
}

var _ = Describe("24", func() {
	It("does an example", func() {
		input := strings.NewReader(`inp w
add z w
mod z 2
div w 2
add y w
mod y 2
div w 2
add x w
mod x 2
div w 2
mod w 2`)

		instructions := readInstructions(input)
		res, err := process(instructions, []int{14})
		Expect(err).NotTo(HaveOccurred())
		Expect(res).To(Equal(map[string]int{"w": 1, "x": 1, "y": 1, "z": 0}))
	})

	It("does part A", func() {
		// input3-7 == input4
		// input2-5 == input5
		// input7+1 == input8
		// input9+5 == input10
		// input11 == input12
		// input6-3 == input13
		// input1+6 == input14
		input, err := os.Open("input24")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		instructions := readInstructions(input)

		in := []int{3, 9, 9, 2, 4, 9, 8, 9, 4, 9, 9, 9, 6, 9}
		res, err := process(instructions, in)
		Expect(err).NotTo(HaveOccurred())
		Expect(res["z"]).To(Equal(0))
	})

	It("does part B", func() {
		// input3-7 == input4
		// input2-5 == input5
		// input7+1 == input8
		// input9+5 == input10
		// input11 == input12
		// input6-3 == input13
		// input1+6 == input14
		input, err := os.Open("input24")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		instructions := readInstructions(input)

		in := []int{1, 6, 8, 1, 1, 4, 1, 2, 1, 6, 1, 1, 1, 7}
		res, err := process(instructions, in)
		Expect(err).NotTo(HaveOccurred())
		Expect(res["z"]).To(Equal(0))
	})
})

func decr(nums []int) []int {
	l := len(nums) - 1
	for l >= 0 {
		if nums[l] > 1 {
			nums[l]--
			return nums
		}
		nums[l] = 9
		l--
	}

	return nil
}

func readInstructions(input io.Reader) []instruction {
	scanner := bufio.NewScanner(input)
	instructions := []instruction{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		instr := instruction{op: parts[0], arg1: parts[1]}
		if len(parts) == 3 {
			instr.arg2 = parts[2]
		}
		instructions = append(instructions, instr)
	}

	return instructions
}

func value(registers map[string]int, s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return registers[s]
	}

	return v
}

func process(instructions []instruction, input []int) (map[string]int, error) {
	registers := map[string]int{}
	for _, instr := range instructions {
		switch instr.op {
		case "inp":
			registers[instr.arg1] = input[0]
			input = input[1:]
		case "add":
			registers[instr.arg1] += value(registers, instr.arg2)
		case "mul":
			registers[instr.arg1] *= value(registers, instr.arg2)
		case "div":
			v := value(registers, instr.arg2)
			if v == 0 {
				return nil, fmt.Errorf("divide by zero")
			}
			registers[instr.arg1] /= value(registers, instr.arg2)
		case "mod":
			v1 := registers[instr.arg1]
			if v1 < 0 {
				return nil, errors.New("a%b, a < 0")
			}
			v2 := value(registers, instr.arg2)
			if v2 <= 0 {
				return nil, errors.New("a%b, b <= 0")
			}
			registers[instr.arg1] %= v2
		case "eql":
			if registers[instr.arg1] == registers[instr.arg2] {
				registers[instr.arg1] = 1
			} else {
				registers[instr.arg1] = 0
			}
		}
		// fmt.Printf("%d: %s %s %s %v\n", i+1, instr.op, instr.arg1, instr.arg2, registers)
	}

	return registers, nil
}
