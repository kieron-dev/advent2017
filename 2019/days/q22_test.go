package days_test

import (
	"fmt"
	"io"
	"os"

	"github.com/kieron-pivotal/advent2017/advent2019/cards"
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

	FIt("can do part B", func() {
		deck := cards.NewDeck(119315717514047)
		deck.SetShuffle(shuffles)

		period := deck.Period(2020)
		fmt.Printf("period = %+v\n", period)
		rem := 101741582076661 % period

		card := 2020
		for i := 0; i < rem; i++ {
			card = deck.CardAt(card)
		}

		Expect(card).To(Equal(-1))
	})
})
