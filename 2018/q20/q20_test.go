package q20_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kieron-pivotal/advent2017/2018/q20"
)

var _ = Describe("Q20", func() {
	It("do a simple graph without branch points", func() {
		p := q20.NewPlan(strings.NewReader("^WNE$"))
		p.ProcessRegex()
		Expect(p.FurthestRoom()).To(Equal(3))
	})

	It("can do a graph with a single branch point", func() {
		p := q20.NewPlan(strings.NewReader("^E(N|SE)E$"))
		p.ProcessRegex()
		furthest, _ := p.FurthestRoom()
		Expect(furthest).To(Equal(4))
	})
})
