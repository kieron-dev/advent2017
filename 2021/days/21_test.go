package days_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("21", func() {
	It("does part A", func() {
		die := NewDeterministicDie(100)

		player1 := NewPlayer(4)
		player2 := NewPlayer(9)

		losingScore := 0
		for {
			player1.Move(die)
			if player1.score >= 1000 {
				losingScore = player2.score
				break
			}
			player2.Move(die)
			if player2.score >= 1000 {
				losingScore = player1.score
				break
			}
		}

		Expect(losingScore * die.rolls).To(Equal(903630))
	})

	It("does part B", func() {
		diceSumFreqs := map[int]int{}
		for i := 1; i < 4; i++ {
			for j := 1; j < 4; j++ {
				for k := 1; k < 4; k++ {
					diceSumFreqs[i+j+k]++
				}
			}
		}

		n := numWins(diceSumFreqs, 3, 0, 8, 0)
		Expect(n).To(Equal(303121579983974))
	})
})

var memo map[int]int

func numWins(diceSumFreqs map[int]int, posA, scoreA, posB, scoreB int) int {
	if memo == nil {
		memo = map[int]int{}
	}

	key := posA
	key = key*100 + scoreA
	key = key*100 + posB
	key = key*100 + scoreB

	if val, ok := memo[key]; ok {
		return val
	}

	sum := 0
	for diceA, countA := range diceSumFreqs {
		newPosA := (posA + diceA) % 10
		newScoreA := scoreA + newPosA + 1
		if newScoreA >= 21 {
			sum += countA
			continue
		}

		for diceB, countB := range diceSumFreqs {
			newPosB := (posB + diceB) % 10
			newScoreB := scoreB + newPosB + 1
			if newScoreB >= 21 {
				continue
			}
			sum += countA * countB * numWins(diceSumFreqs, newPosA, newScoreA, newPosB, newScoreB)
		}
	}

	memo[key] = sum
	return sum
}

const BoardSize = 10

type Player struct {
	pos   int
	score int
}

func (p Player) Position() int {
	return p.pos + 1
}

func NewPlayer(pos int) *Player {
	return &Player{
		pos: pos - 1,
	}
}

func (p *Player) Move(die *DeterministicDie) {
	for i := 0; i < 3; i++ {
		n := die.Roll()
		p.pos = (p.pos + n) % BoardSize
	}
	p.score += p.Position()
}

type DeterministicDie struct {
	next  int
	limit int
	rolls int
}

func NewDeterministicDie(limit int) *DeterministicDie {
	return &DeterministicDie{
		next:  1,
		limit: limit,
	}
}

func (d *DeterministicDie) Roll() int {
	n := d.next
	d.next = (d.next + 1) % d.limit
	d.rolls++
	return n
}
