package life_test

import (
	"github.com/kieron-dev/advent2017/advent2019/life"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var (
	state1 = `....#
#..#.
#..##
..#..
#....`

	state2 = `#..#.
####.
###.#
##.##
.##..`

	state3 = `#####
....#
....#
...#.
#.###`
)

var _ = Describe("Single Chart Life", func() {
	DescribeTable("string to life, and back again", func(input string, equivInt int) {
		state := life.New(input)

		Expect(int(state)).To(Equal(equivInt))
		Expect(life.Life(equivInt).Chart()).To(Equal(input))
	},

		Entry("zero", ".....\n.....\n.....\n.....\n.....", 0),
		Entry("one", "#....\n.....\n.....\n.....\n.....", 1),
		Entry("two", ".#...\n.....\n.....\n.....\n.....", 2),
		Entry("thirty-two", ".....\n#....\n.....\n.....\n.....", 32),
		Entry("2^26 - 1", "#####\n#####\n#####\n#####\n#####", 1<<25-1),
	)

	DescribeTable("count neighbours", func(input string, r, c, count int) {
		l := life.New(input)
		Expect(l.Neighbours(0, 0)).To(Equal(1))
	},

		Entry("state1 0,0", state1, 0, 0, 1),
		Entry("state1 2,2", state1, 2, 2, 3),
		Entry("state1 4,1", state1, 4, 1, 2),
		Entry("state2 1,1", state2, 1, 1, 4),
	)

	DescribeTable("evolution", func(input, output string) {
		initial := life.New(input)
		next := initial.Evolve()
		Expect(next.Chart()).To(Equal(output))
	},

		Entry("state1 -> state2", state1, state2),
		Entry("state2 -> state3", state2, state3),
	)

	It("determines the first repeat", func() {
		l := life.New(state1)
		repeat := l.FirstRepeat()
		Expect(int(repeat)).To(Equal(2129920))
	})
})

var _ = Describe("infinitely stacked life", func() {
	var line life.Line

	BeforeEach(func() {
		line = life.NewLine(state1)
	})

	It("creates new charts while evolving", func() {
		for i := 0; i < 10; i++ {
			line = line.Evolve()
		}
		Expect(line.CountBugs()).To(Equal(99))
	})
})
