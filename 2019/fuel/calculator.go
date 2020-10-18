package fuel

import (
	"io"
	"strconv"
	"strings"

	"github.com/kieron-dev/advent2017/advent2019"
)

type NodeCount struct {
	node  *Node
	count int
}

type Node struct {
	name     string
	num      int
	parents  []NodeCount
	children []*Node
	depth    int
}

func NewNode(name string) *Node {
	node := Node{
		name: name,
	}
	return &node
}

type Calculator struct {
	nodes map[string]*Node
}

func NewCalculator() *Calculator {
	c := Calculator{
		nodes: map[string]*Node{},
	}
	return &c
}

func (c *Calculator) GetNode(name string) *Node {
	n, ok := c.nodes[name]
	if ok {
		return n
	}
	newNode := NewNode(name)
	c.nodes[name] = newNode
	return newNode
}

func (c *Calculator) SetProgram(r io.Reader) {
	fr := advent2019.FileReader{}
	fr.Each(r, func(line string) {
		c.AddEquation(line)
	})
	c.nodes["ORE"].num = 1
	c.setDepths(c.nodes["ORE"], 0)
}

func (c *Calculator) setDepths(node *Node, depth int) {
	node.depth = depth
	for _, child := range node.children {
		c.setDepths(child, depth+1)
	}
}

func (c *Calculator) AddEquation(line string) {
	sides := strings.Split(line, " => ")
	if len(sides) != 2 {
		panic("strange equation line: " + line)
	}
	inputs := map[string]int{}
	for _, part := range strings.Split(sides[0], ",") {
		part = strings.TrimSpace(part)
		constituents := strings.Split(part, " ")
		if len(constituents) != 2 {
			panic("didn't expect this: " + part)
		}
		n, err := strconv.Atoi(constituents[0])
		if err != nil {
			panic(err)
		}
		item := constituents[1]
		inputs[item] = n
	}
	constituents := strings.Split(strings.TrimSpace(sides[1]), " ")
	if len(constituents) != 2 {
		panic("didn't expect this: " + sides[1])
	}
	resultN, err := strconv.Atoi(constituents[0])
	if err != nil {
		panic(err)
	}
	resultItem := constituents[1]

	c.AddEdges(inputs, resultItem, resultN)
}

func (c *Calculator) AddEdges(inputs map[string]int, outputName string, outputNum int) {
	targetNode := c.GetNode(outputName)
	targetNode.num = outputNum
	for sourceName, sourceN := range inputs {
		sourceNode := c.GetNode(sourceName)
		sourceNode.children = append(sourceNode.children, targetNode)
		targetNode.parents = append(targetNode.parents, NodeCount{sourceNode, sourceN})
	}
}

func (c *Calculator) NodeCount() int {
	return len(c.nodes)
}

func (c *Calculator) FuelForOre(ore int) int {
	unitOre, unitOvers := c.OreForFuel(map[string]int{})

	mult := ore / unitOre

	fuel := mult
	overs := map[string]int{"ORE": ore - mult*unitOre}
	for k, v := range unitOvers {
		overs[k] = v * mult
	}
	c.reduce(overs)

	for overs["ORE"] >= unitOre {
		mult = overs["ORE"] / unitOre
		fuel += mult
		overs["ORE"] -= mult * unitOre
		for k, v := range unitOvers {
			overs[k] += v * mult
		}
		c.reduce(overs)
	}

	return fuel
}

func (c *Calculator) OreForFuel(overs map[string]int) (ore int, leftover map[string]int) {
	reqs := map[string]int{"FUEL": 1}

	for {
		if _, oreInMap := reqs["ORE"]; oreInMap && len(reqs) == 1 {
			break
		}
		for name, amount := range reqs {
			node := c.nodes[name]
			if len(node.parents) == 0 {
				continue
			}
			if overs[name] >= amount {
				overs[name] -= amount
				if overs[name] == 0 {
					delete(overs, name)
				}
				delete(reqs, name)
				continue
			}

			required := amount - overs[name]
			delete(overs, name)

			unit := node.num
			toTake := required / unit
			if toTake*unit < required {
				toTake++
			}
			over := toTake*unit - required
			delete(reqs, name)
			if over > 0 {
				overs[name] += over
			}
			for _, parent := range node.parents {
				reqs[parent.node.name] += toTake * parent.count
			}
		}
	}

	return reqs["ORE"], overs
}

func (c *Calculator) reduce(overs map[string]int) {
	for {
		didReduction := false

		for name, num := range overs {
			node := c.nodes[name]
			if len(node.parents) == 0 {
				continue
			}
			if node.num > num {
				continue
			}

			didReduction = true
			toRemove := num / node.num
			over := overs[name] - toRemove*node.num
			if over == 0 {
				delete(overs, name)
			} else {
				overs[name] = over
			}
			for _, p := range node.parents {
				overs[p.node.name] += p.count * toRemove
			}
		}

		if !didReduction {
			break
		}
	}
}
