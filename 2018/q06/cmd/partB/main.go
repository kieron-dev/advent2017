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

	count := 0
	limit := 10000

	for y := bottom; y <= top; y++ {
		for x := left; x <= right; x++ {
			coord := Coord{x: x, y: y}
			total := 0
			for _, c := range coords {
				total += coord.Distance(c)
			}
			if total < limit {
				count++
			}
		}
	}

	fmt.Printf("area < %d is %d", limit, count)

}
