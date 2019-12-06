package advent2019_test

import (
	"github.com/kieron-pivotal/advent2017/advent2019"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("Graph", func() {

	var (
		g *advent2019.Graph
	)

	BeforeEach(func() {
		g = advent2019.NewGraph()
	})

	It("can add an edge", func() {
		parent := "aaa"
		child := "bbb"
		g.AddEdge(parent, child)
		Expect(g.NodeCount()).To(Equal(2))
	})

	It("gets the depth of a node", func() {
		g.AddEdge("aaa", "bbb")
		g.AddEdge("bbb", "ccc")
		Expect(g.Depth("aaa", "ccc")).To(Equal(2))
	})

	It("can walk the tree", func() {
		g.AddEdge("aaa", "bbb")
		g.AddEdge("bbb", "ccc")
		g.AddEdge("bbb", "ddd")
		edges := []string{}
		g.Walk("aaa", func(n *advent2019.Node) {
			edges = append(edges, n.Name())
		})
		Expect(edges).To(Equal([]string{"ccc", "ddd", "bbb", "aaa"}))
	})

	It("can sum the depths", func() {
		g.AddEdge("aaa", "bbb")
		g.AddEdge("bbb", "ccc")
		g.AddEdge("bbb", "ddd")
		g.SetDepths("aaa")
		sum := 0
		g.Walk("aaa", func(n *advent2019.Node) {
			sum += n.Depth()
		})
		Expect(sum).To(Equal(5))
	})

	It("can climb tree", func() {
		g.AddEdge("aaa", "bbb")
		g.AddEdge("bbb", "ccc")
		g.AddEdge("bbb", "ddd")
		nodes := []string{}
		g.Climb("ccc", func(n *advent2019.Node) {
			nodes = append(nodes, n.Name())
		})
		Expect(nodes).To(Equal([]string{"ccc", "bbb", "aaa"}))
	})
})
