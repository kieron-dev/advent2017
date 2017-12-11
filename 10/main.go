package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/kieron-pivotal/advent2017/10/hash"
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

	instructionStrs := strings.Split(strings.TrimSpace(string(contents)), ",")
	instructions := []int{}
	for _, s := range instructionStrs {
		n, _ := strconv.Atoi(s)
		instructions = append(instructions, n)
	}

	fmt.Println("Part1:", hash.Compute(instructions, 256))
}
