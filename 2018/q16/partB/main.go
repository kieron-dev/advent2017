package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kieron-pivotal/advent2017/2018/q16"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	possibleOps := map[int][]int{}

	scanner := bufio.NewScanner(f)
	c := q16.NewComputer()
	d := q16.NewComputer()
	var op, in1, in2, in3 int
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n")
		var w, x, y, z int
		if n, _ := fmt.Sscanf(line, "Before: [%d, %d, %d, %d]", &w, &x, &y, &z); n > 0 {
			c.SetRegisters(w, x, y, z)
		} else if n, _ := fmt.Sscanf(line, "After: [%d, %d, %d, %d]", &w, &x, &y, &z); n > 0 {
			d.SetRegisters(w, x, y, z)
			ops := c.MatchingOps(d, in1, in2, in3)
			possibleOps[op] = ops
		} else {
			fmt.Sscanf(line, "%d %d %d %d", &op, &in1, &in2, &in3)
		}
	}
	out := map[int]int{}
	reduce(possibleOps, out)

	solveProg(out)
}

func solveProg(out map[int]int) {
	f, err := os.Open("program")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	c := q16.NewComputer()

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n")
		var op, x, y, z int
		fmt.Sscanf(line, "%d %d %d %d", &op, &x, &y, &z)
		realOp := out[op]
		c.Ops()[realOp](x, y, z)
	}
	fmt.Printf("c = %+v\n", c)
}

func reduce(possibleOps map[int][]int, out map[int]int) {
	for len(out) != len(possibleOps) {
		for op, ops := range possibleOps {
			if len(ops) == 1 {
				out[op] = ops[0]
				for o, l := range possibleOps {
					if o == op {
						continue
					}
					remaining := []int{}
					for _, a := range l {
						if a != ops[0] {
							remaining = append(remaining, a)
						}
					}
					possibleOps[o] = remaining
				}
			}
		}
	}
}
