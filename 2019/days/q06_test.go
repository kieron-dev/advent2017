package days_test

import (
	"os"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q06", func() {
	var (
		file   *os.File
		graph  *advent2019.Graph
		reader advent2019.FileReader
	)

	BeforeEach(func() {
		var err error
		file, err = os.Open("./input06")
		if err != nil {
			panic(err)
		}
		graph = advent2019.NewGraph()
		reader = advent2019.FileReader{}
	})

	It("does part A", func() {
		reader.Each(file, func(line string) {
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

		Expect(sum).To(Equal(224901))
	})

	It("does part B", func() {
		reader.Each(file, func(line string) {
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

		res := i + j - 2
		Expect(res).To(Equal(334))
	})
})
