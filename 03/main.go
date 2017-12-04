package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kieron-pivotal/advent2017/03/memory"
)

func main() {
	usage := fmt.Sprintf("%s <num>", os.Args[0])
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	num, _ := strconv.Atoi(os.Args[1])
	fmt.Println(memory.Distance(num))
}
