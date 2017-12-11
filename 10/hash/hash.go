package hash

func Get(list []int, pos int) int {
	l := len(list)
	return list[pos%l]
}

func Set(list []int, pos, val int) {
	list[pos%len(list)] = val
}

func Reverse(list []int, start, end int) {
	for end > start {
		sVal := Get(list, start)
		eVal := Get(list, end)
		Set(list, end, sVal)
		Set(list, start, eVal)
		start++
		end--
	}
}

func Compute(instructions []int, listLen int) int {
	list := []int{}
	for i := 0; i < listLen; i++ {
		list = append(list, i)
	}

	pos := 0
	skip := 0
	for _, n := range instructions {
		Reverse(list, pos, pos+n-1)
		pos += n + skip
		skip++
	}
	return list[0] * list[1]
}
