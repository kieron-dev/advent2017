package two023_test

import (
	"bytes"
	"os"
	"sort"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type hand struct {
	cards []byte
	bid   int
	typeA handType
	typeB handType
}

func newHand(cards, bid []byte) hand {
	n, err := strconv.Atoi(string(bid))
	Expect(err).NotTo(HaveOccurred())
	h := hand{
		cards: cards,
		bid:   n,
	}
	h.typeA = h.getHandTypeA()
	h.typeB = h.getHandTypeB()
	return h
}

type handType int

const (
	fiveOfKind handType = iota
	fourOfKind
	fullHouse
	threeOfKind
	twoPair
	onePair
	highCard
)

var cardsA = map[byte]int{
	'A': 1,
	'K': 2,
	'Q': 3,
	'J': 4,
	'T': 5,
	'9': 6,
	'8': 7,
	'7': 8,
	'6': 9,
	'5': 10,
	'4': 11,
	'3': 12,
	'2': 13,
}

var cardsB = map[byte]int{
	'A': 1,
	'K': 2,
	'Q': 3,
	'T': 4,
	'9': 5,
	'8': 6,
	'7': 7,
	'6': 8,
	'5': 9,
	'4': 10,
	'3': 11,
	'2': 12,
	'J': 13,
}

func (h hand) getHandTypeA() handType {
	cardSet := map[byte]int{}
	for _, b := range h.cards {
		cardSet[b]++
	}
	prod := 1
	for _, v := range cardSet {
		prod *= v
	}

	switch len(cardSet) {
	case 1:
		return fiveOfKind
	case 2:
		if prod == 4 {
			return fourOfKind
		}
		return fullHouse
	case 3:
		if prod == 3 {
			return threeOfKind
		}
		return twoPair
	case 4:
		return onePair
	default:
		return highCard
	}
}

func (h hand) getHandTypeB() handType {
	cardSet := map[byte]int{}
	for _, b := range h.cards {
		cardSet[b]++
	}
	jokers := cardSet['J']
	if jokers > 0 && jokers < 5 {
		max := 0
		var maxK byte
		for k, v := range cardSet {
			if k == 'J' {
				continue
			}
			if v > max {
				max = v
				maxK = k
			}
		}
		delete(cardSet, 'J')
		cardSet[maxK] += jokers
	}

	prod := 1
	for _, v := range cardSet {
		prod *= v
	}

	switch len(cardSet) {
	case 1:
		return fiveOfKind
	case 2:
		if prod == 4 {
			return fourOfKind
		}
		return fullHouse
	case 3:
		if prod == 3 {
			return threeOfKind
		}
		return twoPair
	case 4:
		return onePair
	default:
		return highCard
	}
}

var _ = Describe("07", func() {
	It("does part A", func() {
		bs, err := os.ReadFile("input07")
		Expect(err).NotTo(HaveOccurred())

		hands := []hand{}
		fs := bytes.Fields(bs)
		for i := 0; i < len(fs); i += 2 {
			h := newHand(fs[i], fs[i+1])
			hands = append(hands, h)
		}

		sort.Slice(hands, func(a, b int) bool {
			handA := hands[a]
			handB := hands[b]
			if handA.typeA == handB.typeA {
				for i := 0; i < len(hands[a].cards); i++ {
					if handA.cards[i] == handB.cards[i] {
						continue
					}
					return cardsA[handA.cards[i]] > cardsA[handB.cards[i]]
				}
			}
			return handA.typeA > handB.typeA
		})

		sum := 0
		for i, h := range hands {
			sum += h.bid * (i + 1)
		}

		Expect(sum).To(Equal(250347426))
	})

	It("does part B", func() {
		bs, err := os.ReadFile("input07")
		Expect(err).NotTo(HaveOccurred())

		hands := []hand{}
		fs := bytes.Fields(bs)
		for i := 0; i < len(fs); i += 2 {
			h := newHand(fs[i], fs[i+1])
			hands = append(hands, h)
		}

		sort.Slice(hands, func(a, b int) bool {
			handA := hands[a]
			handB := hands[b]
			if handA.typeB == handB.typeB {
				for i := 0; i < len(hands[a].cards); i++ {
					if handA.cards[i] == handB.cards[i] {
						continue
					}
					return cardsB[handA.cards[i]] > cardsB[handB.cards[i]]
				}
			}
			return handA.typeB > handB.typeB
		})

		sum := 0
		for i, h := range hands {
			sum += h.bid * (i + 1)
		}

		Expect(sum).To(Equal(251224870))
	})
})
