package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kieron-pivotal/advent2017/07/tree"
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

	treeMap := map[string]*tree.Node{}

	for scanner.Scan() {
		line := scanner.Text()
		name := getName(line)
		weight := getWeight(line)
		node := tree.NewNode(name, weight)
		for _, childName := range getChildren(line) {
			node.AddChild(childName)
		}
		treeMap[name] = node
	}

	fmt.Println(tree.GetRoot(treeMap))
}

func getName(line string) string {
	words := strings.Split(line, " ")
	return words[0]
}

func getWeight(line string) int {
	split1 := strings.Split(line, "(")
	split2 := strings.Split(split1[1], ")")
	n, _ := strconv.Atoi(split2[0])
	return n
}

func getChildren(line string) []string {
	split1 := strings.Split(line, "-> ")
	if len(split1) < 2 {
		return []string{}
	}
	return strings.Split(split1[1], ", ")
}
