package days_test

import (
	"io"
	"os"

	"github.com/kieron-dev/advent2017/advent2019/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q13", func() {
	var (
		g    *game.Game
		prog io.Reader
	)

	BeforeEach(func() {
		g = game.NewGame()
		var err error
		prog, err = os.Open("../days/input13")
		if err != nil {
			panic(err)
		}
	})

	JustBeforeEach(func() {
		g.LoadProgram(prog)
	})

	It("does part A", func() {
		g.Run()

		Expect(g.TileCount(game.Block)).To(Equal(372))
	})

	It("does part B", func() {
		g.Pay()
		score := g.Run()
		Expect(score).To(Equal(19297))
	})
})
