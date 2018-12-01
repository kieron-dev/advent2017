package memory

type Coord [2]int

func Distance(num int) int {
	min := 0
	max := 0
	inc := 8

	n := 1

	for n < num {
		n += inc

		min += 1
		max += 2
		inc += 8
	}

	// fmt.Println(min, max, n)

	d := -1
	dist := max
	for n != num {
		n--
		dist += d
		if dist == min {
			d = 1
		} else if dist == max {
			d = -1
		}
	}

	return dist
}

func WeirdSum(lowerBound int) int {
	grid := map[Coord]int{Coord{0, 0}: 1}
	dirs := []Coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	curDir := 0
	pos := Coord{0, 0}

	nextVal := 1
	radius := 1
	for nextVal <= lowerBound {
		pos = Coord{pos[0] + dirs[curDir][0], pos[1] + dirs[curDir][1]}
		nextVal = 0
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if i == 0 && j == 0 {
					continue
				}
				nextVal += grid[Coord{pos[0] + i, pos[1] + j}]
			}
		}
		grid[pos] = nextVal
		if (pos == Coord{radius, -radius + 1} || pos == Coord{radius, radius} ||
			pos == Coord{-radius, radius} || pos == Coord{-radius, -radius}) {
			curDir = (curDir + 1) % 4
			if curDir == 0 {
				radius++
			}
		}
	}
	return nextVal
}
