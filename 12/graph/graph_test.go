package graph_test

import (
	"github.com/kieron-pivotal/advent2017/12/graph"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Graph", func() {
	It("counts a simple directed tree", func() {
		g := graph.New()
		g.LinkNodes(0, []int{1, 2, 3})
		Expect(g.Size(0)).To(Equal(4))
	})

	It("counts a tree with reverse links", func() {
		g := graph.New()
		g.LinkNodes(0, []int{1, 2, 3})
		g.LinkNodes(1, []int{0})
		g.LinkNodes(2, []int{0})
		g.LinkNodes(3, []int{0})
		Expect(g.Size(0)).To(Equal(4))
	})

	It("counts a tree with another reverse links", func() {
		g := graph.New()
		g.LinkNodes(0, []int{1, 2, 3})
		g.LinkNodes(1, []int{0})
		g.LinkNodes(2, []int{0})
		g.LinkNodes(3, []int{0})
		g.LinkNodes(1, []int{4})
		g.LinkNodes(4, []int{1})
		Expect(g.Size(0)).To(Equal(5))
	})

	It("counts a forest with reverse links", func() {
		g := graph.New()
		g.LinkNodes(0, []int{1, 2, 3})
		g.LinkNodes(1, []int{0})
		g.LinkNodes(2, []int{0})
		g.LinkNodes(3, []int{0})
		g.LinkNodes(1, []int{4})
		g.LinkNodes(4, []int{1})
		g.LinkNodes(10, []int{11})
		g.LinkNodes(11, []int{10})
		Expect(g.Size(0)).To(Equal(5))
		Expect(g.Size(10)).To(Equal(2))
	})

	It("reverse links are now implicit", func() {
		g := graph.New()
		g.LinkNodes(0, []int{1, 2, 3})
		g.LinkNodes(1, []int{4})
		g.LinkNodes(10, []int{11})
		Expect(g.Size(0)).To(Equal(5))
		Expect(g.Size(10)).To(Equal(2))
	})

	It("counts trees", func() {
		g := graph.New()
		g.LinkNodes(0, []int{1, 2, 3})
		Expect(g.Groups()).To(Equal(1))

		g.LinkNodes(1, []int{4})
		Expect(g.Groups()).To(Equal(1))

		g.LinkNodes(10, []int{11})
		Expect(g.Groups()).To(Equal(2))
	})
})
