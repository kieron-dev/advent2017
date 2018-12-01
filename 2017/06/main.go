package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/kieron-pivotal/advent2017/06/memory"
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

	arr := strings.Split(strings.TrimSpace(string(contents)), "\t")
	narr := []int{}
	for _, nstr := range arr {
		n, _ := strconv.Atoi(nstr)
		narr = append(narr, n)
	}

	fmt.Println(memory.ReallocFirstCyclePos(narr))
}
