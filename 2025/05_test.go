package twentytwentyfive_test

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

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

var ex05 = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func Test05a(t *testing.T) {
	real, err := os.Open("input05")
	Check(err)

	type tc struct {
		in       io.Reader
		expected int
	}

	tcs := map[string]tc{
		"ex01": {
			in:       strings.NewReader(ex05),
			expected: 3,
		},
		"real": {
			in:       real,
			expected: 664,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, countFresh(tc.in))
		})
	}
}

func Test05b(t *testing.T) {
	real, err := os.Open("input05")
	Check(err)

	type tc struct {
		in       io.Reader
		expected int
	}

	tcs := map[string]tc{
		"ex01": {
			in:       strings.NewReader(ex05),
			expected: 14,
		},
		"real": {
			in:       real,
			expected: 350780324308385,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, countFreshInRanges(tc.in))
		})
	}
}

func countFresh(in io.Reader) int {
	ranges := [][2]int{}
	mode := "ranges"
	count := 0

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if mode == "ranges" {
			if line == "" {
				mode = "data"
				continue
			}
			bits := strings.Split(line, "-")
			l, err := strconv.Atoi(bits[0])
			Check(err)
			r, err := strconv.Atoi(bits[1])
			Check(err)
			ranges = append(ranges, [2]int{l, r})
			continue
		}

		n, err := strconv.Atoi(line)
		Check(err)
		for _, rng := range ranges {
			if n >= rng[0] && n <= rng[1] {
				count++
				break
			}
		}
	}

	return count
}

func countFreshInRanges(in io.Reader) int {
	ranges := [][2]int{}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" && len(ranges) > 0 {
			break
		}
		bits := strings.Split(line, "-")
		l, err := strconv.Atoi(bits[0])
		Check(err)
		r, err := strconv.Atoi(bits[1])
		Check(err)
		ranges = append(ranges, [2]int{l, r})
	}

	sort.Slice(ranges, func(a, b int) bool {
		return ranges[a][0] < ranges[b][0]
	})
	//	fmt.Println(ranges)

	last := -1
	ignore := map[int]bool{}
	for i, rng := range ranges {
		lastRange := [2]int{-1, -1}
		if last > -1 {
			lastRange = ranges[last]
		}

		l := rng[0]
		r := rng[1]
		if l <= lastRange[1] && r <= lastRange[1] {
			ignore[i] = true
			continue
		}
		if l <= lastRange[1] {
			ranges[last][1] = r
			ignore[i] = true
			continue
		}
		last = i
	}
	//	fmt.Println(ranges)

	count := 0
	for i, r := range ranges {
		if ignore[i] {
			continue
		}
		count += r[1] - r[0] + 1
	}
	return count
}
