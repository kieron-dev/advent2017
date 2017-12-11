package tree

import "fmt"

func GetRoot(treeMap map[string]*Node, node *Node) string {
	setParents(treeMap)
	for node.Parent != nil {
		node = node.Parent
	}
	return node.Name
}

func setParents(treeMap map[string]*Node) {
	for _, node := range treeMap {
		for _, childStr := range node.Children {
			treeMap[childStr].Parent = node
		}
	}
}

func setWeights(treeMap map[string]*Node, node *Node) int {
	w := node.Weight
	weights := map[int]bool{}
	for _, childStr := range node.Children {
		weight := setWeights(treeMap, treeMap[childStr])
		weights[weight] = true
		w += weight
	}
	if len(weights) < 2 {
		node.Balanced = true
	}
	node.SubTreeWeight = w
	node.SubTreeWeight = w
	return node.SubTreeWeight
}

func GetWrongWeight(treeMap map[string]*Node, start *Node) int {
	node := treeMap[GetRoot(treeMap, start)]
	setWeights(treeMap, node)
	for {
		prev := node
		for _, childStr := range node.Children {
			child := treeMap[childStr]
			if !child.Balanced {
				node = child
				break
			}
		}
		if prev == node {
			break
		}
	}
	for _, childStr := range node.Children {
		child := treeMap[childStr]
		fmt.Printf("Node '%s', weight %d, subtreeWeight %d\n", child.Name, child.Weight, child.SubTreeWeight)
	}
	return 0

}
