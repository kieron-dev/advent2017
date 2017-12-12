package graph

type Graph struct {
	nodes map[int]Node
}

func New() *Graph {
	g := Graph{}
	g.nodes = map[int]Node{}
	return &g
}

type Node struct {
	id       int
	children []int
}

func (g *Graph) AddNode(id int, children []int) {
	g.nodes[id] = Node{id: id, children: children}
	// for _, child := range children {
	// 	childNode, ok := g.nodes[child]
	// 	if !ok {
	// 		g.nodes[child] = Node{id: child}
	// 	}
	// 	g.nodes[child].children = append(g.nodes[child].children, []int{id})
	// }
}

func (g *Graph) Size(containing int) int {
	visited := map[int]bool{}
	queue := []int{containing}
	count := 0

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		if _, ok := visited[next]; ok {
			continue
		}
		visited[next] = true

		node := g.nodes[next]
		count++
		queue = append(queue, node.children...)
	}
	return count
}

func (g *Graph) Groups() int {
	visited := map[int]bool{}
	count := 0
	for id, _ := range g.nodes {
		if _, ok := visited[id]; ok {
			continue
		}
		count++

		queue := []int{id}

		for len(queue) > 0 {
			next := queue[0]
			queue = queue[1:]
			if _, ok := visited[next]; ok {
				continue
			}
			visited[next] = true

			node := g.nodes[next]
			queue = append(queue, node.children...)
		}
	}
	return count
}
