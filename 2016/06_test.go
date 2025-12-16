package twentysixteen_test

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var example = `
eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar
`

func Test06a(t *testing.T) {
	in, err := os.Open("input06")
	if err != nil {
		panic(err)
	}

	type testcase struct {
		in       io.Reader
		expected string
		lineLen  int
	}
	testcases := map[string]testcase{
		"ex01": {
			in:       strings.NewReader(example),
			lineLen:  6,
			expected: "easter",
		},
		"real": {
			in:       in,
			lineLen:  8,
			expected: "liwvqppc",
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, ErrorCorrectA(tc.in, tc.lineLen))
		})
	}
}

func Test06b(t *testing.T) {
	in, err := os.Open("input06")
	if err != nil {
		panic(err)
	}

	type testcase struct {
		in       io.Reader
		expected string
		lineLen  int
	}
	testcases := map[string]testcase{
		"ex01": {
			in:       strings.NewReader(example),
			lineLen:  6,
			expected: "advent",
		},
		"real": {
			in:       in,
			lineLen:  8,
			expected: "caqfbzlh",
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, ErrorCorrectB(tc.in, tc.lineLen))
		})
	}
}

func ErrorCorrectA(in io.Reader, lineLen int) string {
	freqs := make([]map[byte]int, lineLen)
	for i := range freqs {
		freqs[i] = make(map[byte]int)
	}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		for i := range line {
			freqs[i][line[i]]++
		}
	}

	ret := make([]byte, lineLen)
	for i := range ret {
		maxFreq := 0
		var maxFreqR byte

		for r, f := range freqs[i] {
			if f > maxFreq {
				maxFreq = f
				maxFreqR = r
			}
		}

		ret[i] = maxFreqR
	}
	return string(ret)
}

func ErrorCorrectB(in io.Reader, lineLen int) string {
	freqs := make([]map[byte]int, lineLen)
	for i := range freqs {
		freqs[i] = make(map[byte]int)
	}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		for i := range line {
			freqs[i][line[i]]++
		}
	}

	ret := make([]byte, lineLen)
	for i := range ret {
		minFreq := 10000
		var minFreqR byte

		for r, f := range freqs[i] {
			if f < minFreq {
				minFreq = f
				minFreqR = r
			}
		}

		ret[i] = minFreqR
	}
	return string(ret)
}
