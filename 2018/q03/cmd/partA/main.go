package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	moreThan1 := 0
	used := map[string]int{}

	fillRect := func(x, y, w, h int) {
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				key := fmt.Sprintf("%d,%d", i, j)
				if used[key] == 1 {
					moreThan1++
				}
				used[key]++
			}
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		var n, x, y, w, h int
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &n, &x, &y, &w, &h)
		fillRect(x, y, w, h)
	}

	fmt.Println(moreThan1)
}
