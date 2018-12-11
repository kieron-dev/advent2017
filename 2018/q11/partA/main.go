package main

import (
	"fmt"

	"github.com/kieron-pivotal/advent2017/2018/q11"
)

func main() {
	grid := q11.NewGrid(5468, 300, 300)
	r, c := grid.Largest3x3Cell()
	fmt.Printf("Largest Cell: (%d, %d)\n", c, r)

	var p, n int
	p, r, c, n = grid.LargestCell()
	fmt.Printf("Largest NxN Cell: (%d, %d), Size %d (Power %d)\n", c, r, n, p)
}
