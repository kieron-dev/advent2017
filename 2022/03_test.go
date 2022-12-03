package two022_test

import (
	"bufio"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("03", func() {
	It("does part A", func() {
		f, err := os.Open("input03")
		Expect(err).NotTo(HaveOccurred())

		scanner := bufio.NewScanner(f)

		sum := 0
		for scanner.Scan() {
			line := scanner.Text()
			l := len(line)

			leftSet := map[byte]bool{}
			for i := 0; i < l/2; i++ {
				leftSet[line[i]] = true
			}

			var common byte
			for i := l / 2; i < l; i++ {
				if leftSet[line[i]] {
					common = line[i]
					break
				}
			}

			Expect(common).ToNot(BeZero())
			if common >= 'a' && common <= 'z' {
				sum += int(1 + common - 'a')
			} else {
				sum += int(27 + common - 'A')
			}
		}

		Expect(sum).To(Equal(8053))
	})

	It("does part B", func() {
		f, err := os.Open("input03")
		Expect(err).NotTo(HaveOccurred())

		scanner := bufio.NewScanner(f)

		sum := 0
		i := 0
		var itemsSet map[rune]int
		for scanner.Scan() {
			line := scanner.Text()

			var common rune
			switch i % 3 {
			case 0:
				itemsSet = map[rune]int{}
				for _, r := range line {
					itemsSet[r] = 1
				}
			case 1:
				for _, r := range line {
					if itemsSet[r] == 1 {
						itemsSet[r] = 2
					}
				}
			case 2:
				for _, r := range line {
					if itemsSet[r] == 2 {
						common = r
						if common >= 'a' && common <= 'z' {
							sum += int(1 + common - 'a')
						} else {
							sum += int(27 + common - 'A')
						}
						break
					}
				}
			}
			i++
		}

		Expect(sum).To(Equal(2425))
	})
})
