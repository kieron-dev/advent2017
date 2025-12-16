package twentytwentyfive_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test07a(t *testing.T) {
	real, err := os.ReadFile("input07")
	Check(err)
	type tc struct {
		in       []byte
		expected int
	}

	tcs := map[string]tc{
		"ex01": {
			in: []byte(`.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`),
			expected: 21,
		},
		"real": {
			in:       real,
			expected: 1516,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, splitA(tc.in))
		})
	}
}

func Test07b(t *testing.T) {
	real, err := os.ReadFile("input07")
	Check(err)
	type tc struct {
		in       []byte
		expected int
	}

	tcs := map[string]tc{
		"ex01": {
			in: []byte(`.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`),
			expected: 40,
		},
		"real": {
			in:       real,
			expected: 0,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, splitB(tc.in))
		})
	}
}

func splitA(in []byte) int {
	grid := bytes.Fields(in)
	sPos := bytes.Index(grid[0], []byte("S"))
	beams := map[int]bool{sPos: true}
	splitCount := 0

	for r := 1; r < len(grid); r++ {
		newBeams := map[int]bool{}
		for b := range beams {
			if grid[r][b] == '^' {
				splitCount++
				if b-1 > -1 {
					newBeams[b-1] = true
				}
				if b+1 < len(grid[r]) {
					newBeams[b+1] = true
				}
			} else {
				newBeams[b] = true
			}
		}
		beams = newBeams
	}

	return splitCount
}

func splitB(in []byte) int {
	grid := bytes.Fields(in)
	sPos := bytes.Index(grid[0], []byte("S"))
	beams := map[int]bool{sPos: true}
	grid[0][sPos] = '|'

	for r := 1; r < len(grid); r++ {
		newBeams := map[int]bool{}
		for b := range beams {
			if grid[r][b] == '^' {
				if b-1 > -1 {
					newBeams[b-1] = true
				}
				if b+1 < len(grid[r]) {
					newBeams[b+1] = true
				}
			} else {
				newBeams[b] = true
			}
		}
		for pos := range newBeams {
			grid[r][pos] = '|'
		}
		beams = newBeams
	}

	counts := make([][]int, len(grid))
	for i := range counts {
		counts[i] = make([]int, len(grid[0]))
	}

	for r := len(counts) - 1; r >= 0; r-- {
		for c, t := range grid[r] {
			if t != '|' {
				continue
			}
			if r == len(counts)-1 {
				counts[r][c] = 1
				continue
			}

			if grid[r+1][c] == '^' {
				sum := 0
				if c-1 >= 0 {
					sum += counts[r+1][c-1]
				}
				if c+1 < len(counts[0]) {
					sum += counts[r+1][c+1]
				}
				counts[r][c] = sum
				continue
			}
			counts[r][c] = counts[r+1][c]
		}
	}

	return counts[0][sPos]
}
