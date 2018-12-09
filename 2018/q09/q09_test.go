package q09_test

import (
	"github.com/kieron-pivotal/advent2017/2018/q09"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q09", func() {

	DescribeTable("Highest score", func(numPlayers, numMarbles, expectedHighestScore int) {
		game := q09.NewGame(numPlayers, numMarbles)
		Expect(game.Play()).To(Equal(expectedHighestScore))
	},

		Entry("quick", 9, 25, 32),

		Entry("ex 1", 10, 1618, 8317),
		Entry("ex 2", 13, 7999, 146373),
		Entry("ex 3", 17, 1104, 2764),
		Entry("ex 4", 21, 6111, 54718),
		Entry("ex 5", 30, 5807, 37305),
	)

})
