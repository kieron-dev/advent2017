package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kieron-pivotal/advent2017/16/dance"
)

func main() {
	usage := fmt.Sprintf("%s <inputPath>", os.Args[0])
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(os.Args[1])
	contents := strings.TrimSpace(string(bytes))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	d := dance.New(16)
	moves := strings.Split(contents, ",")
	d.ProcessMoves(moves)

	fmt.Println("Part1:", d.String())

	i := 1
	for !d.IsOriginalOrder() {
		d.ProcessMoves(moves)
		i++
	}
	fmt.Println("got back to ID in", i, "moves")

	d = dance.New(16)
	aBillion := 1e9 % i
	for i := 0; i < aBillion; i++ {
		d.ProcessMoves(moves)
	}
	fmt.Println("Part2:", d.String())
}
