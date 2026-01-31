package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/memory"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("15", func() {
	var (
		data *os.File
		game memory.Game
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input15")
		Expect(err).NotTo(HaveOccurred())

		game = memory.NewGame()
		game.Load(data)
	})

	It("does part A", func() {
		Expect(game.Get(2020)).To(Equal(700))
	})

	It("does part B", func() {
		Expect(game.Get(30000000)).To(Equal(51358))
	})
})
