package days_test

import (
	"github.com/kieron-dev/adventofcode/2020/cups"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("23", func() {
	var (
		data string
		game cups.Game
	)

	BeforeEach(func() {
		data = "538914762"
		game = cups.NewGame()
		game.Load(data, false)
	})

	It("does part A", func() {
		for i := 0; i < 100; i++ {
			game.Play()
		}

		Expect(game.Cups()).To(Equal("54327968"))
	})

	It("does part B", func() {
		game.Load(data, true)
		for i := 0; i < 10000000; i++ {
			game.Play()
		}

		Expect(game.After1Prod()).To(Equal(157410423276))
	})
})
