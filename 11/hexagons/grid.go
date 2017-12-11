package hexagons

func Distance(steps []string) int {
	stepMap := map[string]int{}
	for _, s := range steps {
		stepMap[s]++
	}

	x, y := 0, 0

	y += 2 * stepMap["n"]
	y -= 2 * stepMap["s"]

	y += stepMap["nw"] + stepMap["ne"]
	y -= stepMap["sw"] + stepMap["se"]
	x += stepMap["ne"] + stepMap["se"]
	x -= stepMap["nw"] + stepMap["sw"]

	count := x
	if count < 0 {
		count = -x
	}
	if y < 0 {
		y = -y
	}
	if y > count {
		return count + (y-count)/2
	}
	return count

}

func Furthest(steps []string) int {
	furthest := 0
	for i := 1; i <= len(steps); i++ {
		dist := Distance(steps[:i])
		if dist > furthest {
			furthest = dist
		}
	}
	return furthest
}
