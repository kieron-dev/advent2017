package two023_test

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

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

func patternToGroups(pattern string) []int {
	parts := strings.FieldsFunc(pattern, func(r rune) bool {
		return r == '.'
	})

	lens := []int{}
	for _, p := range parts {
		lens = append(lens, len(p))
	}

	return lens
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

func (a arrangement) possibilities() int {
	if len(a.holes) == 0 || a.missing == 0 {
		return 1
	}

	posses := 0

	for i := 0; i < choose(len(a.holes), a.missing); i++ {
		choice := combination(len(a.holes), a.missing, i)
		s1 := bytes.ReplaceAll([]byte(a.pattern), []byte("?"), []byte("."))
		for _, n := range choice {
			s1[a.holes[n-1]] = '#'
			if sliceEqual(a.groups, patternToGroups(string(s1))) {
				posses++
			}
		}
	}

	return posses
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

// pick p from set of size n: iteration x
func combination(n, p, x int) []int {
	x++
	if p == 1 {
		return []int{x}
	}

	res := make([]int, p)
	var r, k int

	for i := 0; i < p-1; i++ {
		if i != 0 {
			res[i] = res[i-1]
		}
		for {
			res[i]++
			r = choose(n-res[i], p-(i+1))
			k += r
			if k >= x {
				break
			}
		}
		k -= r
	}
	res[p-1] = res[p-2] + x - k

	return res
}

// n C k
func choose(n, k int) int {
	return fact(n) / fact(k) / fact(n-k)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
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
	FIt("does part A", func() {
		var a int
		for _, arr := range loadArrangements("input12") {
			a += arr.possibilities()
		}

		Expect(a).To(Equal(8270))
	})

	XIt("does part B", func() {
		var a int
		for _, arr := range loadArrangements("input12a") {
			arr.unfold()
			poss := arr.possibilities()
			a += poss
			fmt.Printf("arr = %+v\n", arr)
			fmt.Printf("poss = %+v\n", poss)
		}

		Expect(a).To(Equal(8270))
	})
})
