package twentysixteen_test

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

func Test08(t *testing.T) {
	in, err := os.Open("input08")
	if err != nil {
		panic(err)
	}
	type tc struct {
		in         io.Reader
		expected   int
		rows, cols int
	}
	tcs := map[string]tc{
		"ex01": {
			in: strings.NewReader(`
				rect 3x2
				rotate column x=1 by 1
				rotate row y=0 by 4
				rotate column x=1 by 1`),
			expected: 6,
			rows:     3,
			cols:     7,
		},
		"real": {
			in:       in,
			rows:     6,
			cols:     50,
			expected: 110,
		},
	}
	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, countLit(tc.rows, tc.cols, tc.in))
		})
	}
}

func countLit(rows, cols int, in io.Reader) int {
	grid := make([][]bool, rows)
	for i := range grid {
		grid[i] = make([]bool, cols)
	}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		switch {
		case strings.Contains(line, "rect"):
			bits := strings.Fields(line)
			dims := strings.Split(bits[1], "x")
			cols, err := strconv.Atoi(dims[0])
			if err != nil {
				panic(err)
			}
			rows, err := strconv.Atoi(dims[1])
			if err != nil {
				panic(err)
			}
			for r := range rows {
				for c := range cols {
					grid[r][c] = true
				}
			}
		case strings.Contains(line, "column"):
			bits := strings.Split(line, "=")
			details := strings.Split(bits[1], " by ")
			column, err := strconv.Atoi(details[0])
			if err != nil {
				panic(err)
			}
			rotation, err := strconv.Atoi(details[1])
			if err != nil {
				panic(err)
			}
			newColumn := make([]bool, rows)
			for r := range newColumn {
				newColumn[(r+rotation)%rows] = grid[r][column]
			}
			for r, val := range newColumn {
				grid[r][column] = val
			}
		case strings.Contains(line, "row"):
			bits := strings.Split(line, "=")
			details := strings.Split(bits[1], " by ")
			row, err := strconv.Atoi(details[0])
			if err != nil {
				panic(err)
			}
			rotation, err := strconv.Atoi(details[1])
			if err != nil {
				panic(err)
			}
			newRow := make([]bool, cols)
			for c := range newRow {
				newRow[(c+rotation)%cols] = grid[row][c]
			}
			grid[row] = newRow
		}
	}

	sum := 0
	for r := range grid {
		for _, v := range grid[r] {
			if v {
				sum++
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

	return sum
}
