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

	count := 0

	scanner := bufio.NewScanner(f)
	c := q16.NewComputer()
	d := q16.NewComputer()
	var in1, in2, in3 int
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n")
		var w, x, y, z int
		if n, _ := fmt.Sscanf(line, "Before: [%d, %d, %d, %d]", &w, &x, &y, &z); n > 0 {
			c.SetRegisters(w, x, y, z)
		} else if n, _ := fmt.Sscanf(line, "After: [%d, %d, %d, %d]", &w, &x, &y, &z); n > 0 {
			d.SetRegisters(w, x, y, z)
			opCount := c.NumOps(d, in1, in2, in3)
			fmt.Println(c, d, in1, in2, in3, "-", opCount)
			if opCount > 2 {
				count++
			}
		} else {
			fmt.Sscanf(line, "%d %d %d %d", &w, &in1, &in2, &in3)
		}
	}
	fmt.Printf("count = %+v\n", count)
}
