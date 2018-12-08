package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type node struct {
	label      string
	children   []*node
	inProgress bool
}

type worker struct {
	task     string
	complete int
}

func main() {
	nodes := map[string]*node{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		var dependedOn, dependsOn string
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &dependedOn, &dependsOn)
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

	t := 0
	numWorkers := 5
	extraJobTime := 60
	var workers []*worker
	for i := 0; i < numWorkers; i++ {
		workers = append(workers, &worker{})
	}

	for len(nodes) > 0 {
		justFinished := GetJustFinished(workers, t)
		for _, finished := range justFinished {
			RemoveNode(nodes, nodes[finished])
		}
		leaves := GetLeaves(nodes)
		SortNodes(leaves)
		availableWorkers := AvailableWorkers(workers, t, len(leaves))
		AssignWorkers(leaves, availableWorkers, t, extraJobTime)
		t++
	}
	fmt.Printf("total time = %+v\n", t-1)

}

func AssignWorkers(nodes []*node, workers []*worker, t, extraJobTime int) {
	l := len(nodes)
	if len(workers) < l {
		l = len(workers)
	}
	for i := 0; i < l; i++ {
		workers[i].task = nodes[i].label
		workers[i].complete = t + extraJobTime + int(nodes[i].label[0]-'A'+1)
		nodes[i].inProgress = true
	}
}

func GetJustFinished(workers []*worker, t int) []string {
	var out []string
	for _, w := range workers {
		if w.task != "" && w.complete == t {
			out = append(out, w.task)
			w.task = ""
			w.complete = 0
		}
	}
	sort.Strings(out)
	return out
}

func AvailableWorkers(workers []*worker, t int, required int) []*worker {
	out := []*worker{}
	for _, w := range workers {
		if len(out) == required {
			break
		}
		if w.complete <= t {
			out = append(out, w)
		}
	}
	return out
}

func GetLeaves(nodes map[string]*node) []*node {
	out := []*node{}
	for _, n := range nodes {
		if n.inProgress {
			continue
		}
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
