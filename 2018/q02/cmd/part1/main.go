package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	twos := 0
	threes := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		code := scanner.Text()
		if hasNum(code, 2) {
			twos++
		}
		if hasNum(code, 3) {
			threes++
		}
	}
	fmt.Println(twos * threes)
}

func hasNum(code string, num int) bool {
	letters := map[rune]int{}
	for _, r := range code {
		letters[r]++
	}
	for _, c := range letters {
		if c == num {
			return true
		}
	}
	return false
}
