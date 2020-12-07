package tree_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/tree"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("parsing",
	func(rule, subject string, expChildren []string, expCounts []int) {
		b := tree.NewBags()

		Expect(b.GetSubject(rule)).To(Equal(subject))

		var children []string
		var childCounts []int

		for _, child := range b.GetChildren(rule) {
			children = append(children, child.Name)
			childCounts = append(childCounts, child.Count)
		}

		Expect(children).To(ConsistOf(expChildren))
		Expect(childCounts).To(ConsistOf(expCounts))
	},

	Entry("1", "bright white bags contain 1 shiny gold bag.", "bright white", []string{"shiny gold"}, []int{1}),
	Entry("2", "light red bags contain 1 bright white bag, 2 muted yellow bags.", "light red", []string{"bright white", "muted yellow"}, []int{1, 2}),
	Entry("3", "faded blue bags contain no other bags.", "faded blue", []string{}, []int{}),
)

var _ = Describe("Bags", func() {
	var (
		data io.Reader
		bags tree.Bags
	)

	BeforeEach(func() {
		data = strings.NewReader(`
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`)
		bags = tree.NewBags()
		bags.Load(data)
	})

	It("can count ancestors", func() {
		Expect(bags.NumOuterContaining("shiny gold")).To(Equal(4))
	})

	It("can count bags inside", func() {
		Expect(bags.BagsInside("shiny gold")).To(Equal(32))
	})
})
