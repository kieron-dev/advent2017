package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type node struct {
	label    string
	children []*node
}

func main() {
	nodes := map[string]*node{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		var dependedOn, dependsOn string
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &dependedOn, &dependsOn)
		fmt.Printf("%s -> %s\n", dependsOn, dependedOn)
		var dependedOnNode, dependsOnNode *node
		var ok bool
		dependedOnNode, ok = nodes[dependedOn]
		if !ok {
			dependedOnNode = &node{label: dependedOn}
			nodes[dependedOn] = dependedOnNode
		}

		dependsOnNode, ok = nodes[dependsOn]
		if !ok {
			dependsOnNode = &node{label: dependsOn}
			nodes[dependsOn] = dependsOnNode
		}

		dependsOnNode.children = append(dependsOnNode.children, dependedOnNode)
	}

	output := []string{}

	for len(nodes) > 0 {
		leaves := GetLeaves(nodes)
		SortNodes(leaves)
		output = append(output, leaves[0].label)
		RemoveNode(nodes, leaves[0])
	}

	for _, o := range output {
		fmt.Printf("%s", o)
	}
	fmt.Println()
}

func GetLeaves(nodes map[string]*node) []*node {
	out := []*node{}
	for _, n := range nodes {
		if len(n.children) == 0 {
			out = append(out, n)
		}
	}
	return out
}

func RemoveNode(nodes map[string]*node, remove *node) {
	for _, n := range nodes {
		newChildren := []*node{}
		for _, c := range n.children {
			if c != remove {
				newChildren = append(newChildren, c)
			}
		}
		n.children = newChildren
	}
	delete(nodes, remove.label)
}

func SortNodes(nodes []*node) {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].label < nodes[j].label
	})
}
