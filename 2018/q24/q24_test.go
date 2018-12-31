package q24_test

import (
	"io"
	"strings"

	"github.com/kieron-pivotal/advent2017/2018/q24"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q24", func() {

	var (
		ex         io.Reader
		noBrackers io.Reader
	)

	BeforeEach(func() {
		ex = strings.NewReader(`Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4`)

		noBrackers = strings.NewReader(`Infection:
2 units each with 3 hit points with an attack that does 5 pooing damage at initiative 7`)
	})

	It("can load the groups", func() {

		s := q24.NewSystem(ex)

		Expect(s.ImmuneGroups).To(HaveLen(2))
		Expect(s.InfectionGroups).To(HaveLen(2))

		immuneGroup1 := s.ImmuneGroups[0]
		immuneGroup2 := s.ImmuneGroups[1]
		Expect(immuneGroup1.Units).To(Equal(17))
		Expect(immuneGroup2.HitPoints).To(Equal(1274))
		Expect(immuneGroup1.WeakTo).To(Equal([]string{"radiation", "bludgeoning"}))
		Expect(immuneGroup2.ImmuneTo).To(Equal([]string{"fire"}))

		infectionGroup1 := s.InfectionGroups[0]
		infectionGroup2 := s.InfectionGroups[1]
		Expect(infectionGroup1.DamagePoints).To(Equal(116))
		Expect(infectionGroup1.DamageType).To(Equal("bludgeoning"))
		Expect(infectionGroup2.Initiative).To(Equal(4))
	})

	It("can load groups without weakness or immunities", func() {
		s := q24.NewSystem(noBrackers)

		Expect(s.InfectionGroups[0].WeakTo).To(HaveLen(0))
		Expect(s.InfectionGroups[0].ImmuneTo).To(HaveLen(0))
	})

	It("can calculate effective power", func() {
		s := q24.NewSystem(ex)
		Expect(s.ImmuneGroups[0].EffectivePower()).To(Equal(17 * 4507))
	})

	It("can sort groups by effective power and initiative", func() {
		s := q24.NewSystem(ex)
		s.SortGroups()
		Expect(s.ImmuneGroups[0].Units).To(Equal(17))
		Expect(s.InfectionGroups[0].Units).To(Equal(801))
	})

	It("can calculate targets", func() {
		s := q24.NewSystem(ex)
		targets := s.CalcTargets()
		for att, def := range targets {
			if att.Units == 17 {
				Expect(def.Units).To(Equal(4485))
			}
			if att.Units == 989 {
				Expect(def.Units).To(Equal(801))
			}
			if att.Units == 801 {
				Expect(def.Units).To(Equal(17))
			}
			if att.Units == 4485 {
				Expect(def.Units).To(Equal(989))
			}
		}
	})

	It("can correctly do a wave of attacks", func() {
		s := q24.NewSystem(ex)
		s.Attack()
		for _, g := range s.ImmuneGroups {
			if g.Initiative == 2 {
				Expect(g.Units).To(Equal(0))
			}
			if g.Initiative == 3 {
				Expect(g.Units).To(Equal(905))
			}
		}

		for _, g := range s.InfectionGroups {
			if g.Initiative == 1 {
				Expect(g.Units).To(Equal(797))
			}
			if g.Initiative == 4 {
				Expect(g.Units).To(Equal(4434))
			}
		}
	})

	It("can go to the end", func() {
		s := q24.NewSystem(ex)
		s.EliminateEnemy()
		Expect(s.ImmuneUnits()).To(Equal(0))
		Expect(s.InfectionUnits()).To(Equal(5216))
	})

})
