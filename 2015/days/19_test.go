package days_test

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("19", func() {
	It("does part A", func() {
		start, substitutions, _ := getInput19()
		newMolecules := map[string]bool{}

		for from, tos := range substitutions {
			for _, to := range tos {
				for _, idx := range IndexAll(start, from) {
					newMolecule := start[:idx] + to + start[idx+len(from):]
					newMolecules[newMolecule] = true
				}
			}
		}

		Expect(len(newMolecules)).To(Equal(509))
	})

	It("does part B", func() {
		start, _, revs := getInput19()
		orderedRevs := []string{}
		for f := range revs {
			orderedRevs = append(orderedRevs, f)
		}
		sort.Slice(orderedRevs, func(i, j int) bool {
			return len(orderedRevs[i]) > len(orderedRevs[j])
		})

		val := 0
		for start != "e" {
			any := false
			for from, to := range revs {
				idx := LastIndex(start, from)
				if idx > -1 {
					any = true
					start = start[:idx] + to + start[idx+len(from):]
					val++
				}
			}
			if !any {
				Fail("failed to terminate")
				break
			}
		}

		Expect(val).To(Equal(195))
	})
})

func LastIndex(s, sub string) int {
	all := IndexAll(s, sub)
	if len(all) == 0 {
		return -1
	}
	return all[len(all)-1]
}

func IndexAll(s string, sub string) []int {
	res := []int{}
	last := 0
	for {
		idx := strings.Index(s[last:], sub)
		if idx < 0 {
			break
		}
		res = append(res, last+idx)
		last += idx + 1
	}

	return res
}

func getInput19() (string, map[string][]string, map[string]string) {
	input, err := os.Open("input19")
	Expect(err).NotTo(HaveOccurred())
	defer input.Close()

	substitutions := map[string][]string{}
	revSubs := map[string]string{}
	start := ""
	re := regexp.MustCompile(`(\w+) => (\w+)`)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			substitutions[matches[1]] = append(substitutions[matches[1]], matches[2])
			revSubs[matches[2]] = matches[1]
		} else if line != "" {
			start = line
		}
	}
	return start, substitutions, revSubs
}
