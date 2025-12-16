package twentysixteen_test

import (
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Coord struct {
	x, y int
}

func NewCoord(x, y int) Coord {
	return Coord{x: x, y: y}
}

func (c Coord) Move(direction Coord, count int) Coord {
	ret := c
	ret.x += direction.x * count
	ret.y += direction.y * count

	return ret
}

func (c Coord) MoveStepwise(direction Coord, count int) []Coord {
	p := c
	ret := make([]Coord, 0, count)

	for range count {
		p = p.Move(direction, 1)
		ret = append(ret, p)
	}

	return ret
}

func (c Coord) GridDist() int {
	d := 0
	if c.x < 0 {
		d -= c.x
	} else {
		d += c.x
	}
	if c.y < 0 {
		d -= c.y
	} else {
		d += c.y
	}
	return d
}

// N, W, S, E
var directions = []Coord{
	NewCoord(0, 1),
	NewCoord(-1, 0),
	NewCoord(0, -1),
	NewCoord(1, 0),
}

func Test01a(t *testing.T) {
	testInput, err := os.ReadFile("input01")
	if err != nil {
		panic(err)
	}
	for name, tt := range map[string]struct {
		input    []byte
		expected int
	}{
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
	} {
		t.Run(name, func(t *testing.T) {
			steps := bytes.Split(tt.input, []byte(", "))

			directionIdx := 0
			pos := NewCoord(0, 0)
			for _, step := range steps {
				step := strings.TrimSpace(string(step))
				if step[0] == 'L' {
					directionIdx = (directionIdx + 1) % 4
				} else {
					directionIdx = (directionIdx + 3) % 4
				}
				count, err := strconv.Atoi(step[1:])
				if err != nil {
					panic(err)
				}
				pos = pos.Move(directions[directionIdx], count)
			}

			assert.Equal(t, tt.expected, pos.GridDist())
		})
	}
}

func Test01b(t *testing.T) {
	testInput, err := os.ReadFile("input01")
	if err != nil {
		panic(err)
	}
	for name, tt := range map[string]struct {
		input    []byte
		expected int
	}{
		"ex1": {
			input:    []byte("R8, R4, R4, R8"),
			expected: 4,
		},
		"real": {
			input:    testInput,
			expected: 113,
		},
	} {
		t.Run(name, func(t *testing.T) {
			steps := bytes.Split(tt.input, []byte(", "))

			directionIdx := 0
			pos := NewCoord(0, 0)
			visited := map[Coord]bool{pos: true}

		outer:
			for _, step := range steps {
				step := strings.TrimSpace(string(step))
				if step[0] == 'L' {
					directionIdx = (directionIdx + 1) % 4
				} else {
					directionIdx = (directionIdx + 3) % 4
				}
				count, err := strconv.Atoi(step[1:])
				if err != nil {
					panic(err)
				}
				paces := pos.MoveStepwise(directions[directionIdx], count)
				for _, pos = range paces {
					if visited[pos] {
						break outer
					}
					visited[pos] = true
				}
			}

			assert.Equal(t, tt.expected, pos.GridDist())
		})
	}
}

func TestCoordMapKey(t *testing.T) {
	m := map[Coord]bool{}
	p1 := NewCoord(0, 1)
	p2 := NewCoord(0, 1)
	p3 := NewCoord(1, 1)
	m[p1] = true
	assert.True(t, m[p2])
	assert.False(t, m[p3])
}
