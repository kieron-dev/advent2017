package twentytwentyfive_test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type shape []string

func (s shape) size() int {
	size := 0
	for _, l := range s {
		size += strings.Count(l, "#")
	}

	return size
}

type instruction struct {
	rows        int
	cols        int
	shapeCounts []int
}

func (i instruction) possible(sizes []int) bool {
	minSquares := 0
	total := 0
	for n, count := range i.shapeCounts {
		minSquares += count * sizes[n]
		total += count
	}
	if minSquares > i.rows*i.cols {
		return false
	}

	if total*9 <= i.rows*i.cols {
		return true
	}

	fmt.Println("damn", i.rows*i.cols, total*9)
	return true
}

func Test12a(t *testing.T) {
	in, err := os.Open("input12")
	Check(err)

	var shapes []shape
	var sizes []int
	var instructions []instruction
	var work shape

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		switch {
		case strings.Contains(line, "x"):
			left, right, found := strings.Cut(line, ":")
			assert.True(t, found)
			bits := strings.Split(left, "x")
			rows, err := strconv.Atoi(bits[0])
			Check(err)
			cols, err := strconv.Atoi(bits[1])
			Check(err)
			right = strings.TrimSpace(right)
			var vals []int
			for _, str := range strings.Fields(right) {
				n, err := strconv.Atoi(str)
				Check(err)
				vals = append(vals, n)
			}
			instructions = append(instructions, instruction{
				rows:        rows,
				cols:        cols,
				shapeCounts: vals,
			})
		case strings.Contains(line, ":"):
			work = shape{}
		case strings.Contains(line, "#"):
			work = append(work, line)
		case line == "":
			shapes = append(shapes, work)
			sizes = append(sizes, work.size())
		}
	}

	ans := 0
	for n, i := range instructions {
		fmt.Println(n)
		if i.possible(sizes) {
			ans++
		}
	}

	assert.Equal(t, -1, ans)
}
