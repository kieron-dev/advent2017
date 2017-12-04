package captcha

import "strconv"

func Decode(input string) int {
	if len(input) == 0 {
		return 0
	}

	sum := 0

	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			val, _ := strconv.Atoi(string(input[i]))
			sum += val
		}
	}
	if input[len(input)-1] == input[0] {
		val, _ := strconv.Atoi(string(input[0]))
		sum += val
	}
	return sum
}
