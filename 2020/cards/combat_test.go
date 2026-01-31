package cards_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/cards"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Combat", func() {
	var (
		data   io.Reader
		combat cards.Combat
	)

	BeforeEach(func() {
		data = strings.NewReader(`
Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
`)

		combat = cards.NewCombat()
		combat.Load(data)
	})

	It("can play a round", func() {
		combat.Play()
		Expect(combat.Stack(1)).To(Equal([]int{2, 6, 3, 1, 9, 5}))
		Expect(combat.Stack(2)).To(Equal([]int{8, 4, 7, 10}))
	})

	It("can play till there's a winner", func() {
		for combat.Winner() == 0 {
			combat.Play()
		}

		Expect(combat.Score(combat.Winner())).To(Equal(306))
	})

	Context("recursive combat", func() {
		BeforeEach(func() {
			combat.SetRecursive(true)
		})

		It("terminates with player 1 win on infinite loop", func() {
			combat.SetStacks([]int{43, 19}, []int{2, 29, 14})
			for combat.Winner() == 0 {
				combat.Play()
			}

			Expect(combat.Winner()).To(Equal(1))
		})

		It("recursively ends with correct score", func() {
			for combat.Winner() == 0 {
				combat.Play()
			}

			Expect(combat.Score(combat.Winner())).To(Equal(291))
		})
	})
})
