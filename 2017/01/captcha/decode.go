package captcha

import "strconv"

func Decode(input string) int {
	if len(input) == 0 {
		return 0
	}

	sum := 0
	l := len(input)

	for i := 0; i < l; i++ {
		if input[i] == input[(i+l/2)%l] {
			val, _ := strconv.Atoi(string(input[i]))
			sum += val
		}
	}
	return sum
}
