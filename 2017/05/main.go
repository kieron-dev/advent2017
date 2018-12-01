package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/kieron-pivotal/advent2017/05/instructions"
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

	arr := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		arr = append(arr, num)
	}

	fmt.Println(instructions.Count(arr))
}
