package twentytwentyfive_test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinButtonPresses(t *testing.T) {
	type tc struct {
		required         string
		buttonSchematics [][]int
		expected         int
	}

	tcs := map[string]tc{
		"ex01": {
			required:         ".##.",
			buttonSchematics: [][]int{{3}, {1, 3}, {2}, {2, 3}, {0, 2}, {0, 1}},
			expected:         2,
		},
		"ex02": {
			required:         "...#.",
			buttonSchematics: [][]int{{0, 2, 3, 4}, {2, 3}, {0, 4}, {0, 1, 2}, {1, 2, 3, 4}},
			expected:         3,
		},
		"ex03": {
			required:         ".###.#",
			buttonSchematics: [][]int{{0, 1, 2, 3, 4}, {0, 3, 4}, {0, 1, 2, 4, 5}, {1, 2}},
			expected:         2,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, minPresses(tc.required, tc.buttonSchematics))
		})
	}
}

func TestMinJoltagePresses(t *testing.T) {
	type tc struct {
		required         []int
		buttonSchematics [][]int
		expected         int
	}

	tcs := map[string]tc{
		"ex01": {
			required:         []int{3, 5, 4, 7},
			buttonSchematics: [][]int{{3}, {1, 3}, {2}, {2, 3}, {0, 2}, {0, 1}},
			expected:         10,
		},
		"ex02": {
			required:         []int{7, 5, 12, 7, 2},
			buttonSchematics: [][]int{{0, 2, 3, 4}, {2, 3}, {0, 4}, {0, 1, 2}, {1, 2, 3, 4}},
			expected:         12,
		},
		"ex03": {
			required:         []int{10, 11, 11, 5, 10, 5},
			buttonSchematics: [][]int{{0, 1, 2, 3, 4}, {0, 3, 4}, {0, 1, 2, 4, 5}, {1, 2}},
			expected:         11,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, minJoltagePresses(tc.required, tc.buttonSchematics))
		})
	}
}

var maxN = 100000

var cache = map[string]int{}

func minJoltagePresses(required []int, schematics [][]int) int {
	// fmt.Printf("%v\n", required)
	key := fmt.Sprintf("%v", required)
	if v, ok := cache[key]; ok {
		return v
	}

	done := true
	for _, r := range required {
		if r > 0 {
			done = false
			break
		}
	}
	if done {
		cache[key] = 0
		return 0
	}

	var pattern strings.Builder
	for _, r := range required {
		if r%2 == 1 {
			pattern.WriteByte('#')
		} else {
			pattern.WriteByte('.')
		}
	}

	minCount := maxN
outer:
	for _, combs := range getSolnsA(pattern.String(), schematics) {
		count := len(combs)
		joltages := make([]int, len(required))
		copy(joltages, required)
		for _, p := range combs {
			for _, n := range schematics[p] {
				joltages[n]--
				if joltages[n] < 0 {
					continue outer
				}
			}
		}

		for n, j := range joltages {
			joltages[n] = j / 2
		}

		count += 2 * minJoltagePresses(joltages, schematics)
		if count < minCount {
			minCount = count
		}
	}

	cache[key] = minCount
	return minCount
}

func Test10a(t *testing.T) {
	real, err := os.Open("input10")
	Check(err)
	type tc struct {
		in       io.Reader
		expected int
	}

	tcs := map[string]tc{
		"ex01": {
			in: strings.NewReader(`[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
`),
			expected: 7,
		},
		"real": {
			in:       real,
			expected: 409,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, totalMinPresses(tc.in))
		})
	}
}

func Test10b(t *testing.T) {
	real, err := os.Open("input10")
	Check(err)
	type tc struct {
		in       io.Reader
		expected int
	}

	tcs := map[string]tc{
		"ex01": {
			in: strings.NewReader(`
			    [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
				[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
				[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
				`),
			expected: 33,
		},
		"ex02": {
			in: strings.NewReader(`
[###.#...#.] (0,1,4,5,6,8,9) (1,3,4,7,8,9) (1,6,7,8) (0,2,3,5,7,8,9) (6,8,9) (1,3,4) (1,4,5) (1,2,6,8) (4,7,9) (0,2,3,4,5,6,7,9) (0,1,2,4,5,6,7) (4,6) (0,1,2,3,5,6,7,9) {46,102,50,59,84,57,75,80,62,55}
			`),
			expected: 123,
		},
		"real": {
			in:       real,
			expected: 15489,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, totalMinJoltagePresses(tc.in))
		})
	}
}

func totalMinPresses(in io.Reader) int {
	sum := 0
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		lastSqBr := strings.Index(line, "]")
		required := line[1:lastSqBr]

		firstRoundBr := strings.Index(line, "(")
		lastRoundBr := strings.LastIndex(line, ")")
		fields := strings.FieldsFunc(line[firstRoundBr:lastRoundBr], func(r rune) bool {
			if r == '(' || r == ')' || r == ' ' {
				return true
			}
			return false
		})
		schematics := [][]int{}
		for _, field := range fields {
			bits := strings.Split(field, ",")
			bitInts := []int{}
			for _, b := range bits {
				n, err := strconv.Atoi(b)
				Check(err)
				bitInts = append(bitInts, n)
			}
			schematics = append(schematics, bitInts)
		}
		sum += minPresses(required, schematics)
	}

	return sum
}

var cacheA = map[string][][]int{}

func getSolnsA(required string, schematics [][]int) [][]int {
	if v, ok := cacheA[required]; ok {
		return v
	}

	res := [][]int{}
	for i := range 1 << len(schematics) {
		pattern := strings.Repeat(".", len(required))
		p := 0
		n := i
		active := []int{}
		for n > 0 {
			if n%2 == 1 {
				pattern = applySchematic(pattern, schematics[p])
				active = append(active, p)
			}
			p++
			n /= 2
		}

		if pattern == required {
			res = append(res, active)
		}
	}

	cacheA[required] = res
	return res
}

func minPresses(required string, schematics [][]int) int {
	minLen := len(schematics) + 1
	cacheA = map[string][][]int{}
	for _, soln := range getSolnsA(required, schematics) {
		if len(soln) < minLen {
			minLen = len(soln)
		}
	}

	return minLen
}

func applySchematic(cur string, sc []int) string {
	bs := []byte(cur)
	for _, n := range sc {
		if bs[n] == '.' {
			bs[n] = '#'
		} else {
			bs[n] = '.'
		}
	}
	return string(bs)
}

func totalMinJoltagePresses(in io.Reader) int {
	sum := 0
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		firstRoundBr := strings.Index(line, "(")
		lastRoundBr := strings.LastIndex(line, ")")
		fields := strings.FieldsFunc(line[firstRoundBr:lastRoundBr], func(r rune) bool {
			if r == '(' || r == ')' || r == ' ' {
				return true
			}
			return false
		})
		schematics := [][]int{}
		for _, field := range fields {
			bits := strings.Split(field, ",")
			bitInts := []int{}
			for _, b := range bits {
				n, err := strconv.Atoi(b)
				Check(err)
				bitInts = append(bitInts, n)
			}
			schematics = append(schematics, bitInts)
		}

		firstCurlyBr := strings.Index(line, "{")
		lastCurlyBr := strings.LastIndex(line, "}")
		required := []int{}
		for _, field := range strings.Split(line[firstCurlyBr+1:lastCurlyBr], ",") {
			n, err := strconv.Atoi(field)
			Check(err)
			required = append(required, n)
		}

		cache = map[string]int{}
		cacheA = map[string][][]int{}
		sum += minJoltagePresses(required, schematics)
	}

	return sum
}
