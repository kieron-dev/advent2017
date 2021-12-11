package days_test

import (
	"bufio"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("05", func() {
	It("does part A", func() {
		file, err := os.Open("input05")
		Expect(err).NotTo(HaveOccurred())
		defer file.Close()

		niceCount := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			vowels := 0

			for _, v := range []string{"a", "e", "i", "o", "u"} {
				vowels += strings.Count(line, v)
			}

			if vowels < 3 {
				continue
			}

			double := false
			last := rune('-')
			for _, r := range line {
				if r == last {
					double = true
					break
				}
				last = r
			}

			if !double {
				continue
			}

			bad := false
			for _, s := range []string{"ab", "cd", "pq", "xy"} {
				if strings.Contains(line, s) {
					bad = true
					break
				}
			}

			if !bad {
				niceCount++
			}

		}
		Expect(niceCount).To(Equal(236))
	})

	It("does part B", func() {
		file, err := os.Open("input05")
		Expect(err).NotTo(HaveOccurred())
		defer file.Close()

		niceCount := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			if isNice(line) {
				niceCount++
			}
		}

		Expect(niceCount).To(Equal(51))
	})

	DescribeTable("isNice", func(line string, res bool) {
		Expect(isNice(line)).To(Equal(res))
	},

		Entry("1", "qjhvhtzxzqqjkmpb", true),
		Entry("2", "xxyxx", true),
		Entry("3", "uurcxstgmygtbstg", false),
		Entry("4", "ieodomkazucvgmuy", false),
		Entry("5", "aaa", false),
		Entry("6", "aaaa", true),
		Entry("7", "xyxy", true),
	)
})

func isNice(line string) bool {
	hasDoublePair := false
	pairs := map[string]int{}
	for i := 0; i < len(line)-1; i++ {
		s := line[i : i+2]
		pos, ok := pairs[s]
		if ok {
			if i-pos > 1 {
				hasDoublePair = true
				break
			}
		} else {
			pairs[s] = i
		}
	}
	if !hasDoublePair {
		return false
	}

	hasSandwich := false
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			hasSandwich = true
			break
		}
	}

	return hasSandwich
}
