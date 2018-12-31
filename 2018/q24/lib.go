package q24

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Group struct {
	Units        int
	HitPoints    int
	WeakTo       []string
	ImmuneTo     []string
	DamageType   string
	DamagePoints int
	Initiative   int
}

func (g *Group) EffectivePower() int {
	units := g.Units
	if units < 0 {
		units = 0
	}
	return units * g.DamagePoints
}

func (g *Group) DamageBy(damagePoints int, damageType string) int {
	if g.IsWeakTo(damageType) {
		return damagePoints * 2
	}
	if g.IsImmuneTo(damageType) {
		return 0
	}
	return damagePoints
}

func (g *Group) Attack(damagePoints int, damageType string) {
	if g.IsImmuneTo(damageType) {
		return
	}
	hitPoints := g.HitPoints
	if g.IsWeakTo(damageType) {
		damagePoints *= 2
	}
	unitsHit := damagePoints / hitPoints
	newUnits := g.Units - unitsHit
	if newUnits < 0 {
		newUnits = 0
	}
	g.Units = newUnits
}

func (g *Group) IsImmuneTo(damageType string) bool {
	for _, immune := range g.ImmuneTo {
		if immune == damageType {
			return true
		}
	}
	return false
}

func (g *Group) IsWeakTo(damageType string) bool {
	for _, weak := range g.WeakTo {
		if weak == damageType {
			return true
		}
	}
	return false
}

type System struct {
	ImmuneGroups    []*Group
	InfectionGroups []*Group
}

func NewSystem(in io.Reader) *System {
	s := System{}
	s.ImmuneGroups = []*Group{}
	s.InfectionGroups = []*Group{}

	scanner := bufio.NewScanner(in)
	var currentSet *[]*Group

	mainRE := regexp.MustCompile(`(\d+) units each with (\d+) hit points (\(.*\) )?with an attack that does (\d+) (\w+) damage at initiative (\d+)`)
	weakRE := regexp.MustCompile(`\(.*weak to ([^;)]+).*`)
	immuneRE := regexp.MustCompile(`\(.*immune to ([^;)]+).*`)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, "\n")

		switch line {
		case "Immune System:":
			currentSet = &s.ImmuneGroups
		case "Infection:":
			currentSet = &s.InfectionGroups
		case "":
		default:
			g := Group{}
			matches := mainRE.FindStringSubmatch(line)

			g.Units = atoi(matches[1])
			g.HitPoints = atoi(matches[2])
			g.DamagePoints = atoi(matches[4])
			g.DamageType = matches[5]
			g.Initiative = atoi(matches[6])

			if matches[3] != "" {
				weakMatches := weakRE.FindStringSubmatch(matches[3])
				if len(weakMatches) == 2 {
					g.WeakTo = strings.Split(weakMatches[1], ", ")
				}
				immuneMatches := immuneRE.FindStringSubmatch(matches[3])
				if len(immuneMatches) == 2 {
					g.ImmuneTo = strings.Split(immuneMatches[1], ", ")
				}
			}

			*currentSet = append(*currentSet, &g)
		}
	}

	return &s
}

func (s *System) BoostImmuneSystem(boost int) {
	for _, immuneGroup := range s.ImmuneGroups {
		immuneGroup.DamagePoints += boost
	}
}

func (s *System) CalcTargets() map[*Group]*Group {
	s.SortGroups()
	selections := map[*Group]*Group{}
	infectionSelected := map[*Group]bool{}
	immuneSelected := map[*Group]bool{}

	for _, immuneGroup := range s.ImmuneGroups {
		if immuneGroup.Units == 0 {
			continue
		}
		maxDamage := 0
		var maxDamageGroup *Group
		effectivePower := immuneGroup.EffectivePower()

		for _, infectionGroup := range s.InfectionGroups {
			if infectionSelected[infectionGroup] || infectionGroup.Units == 0 {
				continue
			}
			possibleDamage := infectionGroup.DamageBy(effectivePower, immuneGroup.DamageType)
			if possibleDamage > maxDamage {
				maxDamage = possibleDamage
				maxDamageGroup = infectionGroup
			}
		}

		if maxDamage > 0 {
			selections[immuneGroup] = maxDamageGroup
			infectionSelected[maxDamageGroup] = true
		}
	}

	for _, infectionGroup := range s.InfectionGroups {
		if infectionGroup.Units == 0 {
			continue
		}
		maxDamage := 0
		var maxDamageGroup *Group
		effectivePower := infectionGroup.EffectivePower()

		for _, immuneGroup := range s.ImmuneGroups {
			if immuneSelected[immuneGroup] || immuneGroup.Units == 0 {
				continue
			}
			possibleDamage := immuneGroup.DamageBy(effectivePower, infectionGroup.DamageType)
			if possibleDamage > maxDamage {
				maxDamage = possibleDamage
				maxDamageGroup = immuneGroup
			}
		}

		if maxDamage > 0 {
			selections[infectionGroup] = maxDamageGroup
			immuneSelected[maxDamageGroup] = true
		}
	}

	return selections
}

func (s *System) SortGroups() {
	ordering := func(groups []*Group) func(i, j int) bool {
		return func(i, j int) bool {
			a := groups[i]
			b := groups[j]
			if a.EffectivePower() == b.EffectivePower() {
				return a.Initiative > b.Initiative
			}
			return a.EffectivePower() > b.EffectivePower()
		}
	}
	sort.Slice(s.ImmuneGroups, ordering(s.ImmuneGroups))
	sort.Slice(s.InfectionGroups, ordering(s.InfectionGroups))
}

func (s *System) Attack() {
	targets := s.CalcTargets()
	attackers := []*Group{}
	for a, _ := range targets {
		attackers = append(attackers, a)
	}

	sort.Slice(attackers, func(i, j int) bool {
		return attackers[i].Initiative > attackers[j].Initiative
	})

	for _, a := range attackers {
		t := targets[a]
		t.Attack(a.EffectivePower(), a.DamageType)
	}
}

func (s *System) ImmuneUnits() int {
	c := 0
	for _, g := range s.ImmuneGroups {
		c += g.Units
	}
	return c
}

func (s *System) InfectionUnits() int {
	c := 0
	for _, g := range s.InfectionGroups {
		c += g.Units
	}
	return c
}

func (s *System) EliminateEnemy() {
	prevImmune := s.ImmuneUnits()
	prevInfection := s.InfectionUnits()
	for prevImmune > 0 && prevInfection > 0 {
		s.Attack()
		immune := s.ImmuneUnits()
		infection := s.InfectionUnits()
		if immune == prevImmune && infection == prevInfection {
			fmt.Println("Stalemate")
			break
		}
		prevImmune, prevInfection = immune, infection
	}
}

func (s *System) PrintGroups() {
	fmt.Println("Immune groups:")
	for _, g := range s.ImmuneGroups {
		fmt.Printf("Units %d, Power %d, Initiative %d\n", g.Units, g.EffectivePower(), g.Initiative)
	}
	fmt.Println("\nInfection groups:")
	for _, g := range s.InfectionGroups {
		fmt.Printf("Units %d, Power %d, Initiative %d\n", g.Units, g.EffectivePower(), g.Initiative)
	}
}

func atoi(ascii string) int {
	res, err := strconv.Atoi(ascii)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
