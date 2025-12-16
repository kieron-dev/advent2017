package twentytwentyfive_test

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ex06 = `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`

func Test06a(t *testing.T) {
	real, err := os.Open("input06")
	Check(err)

	type tc struct {
		in       io.Reader
		expected int
	}

	tcs := map[string]tc{
		"ex01": {
			in:       strings.NewReader(ex06),
			expected: 4277556,
		},
		"real": {
			in:       real,
			expected: 4805473544166,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, calcA(tc.in))
		})
	}
}

func Test06b(t *testing.T) {
	real, err := os.Open("input06")
	Check(err)

	type tc struct {
		in       io.Reader
		expected int
	}

	tcs := map[string]tc{
		"ex01": {
			in:       strings.NewReader(ex06),
			expected: 3263827,
		},
		"real": {
			in:       real,
			expected: 8907730960817,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, calcB(tc.in))
		})
	}
}

func calcA(in io.Reader) int {
	parts := [][]string{}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts = append(parts, strings.Fields(line))
	}

	cols := len(parts[0])
	numCount := len(parts)

	sum := 0
	for c := range cols {
		op := parts[numCount-1][c]
		op = strings.TrimSpace(op)
		m := 0
		if op == "*" {
			m = 1
		}
		for i := range numCount - 1 {
			n, err := strconv.Atoi(parts[i][c])
			Check(err)
			if strings.TrimSpace(op) == "*" {
				m *= n
			} else {
				m += n
			}
		}
		sum += m
	}
	return sum
}

func calcB(in io.Reader) int {
	lines := []string{}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	opLineIdx := len(lines) - 1
	opLine := lines[opLineIdx]

	end := len(opLine)
	var start int
	sum := 0
	for i := len(opLine) - 1; i >= 0; i-- {
		start = i
		if opLine[i] != ' ' {
			sum += calc(lines, start, end)
			end = i
		}
	}

	return sum
}

func calc(lines []string, start, end int) int {
	opLineIdx := len(lines) - 1
	op := lines[opLineIdx][start]

	nums := []int{}
	for i := start; i < end; i++ {
		n := 0
		for r := range opLineIdx {
			if lines[r][i] == ' ' {
				continue
			}
			n = 10*n + int(lines[r][i]-'0')
		}
		if n != 0 {
			nums = append(nums, n)
		}
	}

	res := 0
	if op == '*' {
		res = 1
	}
	for _, n := range nums {
		if op == '*' {
			res *= n
		} else {
			res += n
		}
	}

	return res
}
