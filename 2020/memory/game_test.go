package memory_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/memory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	var (
		data io.Reader
		game memory.Game
	)

	BeforeEach(func() {
		data = strings.NewReader(`0,3,6`)
		game = memory.NewGame()
		game.Load(data)
	})

	DescribeTable("get nth num", func(n, expected int) {
		Expect(game.Get(n)).To(Equal(expected))
	},

		Entry("1", 1, 0),
		Entry("2", 2, 3),
		Entry("3", 3, 6),

		Entry("4", 4, 0),
		Entry("5", 5, 3),
		Entry("6", 6, 3),
		Entry("7", 7, 1),
		Entry("8", 8, 0),
		Entry("9", 9, 4),
		Entry("10", 10, 0),
	)
})
