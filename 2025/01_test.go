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

var tc01 = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func TestDay01a(t *testing.T) {
	type testcase struct {
		in       io.Reader
		expected int
	}

	real, err := os.Open("input01")
	if err != nil {
		panic(err)
	}

	testcases := map[string]testcase{
		"ex01": {
			in:       strings.NewReader(tc01),
			expected: 3,
		},
		"real": {
			in:       real,
			expected: 1086,
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, countZerosA(tc.in))
		})
	}
}

func TestDay01b(t *testing.T) {
	type testcase struct {
		in       io.Reader
		expected int
	}

	real, err := os.Open("input01")
	if err != nil {
		panic(err)
	}

	testcases := map[string]testcase{
		"ex01": {
			in:       strings.NewReader(tc01),
			expected: 6,
		},
		"real": {
			in:       real,
			expected: 6268,
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, countZerosB(tc.in))
		})
	}
}

func countZerosA(in io.Reader) int {
	p := 50
	count := 0

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		n, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		if line[0] == 'L' {
			p = (p - n + 100) % 100
		} else {
			p = (p + n) % 100
		}

		if p == 0 {
			count++
		}

	}
	return count
}

func countZerosB(in io.Reader) int {
	p := 50
	count := 0

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Println(p, line, count)
		n, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		prev := p
		if line[0] == 'L' {
			count += n / 100
			p -= (n % 100)
			if p <= 0 && prev > 0 {
				count++
			}
		} else {
			count += n / 100
			p += (n % 100)
			if p >= 100 && prev > 0 {
				count++
			}
		}
		p = p % 100
		if p < 0 {
			p += 100
		}
	}
	return count
}
