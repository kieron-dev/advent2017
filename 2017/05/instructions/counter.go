package instructions

func Count(steps []int) int {
	pos := 0
	l := len(steps)
	c := 0

	for pos >= 0 && pos < l {
		cur := steps[pos]
		if cur > 2 {
			steps[pos]--
		} else {
			steps[pos]++
		}
		pos += cur
		c++
	}
	return c
}
