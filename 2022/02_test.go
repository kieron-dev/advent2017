package two022_test

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("02", func() {
	It("does part A", func() {
		f, err := os.Open("input02")
		Expect(err).NotTo(HaveOccurred())

		scanner := bufio.NewScanner(f)

		score := 0

		for scanner.Scan() {
			line := scanner.Text()
			var them, you byte
			fmt.Sscanf(line, "%c %c", &them, &you)

			score += int(you-'W') + matchScore(them, you)
		}

		Expect(score).To(Equal(15572))
	})

	It("does part B", func() {
		f, err := os.Open("input02")
		Expect(err).NotTo(HaveOccurred())

		scanner := bufio.NewScanner(f)

		score := 0

		for scanner.Scan() {
			line := scanner.Text()
			var them, res byte
			fmt.Sscanf(line, "%c %c", &them, &res)

			switch res {
			case 'X':
				score += losePieceFor(them)
			case 'Y':
				score += drawPieceFor(them)
				score += 3
			case 'Z':
				score += winPieceFor(them)
				score += 6
			}
		}

		Expect(score).To(Equal(16098))
	})
})

func losePieceFor(them byte) int {
	switch them {
	case 'A':
		return 3
	case 'B':
		return 1
	case 'C':
		return 2
	}
	return 0
}

func drawPieceFor(them byte) int {
	switch them {
	case 'A':
		return 1
	case 'B':
		return 2
	case 'C':
		return 3
	}
	return 0
}

func winPieceFor(them byte) int {
	switch them {
	case 'A':
		return 2
	case 'B':
		return 3
	case 'C':
		return 1
	}
	return 0
}

func matchScore(them, you byte) int {
	switch {
	case you == 'X' && them == 'C':
		fallthrough
	case you == 'Y' && them == 'A':
		fallthrough
	case you == 'Z' && them == 'B':
		return 6
	case you == 'X' && them == 'A':
		fallthrough
	case you == 'Y' && them == 'B':
		fallthrough
	case you == 'Z' && them == 'C':
		return 3
	}
	return 0
}
