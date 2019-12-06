package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019"
)

func main() {
	reader := advent2019.FileReader{}
	graph := advent2019.NewGraph()

	reader.Each(os.Stdin, func(line string) {
		parts := strings.Split(line, ")")
		if len(parts) != 2 {
			panic("eh?")
		}
		graph.AddEdge(parts[0], parts[1])
	})

	youNodes := []string{}
	sanNodes := []string{}
	sanMap := map[string]bool{}
	graph.Climb("YOU", func(n *advent2019.Node) {
		youNodes = append(youNodes, n.Name())
	})
	graph.Climb("SAN", func(n *advent2019.Node) {
		sanNodes = append(sanNodes, n.Name())
		sanMap[n.Name()] = true
	})

	var i int
	var o string
	for i, o = range youNodes {
		if sanMap[o] {
			break
		}
	}

	var j int
	var p string
	for j, p = range sanNodes {
		if p == o {
			break
		}
	}

	fmt.Printf("res = %d\n", i+j-2)

}
