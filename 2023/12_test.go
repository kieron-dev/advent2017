package two023_test

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var memo map[string]int

type arrangement struct {
	pattern string
	groups  []int
	holes   []int
	missing int
}

func (a *arrangement) unfold() {
	origPattern := a.pattern
	origGroups := make([]int, len(a.groups))
	copy(origGroups, a.groups)

	for i := 0; i < 4; i++ {
		a.pattern += "?" + origPattern
		a.groups = append(a.groups, origGroups...)
	}

	a.holes = nil
	for i := range a.pattern {
		if a.pattern[i] == '?' {
			a.holes = append(a.holes, i)
		}
	}
	a.missing = sum(a.groups) - strings.Count(a.pattern, "#")
}

func loadArrangement(line string) arrangement {
	parts := strings.Split(line, " ")
	Expect(parts).To(HaveLen(2))

	arr := arrangement{}
	arr.pattern = parts[0]
	arr.groups = alisttoi(strings.ReplaceAll(parts[1], ",", " "))
	for i := range line {
		if line[i] == '?' {
			arr.holes = append(arr.holes, i)
		}
	}
	arr.missing = sum(arr.groups) - strings.Count(line, "#")

	return arr
}

func (a arrangement) chomp() (arrangement, error) {
	firstQ := strings.Index(a.pattern, "?")
	pattern := a.pattern

	if firstQ > -1 {
		precedingDot := strings.LastIndex(pattern[:firstQ], ".")
		if precedingDot == -1 {
			return a, nil
		}
		pattern = pattern[:precedingDot+1]
	}

	groups := calcGroups(pattern)
	if sliceHasPrefix(a.groups, groups) {
		return arrangement{pattern: a.pattern[len(pattern):], groups: a.groups[len(groups):]}, nil
	}

	return arrangement{}, errors.New("group mismatch")
}

func (a arrangement) toKey() string {
	return fmt.Sprintf("%s:%v", a.pattern, a.groups)
}

func (a arrangement) possibilities() int {
	key := a.toKey()
	if v, ok := memo[key]; ok {
		return v
	}

	var err error
	a, err = a.chomp()
	if err != nil {
		memo[key] = 0
		return 0
	}

	if len(a.groups) == 0 {
		if strings.Count(a.pattern, "#") == 0 {
			memo[key] = 1
			return 1
		} else {
			memo[key] = 0
			return 0
		}
	}

	if groupsToMinLen(a.groups) > len(a.pattern) {
		memo[key] = 0
		return 0
	}

	a.holes = nil
	for i := range a.pattern {
		if a.pattern[i] == '?' {
			a.holes = append(a.holes, i)
		}
	}
	a.missing = sum(a.groups) - strings.Count(a.pattern, "#")

	holes := len(a.holes)

	if holes < a.missing {
		return 0
	}

	firstHole := a.holes[0]
	a.pattern = a.pattern[:firstHole] + "#" + a.pattern[firstHole+1:]

	count := a.possibilities()
	a.pattern = a.pattern[:firstHole] + "." + a.pattern[firstHole+1:]

	count += a.possibilities()

	memo[key] = count
	return count
}

func groupsToMinLen(groups []int) int {
	return sum(groups) + len(groups) - 1
}

func calcGroups(s string) []int {
	var a []int
	for _, f := range strings.FieldsFunc(s, func(r rune) bool {
		return r == '.'
	}) {
		a = append(a, len(f))
	}

	return a
}

func sliceHasPrefix(s, p []int) bool {
	if len(p) > len(s) {
		return false
	}

	for i := range p {
		if s[i] != p[i] {
			return false
		}
	}

	return true
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func sum(a []int) int {
	var r int
	for _, i := range a {
		r += i
	}

	return r
}

func loadArrangements(filename string) []arrangement {
	f, err := os.Open(filename)
	Expect(err).NotTo(HaveOccurred())
	defer f.Close()

	arrs := []arrangement{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		arrs = append(arrs, loadArrangement(line))
	}

	return arrs
}

var _ = Describe("12", func() {
	BeforeEach(func() {
		memo = map[string]int{}
	})

	DescribeTable("chomp", func(pattern string, groups []int, resPattern string, resGroups []int, shouldFail bool) {
		a := arrangement{
			pattern: pattern,
			groups:  groups,
		}

		var err error
		a, err = a.chomp()
		if shouldFail {
			Expect(err).To(HaveOccurred())
			return
		}

		Expect(err).NotTo(HaveOccurred())
		Expect(a.pattern).To(Equal(resPattern))
		Expect(a.groups).To(Equal(resGroups))
	},
		Entry("#", "#", []int{1}, "", []int{}, false),
		Entry("#.?", "#.?", []int{1}, "?", []int{}, false),
		Entry("#.#.?", "#.#.?", []int{1}, "", []int{}, true),
		Entry("#.##.?", "#.##.?", []int{1, 2, 3}, "?", []int{3}, false),
		Entry("#.##.?", "#.##.?", []int{1, 3, 3}, "", []int{}, true),
		Entry("?", "?", []int{1}, "?", []int{1}, false),
		Entry("..#??..#??.?.", "..#??..#??.?.", []int{3, 1, 1, 1}, "#??..#??.?.", []int{3, 1, 1, 1}, false),
		Entry("#.?", "#.?", []int{2}, "?", []int{}, true),
	)

	It("does part A", func() {
		var a int
		for _, arr := range loadArrangements("input12") {
			a += arr.possibilities()
		}

		Expect(a).To(Equal(8270))
	})

	It("does part B", func() {
		var a int
		for _, arr := range loadArrangements("input12") {
			arr.unfold()
			poss := arr.possibilities()
			a += poss
		}

		Expect(a).To(Equal(204640299929836))
	})
})
