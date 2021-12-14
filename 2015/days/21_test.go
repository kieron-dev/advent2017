package days_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type rpgChar struct {
	hitPoints int
	damage    int
	armor     int
}

func (r *rpgChar) receiveHit(o *rpgChar) {
	r.hitPoints -= (o.damage - r.armor)
}

func (r *rpgChar) fightOpponent() bool {
	opponent := &rpgChar{
		hitPoints: 109,
		damage:    8,
		armor:     2,
	}
	r.hitPoints = 100

	for {
		if r.hitPoints > 0 {
			opponent.receiveHit(r)
		} else {
			break
		}
		if opponent.hitPoints > 0 {
			r.receiveHit(opponent)
		} else {
			break
		}
	}

	return r.hitPoints > 0
}

type rpgItem struct {
	name   string
	cost   int
	armor  int
	damage int
}

var _ = Describe("21", func() {
	weapons := []rpgItem{
		{name: "Dagger", cost: 8, damage: 4},
		{name: "Shortsword", cost: 10, damage: 5},
		{name: "Warhammer", cost: 25, damage: 6},
		{name: "Longsword", cost: 40, damage: 7},
		{name: "Greataxe", cost: 74, damage: 8},
	}
	armors := []rpgItem{
		{name: "Birthday suit"},
		{name: "Leather", cost: 13, armor: 1},
		{name: "Chainmail", cost: 31, armor: 2},
		{name: "Splintmail", cost: 53, armor: 3},
		{name: "Bandedmail", cost: 75, armor: 4},
		{name: "Platemail", cost: 102, armor: 5},
	}
	rings := []rpgItem{
		{name: "Damage +1", cost: 25, damage: 1},
		{name: "Damage +2", cost: 50, damage: 2},
		{name: "Damage +3", cost: 100, damage: 3},
		{name: "Armor +1", cost: 20, damage: 1},
		{name: "Armor +2", cost: 40, damage: 2},
		{name: "Armor +3", cost: 80, damage: 3},
	}

	It("does part A", func() {
		me := &rpgChar{}

		min := 100000
		var cost int

		for _, weapon := range weapons {
			me.damage = weapon.damage
			cost += weapon.cost

			for _, armor := range armors {
				me.armor = armor.armor
				cost += armor.cost

				if me.fightOpponent() && cost < min {
					min = cost
				}

				for i, ring1 := range rings {
					cost += ring1.cost
					me.armor += ring1.armor
					me.damage += ring1.damage

					if me.fightOpponent() && cost < min {
						min = cost
					}

					for j := i + 1; j < len(rings); j++ {
						ring2 := rings[j]
						cost += ring2.cost
						me.armor += ring2.armor
						me.damage += ring2.damage

						if me.fightOpponent() && cost < min {
							min = cost
						}
						cost -= ring2.cost
						me.armor -= ring2.armor
						me.damage -= ring2.damage
					}

					cost -= ring1.cost
					me.armor -= ring1.armor
					me.damage -= ring1.damage

				}

				cost -= armor.cost
			}

			cost -= weapon.cost

		}

		Expect(min).To(Equal(111))
	})

	It("does part B", func() {
		me := &rpgChar{}

		max := 0
		var cost int

		for _, weapon := range weapons {
			me.damage = weapon.damage
			cost += weapon.cost

			for _, armor := range armors {
				me.armor = armor.armor
				cost += armor.cost

				if !me.fightOpponent() && cost > max {
					max = cost
				}

				for i, ring1 := range rings {
					cost += ring1.cost
					me.armor += ring1.armor
					me.damage += ring1.damage

					if !me.fightOpponent() && cost > max {
						max = cost
					}

					for j := i + 1; j < len(rings); j++ {
						ring2 := rings[j]
						cost += ring2.cost
						me.armor += ring2.armor
						me.damage += ring2.damage

						if !me.fightOpponent() && cost > max {
							max = cost
						}
						cost -= ring2.cost
						me.armor -= ring2.armor
						me.damage -= ring2.damage
					}
					cost -= ring1.cost
					me.armor -= ring1.armor
					me.damage -= ring1.damage
				}
				cost -= armor.cost
			}
			cost -= weapon.cost
		}

		Expect(max).To(Equal(188))
	})
})
