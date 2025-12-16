package twentysixteen_test

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test02a(t *testing.T) {
	input, err := os.Open("input02")
	if err != nil {
		panic(err)
	}

	for name, tt := range map[string]struct {
		input    io.Reader
		expected int
	}{
		"ex01": {
			input: strings.NewReader(`ULL
RRDDD
LURDL
UUUUD`),
			expected: 1985,
		},
		"actual": {
			input:    input,
			expected: 38961,
		},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.expected, day02a(tt.input))
		})
	}
}

func Test02b(t *testing.T) {
	input, err := os.Open("input02")
	if err != nil {
		panic(err)
	}

	for name, tt := range map[string]struct {
		input    io.Reader
		expected string
	}{
		"ex01": {
			input: strings.NewReader(`ULL
RRDDD
LURDL
UUUUD`),
			expected: "5DB3",
		},
		"actual": {
			input:    input,
			expected: "46C92",
		},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.expected, day02b(tt.input))
		})
	}
}

// keypad
// 1 2 3
// 4 5 6
// 7 8 9
var movements = map[rune]Coord{
	'U': NewCoord(0, 1),
	'D': NewCoord(0, -1),
	'L': NewCoord(-1, 0),
	'R': NewCoord(1, 0),
}

func (c Coord) Clamp(max Coord) Coord {
	ret := c
	if ret.x > max.x {
		ret.x = max.x
	}
	if ret.x < 0 {
		ret.x = 0
	}

	if ret.y > max.y {
		ret.y = max.y
	}
	if ret.y < 0 {
		ret.y = 0
	}

	return ret
}

func day02a(in io.Reader) int {
	max := NewCoord(2, 2)
	scanner := bufio.NewScanner(in)
	pos := NewCoord(1, 1)
	ret := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		for _, r := range line {
			pos = pos.Move(movements[r], 1)
			pos = pos.Clamp(max)
		}

		v := 3*(2-pos.y) + pos.x + 1
		ret = 10*ret + v

	}
	return ret
}

var keypad = []string{
	"  1  ",
	" 234 ",
	"56789",
	" ABC ",
	"  D  ",
}

func day02b(in io.Reader) string {
	max := NewCoord(4, 4)
	scanner := bufio.NewScanner(in)
	pos := NewCoord(1, 1)
	ret := []byte{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		for _, r := range line {
			npos := pos.Move(movements[r], 1)
			npos = npos.Clamp(max)
			if keypad[4-npos.y][npos.x] == ' ' {
				continue
			}
			pos = npos
		}

		ret = append(ret, keypad[4-pos.y][pos.x])

	}
	return string(ret)
}
