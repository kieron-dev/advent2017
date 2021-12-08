package days_test

import (
	"bufio"
	"os"
	"sort"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("08", func() {
	lcdMap := map[string]int{
		"abcefg":  0,
		"cf":      1,
		"acdeg":   2,
		"acdfg":   3,
		"bcdf":    4,
		"abdfg":   5,
		"abdefg":  6,
		"acf":     7,
		"abcdefg": 8,
		"abcdfg":  9,
	}

	It("does the example", func() {
		line := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
		parts := strings.Split(line, " | ")
		Expect(parts).To(HaveLen(2))

		trans := getTranslation(parts[0])

		digits := toDigits(parts[1], trans, lcdMap)
		Expect(digits).To(Equal([]int{5, 3, 5, 3}))
	})

	It("does part A", func() {
		input, err := os.Open("input08")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		count := 0
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, " | ")
			Expect(parts).To(HaveLen(2))

			trans := getTranslation(parts[0])

			for _, d := range toDigits(parts[1], trans, lcdMap) {
				if d == 1 || d == 4 || d == 7 || d == 8 {
					count++
				}
			}
		}

		Expect(count).To(Equal(381))
	})

	It("does part B", func() {
		input, err := os.Open("input08")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		count := 0
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, " | ")
			Expect(parts).To(HaveLen(2))

			trans := getTranslation(parts[0])

			num := 0
			for _, d := range toDigits(parts[1], trans, lcdMap) {
				num = 10*num + d
			}
			count += num
		}

		Expect(count).To(Equal(1023686))
	})
})

func toDigits(words string, trans map[rune]rune, lcdMap map[string]int) []int {
	digits := ""
	for _, r := range words {
		if r == ' ' {
			digits += " "
			continue
		}

		digits += string(trans[r])
	}

	res := []int{}
	for _, digit := range strings.Split(digits, " ") {
		slice := []rune(digit)
		sort.Slice(slice, func(a, b int) bool {
			return slice[a] < slice[b]
		})
		res = append(res, lcdMap[string(slice)])
	}

	return res
}

func getTranslation(words string) map[rune]rune {
	actualFreqs := map[rune]int{}
	for _, r := range words {
		if r == ' ' {
			continue
		}
		actualFreqs[r]++
	}

	actual6Freqs := map[rune]int{}
	oneContents := ""
	for _, word := range strings.Split(words, " ") {
		if len(word) == 2 {
			oneContents = word
			continue
		}
		if len(word) != 6 {
			continue
		}
		for _, r := range word {
			actual6Freqs[r]++
		}
	}

	trans := map[rune]rune{}
	for k, v := range actualFreqs {
		switch v {
		case 4:
			trans[k] = 'e'
		case 6:
			trans[k] = 'b'
		case 9:
			trans[k] = 'f'
		case 7:
			if actual6Freqs[k] == 3 {
				trans[k] = 'g'
			} else {
				trans[k] = 'd'
			}
		case 8:
			if strings.Contains(oneContents, string(k)) {
				trans[k] = 'c'
			} else {
				trans[k] = 'a'
			}
		}
	}
	return trans
}

func intersect(a, b []byte) []byte {
	res := []byte{}
	m := map[byte]bool{}
	for _, c := range a {
		m[c] = true
	}
	for _, c := range b {
		if m[c] {
			res = append(res, c)
		}
	}

	return res
}
