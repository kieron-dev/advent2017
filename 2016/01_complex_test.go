package twentysixteen_test

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplex01a(t *testing.T) {
	testInput, err := os.ReadFile("input01")
	if err != nil {
		panic(err)
	}
	type testcase struct {
		input    []byte
		expected int
	}
	testcases := map[string]testcase{
		"ex1": {
			input:    []byte("R2, L3"),
			expected: 5,
		},
		"ex2": {
			input:    []byte("R2, R2, R2"),
			expected: 2,
		},
		"ex3": {
			input:    []byte("R5, L5, R5, R3"),
			expected: 12,
		},
		"real": {
			input:    testInput,
			expected: 234,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, complexRouteA(tc.input))
		})
	}
}

func TestComplex01b(t *testing.T) {
	testInput, err := os.ReadFile("input01")
	if err != nil {
		panic(err)
	}
	type testcase struct {
		input    []byte
		expected int
	}
	testcases := map[string]testcase{
		"ex1": {
			input:    []byte("R8, R4, R4, R8"),
			expected: 4,
		},
		"real": {
			input:    testInput,
			expected: 113,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, complexRouteB(tc.input))
		})
	}
}

func complexRouteA(in []byte) int {
	dir := 1i
	var pos complex128 = 0
	inStr := string(in)
	for _, part := range strings.Split(inStr, ", ") {
		part = strings.TrimSpace(part)
		if part[0] == 'L' {
			dir *= 1i
		} else {
			dir *= -1i
		}
		dist, err := strconv.Atoi(part[1:])
		if err != nil {
			panic(err)
		}
		pos += dir * complex(float64(dist), 0)
	}

	return int(abs(real(pos)) + abs(imag(pos)))
}

func complexRouteB(in []byte) int {
	dir := 1i
	var pos complex128 = 0
	inStr := string(in)
	visited := map[complex128]bool{0: true}
outer:
	for _, part := range strings.Split(inStr, ", ") {
		part = strings.TrimSpace(part)
		if part[0] == 'L' {
			dir *= 1i
		} else {
			dir *= -1i
		}
		dist, err := strconv.Atoi(part[1:])
		if err != nil {
			panic(err)
		}
		for range dist {
			pos += dir
			if visited[pos] {
				break outer
			}
			visited[pos] = true
		}
	}

	return int(abs(real(pos)) + abs(imag(pos)))
}

func abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n
}
