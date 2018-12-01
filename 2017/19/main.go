package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kieron-pivotal/advent2017/19/route"
)

func main() {
	usage := fmt.Sprintf("%s <inputPath>", os.Args[0])
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	m := route.New(lines)
	m.Walk()

	fmt.Println("Part1:", m.GetLetters())
	fmt.Println("Part2:", m.GetStepCount())

}
