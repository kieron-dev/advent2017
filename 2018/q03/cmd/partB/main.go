package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	used := map[string]int{}

	fillRect := func(x, y, w, h int) {
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				key := fmt.Sprintf("%d,%d", i, j)
				used[key]++
			}
		}
	}

	checkRect := func(x, y, w, h int) bool {
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				key := fmt.Sprintf("%d,%d", i, j)
				if used[key] != 1 {
					return false
				}
			}
		}
		return true
	}

	for _, line := range lines {
		var n, x, y, w, h int
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &n, &x, &y, &w, &h)
		fillRect(x, y, w, h)
	}

	for _, line := range lines {
		var n, x, y, w, h int
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &n, &x, &y, &w, &h)
		if checkRect(x, y, w, h) {
			fmt.Println(n)
		}
	}
}
