package game_test

import (
	"io"
	"os"

	"github.com/kieron-pivotal/advent2017/advent2019/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {

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

	It("can count tiles added", func() {
		g.Run()
		Expect(g.TileCount(game.Block)).To(Equal(-1))
	})

})
