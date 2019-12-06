package advent2019

type Graph struct {
	nodes map[string]*Node
}

func NewGraph() *Graph {
	g := Graph{}
	g.nodes = map[string]*Node{}
	return &g
}

func (g *Graph) AddEdge(parent, child string) {
	var parentNode, childNode *Node
	var ok bool
	if parentNode, ok = g.nodes[parent]; !ok {
		parentNode = g.NewNode(parent)
	}
	if childNode, ok = g.nodes[child]; !ok {
		childNode = g.NewNode(child)
	}
	parentNode.AddChild(childNode)
	childNode.AddParent(parentNode)
}

func (g *Graph) NodeCount() int {
	return len(g.nodes)
}

func (g *Graph) Depth(root, node string) int {
	rootNode, ok := g.nodes[root]
	if !ok {
		panic("root node not found: " + root)
	}
	descNode, ok := g.nodes[node]
	if !ok {
		panic("descendent node not found: " + node)
	}

	g.setDepths(rootNode, 0)

	return descNode.Depth()
}

func (g *Graph) SetDepths(root string) {
	rootNode, ok := g.nodes[root]
	if !ok {
		panic("root node not found: " + root)
	}
	g.setDepths(rootNode, 0)
}

func (g *Graph) setDepths(root *Node, depth int) {
	for _, child := range root.children {
		g.setDepths(child, depth+1)
	}
	root.depth = depth
}

func (g *Graph) NewNode(name string) *Node {
	n := Node{
		name:     name,
		children: []*Node{},
	}
	g.nodes[name] = &n
	return &n
}

func (g *Graph) Walk(rootName string, fn func(n *Node)) {
	root, ok := g.nodes[rootName]
	if !ok {
		panic("no root node: " + rootName)
	}
	g.walk(root, fn)
}

func (g *Graph) walk(node *Node, fn func(n *Node)) {
	for _, child := range node.children {
		g.walk(child, fn)
	}
	fn(node)
}

func (g *Graph) Climb(from string, fn func(n *Node)) {
	fromNode, ok := g.nodes[from]
	if !ok {
		panic("no from node: " + from)
	}
	g.climb(fromNode, fn)
}

func (g *Graph) climb(from *Node, fn func(n *Node)) {
	fn(from)
	if from.parent != nil {
		g.climb(from.parent, fn)
	}
}

// NODE

type Node struct {
	name     string
	parent   *Node
	children []*Node
	depth    int
}

func (n *Node) AddChild(child *Node) {
	n.children = append(n.children, child)
}

func (n *Node) AddParent(parent *Node) {
	n.parent = parent
}

func (n *Node) Name() string {
	return n.name
}

func (n *Node) Depth() int {
	return n.depth
}
