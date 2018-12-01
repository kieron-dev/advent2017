package passphrase

import (
	"sort"
	"strings"
)

func Check(lines []string) int {
	sum := 0
	for _, line := range lines {
		used := map[string]bool{}
		words := strings.Split(line, " ")
		ok := true
		for _, word := range words {

			word = sortStr(word)
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

func sortStr(str string) string {
	ints := []int{}
	for _, r := range str {
		ints = append(ints, int(r-'a'))
	}

	sort.Ints(ints)
	res := ""
	for _, i := range ints {
		res += string('a' + i)
	}
	return res
}
