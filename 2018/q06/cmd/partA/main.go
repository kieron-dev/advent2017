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
	x := from.x - c.y
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

	left := 10000000
	bottom := 10000000
	right := -10000000
	top := -10000000

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

	fmt.Println(coords)
	fmt.Printf("left = %+v\n", left)
	fmt.Printf("right = %+v\n", right)
	fmt.Printf("top = %+v\n", top)
	fmt.Printf("bottom = %+v\n", bottom)

	grid := map[Coord]int{}
	for x := left; x <= right; x++ {
		for y := bottom; y <= top; y++ {
			coord := Coord{x: x, y: y}
			minD := 1000000000
			minI := -1
			for i, c := range coords {
				d := coord.Distance(c)
				if d < minD {
					minD = d
					minI = i
				}
			}
			grid[coord] = minI
		}
	}
	fmt.Printf("grid = %+v\n", grid)
}
