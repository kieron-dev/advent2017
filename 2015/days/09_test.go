package days_test

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Town struct {
	name      string
	distances map[*Town]int
}

var _ = Describe("09", func() {
	It("does part A & B", func() {
		file, err := os.Open("input09")
		Expect(err).NotTo(HaveOccurred())
		defer file.Close()

		towns := map[string]*Town{}

		scanner := bufio.NewScanner(file)
		re := regexp.MustCompile(`^(.*) to (.*) = (.*)$`)

		for scanner.Scan() {
			line := scanner.Text()

			matches := re.FindStringSubmatch(line)
			Expect(matches).ToNot(BeNil())

			from, ok := towns[matches[1]]
			if !ok {
				from = &Town{
					name:      matches[1],
					distances: map[*Town]int{},
				}
				towns[matches[1]] = from
			}
			to, ok := towns[matches[2]]
			if !ok {
				to = &Town{
					name:      matches[2],
					distances: map[*Town]int{},
				}
				towns[matches[2]] = to
			}
			dist, err := strconv.Atoi(matches[3])
			Expect(err).NotTo(HaveOccurred())

			from.distances[to] = dist
			to.distances[from] = dist
		}

		townSlice := []*Town{}
		for _, town := range towns {
			townSlice = append(townSlice, town)
		}

		ch := permutations(townSlice)
		shortest := 1000000
		longest := 0
		for perm := range ch {
			dist := 0
			for i := 0; i < len(perm)-1; i++ {
				dist += perm[i].distances[perm[i+1]]
			}

			if dist < shortest {
				shortest = dist
			}

			if dist > longest {
				longest = dist
			}
		}

		Expect(shortest).To(Equal(207))
		Expect(longest).To(Equal(804))
	})
})

func permutations(towns []*Town) chan ([]*Town) {
	ch := make(chan ([]*Town), 1024)
	go func() {
		generate(len(towns), towns, ch)
		close(ch)
	}()

	return ch
}

func generate(k int, towns []*Town, ch chan ([]*Town)) {
	if k == 1 {
		townsCopy := make([]*Town, len(towns))
		copy(townsCopy, towns)
		ch <- townsCopy
		return
	}

	generate(k-1, towns, ch)

	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			towns[i], towns[k-1] = towns[k-1], towns[i]
		} else {
			towns[0], towns[k-1] = towns[k-1], towns[0]
		}
		generate(k-1, towns, ch)
	}
}
