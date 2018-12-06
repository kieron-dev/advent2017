package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	x int
	y int
}

func (c Coord) Distance(from Coord) int {
	x := from.x - c.x
	if x < 0 {
		x = -x
	}
	y := from.y - c.y
	if y < 0 {
		y = -y
	}
	return x + y
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	coords := []Coord{}
	const largeNum = 10000000

	left := largeNum
	bottom := largeNum
	right := -largeNum
	top := -largeNum

	for scanner.Scan() {
		line := scanner.Text()
		var x, y int
		fmt.Sscanf(line, "%d, %d", &x, &y)
		coord := Coord{x: x, y: y}
		coords = append(coords, coord)
		if x < left {
			left = x
		}
		if x > right {
			right = x
		}
		if y < bottom {
			bottom = y
		}
		if y > top {
			top = y
		}
	}

	counts := map[int]int{}

	grid := map[Coord]int{}
	for y := bottom; y <= top; y++ {
		for x := left; x <= right; x++ {
			coord := Coord{x: x, y: y}
			minD := largeNum
			var minI int
			var minCount int
			for i, c := range coords {
				d := coord.Distance(c)
				if d < minD {
					minCount = 1
					minD = d
					minI = i
				} else if d == minD {
					minCount++
				}
			}
			if minCount == 1 {
				grid[coord] = minI
			} else {
				grid[coord] = -1
			}
			counts[grid[coord]]++
		}
	}

	edgeVals := map[int]bool{-1: true}
	for y := top; y <= bottom; y++ {
		edgeVals[grid[Coord{x: left, y: y}]] = true
		edgeVals[grid[Coord{x: right, y: y}]] = true
	}
	for x := left; x <= right; x++ {
		edgeVals[grid[Coord{x: x, y: top}]] = true
		edgeVals[grid[Coord{x: x, y: bottom}]] = true
	}

	max := 0
	for i, c := range counts {
		if edgeVals[i] {
			continue
		}
		if c > max {
			max = c
		}
	}
	fmt.Printf("Largest non-infinite area is %d\n", max)

}
