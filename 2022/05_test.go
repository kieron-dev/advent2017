package two022_test

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("05", func() {
	It("does part A", func() {
		stacks := loadStacks()

		f, err := os.Open("input05")
		Expect(err).NotTo(HaveOccurred())
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}

			var n, from, to int

			_, err := fmt.Sscanf(line, "move %d from %d to %d", &n, &from, &to)
			Expect(err).NotTo(HaveOccurred())

			for i := 0; i < n; i++ {
				stacks[to-1] = append(stacks[to-1], pop(&stacks[from-1]))
			}
		}

		for i := range stacks {
			fmt.Printf("%s", stacks[i][len(stacks[i])-1])
		}
		fmt.Println()
	})

	It("does part B", func() {
		stacks := loadStacks()

		f, err := os.Open("input05")
		Expect(err).NotTo(HaveOccurred())
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}

			var n, from, to int

			_, err := fmt.Sscanf(line, "move %d from %d to %d", &n, &from, &to)
			Expect(err).NotTo(HaveOccurred())

			stacks[to-1] = append(stacks[to-1], popN(&stacks[from-1], n)...)
		}

		for i := range stacks {
			fmt.Printf("%s", stacks[i][len(stacks[i])-1])
		}
		fmt.Println()
	})

	It("can pop 2", func() {
		s := []string{"a", "b", "c"}
		v := popN(&s, 2)
		Expect(v).To(ConsistOf("b", "c"))
		Expect(s).To(ConsistOf("a"))
	})
})

func popN(s *[]string, n int) []string {
	stack := *s
	l := len(stack)
	v := stack[l-n:]
	*s = stack[:l-n]

	return v
}

func pop(s *[]string) string {
	stack := *s
	v := stack[len(stack)-1]
	*s = stack[:len(stack)-1]

	return v
}

func loadStacks() [][]string {
	stackF, err := os.Open("input05.1")
	Expect(err).NotTo(HaveOccurred())
	scanner := bufio.NewScanner(stackF)
	stacks := make([][]string, 9)
	for i := range stacks {
		stacks[i] = []string{}
	}

	for scanner.Scan() {
		line := scanner.Text()
		for i := range line {
			if line[i] != ' ' {
				stacks[i] = append(stacks[i], string(line[i]))
			}
		}
	}
	return stacks
}
