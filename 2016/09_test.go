package twentysixteen_test

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test09a(t *testing.T) {
	real, err := os.ReadFile("input09")
	if err != nil {
		panic(err)
	}
	type tc struct {
		in       string
		expected string
	}
	tcs := map[string]tc{
		"ex01": {
			in:       "ADVENT",
			expected: "ADVENT",
		},
		"ex02": {
			in:       "A(1x5)BC",
			expected: "ABBBBBC",
		},
		"ex03": {
			in:       "(3x3)XYZ",
			expected: "XYZXYZXYZ",
		},
		"ex04": {
			in:       "(6x1)(1x3)A",
			expected: "(1x3)A",
		},
		"ex05": {
			in:       "X(8x2)(3x3)ABCY",
			expected: "X(3x3)ABC(3x3)ABCY",
		},
		"real": {
			in:       strings.TrimSpace(string(real)),
			expected: "",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			actual := decompress(tc.in)
			assert.Equal(t, tc.expected, actual, "len %d", len(actual))
		})
	}
}

func Test09b(t *testing.T) {
	real, err := os.ReadFile("input09")
	if err != nil {
		panic(err)
	}
	type tc struct {
		in       string
		expected int
	}
	tcs := map[string]tc{
		"ex01": {
			in:       "ADVENT",
			expected: 6,
		},
		"ex02": {
			in:       "(3x3)XYZ",
			expected: 9,
		},
		"ex03": {
			in:       "X(8x2)(3x3)ABCY",
			expected: 20,
		},
		"ex04": {
			in:       "(27x12)(20x12)(13x14)(7x10)(1x12)A",
			expected: 241920,
		},
		"ex05": {
			in:       "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN",
			expected: 445,
		},
		"real": {
			in:       strings.TrimSpace(string(real)),
			expected: 10780403063,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			actual := decompress2(tc.in)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

// note there is a good-nexting assumption here which seems to work
func decompress2(in string) int {
	var out int

	var commandStr strings.Builder
	var dataLen int
	var dataRepeat int

	mode := normal

	for i := 0; i < len(in); i++ {
		b := in[i]

		switch mode {
		case normal:
			if b == '(' {
				mode = command
				commandStr = strings.Builder{}
				continue
			}
			out++
		case command:
			if b == ')' {
				mode = normal
				cmd := commandStr.String()
				bits := strings.Split(cmd, "x")
				var err error
				dataLen, err = strconv.Atoi(bits[0])
				Check(err)
				dataRepeat, err = strconv.Atoi(bits[1])
				Check(err)
				recLen := decompress2(in[i+1 : i+dataLen+1])
				out += dataRepeat * recLen
				i += dataLen
				continue
			}
			commandStr.WriteByte(b)
		}
	}

	return out
}

type decompressMode int

const (
	normal decompressMode = iota
	command
	data
)

func decompress(in string) string {
	var out strings.Builder
	var commandStr strings.Builder
	var dataStr strings.Builder
	var dataLen int
	var dataRepeat int

	mode := normal

	for _, r := range in {
		switch mode {
		case normal:
			if r == '(' {
				mode = command
				commandStr = strings.Builder{}
				continue
			}
			out.WriteRune(r)
		case command:
			if r == ')' {
				mode = data
				dataStr = strings.Builder{}
				cmd := commandStr.String()
				bits := strings.Split(cmd, "x")
				var err error
				dataLen, err = strconv.Atoi(bits[0])
				if err != nil {
					panic(err)
				}
				dataRepeat, err = strconv.Atoi(bits[1])
				if err != nil {
					panic(err)
				}
				continue
			}
			commandStr.WriteRune(r)
		case data:
			dataStr.WriteRune(r)
			dataLen--
			if dataLen == 0 {
				mode = normal
				out.WriteString(strings.Repeat(dataStr.String(), dataRepeat))
			}
		}
	}

	return out.String()
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
