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

	sum := 0
	graph.SetDepths("COM")
	graph.Walk("COM", func(n *advent2019.Node) {
		sum += n.Depth()
	})

	fmt.Printf("sum = %+v\n", sum)
}
