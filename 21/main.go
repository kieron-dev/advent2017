package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/kieron-pivotal/advent2017/21/patterns"
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

	re := regexp.MustCompile("(.*) => (.*)")

	art := patterns.New()

	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		art.AddRule(match[1], match[2])
	}

	for i := 0; i < 5; i++ {
		art.Advance()
	}

	fmt.Println("Part1:", art.OnCount())

	for i := 5; i < 18; i++ {
		art.Advance()
	}
	fmt.Println("Part2:", art.OnCount())
}
