package q11

import "fmt"

type Grid struct {
	serial int
	rows   int
	cols   int
	cells  [][]int
	dp     []map[string]int
}

func NewGrid(serial, rows, cols int) *Grid {
	g := Grid{serial: serial, rows: rows, cols: cols}
	g.cells = make([][]int, rows)
	for i := 0; i < rows; i++ {
		g.cells[i] = make([]int, cols)
	}

	g.dp = make([]map[string]int, rows+1)
	g.dp[1] = make(map[string]int, 300*300)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			g.cells[r][c] = g.CellPower(r, c)
			g.dp[1][fmt.Sprintf("%d,%d", r, c)] = g.cells[r][c]
		}
	}

	return &g
}

func (g *Grid) CellPower(row, col int) int {
	rackId := (col + 1) + 10
	p := rackId * (row + 1)
	p += g.serial
	p *= rackId
	p /= 100
	p = p % 10
	p -= 5
	return p
}

func (g *Grid) LargestCell() (p, r, c, n int) {
	maxN := 1
	maxPower, maxR, maxC := g.LargestNxNCell(1)

	for n := 2; n < g.rows+1; n++ {
		g.dp[n] = make(map[string]int, (300-n)*(300-n))
		for r := 0; r < g.rows-n; r++ {
			for c := 0; c < g.cols-n; c++ {
				rowcol := fmt.Sprintf("%d,%d", r, c)
				p := g.dp[n-1][rowcol]
				for i := 0; i < n; i++ {
					p += g.cells[r+i][c+n-1]
					p += g.cells[r+n-1][c+i]
				}
				p -= g.cells[r+n-1][c+n-1]
				if p > maxPower {
					maxPower = p
					maxR = r
					maxC = c
					maxN = n
				}
				g.dp[n][rowcol] = p
			}
		}
	}
	return maxPower, maxR + 1, maxC + 1, maxN
}

func (g *Grid) LargestNxNCell(n int) (p, r, c int) {
	maxPower := n*n*-9 - 1
	var maxR, maxC int
	for r := 0; r < g.rows-n; r++ {
		for c := 0; c < g.cols-n; c++ {
			total := 0
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					total += g.cells[r+i][c+j]
				}
			}
			if total > maxPower {
				maxPower = total
				maxR = r
				maxC = c
			}
		}
	}
	return maxPower, maxR + 1, maxC + 1
}

func (g *Grid) Largest3x3Cell() (r, c int) {
	_, r, c = g.LargestNxNCell(3)
	return r, c
}
