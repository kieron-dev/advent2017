package two022_test

import (
	"bytes"
	"io/ioutil"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("12", func() {
	It("does part A", func() {
		bs, err := ioutil.ReadFile("input12")
		Expect(err).NotTo(HaveOccurred())

		lines := bytes.Fields(bs)

		var startR, startC int

		var row []byte
	outer:
		for startR, row = range lines {
			var b byte
			for startC, b = range row {
				if b == 'S' {
					break outer
				}
			}
		}

		queue := []Coord{NewCoord(startC, startR)}
		distances := map[Coord]int{}
		visited := map[Coord]bool{}

		var steps int
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			if visited[cur] {
				continue
			}

			if lines[cur.Y][cur.X] == 'E' {
				steps = distances[cur]
				break
			}

			for _, d := range []string{"U", "D", "L", "R"} {
				next := cur.Move(d, 1)

				if next.X < 0 || next.X > len(lines[0])-1 || next.Y < 0 || next.Y > len(lines)-1 {
					continue
				}

				if visited[next] {
					continue
				}

				curEl := lines[cur.Y][cur.X]
				if lines[cur.Y][cur.X] == 'S' {
					curEl = 'a'
				}

				nextEl := lines[next.Y][next.X]
				if lines[next.Y][next.X] == 'E' {
					nextEl = 'z'
				}

				if int(nextEl)-int(curEl) > 1 {
					continue
				}

				distances[next] = distances[cur] + 1
				queue = append(queue, next)
			}

			visited[cur] = true

		}

		Expect(steps).To(Equal(456))
	})

	It("does part B", func() {
		bs, err := ioutil.ReadFile("input12")
		Expect(err).NotTo(HaveOccurred())

		lines := bytes.Fields(bs)

		var startR, startC int

		var row []byte
	outer:
		for startR, row = range lines {
			var b byte
			for startC, b = range row {
				if b == 'E' {
					break outer
				}
			}
		}

		queue := []Coord{NewCoord(startC, startR)}
		distances := map[Coord]int{}
		visited := map[Coord]bool{}

		var steps int
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			if visited[cur] {
				continue
			}

			if lines[cur.Y][cur.X] == 'a' {
				steps = distances[cur]
				break
			}

			for _, d := range []string{"U", "D", "L", "R"} {
				next := cur.Move(d, 1)

				if next.X < 0 || next.X > len(lines[0])-1 || next.Y < 0 || next.Y > len(lines)-1 {
					continue
				}

				if visited[next] {
					continue
				}

				curEl := lines[cur.Y][cur.X]
				if lines[cur.Y][cur.X] == 'E' {
					curEl = 'z'
				}

				nextEl := lines[next.Y][next.X]
				if lines[next.Y][next.X] == 'S' {
					nextEl = 'a'
				}

				if int(curEl)-int(nextEl) > 1 {
					continue
				}

				distances[next] = distances[cur] + 1
				queue = append(queue, next)
			}

			visited[cur] = true

		}

		Expect(steps).To(Equal(454))
	})
})
