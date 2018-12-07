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
	roots := map[string]bool{}

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
		roots[dependedOn] = false

		dependsOnNode, ok = nodes[dependsOn]
		if !ok {
			dependsOnNode = &node{label: dependsOn}
			nodes[dependsOn] = dependsOnNode
			roots[dependsOn] = true
		}

		dependsOnNode.children = append(dependsOnNode.children, dependedOnNode)

	}

	fmt.Printf("roots = %+v\n", roots)

	rootNodes := []*node{}
	for n, isRoot := range roots {
		if !isRoot {
			continue
		}
		rootNodes = append(rootNodes, nodes[n])
	}

	SortNodes(rootNodes)
	order := []string{}
	for _, n := range rootNodes {
		fmt.Printf("n = %+v\n", n)
		n.DFS(&order, map[*node]bool{})
	}

	for _, o := range order {
		fmt.Print(o)
	}
	fmt.Println()
	fmt.Printf("order = %+v\n", order)

}

func (n *node) DFS(order *[]string, visited map[*node]bool) {
	SortNodes(n.children)
	for _, c := range n.children {
		if visited[c] {
			continue
		}
		c.DFS(order, visited)
	}
	// fmt.Println(n.label)
	*order = append(*order, n.label)
	visited[n] = true
}

func SortNodes(nodes []*node) {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].label < nodes[j].label
	})
}
