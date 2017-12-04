package memory

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
