package twentytwentyfive_test

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ex04 = `
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`

func TestDay04a(t *testing.T) {
	real, err := os.Open("input04")
	if err != nil {
		panic(err)
	}

	type tc struct {
		in       io.Reader
		expected int
	}
	tcs := map[string]tc{
		"ex01": {
			in:       strings.NewReader(ex04),
			expected: 13,
		},
		"real": {
			in:       real,
			expected: 1587,
		},
	}
	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, countAccessible(tc.in))
		})
	}
}

func TestDay04b(t *testing.T) {
	real, err := os.Open("input04")
	if err != nil {
		panic(err)
	}

	type tc struct {
		in       io.Reader
		expected int
	}
	tcs := map[string]tc{
		"ex01": {
			in:       strings.NewReader(ex04),
			expected: 43,
		},
		"real": {
			in:       real,
			expected: 8946,
		},
	}
	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, countRemovable(tc.in))
		})
	}
}

func countAccessible(in io.Reader) int {
	grid := inToGrid(in)
	return len(getAccessible(grid))
}

func countRemovable(in io.Reader) int {
	grid := inToGrid(in)
	sum := 0
	for {
		accessible := getAccessible(grid)
		if len(accessible) == 0 {
			break
		}
		sum += len(accessible)
		for _, coord := range accessible {
			grid[coord[0]][coord[1]] = '.'
		}
	}

	return sum
}

func inToGrid(in io.Reader) [][]byte {
	grid := [][]byte{}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		grid = append(grid, []byte(line))
	}

	return grid
}

func getAccessible(grid [][]byte) [][2]int {
	ret := [][2]int{}
	for r := range grid {
		for c, v := range grid[r] {
			if v == '.' {
				continue
			}
			count := 0
			for x := range 3 {
				for y := range 3 {
					tr := r - 1 + x
					tc := c - 1 + y
					if tr == r && tc == c {
						continue
					}
					if tr < 0 || tc < 0 || tr >= len(grid) || tc >= len(grid[tr]) {
						continue
					}
					if grid[tr][tc] == '@' {
						count++
					}
				}
			}
			if count < 4 {
				ret = append(ret, [2]int{r, c})
			}
		}
	}

	return ret
}
