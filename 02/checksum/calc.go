package checksum

func RowVal(input []int) int {
	mn := 1 << 32
	mx := 0
	for _, n := range input {
		if n < mn {
			mn = n
		}
		if n > mx {
			mx = n
		}
	}
	return mx - mn
}

func Calc(input [][]int) int {
	sum := 0
	for _, r := range input {
		sum += RowVal(r)
	}
	return sum
}
