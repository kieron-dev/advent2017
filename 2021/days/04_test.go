package days_test

import (
	"bufio"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const bingoCardSize = 5

type bingoCard struct {
	nums   [bingoCardSize * bingoCardSize]int
	curRow int
}

func (b *bingoCard) AddLine(nums []int) {
	Expect(nums).To(HaveLen(bingoCardSize))
	for i, n := range nums {
		b.nums[b.curRow*bingoCardSize+i] = n
	}

	b.curRow++
}

func (b *bingoCard) Mark(num int) {
	for i := 0; i < len(b.nums); i++ {
		if b.nums[i] == num {
			b.nums[i] = -num
		}
	}
}

func (b bingoCard) IsWinning() bool {
	for x := 0; x < bingoCardSize; x++ {
		allNegRow := true
		allNegCol := true
		for y := 0; y < bingoCardSize; y++ {
			if b.nums[x*bingoCardSize+y] > 0 {
				allNegRow = false
			}
			if b.nums[y*bingoCardSize+x] > 0 {
				allNegCol = false
			}
		}
		if allNegRow || allNegCol {
			return true
		}
	}

	return false
}

func (b bingoCard) Score() int {
	sum := 0

	for _, n := range b.nums {
		if n > 0 {
			sum += n
		}
	}

	return sum
}

func (c bingoCard) IsFull() bool {
	return c.curRow == bingoCardSize
}

var _ = Describe("04", func() {
	It("does part A", func() {
		nums, cards := loadDay04("input04")

		var score int
		var num int
	outer:
		for _, n := range nums {
			for _, c := range cards {
				c.Mark(n)
				if c.IsWinning() {
					score = c.Score()
					num = n
					break outer
				}

			}
		}

		Expect(score * num).To(Equal(27027))
	})

	It("does part B", func() {
		nums, cards := loadDay04("input04")

		var score int
		var num int
		wins := map[int]bool{}
	outer:
		for _, n := range nums {
			for i, c := range cards {
				if wins[i] {
					continue
				}
				c.Mark(n)
				if c.IsWinning() {
					wins[i] = true
					if len(wins) == len(cards) {
						score = c.Score()
						num = n
						break outer
					}
				}

			}
		}

		Expect(score * num).To(Equal(36975))
	})
})

func loadDay04(filename string) ([]int, []*bingoCard) {
	in, err := os.Open(filename)
	Expect(err).NotTo(HaveOccurred())
	defer in.Close()

	nums := []int{}
	cards := []*bingoCard{}

	curCard := &bingoCard{}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, ",") {
			nums = parseNumList(line, ",")
			continue
		}

		if line == "" {
			continue
		}

		line = strings.ReplaceAll(line, "  ", " ")
		line = strings.TrimSpace(line)
		nums := parseNumList(line, " ")
		curCard.AddLine(nums)
		if curCard.IsFull() {
			cards = append(cards, curCard)
			curCard = &bingoCard{}
		}
	}

	return nums, cards
}

func parseNumList(line, sep string) []int {
	nums := []int{}
	for _, a := range strings.Split(line, sep) {
		nums = append(nums, AToI(a))
	}

	return nums
}
