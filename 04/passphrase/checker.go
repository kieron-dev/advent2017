package passphrase

import "strings"

func Check(lines []string) int {
	sum := 0
	for _, line := range lines {
		used := map[string]bool{}
		words := strings.Split(line, " ")
		ok := true
		for _, word := range words {
			if used[word] {
				ok = false
				break
			}
			used[word] = true
		}
		if ok {
			sum++
		}
	}
	return sum
}
