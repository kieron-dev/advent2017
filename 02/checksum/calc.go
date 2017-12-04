package checksum

func RowVal(input []int) int {
	for i, n := range input {
		for j := i + 1; j < len(input); j++ {
			mn, mx := n, input[j]
			if mn > mx {
				mn, mx = mx, mn
			}
			if mx%mn == 0 {
				return mx / mn
			}
		}
	}
	return 0
}

func Calc(input [][]int) int {
	sum := 0
	for _, r := range input {
		sum += RowVal(r)
	}
	return sum
}
