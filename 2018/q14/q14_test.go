package q14_test

import (
	"github.com/kieron-pivotal/advent2017/2018/q14"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q14", func() {
	var (
		r *q14.Recipes
	)

	BeforeEach(func() {
		r = q14.NewRecipes()
	})

	Context("basic", func() {
		It("can add new recipes", func() {
			r.Step()
			Expect(r.Length).To(Equal(4))
			Expect(r.End.Score).To(Equal(0))
			Expect(r.End.Left.Score).To(Equal(1))
		})

		It("can move elves after step", func() {
			r.Step()
			r.Step()
			Expect(r.ElfARecipe).To(Equal(r.End.Left))
			Expect(r.ElfBRecipe).To(Equal(r.End.Left.Left))
		})

		DescribeTable("can get ten scores after given length of prefix", func(prefixLen int, expected string) {
			Expect(r.ScoresAfter(prefixLen)).To(Equal(expected))
		},
			Entry("9", 9, "5158916779"),
			Entry("5", 5, "0124515891"),
			Entry("18", 18, "9251071085"),
			Entry("2018", 2018, "5941429882"),
		)

	})

	Context("part II", func() {
		DescribeTable("can count recipes until sequence appears", func(sequence string, count int) {
			Expect(r.ScoresBefore(sequence)).To(Equal(count))
		},
			Entry("1", "51589", 9),
			Entry("1", "01245", 5),
			Entry("1", "92510", 18),
			Entry("1", "59414", 2018),
		)
	})
})
