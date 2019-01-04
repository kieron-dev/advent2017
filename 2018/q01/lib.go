package q01

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type Calibrator struct {
	nums []int
}

func NewCalibrator(in io.Reader) *Calibrator {
	c := Calibrator{}
	c.LoadFile(in)

	return &c
}

func (c *Calibrator) LoadFile(in io.Reader) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, "\n")
		c.nums = append(c.nums, atoi(line))
	}
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func (c *Calibrator) Add() int {
	s := 0
	for _, n := range c.nums {
		s += n
	}
	return s
}

func (c *Calibrator) FirstRepeat() int {
	visited := map[int]bool{}
	s := 0
	for {
		for _, n := range c.nums {
			s += n
			if visited[s] {
				return s
			}
			visited[s] = true
		}
	}
}
