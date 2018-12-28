package q20

import (
	"io"
	"io/ioutil"
	"log"
)

type Coord struct {
	X int
	Y int
}

func C(x, y int) Coord {
	return Coord{X: x, Y: y}
}

func (c Coord) Move(dir rune) Coord {
	coord := c
	switch dir {
	case 'N':
		coord.Y++
	case 'S':
		coord.Y--
	case 'W':
		coord.X--
	case 'E':
		coord.X++
	}
	return coord
}

type Node struct {
	Coord    Coord
	Children []*Node
}

type Plan struct {
	Regex string
	Root  *Node
	Nodes map[Coord]*Node

	WorkStack   [][]Coord
	BranchStack [][]Coord
}

func NewPlan(in io.Reader) *Plan {
	p := Plan{}

	bs, err := ioutil.ReadAll(in)
	if err != nil {
		log.Fatal("couldn't read input ", err)
	}

	re := string(bs)
	p.Regex = re

	p.Root = &Node{Coord: C(0, 0)}
	p.Nodes = map[Coord]*Node{C(0, 0): p.Root}

	return &p
}

func (p *Plan) AddChild(parentCoord, childCoord Coord) *Node {
	parent := p.Nodes[parentCoord]
	var child *Node
	if node, ok := p.Nodes[childCoord]; ok {
		child = node
	} else {
		child = &Node{Coord: childCoord}
		p.Nodes[childCoord] = child
	}
	parent.Children = append(parent.Children, child)
	return child
}

func (p *Plan) ProcessRegex() {
	for _, r := range p.Regex {
		switch r {
		case '^':
			p.WorkStack = [][]Coord{{C(0, 0)}}
		case 'N':
			fallthrough
		case 'S':
			fallthrough
		case 'W':
			fallthrough
		case 'E':
			for i, c := range p.WorkStack[len(p.WorkStack)-1] {
				next := c.Move(r)
				p.WorkStack[len(p.WorkStack)-1][i] = next
				p.AddChild(c, next)
			}
		case '(':
			latestWork := p.WorkStack[len(p.WorkStack)-1]
			latestCopy := make([]Coord, len(latestWork))
			copy(latestCopy, latestWork)
			p.WorkStack = append(p.WorkStack, latestCopy)
			p.BranchStack = append(p.BranchStack, []Coord{})
		case '|':
			p.AppendWorkStackHeadToBranchStackHead()
			latestWork := p.WorkStack[len(p.WorkStack)-2]
			latestCopy := make([]Coord, len(latestWork))
			copy(latestCopy, latestWork)
			p.WorkStack[len(p.WorkStack)-1] = latestCopy
		case ')':
			p.AppendWorkStackHeadToBranchStackHead()
			p.WorkStack = p.WorkStack[:len(p.WorkStack)-1]
			p.WorkStack[len(p.WorkStack)-1] = p.BranchStack[len(p.BranchStack)-1]
			p.BranchStack = p.BranchStack[:len(p.BranchStack)-1]
		}
	}
}

func (p *Plan) AppendWorkStackHeadToBranchStackHead() {
	coords := map[Coord]bool{}
	for _, c := range p.BranchStack[len(p.BranchStack)-1] {
		coords[c] = true
	}
	for _, c := range p.WorkStack[len(p.WorkStack)-1] {
		coords[c] = true
	}
	newSlice := []Coord{}
	for c, _ := range coords {
		newSlice = append(newSlice, c)
	}
	p.BranchStack[len(p.BranchStack)-1] = newSlice

}

func (p *Plan) FurthestRoom() (maxDepth, countOver999 int) {
	visited := map[Coord]bool{}
	stack := []Coord{C(0, 0)}
	depths := map[Coord]int{C(0, 0): 0}

	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]
		if visited[cur] {
			continue
		}

		for _, c := range p.Nodes[cur].Children {
			if visited[c.Coord] {
				continue
			}
			curDepth, isSet := depths[c.Coord]
			if !isSet || curDepth > depths[cur]+1 {
				depths[c.Coord] = depths[cur] + 1
			}
			if depths[c.Coord] > maxDepth {
				maxDepth = depths[c.Coord]
			}
			stack = append(stack, c.Coord)
		}

		visited[cur] = true
	}
	for _, d := range depths {
		if d > 999 {
			countOver999++
		}
	}
	return
}
