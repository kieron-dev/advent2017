// Package jolt gives info about Jolt adapter combinations
package jolt

import (
	"bufio"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Organiser struct {
	adapters []int
}

func NewOrganiser() Organiser {
	return Organiser{}
}

func (o *Organiser) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("failed to convert %q to int", line)
		}

		o.adapters = append(o.adapters, n)
	}

	sort.Ints(o.adapters)
}

func (o Organiser) GetDiffs() map[int]int {
	adapters := o.AdaptersWithStartEnd()
	diffs := map[int]int{}

	for i := 0; i < len(adapters)-1; i++ {
		diffs[adapters[i+1]-adapters[i]]++
	}

	return diffs
}

func (o Organiser) AdaptersWithStartEnd() []int {
	adapters := append([]int{0}, o.adapters...)
	return append(adapters, adapters[len(adapters)-1]+3)
}

func (o Organiser) Combinations() int {
	adapters := o.AdaptersWithStartEnd()

	// dp[n] = combinations starting at pos n
	dp := map[int]int{}

	dp[len(adapters)-1] = 1

	for i := len(adapters) - 2; i >= 0; i-- {
		for j := 1; i+j < len(adapters) && adapters[i+j]-adapters[i] < 4; j++ {
			dp[i] += dp[i+j]
		}
	}

	return dp[0]
}
