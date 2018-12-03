package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	codes := []string{}

	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}

	for i := 0; i < len(codes); i++ {
		for j := i + 1; j < len(codes); j++ {
			if diffByOne(codes[i], codes[j]) {
				fmt.Println(codes[i])
				fmt.Println(codes[j])
				break
			}
		}
	}
}

func diffByOne(a, b string) bool {
	bRunes := []rune(b)
	diff := 0
	for i, r := range a {
		if bRunes[i] == r {
			continue
		}
		diff++
		if diff > 1 {
			return false
		}
	}
	return diff == 1
}
