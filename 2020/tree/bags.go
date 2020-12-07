// Package tree does tree stuff
package tree

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Bags struct {
	nodes map[string]*Node
}

type Node struct {
	name     string
	children map[string]int
	parents  []*Node
}

func (n *Node) AddChild(child *Node, count int) {
	n.children[child.name] = count
	child.parents = append(child.parents, n)
}

func (n *Node) VisitAncestors(fn func(n *Node)) {
	for _, p := range n.parents {
		p.VisitAncestors(fn)
		fn(p)
	}
}

func (b *Bags) ContainedBags(bag string) int {
	c := 1
	bagNode := b.GetNode(bag)

	for name, count := range bagNode.children {
		c += count * b.ContainedBags(name)
	}

	return c
}

func (b *Bags) GetNode(name string) *Node {
	if _, ok := b.nodes[name]; !ok {
		b.nodes[name] = &Node{
			name:     name,
			children: map[string]int{},
		}
	}

	return b.nodes[name]
}

func NewBags() Bags {
	return Bags{
		nodes: map[string]*Node{},
	}
}

func (b *Bags) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		b.AddRule(line)
	}
}

func (b *Bags) AddRule(rule string) {
	node := b.GetNode(b.GetSubject(rule))

	for _, child := range b.GetChildren(rule) {
		childNode := b.GetNode(child.Name)
		node.AddChild(childNode, child.Count)
	}
}

func (b *Bags) GetSubject(rule string) string {
	items := strings.Split(rule, "bags")
	return strings.TrimSpace(items[0])
}

type Child struct {
	Name  string
	Count int
}

func (b *Bags) GetChildren(rule string) []Child {
	items := strings.Split(rule, "contain")
	right := strings.TrimSpace(items[1])

	if right == "no other bags." {
		return nil
	}

	bags := strings.Split(right, ",")

	var children []Child

	for _, desc := range bags {
		children = append(children, parseDesc(desc))
	}

	return children
}

func (b *Bags) NumOuterContaining(name string) int {
	node := b.GetNode(name)

	counts := map[string]bool{}
	node.VisitAncestors(func(n *Node) {
		counts[n.name] = true
	})

	return len(counts)
}

func (b *Bags) BagsInside(name string) int {
	return b.ContainedBags(name) - 1
}

var re = regexp.MustCompile(`(\d)+ (.*) bags?`)

func parseDesc(desc string) Child {
	matches := re.FindStringSubmatch(desc)

	if len(matches) != 3 {
		log.Fatalf("regexp failed on %q", desc)
	}

	count, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatalf("expected an int, got %q", matches[1])
	}

	return Child{
		Name:  matches[2],
		Count: count,
	}
}
