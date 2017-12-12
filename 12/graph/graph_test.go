package graph_test

import (
	"github.com/kieron-pivotal/advent2017/12/graph"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Graph", func() {
	It("counts a tree", func() {
		g := graph.New()
		g.AddNode(0, []int{1, 2, 3})
		Expect(g.Size(0)).To(Equal(4))

		g.AddNode(1, []int{0})
		g.AddNode(2, []int{0})
		g.AddNode(3, []int{0})
		Expect(g.Size(0)).To(Equal(4))

		g.AddNode(1, []int{4})
		Expect(g.Size(0)).To(Equal(5))

		g.AddNode(10, []int{11})
		Expect(g.Size(0)).To(Equal(5))
		Expect(g.Size(10)).To(Equal(2))
	})

	FIt("counts trees", func() {
		g := graph.New()
		g.AddNode(0, []int{1, 2, 3})
		Expect(g.Groups()).To(Equal(1))

		g.AddNode(1, []int{0})
		g.AddNode(2, []int{0})
		g.AddNode(3, []int{0})
		Expect(g.Groups()).To(Equal(1))

		g.AddNode(1, []int{4})
		Expect(g.Groups()).To(Equal(1))

		g.AddNode(10, []int{11})
		Expect(g.Groups()).To(Equal(2))
	})
})
