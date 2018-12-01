package tree

type Node struct {
	Name          string
	Weight        int
	Children      []string
	Parent        *Node
	Balanced      bool
	SubTreeWeight int
}

func NewNode(name string, weight int) *Node {
	node := Node{Name: name, Weight: weight}
	return &node
}

func (n *Node) AddChild(childName string) {
	n.Children = append(n.Children, childName)
}
