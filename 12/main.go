package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kieron-pivotal/advent2017/12/graph"
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

	g := graph.New()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		idChildren := strings.Split(scanner.Text(), " <-> ")
		idStr := idChildren[0]
		id, _ := strconv.Atoi(idStr)
		children := []int{}
		for _, c := range strings.Split(idChildren[1], ", ") {
			childId, _ := strconv.Atoi(c)
			children = append(children, childId)
		}
		g.AddNode(id, children)
	}

	fmt.Println("Part1:", g.Size(0))
	fmt.Println("Part2:", g.Groups())
}
