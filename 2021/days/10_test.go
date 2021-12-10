package days_test

import (
	"bufio"
	"os"
	"sort"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("10", func() {
	It("does part A", func() {
		input, err := os.Open("input10")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		score := 0
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.TrimSpace(line)

			errChar, _ := firstErr(line)
			switch errChar {
			case ')':
				score += 3
			case ']':
				score += 57
			case '}':
				score += 1197
			case '>':
				score += 25137
			}
		}

		Expect(score).To(Equal(370407))
	})

	It("does part B", func() {
		input, err := os.Open("input10")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		scores := []int{}
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.TrimSpace(line)

			errChar, stack := firstErr(line)
			if errChar != ' ' {
				continue
			}
			scores = append(scores, scoreStack(stack))
		}

		sort.Ints(scores)
		mid := scores[len(scores)/2]
		Expect(mid).To(Equal(3249889609))
	})
})

var bracketMatches = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
	'>': '<',
}

func scoreStack(runes []rune) int {
	score := 0
	str := " ([{<"
	for _, r := range runes {
		score = score*5 + strings.Index(str, string(r))
	}

	return score
}

func firstErr(line string) (rune, []rune) {
	stack := []rune{}

	for _, r := range line {
		if r == '(' || r == '{' || r == '[' || r == '<' {
			stack = append([]rune{r}, stack...)
		}
		if r == ')' || r == '}' || r == ']' || r == '>' {
			if len(stack) == 0 {
				return r, nil
			}
			popped := stack[0]
			stack = stack[1:]
			if bracketMatches[r] != popped {
				return r, nil
			}
		}
	}
	return ' ', stack
}
