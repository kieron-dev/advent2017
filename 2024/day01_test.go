package twenty24

import (
	"bytes"
	"os"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getInput01(t *testing.T) ([]int, []int) {
	bs, err := os.ReadFile("input01")
	assert.NoError(t, err)

	fs := bytes.Fields(bs)

	var left, right []int
	for i, f := range fs {
		n := bsToI(t, f)
		if i%2 == 0 {
			left = append(left, n)
		} else {
			right = append(right, n)
		}
	}

	return left, right
}

func Test01A(t *testing.T) {
	left, right := getInput01(t)

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := range left {
		sum += abs(left[i] - right[i])
	}

	assert.Equal(t, 1603498, sum)
}

func Test01B(t *testing.T) {
	left, right := getInput01(t)

	rightCounts := map[int]int{}
	for _, n := range right {
		rightCounts[n]++
	}

	sum := 0
	for _, n := range left {
		sum += n * rightCounts[n]
	}

	assert.Equal(t, 25574739, sum)
}

func bsToI(t *testing.T, in []byte) int {
	i, err := strconv.Atoi(string(in))
	assert.NoError(t, err)
	return i
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
