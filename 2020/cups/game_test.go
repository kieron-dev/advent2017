package cups_test

import (
	"github.com/kieron-dev/adventofcode/2020/cups"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	var (
		data string
		game cups.Game
	)

	BeforeEach(func() {
		data = "389125467"
		game = cups.NewGame()
		game.Load(data, false)
	})

	It("can do 1 round", func() {
		game.Play()
		Expect(game.Cups()).To(Equal("54673289"))
	})

	It("can do 10 rounds", func() {
		for i := 0; i < 10; i++ {
			game.Play()
		}
		Expect(game.Cups()).To(Equal("92658374"))
	})

	Context("load till a million", func() {
		BeforeEach(func() {
			game.Load(data, true)
		})

		It("can play this way 10 million times", func() {
			for i := 0; i < 10000000; i++ {
				game.Play()
			}
			Expect(game.After1Prod()).To(Equal(149245887792))
		})
	})
})
