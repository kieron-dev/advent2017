package days_test

import (
	"bufio"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("14", func() {
	It("does part A", func() {
		s, m := load14()
		pairs := toPairs(s)

		for i := 0; i < 10; i++ {
			pairs = processPairs(pairs, m)
		}

		f := freqs(pairs)
		min := f[s[0]]
		max := f[s[0]]
		f[s[len(s)-1]]++
		for _, v := range f {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}

		Expect(max - min).To(Equal(2233))
	})

	It("does part B", func() {
		s, m := load14()
		pairs := toPairs(s)

		for i := 0; i < 40; i++ {
			pairs = processPairs(pairs, m)
		}

		f := freqs(pairs)
		f[s[len(s)-1]]++
		min := f[s[0]]
		max := f[s[0]]
		for _, v := range f {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}

		Expect(max - min).To(Equal(2884513602164))
	})
})

func processPairs(pairs map[string]int, subs map[string]string) map[string]int {
	newPairs := map[string]int{}
	for k, v := range pairs {
		sub, ok := subs[k]
		if !ok {
			newPairs[k] = v
			continue
		}
		first := k[0:1] + sub
		second := sub + k[1:2]
		newPairs[first] += v
		newPairs[second] += v
	}

	return newPairs
}

func toPairs(s string) map[string]int {
	res := map[string]int{}

	for i := 0; i < len(s)-1; i++ {
		res[s[i:i+2]]++
	}

	return res
}

func freqs(pairs map[string]int) map[byte]int {
	res := map[byte]int{}
	for k, v := range pairs {
		res[k[0]] += v
	}

	return res
}

func load14() (string, map[string]string) {
	input, err := os.Open("input14")
	Expect(err).NotTo(HaveOccurred())
	defer input.Close()

	subs := map[string]string{}
	var start string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.Contains(line, "->") {
			parts := strings.Split(line, " -> ")
			Expect(parts).To(HaveLen(2))
			subs[parts[0]] = parts[1]
			continue
		}

		start = line
	}

	return start, subs
}
