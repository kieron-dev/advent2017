package two023_test

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type mapping struct {
	source int
	target int
	span   int
}

func (m mapping) translate(from int) (int, bool) {
	if from >= m.source && from < m.source+m.span {
		return m.target + from - m.source, true
	}

	return from, false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (m mapping) translateRange(start, length int) (remaining [][2]int, mapped [2]int) {
	if start < m.source {
		remainingLen := min(m.source-start, length)
		remaining = append(remaining, [2]int{start, remainingLen})
		if remainingLen == length {
			return
		}
		start += remainingLen
		length -= remainingLen
	}

	if start < m.source+m.span {
		mappedLen := min(m.source+m.span-start, length)
		mapped = [2]int{m.target + start - m.source, mappedLen}
		if mappedLen == length {
			return
		}
		start += mappedLen
		length -= mappedLen
	}

	remaining = append(remaining, [2]int{start, length})

	return
}

type mappings struct {
	from string
	to   string
	maps []mapping
}

func (m mappings) getVal(n int) int {
	for _, ms := range m.maps {
		v, ok := ms.translate(n)
		if ok {
			return v
		}
	}
	return n
}

func (m mappings) getValRange(start [2]int) [][2]int {
	mapped := [][2]int{}

	remaining := [][2]int{start}

	for _, ms := range m.maps {
		newRemaining := [][2]int{}
		for _, pair := range remaining {
			r, m := ms.translateRange(pair[0], pair[1])
			if m[1] > 0 {
				mapped = append(mapped, m)
			}
			for _, b := range r {
				if b[1] > 0 {
					newRemaining = append(newRemaining, b)
				}
			}
		}
		remaining = newRemaining
	}

	return append(remaining, mapped...)
}

var _ = DescribeTable("translate", func(source, target, span, input, expected int, expectedOk bool) {
	m := mapping{source, target, span}

	t, b := m.translate(input)
	Expect(t).To(Equal(expected))
	Expect(b).To(Equal(expectedOk))
},
	Entry("50", 50, 98, 2, 50, 98, true),
	Entry("49", 50, 98, 2, 49, 49, false),
	Entry("51", 50, 98, 2, 51, 99, true),
	Entry("52", 50, 98, 2, 52, 52, false),
)

var _ = DescribeTable("translate range", func(source, target, span, inStart, inSpan int, expectedRemaining [][2]int, expectedMapped [2]int) {
	m := mapping{source, target, span}
	remaining, mapped := m.translateRange(inStart, inSpan)
	Expect(remaining).To(Equal(expectedRemaining))
	Expect(mapped).To(Equal(expectedMapped))
},
	Entry("all before", 50, 98, 2, 48, 2, [][2]int{{48, 2}}, nil),
	Entry("all after", 50, 98, 2, 52, 3, [][2]int{{52, 3}}, nil),
	Entry("all in", 50, 98, 10, 52, 3, nil, [2]int{100, 3}),
	Entry("around", 50, 98, 10, 48, 20, [][2]int{{48, 2}, {60, 8}}, [2]int{98, 10}),
)

func loadMappingses(fileName string) ([]int, map[string]mappings) {
	f, err := os.Open(fileName)
	Expect(err).NotTo(HaveOccurred())
	defer f.Close()

	scanner := bufio.NewScanner(f)
	state := "seeds"
	var seeds []int
	mappingses := map[string]mappings{}
	var ms mappings

	for scanner.Scan() {
		line := scanner.Text()
		switch state {
		case "seeds":
			seeds = alisttoi(strings.Split(line, ":")[1])
			state = "maps"
			continue
		case "maps":
			if line == "" {
				if len(ms.maps) > 0 {
					mappingses[ms.from] = ms
				}
				ms = mappings{}
				continue
			}
			if strings.Contains(line, ":") {
				parts := strings.Split(line, " ")
				parts2 := strings.Split(parts[0], "-")
				ms.from = parts2[0]
				ms.to = parts2[2]
				continue
			}
			nums := alisttoi(line)
			Expect(nums).To(HaveLen(3))
			ms.maps = append(ms.maps, mapping{nums[1], nums[0], nums[2]})
		}
	}
	if len(ms.maps) > 0 {
		mappingses[ms.from] = ms
	}

	return seeds, mappingses
}

var _ = Describe("05", func() {
	It("does part A", func() {
		seeds, mappingses := loadMappingses("input05")

		min := 99999999999
		for _, seed := range seeds {
			mappings := mappingses["seed"]
			v := seed
			for {
				nextVal := mappings.getVal(v)
				nextType := mappings.to
				if nextType == "location" {
					if nextVal < min {
						min = nextVal
					}
					break
				}
				mappings = mappingses[nextType]
				v = nextVal
			}
		}

		Expect(min).To(Equal(214922730))
	})

	It("does part B", func() {
		seeds, mappingses := loadMappingses("input05")

		var ranges [][2]int
		for i := 0; i < len(seeds); i += 2 {
			ranges = append(ranges, [2]int{seeds[i], seeds[i+1]})
		}

		mappings := mappingses["seed"]
		for {
			nextRanges := [][2]int{}
			for _, r := range ranges {
				nextRanges = append(nextRanges, mappings.getValRange(r)...)
			}
			nextType := mappings.to
			mappings = mappingses[nextType]
			ranges = nextRanges
			if nextType == "location" {
				break
			}
		}

		min := 99999999999
		for _, r := range ranges {
			if r[0] < min {
				min = r[0]
			}
		}

		Expect(min).To(Equal(148041808))
	})
})

func alisttoi(s string) []int {
	val := []int{}
	for _, n := range strings.Fields(s) {
		i, err := strconv.Atoi(n)
		Expect(err).NotTo(HaveOccurred())
		val = append(val, i)
	}
	return val
}
