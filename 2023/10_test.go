package two023_test

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type coord [2]int

func (c coord) add(d coord) coord {
	return coord{c[0] + d[0], c[1] + d[1]}
}

func (c coord) mult(n int) coord {
	return coord{c[0] * n, c[1] * n}
}

func (c coord) dist(d coord) int {
	return absDiff(c[0], d[0]) + absDiff(c[1], d[1])
}

func absDiff(a, b int) int {
	d := a - b
	if d < 0 {
		d = -d
	}
	return d
}

type grid struct {
	rows     []string
	start    coord
	loop     map[coord]bool
	origRows int
	origCols int
}

func (g *grid) storeRows(fileName string) {
	f, err := os.Open(fileName)
	Expect(err).NotTo(HaveOccurred())
	defer f.Close()

	g.rows = []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			g.origRows++
			g.origCols = len(line)
			var entry strings.Builder
			entry.WriteString("x")
			for i := range line {
				entry.WriteByte(line[i])
				entry.WriteString("x")
			}
			g.rows = append(g.rows, strings.Repeat("x", entry.Len()))
			g.rows = append(g.rows, entry.String())
		}
	}
	g.rows = append(g.rows, strings.Repeat("x", len(g.rows[0])))
}

func (g *grid) storeStart() {
	for r, line := range g.rows {
		c := strings.Index(line, "S")
		if c > -1 {
			g.start = coord{r, c}
			return
		}
	}
}

func (g *grid) storeLoop() {
	c := g.start
	g.loop = map[coord]bool{}

	for {
		g.loop[c] = true
		n, err := g.next(c)
		if err != nil {
			return
		}
		c = n
	}
}

func (g grid) furthestDist() int {
	l := len(g.loop)
	if l%2 == 1 {
		l++
	}
	return l / 2
}

func (g grid) isValid(c coord) bool {
	return c[0] >= 0 && c[1] >= 0 && c[0] < len(g.rows) && c[1] < len(g.rows[0])
}

func (g grid) val(c coord) byte {
	return g.rows[c[0]][c[1]]
}

func (g grid) next(from coord) (coord, error) {
	above := from.add(coord{-2, 0})
	below := from.add(coord{2, 0})
	left := from.add(coord{0, -2})
	right := from.add(coord{0, 2})

	switch g.val(from) {
	case 'S':
		if g.isValid(above) &&
			(g.val(above) == '|' || g.val(above) == 'F' || g.val(above) == '7') {
			return above, nil
		}
		if g.isValid(below) &&
			(g.val(below) == '|' || g.val(below) == 'L' || g.val(below) == 'J') {
			return below, nil
		}
		if g.isValid(left) &&
			(g.val(left) == '-' || g.val(left) == 'L' || g.val(left) == 'F') {
			return left, nil
		}
		if g.isValid(right) &&
			(g.val(right) == '-' || g.val(right) == '7' || g.val(right) == 'J') {
			return right, nil
		}
	case '|':
		if !g.loop[above] {
			return above, nil
		}
		if !g.loop[below] {
			return below, nil
		}
	case '-':
		if !g.loop[right] {
			return right, nil
		}
		if !g.loop[left] {
			return left, nil
		}
	case 'F':
		if !g.loop[right] {
			return right, nil
		}
		if !g.loop[below] {
			return below, nil
		}
	case 'J':
		if !g.loop[left] {
			return left, nil
		}
		if !g.loop[above] {
			return above, nil
		}
	case '7':
		if !g.loop[left] {
			return left, nil
		}
		if !g.loop[below] {
			return below, nil
		}
	case 'L':
		if !g.loop[above] {
			return above, nil
		}
		if !g.loop[right] {
			return right, nil
		}
	}

	return coord{}, errors.New("the end")
}

func (g grid) insideLoop() int {
	outsideCount := 0
	visited := map[coord]bool{}
	q := []coord{{0, 0}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if visited[cur] {
			continue
		}

		for _, n := range []coord{
			cur.add(coord{0, -1}),
			cur.add(coord{0, 1}),
			cur.add(coord{-1, 0}),
			cur.add(coord{1, 0}),
		} {
			if !g.isValid(n) || visited[n] || g.inLoop(n) {
				continue
			}

			q = append(q, n)
		}

		visited[cur] = true
		if cur[0]%2 == 1 && cur[1]%2 == 1 {
			outsideCount++
		}
	}

	return g.origRows*g.origCols - len(g.loop) - outsideCount
}

func (g grid) inLoop(c coord) bool {
	if g.loop[c] {
		return true
	}

	if g.val(c) == 'x' {
		left := c.add(coord{0, -1})
		right := c.add(coord{0, 1})
		top := c.add(coord{-1, 0})
		bottom := c.add(coord{1, 0})

		if g.loop[left] && g.loop[right] &&
			(g.val(left) == '-' || g.val(left) == 'L' || g.val(left) == 'F' || g.val(left) == 'S') &&
			(g.val(right) == '-' || g.val(right) == 'J' || g.val(right) == '7' || g.val(right) == 'S') {
			return true
		}

		if g.loop[top] && g.loop[bottom] &&
			(g.val(top) == '|' || g.val(top) == '7' || g.val(top) == 'F' || g.val(top) == 'S') &&
			(g.val(bottom) == '|' || g.val(bottom) == 'L' || g.val(bottom) == 'J' || g.val(bottom) == 'S') {
			return true
		}
	}

	return false
}

func newGrid(fileName string) grid {
	g := grid{}

	g.storeRows(fileName)
	g.storeStart()
	g.storeLoop()

	return g
}

func (g grid) print() {
	fmt.Println()
	for _, line := range g.rows {
		fmt.Println(line)
	}
	fmt.Println()
}

var _ = Describe("10", func() {
	It("does part A", func() {
		grid := newGrid("input10")
		Expect(grid.furthestDist()).To(Equal(6942))
	})

	It("does part B", func() {
		grid := newGrid("input10")
		Expect(grid.insideLoop()).To(Equal(297))
	})
})
