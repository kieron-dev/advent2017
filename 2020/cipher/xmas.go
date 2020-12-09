package cipher

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type Xmas struct {
	data []int
}

func NewXmas() Xmas {
	return Xmas{}
}

func (x *Xmas) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("failed to get int from %q", line)
		}

		x.data = append(x.data, n)
	}
}

func (x Xmas) FirstError(size int) int {
	numList := x.data[:size]
	numMap := map[int]int{}
	for _, n := range numList {
		numMap[n]++
	}

	for i := size; i < len(x.data); i++ {
		if !sumOf(x.data[i], numList, numMap) {
			return x.data[i]
		}

		toRemove := numList[0]
		numList = append(numList[1:], x.data[i])
		numMap[toRemove]--
		numMap[x.data[i]]++
	}

	return 0
}

func sumOf(n int, numList []int, numMap map[int]int) bool {
	for _, p := range numList {
		if numMap[n-p] > 0 && p != n-p {
			return true
		}
	}

	return false
}

func (x Xmas) WeaknessNums(size int) []int {
	sum := x.FirstError(size)

	work := 0
	first := 0

	for i := 0; i < len(x.data); i++ {
		work += x.data[i]

		for work > sum {
			work -= x.data[first]
			first++
		}

		if work == sum {
			return x.data[first : i+1]
		}
	}

	return nil
}

func (x Xmas) EncryptionWeakness(size int) int {
	numList := x.WeaknessNums(size)

	min := numList[0]
	max := numList[0]

	for i := 1; i < len(numList); i++ {
		n := numList[i]
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	return min + max
}
