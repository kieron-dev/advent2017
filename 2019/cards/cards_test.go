package cards_test

import (
	"io"
	"strconv"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019/cards"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("shuffles", func(size int, shuffles io.Reader, expectedOrder []int) {
	deck := cards.NewDeck(size)
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

var _ = Describe("big nums", func() {
	var shuffles io.Reader

	BeforeEach(func() {
		shuffles = strings.NewReader(`deal into new stack
						   cut -2
						   deal with increment 7
						   cut 8
						   cut -4
						   deal with increment 7
						   cut 3
						   deal with increment 9
						   deal with increment 3
						   cut -1`)
	})

	It("can deal with the big input", func() {
		deck := cards.NewDeck(119315717514047)
		deck.SetShuffle(shuffles)

		p0 := 0
		p1 := deck.PosOf(p0)
		p2 := deck.PosOf(p1)
		c1 := deck.CardAt(p2)
		c0 := deck.CardAt(c1)

		Expect(p0).To(Equal(c0))
		Expect(p1).To(Equal(c1))
	})
})

var _ = Describe("double shuffle", func() {
	var (
		deck       *cards.Deck
		shuffle    io.Reader
		size       int
		iterations int
	)

	BeforeEach(func() {
		size = 10007
		iterations = 27
		shuffle = strings.NewReader(`cut 2`)
	})

	JustBeforeEach(func() {
		deck = cards.NewDeck(size)
		deck.SetShuffle(shuffle)
	})

	It("gets compound shuffles correct", func() {
		for i := 0; i < size; i++ {
			expectedCard := i
			for j := 0; j < iterations; j++ {
				expectedCard = deck.CardAt(expectedCard)
			}

			card := deck.EquivalentCardAt(i, iterations)
			Expect(card).To(Equal(expectedCard), strconv.Itoa(i))
		}
	})
})
