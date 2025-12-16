package twentysixteen_test

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ToInts(numStrs []string) []int {
	ret := make([]int, 0, len(numStrs))
	for _, s := range numStrs {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ret = append(ret, n)
	}

	return ret
}

func Day03a(in io.Reader) int {
	scanner := bufio.NewScanner(in)
	validCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		numStrs := strings.Fields(line)
		nums := ToInts(numStrs)
		sort.Ints(nums)
		if nums[0]+nums[1] > nums[2] {
			validCount++
		}
	}
	return validCount
}

func Day03b(in io.Reader) int {
	scanner := bufio.NewScanner(in)
	validCount := 0
	var window [3][]int
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		numStrs := strings.Fields(line)
		window[i] = ToInts(numStrs)
		i = (i + 1) % 3

		if i == 0 {
			for c := range 3 {
				sides := make([]int, 0, 3)
				for r := range 3 {
					sides = append(sides, window[r][c])
				}
				sort.Ints(sides)
				if sides[0]+sides[1] > sides[2] {
					validCount++
				}
			}
		}
	}
	return validCount
}

func TestDay03a(t *testing.T) {
	type testcase struct {
		in       io.Reader
		expected int
	}
	in, err := os.Open("input03")
	if err != nil {
		panic(err)
	}
	testcases := map[string]testcase{
		"ex01": {
			in: strings.NewReader(`
3 4 5
1 2 7
7 2 1
				`),
			expected: 1,
		},
		"real": {
			in:       in,
			expected: 983,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Day03a(tc.in))
		})
	}
}

func TestDay03b(t *testing.T) {
	type testcase struct {
		in       io.Reader
		expected int
	}
	in, err := os.Open("input03")
	if err != nil {
		panic(err)
	}
	testcases := map[string]testcase{
		"ex01": {
			in: strings.NewReader(`
3 4 5
1 2 7
7 2 1
2 2 3
6 3 6
5 2 4
				`),
			expected: 3,
		},
		"real": {
			in:       in,
			expected: 1836,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Day03b(tc.in))
		})
	}
}
