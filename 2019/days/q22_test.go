package days_test

import (
	"io"
	"os"

	"github.com/kieron-dev/advent2017/advent2019/cards"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q22", func() {
	var shuffles io.ReadCloser

	BeforeEach(func() {
		var err error
		shuffles, err = os.Open("./input22")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		shuffles.Close()
	})

	It("can do part A", func() {
		deck := cards.NewDeck(10007)
		deck.SetShuffle(shuffles)

		Expect(deck.PosOf(2019)).To(Equal(6794))
	})

	It("can do part B", func() {
		deck := cards.NewDeck(119315717514047)
		deck.SetShuffle(shuffles)
		Expect(deck.EquivalentCardAt(2020, 101741582076661)).To(Equal(70725194521472))
	})
})
