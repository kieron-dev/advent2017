package twentytwentyfive_test

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ex01 = `
987654321111111
811111111111119
234234234234278
818181911112111
`

func Test03a(t *testing.T) {
	real, err := os.Open("input03")
	if err != nil {
		panic(err)
	}
	type TC struct {
		in       io.Reader
		expected int
	}
	tcs := map[string]TC{
		"ex01": {
			in:       strings.NewReader(ex01),
			expected: 357,
		},
		"real": {
			in:       real,
			expected: 17179,
		},
	}
	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, sumJoltage(tc.in))
		})
	}
}

func Test03b(t *testing.T) {
	real, err := os.Open("input03")
	if err != nil {
		panic(err)
	}
	type TC struct {
		in       io.Reader
		expected int
	}
	tcs := map[string]TC{
		"ex01": {
			in:       strings.NewReader(ex01),
			expected: 3121910778619,
		},
		"real": {
			in:       real,
			expected: 170025781683941,
		},
	}
	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, sumJoltageB(tc.in))
		})
	}
}

func sumJoltage(in io.Reader) int {
	scanner := bufio.NewScanner(in)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		max1 := -1
		maxP := -1
		for i := 0; i < len(line)-1; i++ {
			n := int(line[i] - '0')
			if n > max1 {
				max1 = n
				maxP = i
			}
		}
		max2 := -1
		for i := maxP + 1; i < len(line); i++ {
			n := int(line[i] - '0')
			if n > max2 {
				max2 = n
			}
		}
		sum += 10*max1 + max2
	}
	return sum
}

func sumJoltageB(in io.Reader) int {
	scanner := bufio.NewScanner(in)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		maxes := make([]int, 12)
		poses := make([]int, 13)
		poses[0] = -1

		for i := range 12 {
			for j := poses[i] + 1; j < len(line)-11+i; j++ {
				n := int(line[j] - '0')
				if n > maxes[i] {
					maxes[i] = n
					poses[i+1] = j
				}
			}
		}

		joltage := 0
		for i := range 12 {
			joltage = joltage*10 + maxes[i]
		}
		// fmt.Println(line, maxes)
		sum += joltage
	}
	return sum
}
