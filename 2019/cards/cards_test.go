package cards_test

import (
	"io"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019/cards"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("shuffles", func(num int, shuffles io.Reader, expectedOrder []int) {
	deck := cards.NewDeck(num)
	deck.SetShuffle(shuffles)
	Expect(deck.Cards()).To(Equal(expectedOrder))

	var posOf0 int

	for pos, v := range deck.Cards() {
		if v == 0 {
			posOf0 = pos
			break
		}
	}

	Expect(deck.PosOf(0)).To(Equal(posOf0))
	Expect(deck.CardAt(0)).To(Equal(expectedOrder[0]))
},

	Entry("identity", 10,
		strings.NewReader(""),
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),

	Entry("deal into new stack", 10,
		strings.NewReader("deal into new stack"),
		[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}),

	Entry("cut 3", 10,
		strings.NewReader("cut 3"),
		[]int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}),

	Entry("cut -4", 10,
		strings.NewReader("cut -4"),
		[]int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}),

	Entry("deal with increment 3", 10,
		strings.NewReader("deal with increment 3"),
		[]int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}),

	Entry("multi 1", 10,
		strings.NewReader(`deal with increment 7
   					  	   deal into new stack
   						   deal into new stack`),
		[]int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}),

	Entry("multi 2", 10,
		strings.NewReader(`cut 6
						   deal with increment 7
						   deal into new stack`),
		[]int{3, 0, 7, 4, 1, 8, 5, 2, 9, 6}),

	Entry("multi 3", 10,
		strings.NewReader(`deal into new stack
						   cut -2
						   deal with increment 7
						   cut 8
						   cut -4
						   deal with increment 7
						   cut 3
						   deal with increment 9
						   deal with increment 3
						   cut -1`),
		[]int{9, 2, 5, 8, 1, 4, 7, 0, 3, 6}),
)

var _ = Describe("period", func() {
	var deck *cards.Deck
	BeforeEach(func() {
		deck = cards.NewDeck(11)
		deck.SetShuffle(strings.NewReader(`deal into new stack
										   cut -2
										   deal with increment 7
										   cut 8
										   cut -4
										   deal with increment 7
										   cut 3
										   deal with increment 9
										   deal with increment 3
										   cut -1`))
	})

	FIt("calcs the period", func() {
		Expect(deck.Period(5)).To(Equal(-1))
	})
})
