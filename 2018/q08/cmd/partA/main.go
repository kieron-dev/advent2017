package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	metadata []int
	children []*Node
}

func main() {
	nums := GetInput(os.Stdin)
	root, _ := ParseNode(nums, 0)

	sumMetaA := SumMetaA(root)
	fmt.Printf("sumMetaA = %+v\n", sumMetaA)

	sumMetaB := SumMetaB(root)
	fmt.Printf("sumMetaB = %+v\n", sumMetaB)
}

func GetInput(reader io.Reader) []int {
	input, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	sinput := strings.TrimSpace(string(input))
	nums := strings.Split(sinput, " ")
	ints := []int{}
	for _, num := range nums {
		n, err := strconv.Atoi(string(num))
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, n)
	}
	return ints
}

func ParseNode(input []int, start int) (*Node, int) {
	childCount := input[start]
	metaCount := input[start+1]
	pos := start + 2
	node := Node{}
	for i := 0; i < childCount; i++ {
		var child *Node
		child, pos = ParseNode(input, pos)
		node.children = append(node.children, child)
	}
	for i := 0; i < metaCount; i++ {
		node.metadata = append(node.metadata, input[pos])
		pos++
	}
	return &node, pos
}

func SumMetaA(node *Node) int {
	sum := 0
	for _, m := range node.metadata {
		sum += m
	}
	for _, c := range node.children {
		sum += SumMetaA(c)
	}
	return sum
}

func SumMetaB(node *Node) int {
	sum := 0
	if len(node.children) == 0 {
		for _, m := range node.metadata {
			sum += m
		}
		return sum
	}
	for _, m := range node.metadata {
		if m == 0 || m > len(node.children) {
			continue
		}
		sum += SumMetaB(node.children[m-1])
	}
	return sum
}
