package twentytwentyfive_test

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ex09 = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func Test09a(t *testing.T) {
	real, err := os.Open("input09")
	Check(err)

	type tc struct {
		in       io.Reader
		expected int
	}

	tcs := map[string]tc{
		"ex01": {
			in:       strings.NewReader(ex09),
			expected: 50,
		},
		"real": {
			in:       real,
			expected: 4759930955,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, maxArea(tc.in))
		})
	}
}

func Test09b(t *testing.T) {
	real, err := os.Open("input09")
	Check(err)

	type tc struct {
		in       io.Reader
		expected int
	}

	tcs := map[string]tc{
		"ex01": {
			in:       strings.NewReader(ex09),
			expected: 24,
		},
		"real": {
			in:       real,
			expected: 1525241870,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, maxArea2(tc.in))
		})
	}
}

type Coord [2]int

func (c Coord) Add(d Coord) Coord {
	return Coord{c[0] + d[0], c[1] + d[1]}
}

func coordMapping(orig []Coord) map[Coord]Coord {
	xs := map[int]bool{}
	ys := map[int]bool{}

	for _, c := range orig {
		x, y := c[0], c[1]
		xs[x] = true
		ys[y] = true
	}

	var xlist, ylist []int
	for x := range xs {
		xlist = append(xlist, x)
	}
	for y := range ys {
		ylist = append(ylist, y)
	}
	sort.Ints(xlist)
	sort.Ints(ylist)

	origToNewX := map[int]int{}
	for i, x := range xlist {
		origToNewX[x] = i * 1
	}
	origToNewY := map[int]int{}
	for i, y := range ylist {
		origToNewY[y] = i * 1
	}

	out := map[Coord]Coord{}
	for _, c := range orig {
		out[c] = Coord{origToNewX[c[0]], origToNewY[c[1]]}
	}

	return out
}

func maxArea2(in io.Reader) int {
	var coords []Coord
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		bits := strings.Split(line, ",")
		x, err := strconv.Atoi(bits[0])
		Check(err)
		y, err := strconv.Atoi(bits[1])
		Check(err)
		coords = append(coords, Coord{x, y})
	}

	mapping := coordMapping(coords)
	reverse := map[Coord]Coord{}
	var maxX, maxY int
	for o, n := range mapping {
		reverse[n] = o
		if n[0] > maxX {
			maxX = n[0]
		}
		if n[1] > maxY {
			maxY = n[1]
		}
	}

	grid := make([][]byte, maxY+1)
	for r := range grid {
		grid[r] = bytes.Repeat([]byte{'.'}, maxX+1)
	}

	last := mapping[coords[len(coords)-1]]
	for _, c := range coords {
		c = mapping[c]
		grid[c[1]][c[0]] = 'O'
		if last[0] == c[0] {
			minY := min(c[1], last[1])
			maxY := max(c[1], last[1])
			for y := minY + 1; y < maxY; y++ {
				grid[y][c[0]] = '|'
			}
		} else {
			minX := min(c[0], last[0])
			maxX := max(c[0], last[0])
			for x := minX + 1; x < maxX; x++ {
				grid[c[1]][x] = '-'
			}
		}
		last = c
	}

	fillGrid(grid)

	// for _, r := range grid {
	// 	fmt.Println(string(r))
	// }

	maxA := 0
	for i := 0; i < len(coords)-1; i++ {
	outer:
		for j := i + 1; j < len(coords); j++ {
			c1 := coords[i]
			c2 := coords[j]

			m1 := mapping[c1]
			m2 := mapping[c2]

			xMin := min(m1[0], m2[0])
			xMax := max(m1[0], m2[0])
			yMin := min(m1[1], m2[1])
			yMax := max(m1[1], m2[1])

			d1 := c1[0] - c2[0]
			if d1 < 0 {
				d1 = -d1
			}
			d2 := c1[1] - c2[1]
			if d2 < 0 {
				d2 = -d2
			}
			a := (d1 + 1) * (d2 + 1)

			for x := xMin; x <= xMax; x++ {
				for y := yMin; y <= yMax; y++ {
					if grid[y][x] == '.' {
						continue outer
					}
				}
			}

			if a > maxA {
				maxA = a
			}
		}
	}

	return maxA
}

func fillGrid(grid [][]byte) {
	start := Coord{0, 0}
	for grid[start[1]][start[0]] != '-' {
		start[0]++
	}
	for grid[start[1]][start[0]] != '.' {
		start[1]++
	}

	queue := []Coord{start}
	visited := map[Coord]bool{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if visited[cur] {
			continue
		}
		grid[cur[1]][cur[0]] = 'o'
		visited[cur] = true

		for _, sum := range []Coord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nxt := cur.Add(sum)
			if nxt[0] < 0 || nxt[0] > len(grid[0])-1 || nxt[1] < 0 || nxt[1] > len(grid)-1 {
				continue
			}
			if grid[nxt[1]][nxt[0]] != '.' {
				continue
			}
			queue = append(queue, nxt)
		}
	}
}

func maxArea(in io.Reader) int {
	coords := []complex128{}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		bits := strings.Split(line, ",")
		x, err := strconv.Atoi(bits[0])
		Check(err)
		y, err := strconv.Atoi(bits[1])
		Check(err)
		coords = append(coords, complex(float64(x), float64(y)))
	}

	var maxA float64 = 0
	for i := 0; i < len(coords)-1; i++ {
		for j := i + 1; j < len(coords); j++ {
			c1 := coords[i]
			c2 := coords[j]
			d1 := real(c1) - real(c2)
			if d1 < 0 {
				d1 = -d1
			}
			d2 := imag(c1) - imag(c2)
			if d2 < 0 {
				d2 = -d2
			}
			a := (d1 + 1) * (d2 + 1)
			if a > maxA {
				maxA = a
			}
		}
	}

	return int(maxA)
}
