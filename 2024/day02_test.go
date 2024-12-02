package twenty24

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay02A(t *testing.T) {
	f, err := os.Open("input02")
	assert.NoError(t, err)
	defer f.Close()

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		entries := asToIs(t, line)
		if safe(entries) {
			sum++
			continue
		}
	}

	assert.Equal(t, 246, sum)
}

func TestDay02B(t *testing.T) {
	f, err := os.Open("input02")
	assert.NoError(t, err)
	defer f.Close()

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		entries := asToIs(t, line)
		if safe(entries) {
			sum++
			continue
		}

		for i := 0; i < len(entries); i++ {
			if safe(withoutEntry(entries, i)) {
				sum++
				break
			}
		}

	}

	assert.Equal(t, 318, sum)
}

func TestWithoutEntry(t *testing.T) {
	for name, tc := range map[string]struct {
		list     []int
		idx      int
		expected []int
	}{
		"middle": {list: []int{1, 2, 3, 4}, idx: 2, expected: []int{1, 2, 4}},
		"start":  {list: []int{1, 2, 3, 4}, idx: 0, expected: []int{2, 3, 4}},
		"end":    {list: []int{1, 2, 3, 4}, idx: 3, expected: []int{1, 2, 3}},
	} {
		t.Run(name, func(t *testing.T) {
			copy := slices.Clone(tc.list)
			res := withoutEntry(tc.list, tc.idx)
			assert.Equal(t, tc.expected, res)
			assert.Equal(t, copy, tc.list)
		})
	}
}

func withoutEntry(entries []int, idx int) []int {
	res := append([]int{}, entries[:idx]...)
	res = append(res, entries[idx+1:]...)
	return res
}

func safe(entries []int) bool {
	if len(entries) < 2 {
		return true
	}

	last := entries[0]
	isAscending := entries[1] > entries[0]

	for i := 1; i < len(entries); i++ {
		cur := entries[i]
		diff := abs(cur - last)
		if diff < 1 || diff > 3 {
			return false
		}
		if (cur > last) != isAscending {
			return false
		}
		isAscending = cur > last
		last = cur
	}
	return true
}

func aToI(t *testing.T, s string) int {
	n, err := strconv.Atoi(s)
	assert.NoError(t, err)
	return n
}

func asToIs(t *testing.T, line string) []int {
	var out []int
	for _, s := range strings.Fields(line) {
		out = append(out, aToI(t, s))
	}
	return out
}
