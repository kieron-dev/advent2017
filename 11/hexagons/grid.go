package hexagons

import "fmt"

func Distance(steps []string) int {
	fmt.Println(steps)
	stepMap := map[string]int{}
	for _, s := range steps {
		stepMap[s]++
	}

	// NE+NW == N
	upLeftRight := stepMap["nw"]
	if stepMap["ne"] < upLeftRight {
		upLeftRight = stepMap["ne"]
	}
	stepMap["ne"] -= upLeftRight
	stepMap["nw"] -= upLeftRight
	stepMap["n"] += upLeftRight

	// SE+SW == S
	downLeftRight := stepMap["sw"]
	if stepMap["se"] < downLeftRight {
		downLeftRight = stepMap["se"]
	}
	stepMap["se"] -= downLeftRight
	stepMap["sw"] -= downLeftRight
	stepMap["s"] += downLeftRight

	// S + N = 0
	upDown := stepMap["s"]
	if stepMap["n"] < upDown {
		upDown = stepMap["n"]
	}
	stepMap["n"] -= upDown
	stepMap["s"] -= upDown

	fmt.Println(stepMap)
	sum := 0
	for _, v := range stepMap {
		sum += v
	}
	return sum
}
