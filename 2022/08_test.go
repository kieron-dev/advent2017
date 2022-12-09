package two022_test

import (
	"bufio"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("08", func() {
	It("does part A", func() {
		f, err := os.Open("input08")
		Expect(err).NotTo(HaveOccurred())

		scanner := bufio.NewScanner(f)

		forest := []string{}

		for scanner.Scan() {
			line := scanner.Text()
			forest = append(forest, line)
		}

		visibles := make([]bool, len(forest)*len(forest[0]))

		visible := 0
		for r := 0; r < len(forest); r++ {
			max := -1
			for c := 0; c < len(forest[0]); c++ {
				h := atoi(forest[r][c])
				if h > max {
					max = h
					if !visibles[r*len(forest[0])+c] {
						visible++
						visibles[r*len(forest[0])+c] = true
					}
				}
			}

			max = -1
			for c := len(forest[0]) - 1; c >= 0; c-- {
				h := atoi(forest[r][c])
				if h > max {
					max = h
					if !visibles[r*len(forest[0])+c] {
						visible++
						visibles[r*len(forest[0])+c] = true
					}
				}
			}
		}

		for c := 0; c < len(forest[0]); c++ {
			max := -1
			for r := 0; r < len(forest); r++ {
				h := atoi(forest[r][c])
				if h > max {
					max = h
					if !visibles[r*len(forest[0])+c] {
						visible++
						visibles[r*len(forest[0])+c] = true
					}
				}
			}

			max = -1
			for r := len(forest) - 1; r >= 0; r-- {
				h := atoi(forest[r][c])
				if h > max {
					max = h
					if !visibles[r*len(forest[0])+c] {
						visible++
						visibles[r*len(forest[0])+c] = true
					}
				}
			}
		}

		Expect(visible).To(Equal(1803))
	})

	It("does part B", func() {
		f, err := os.Open("input08")
		Expect(err).NotTo(HaveOccurred())

		scanner := bufio.NewScanner(f)

		forest := []string{}

		for scanner.Scan() {
			line := scanner.Text()
			forest = append(forest, line)
		}

		maxScore := 0

		getScore := func(r, c int) int {
			leftCount := 0
			for i := c - 1; i >= 0; i-- {
				leftCount++
				if atoi(forest[r][i]) >= atoi(forest[r][c]) {
					break
				}
			}

			rightCount := 0
			for i := c + 1; i < len(forest[0]); i++ {
				rightCount++
				if atoi(forest[r][i]) >= atoi(forest[r][c]) {
					break
				}
			}

			topCount := 0
			for j := r - 1; j >= 0; j-- {
				topCount++
				if atoi(forest[j][c]) >= atoi(forest[r][c]) {
					break
				}
			}

			botCount := 0
			for j := r + 1; j < len(forest); j++ {
				botCount++
				if atoi(forest[j][c]) >= atoi(forest[r][c]) {
					break
				}
			}

			return leftCount * rightCount * topCount * botCount
		}

		for r := 0; r < len(forest); r++ {
			for c := 0; c < len(forest[0]); c++ {
				score := getScore(r, c)

				if score > maxScore {
					maxScore = score
				}
			}
		}

		Expect(maxScore).To(Equal(268912))
	})
})

func atoi(b byte) int {
	return int(b - '0')
}
