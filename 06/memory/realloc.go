package memory

import (
	"fmt"
	"strings"
)

func ReallocFirstCyclePos(mem []int) int {
	visited := map[string]bool{}

	i := 0
	for {
		key := getKey(mem)
		_, ok := visited[key]
		if ok {
			return i
		}
		visited[key] = true
		i++
		realloc(mem)
	}
}

func getKey(mem []int) string {
	strArr := []string{}
	for _, n := range mem {
		strArr = append(strArr, fmt.Sprintf("%d", n))
	}
	return strings.Join(strArr, "-")
}

func realloc(mem []int) {
	max := 0
	maxPos := -1

	for i, n := range mem {
		if n > max {
			max = n
			maxPos = i
		}
	}

	l := len(mem)
	mem[maxPos] = 0
	pos := maxPos
	for max > 0 {
		pos = (pos + 1) % l
		mem[pos]++
		max--
	}
}
