package two022_test

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Coord struct {
	X, Y int
}

func NewCoord(x, y int) Coord {
	return Coord{X: x, Y: y}
}

var dirs = map[string]Coord{
	"U": {X: 0, Y: 1},
	"D": {X: 0, Y: -1},
	"L": {X: -1, Y: 0},
	"R": {X: 1, Y: 0},
}

func (c Coord) Mult(n int) Coord {
	return Coord{
		X: c.X * n,
		Y: c.Y * n,
	}
}

func (c Coord) Add(d Coord) Coord {
	return Coord{
		X: c.X + d.X,
		Y: c.Y + d.Y,
	}
}

func (c Coord) Move(dir string, moves int) Coord {
	return c.Add(dirs[dir].Mult(moves))
}

var _ = Describe("09", func() {
	It("does part A", func() {
		f, err := os.Open("input09")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		var posH, posT Coord
		visited := map[Coord]bool{{}: true}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			var dir string
			var moves int
			_, err := fmt.Sscanf(line, "%s %d", &dir, &moves)
			Expect(err).NotTo(HaveOccurred())

			for i := 0; i < moves; i++ {
				posH = posH.Move(dir, 1)
				posT = moveTail(posH, posT)
				visited[posT] = true
			}

			// fmt.Printf("%s%d: %v %v %v\n", dir, moves, posH, posT, visited)
		}

		Expect(len(visited)).To(Equal(6337))
	})

	It("does part B", func() {
		f, err := os.Open("input09")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		var knots [10]Coord
		visited := map[Coord]bool{{}: true}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			var dir string
			var moves int
			_, err := fmt.Sscanf(line, "%s %d", &dir, &moves)
			Expect(err).NotTo(HaveOccurred())

			for i := 0; i < moves; i++ {
				knots[0] = knots[0].Move(dir, 1)
				for j := 1; j < 10; j++ {
					npos := moveTail(knots[j-1], knots[j])
					if npos == knots[j] {
						break
					}
					knots[j] = npos
				}
				visited[knots[9]] = true
			}
		}

		Expect(len(visited)).To(Equal(2455))
	})
})

func moveTail(h, t Coord) Coord {
	if touching(h, t) {
		return t
	}

	if h.X == t.X {
		y := t.Y + 1
		if h.Y < t.Y {
			y = t.Y - 1
		}
		return NewCoord(t.X, y)
	}

	if h.Y == t.Y {
		x := t.X + 1
		if h.X < t.X {
			x = t.X - 1
		}
		return NewCoord(x, t.Y)
	}

	x := t.X + 1
	if h.X < t.X {
		x = t.X - 1
	}
	y := t.Y + 1
	if h.Y < t.Y {
		y = t.Y - 1
	}

	return NewCoord(x, y)
}

func touching(h, t Coord) bool {
	if h.X == t.X && h.Y == t.Y {
		return true
	}

	if h.X == t.X && mod(h.Y-t.Y) == 1 {
		return true
	}

	if h.Y == t.Y && mod(h.X-t.X) == 1 {
		return true
	}

	if mod(h.X-t.X) == 1 && mod(h.Y-t.Y) == 1 {
		return true
	}

	return false
}

func mod(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
