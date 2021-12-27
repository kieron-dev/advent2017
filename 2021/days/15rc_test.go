package days_test

import (
	"bytes"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("15rc", func() {
	It("russ cox's version", func() {
		data, err := os.ReadFile("input15")
		Expect(err).NotTo(HaveOccurred())
		grid := bytes.Fields(data)

		repl := func(c, r byte) byte {
			return (c-'1'+r)%9 + '1'
		}

		for i := range grid {
			old := grid[i]
			for r := byte(1); r < 5; r++ {
				for _, c := range old {
					grid[i] = append(grid[i], repl(r, c))
				}
			}
		}

		old := grid
		for i := byte(1); i < 5; i++ {
			for _, row := range old {
				ext := []byte{}
				for _, c := range row {
					ext = append(ext, repl(c, i))
				}
				grid = append(grid, ext)
			}
		}

		n := len(grid)
		dist := make([][]int, n)
		for i := range dist {
			dist[i] = make([]int, n)
			for j := range dist[i] {
				dist[i][j] = 1e9
			}
		}

		type point struct{ x, y int }
		work := make([][]point, 20*n)

		add := func(p point, d int) {
			if p.x < 0 || p.y < 0 || p.x >= n || p.y >= n {
				return
			}
			d += int(grid[p.x][p.y] - '0')
			if dist[p.x][p.y] <= d {
				return
			}
			dist[p.x][p.y] = d
			work[d] = append(work[d], p)
		}

		add(point{0, 0}, -int(grid[0][0]-'0'))

		visit := func(p point) bool {
			d := dist[p.x][p.y]
			if p.x == n-1 && p.y == n-1 {
				return true
			}
			add(point{p.x - 1, p.y}, d)
			add(point{p.x + 1, p.y}, d)
			add(point{p.x, p.y - 1}, d)
			add(point{p.x, p.y + 1}, d)

			return false
		}

	outer:
		for _, w := range work {
			for _, p := range w {
				if visit(p) {
					break outer
				}
			}
		}

		Expect(dist[n-1][n-1]).To(Equal(2814))
	})
})
