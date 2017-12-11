package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kieron-pivotal/advent2017/09/brackets"
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

	fmt.Println("Part1:", brackets.Count(string(contents)))
	fmt.Println("Part2:", brackets.Garbage(string(contents)))
}
