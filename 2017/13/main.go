package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kieron-pivotal/advent2017/13/firewall"
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
	config := map[int]int{}
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), ": ")
		k, _ := strconv.Atoi(vals[0])
		v, _ := strconv.Atoi(vals[1])
		config[k] = v
	}

	fmt.Println("Part1:", firewall.Severity(config, 0))
	fmt.Println("Part2:", firewall.MinDelay(config))
}
