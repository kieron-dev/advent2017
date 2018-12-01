package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kieron-pivotal/advent2017/11/hexagons"
)

func main() {
	usage := fmt.Sprintf("%s <inputPath>", os.Args[0])
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	contents, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	instructions := strings.Split(strings.TrimSpace(string(contents)), ",")

	fmt.Println("Part1:", hexagons.Distance(instructions))
	fmt.Println("Part2:", hexagons.Furthest(instructions))
}
