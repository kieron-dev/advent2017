package q01

import (
	"bufio"
	"io"
	"log"
	"strconv"
)

func SolveA(in io.Reader) int {
	freq := 0
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		freq += num
	}

	return freq
}

func SolveB(in io.Reader) int {
	nums := []int{}

	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	freq := 0
	freqs := map[int]bool{}

	for {
		for _, n := range nums {
			freq += n
			if freqs[freq] {
				return freq
			}
			freqs[freq] = true
		}
	}
}
