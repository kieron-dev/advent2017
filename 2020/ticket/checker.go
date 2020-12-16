// Package ticket validates strange train tickets
package ticket

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type valRange struct {
	from int
	to   int
}

func (v valRange) contains(n int) bool {
	return n >= v.from && n <= v.to
}

type fieldRule struct {
	field  string
	range1 valRange
	range2 valRange
}

func (r fieldRule) validate(n int) bool {
	return r.range1.contains(n) || r.range2.contains(n)
}

type Checker struct {
	fieldRules    []fieldRule
	nearbyTickets [][]int
	yourTicket    []int
}

func NewChecker() Checker {
	return Checker{}
}

type state int

const (
	_ state = iota
	rules
	yourTicket
	nearbyTickets
)

func (c *Checker) Load(data io.Reader) {
	state := rules

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if line == "your ticket:" {
			state = yourTicket
			continue
		}

		if line == "nearby tickets:" {
			state = nearbyTickets
			continue
		}

		switch state {
		case rules:
			c.AddRule(line)
		case yourTicket:
			c.AddYourTicket(line)
		case nearbyTickets:
			c.AddNearbyTicket(line)
		default:
			log.Fatalf("shouldn't be here: %v", state)
		}
	}
}

var ruleRE = regexp.MustCompile(`(.*): (\d+)-(\d+) or (\d+)-(\d+)`)

func (c Checker) RuleCount() int {
	return len(c.fieldRules)
}

func (c Checker) NearbyCount() int {
	return len(c.nearbyTickets)
}

func (c Checker) ErrorRate() int {
	rate := 0

	for _, ticket := range c.nearbyTickets {
	ticketValue:
		for _, val := range ticket {
			for _, rule := range c.fieldRules {
				if rule.validate(val) {
					continue ticketValue
				}
			}
			rate += val
		}
	}

	return rate
}

func (c *Checker) DepartureProduct() int {
	prod := 1

	for field, idx := range c.FieldPositions() {
		if strings.HasPrefix(field, "departure") {
			prod *= c.yourTicket[idx]
		}
	}

	return prod
}

func (c *Checker) FieldPositions() map[string]int {
	// map of ruleIdx => map of fieldIdx => validPossibility
	possibilities := map[int]map[int]bool{}

	for i := range c.fieldRules {
		possibilities[i] = map[int]bool{}
		for j := 0; j < len(c.fieldRules); j++ {
			possibilities[i][j] = true
		}
	}

	c.removeBadTickets()

	for {
		changed := false

		for ruleIdx, possMap := range possibilities {
			for fieldIdx := range possMap {
				for _, ticket := range c.nearbyTickets {
					if !c.fieldRules[ruleIdx].validate(ticket[fieldIdx]) {
						delete(possMap, fieldIdx)
						changed = true
						break
					}
				}
			}
		}

		for ruleIdx, possMap := range possibilities {
			if len(possMap) == 1 {
				var last int
				for k := range possMap {
					last = k
				}

				for j, pm := range possibilities {
					if j != ruleIdx && pm[last] {
						delete(pm, last)
						changed = true
					}
				}
			}
		}

		if !changed {
			break
		}
	}

	res := map[string]int{}

	for ruleIdx, possMap := range possibilities {
		if len(possMap) != 1 {
			log.Fatalf("more than 1 entry in map: %v", possMap)
		}

		var onlyPossibility int
		for k := range possMap {
			onlyPossibility = k
		}

		res[c.fieldRules[ruleIdx].field] = onlyPossibility
	}

	return res
}

func (c *Checker) removeBadTickets() {
	toRemove := []int{}

	for i, ticket := range c.nearbyTickets {
		for _, n := range ticket {
			if c.notValidForAnyField(n) {
				toRemove = append(toRemove, i)
				break
			}
		}
	}

	newNearbyTickets := c.nearbyTickets[:]
	for i, n := range toRemove {
		newNearbyTickets = append(newNearbyTickets[:n-i], newNearbyTickets[n-i+1:]...)
	}

	c.nearbyTickets = newNearbyTickets
}

func (c Checker) notValidForAnyField(n int) bool {
	for _, rule := range c.fieldRules {
		if rule.validate(n) {
			return false
		}
	}

	return true
}

func (c *Checker) AddRule(line string) {
	matches := ruleRE.FindStringSubmatch(line)
	if len(matches) != 6 {
		log.Fatalf("rule regex failed for %q", line)
	}
	name := matches[1]
	from1 := atoi(matches[2])
	to1 := atoi(matches[3])
	from2 := atoi(matches[4])
	to2 := atoi(matches[5])

	c.fieldRules = append(c.fieldRules, fieldRule{
		field: name,
		range1: valRange{
			from: from1,
			to:   to1,
		},
		range2: valRange{
			from: from2,
			to:   to2,
		},
	})
}

func atoi(a string) int {
	n, err := strconv.Atoi(a)
	if err != nil {
		log.Fatalf("string to int failed: %v", err)
	}

	return n
}

func (c *Checker) AddYourTicket(line string) {
	c.yourTicket = c.parseTicket(line)
}

func (c Checker) parseTicket(line string) []int {
	items := strings.Split(line, ",")

	if len(items) != len(c.fieldRules) {
		log.Fatalf("wrong number of ticket fields in %q", line)
	}

	ticket := []int{}

	for _, f := range items {
		n := atoi(f)
		ticket = append(ticket, n)
	}

	return ticket
}

func (c *Checker) AddNearbyTicket(line string) {
	items := strings.Split(line, ",")

	if len(items) != len(c.fieldRules) {
		log.Fatalf("wrong number of ticket fields in %q", line)
	}

	ticket := []int{}

	for _, f := range items {
		n := atoi(f)
		ticket = append(ticket, n)
	}

	c.nearbyTickets = append(c.nearbyTickets, ticket)
}
