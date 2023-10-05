package two022_test

import (
	"bytes"
	"fmt"
	"os"

	lru "github.com/hashicorp/golang-lru/v2"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type (
	blueprint struct {
		index int
		costs map[mineral]map[mineral]int
	}

	state struct {
		stock  [3]int
		robots [3]int
		time   int
		score  int
	}

	mineral int
)

const (
	ore mineral = iota
	clay
	obsidian
	geode
)

func loadBlueprints() []blueprint {
	bs, err := os.ReadFile("input19")
	Expect(err).NotTo(HaveOccurred())

	var idx, oreOre, clayOre, obsidianOre, obsidianClay, geodeOre, geodeObsidian int

	var blueprints []blueprint

	for _, line := range bytes.Split(bs, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		_, err := fmt.Sscanf(string(line), "Blueprint %d: Each ore robot costs %d ore. "+
			"Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. "+
			"Each geode robot costs %d ore and %d obsidian.",
			&idx, &oreOre, &clayOre, &obsidianOre, &obsidianClay, &geodeOre, &geodeObsidian)
		Expect(err).NotTo(HaveOccurred())

		blueprints = append(blueprints, blueprint{
			index: idx,
			costs: map[mineral]map[mineral]int{
				ore:      {ore: oreOre},
				clay:     {ore: clayOre},
				obsidian: {ore: obsidianOre, clay: obsidianClay},
				geode:    {ore: geodeOre, obsidian: geodeObsidian},
			},
		})
	}
	return blueprints
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (s state) timeFor(m mineral, amount int) int {
	t := amount / s.robots[m]
	if amount%s.robots[m] > 0 {
		t++
	}
	return t
}

func (s state) next(m mineral, bp blueprint, deadline int, curMax *int) (state, bool) {
	t := 0
	for _, c := range []mineral{ore, clay, obsidian} {
		if bp.costs[m][c] == 0 {
			continue
		}
		if bp.costs[m][c] > 0 && s.robots[c] == 0 {
			return state{}, false
		}
		t = max(t, s.timeFor(c, bp.costs[m][c]-s.stock[c]))
	}
	if s.time+t > deadline {
		return state{}, false
	}

	next := s
	next.advance(t + 1)
	for _, c := range []mineral{ore, clay, obsidian} {
		next.stock[c] -= bp.costs[m][c]
	}

	if m == geode {
		next.score += deadline - next.time
	} else {
		next.robots[m]++
	}

	remaining := deadline - (s.time + t)
	poss := (1 + remaining) * remaining / 2
	if next.score+poss < *curMax {
		return state{}, false
	}

	return next, true
}

func (s state) maxGeodes(bp blueprint, deadline int, cache *lru.Cache[state, int], curMax *int) int {
	if score, ok := cache.Get(s); ok {
		return score
	}

	maxGeodes := s.score

	for _, c := range []mineral{geode, obsidian, clay, ore} {
		next, ok := s.next(c, bp, deadline, curMax)
		if !ok {
			continue
		}
		score := next.maxGeodes(bp, deadline, cache, curMax)
		if score > maxGeodes {
			maxGeodes = score
		}
	}

	cache.Add(s, maxGeodes)

	if maxGeodes > *curMax {
		*curMax = maxGeodes
	}

	return maxGeodes
}

func (s *state) advance(mins int) {
	for _, c := range []mineral{ore, clay, obsidian} {
		s.stock[c] += s.robots[c] * mins
	}
	s.time += mins
}

var _ = Describe("19", func() {
	var cache *lru.Cache[state, int]

	BeforeEach(func() {
		var err error
		cache, err = lru.New[state, int](2 * 1024 * 1024)
		Expect(err).NotTo(HaveOccurred())
	})

	It("loads the file", func() {
		blueprints := loadBlueprints()
		Expect(blueprints).To(HaveLen(30))
	})

	It("does example 1", func() {
		bp := blueprint{
			index: 1,
			costs: map[mineral]map[mineral]int{
				ore:      {ore: 4},
				clay:     {ore: 2},
				obsidian: {ore: 3, clay: 14},
				geode:    {ore: 2, obsidian: 7},
			},
		}
		s := state{}
		s.robots[ore] = 1
		curMax := 0

		Expect(s.maxGeodes(bp, 24, cache, &curMax)).To(Equal(9))
	})

	It("does example 2", func() {
		bp := blueprint{
			index: 1,
			costs: map[mineral]map[mineral]int{
				ore:      {ore: 2},
				clay:     {ore: 3},
				obsidian: {ore: 3, clay: 8},
				geode:    {ore: 3, obsidian: 12},
			},
		}
		s := state{}
		s.robots[ore] = 1
		curMax := 0

		Expect(s.maxGeodes(bp, 24, cache, &curMax)).To(Equal(12))
	})

	It("does part A", func() {
		blueprints := loadBlueprints()
		sum := 0

		for _, bp := range blueprints {
			var err error
			cache, err = lru.New[state, int](2 * 1024 * 1024)
			Expect(err).NotTo(HaveOccurred())
			s := state{}
			s.robots[ore] = 1
			curMax := 0
			sum += bp.index * s.maxGeodes(bp, 24, cache, &curMax)
			fmt.Printf("%2d: sum = %+v\n", bp.index, sum)
		}

		Expect(sum).To(Equal(1719))
	})

	It("does part B", func() {
		blueprints := loadBlueprints()
		prod := 1

		for i := 0; i < 3; i++ {
			bp := blueprints[i]
			var err error
			cache, err = lru.New[state, int](2 * 1024 * 1024)
			Expect(err).NotTo(HaveOccurred())
			s := state{}
			s.robots[ore] = 1
			curMax := 0
			prod *= s.maxGeodes(bp, 32, cache, &curMax)
			fmt.Printf("prod = %+v\n", prod)
		}

		Expect(prod).To(Equal(19530))
	})
})
