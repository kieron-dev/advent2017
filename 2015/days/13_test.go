package days_test

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type diner struct {
	name        string
	preferences map[string]int
}

func newDiner(name string) *diner {
	return &diner{
		name:        name,
		preferences: map[string]int{},
	}
}

var _ = Describe("13", func() {
	It("does part A", func() {
		input, err := os.Open("input13")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		diners := map[string]*diner{}
		re := regexp.MustCompile(`^(\w+) would (\w+) (\w+) happiness .* (\w+).$`)

		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			matches := re.FindStringSubmatch(line)
			Expect(matches).ToNot(BeNil())
			name := matches[1]
			if diners[name] == nil {
				diners[name] = newDiner(name)
			}
			mult := 1
			if matches[2] == "lose" {
				mult = -1
			}

			num, err := strconv.Atoi(matches[3])
			Expect(err).NotTo(HaveOccurred())

			target := matches[4]
			diners[name].preferences[target] = mult * num
		}

		names := []string{}
		for n := range diners {
			names = append(names, n)
		}

		maxScore := -10000
		for order := range permute(names[1:]) {
			n := score(append([]string{names[0]}, order...), diners)
			if n > maxScore {
				maxScore = n
			}
		}

		Expect(maxScore).To(Equal(733))
	})

	It("does part B", func() {
		input, err := os.Open("input13")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		diners := map[string]*diner{"me": newDiner("me")}
		re := regexp.MustCompile(`^(\w+) would (\w+) (\w+) happiness .* (\w+).$`)

		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			matches := re.FindStringSubmatch(line)
			Expect(matches).ToNot(BeNil())
			name := matches[1]
			if diners[name] == nil {
				diners[name] = newDiner(name)
				diners[name].preferences["me"] = 0
				diners["me"].preferences[name] = 0
			}
			mult := 1
			if matches[2] == "lose" {
				mult = -1
			}

			num, err := strconv.Atoi(matches[3])
			Expect(err).NotTo(HaveOccurred())

			target := matches[4]
			diners[name].preferences[target] = mult * num
		}

		names := []string{}
		for n := range diners {
			names = append(names, n)
		}

		maxScore := -10000
		for order := range permute(names[1:]) {
			n := score(append([]string{names[0]}, order...), diners)
			if n > maxScore {
				maxScore = n
			}
		}

		Expect(maxScore).To(Equal(725))
	})
})

func score(order []string, diners map[string]*diner) int {
	n := diners[order[len(order)-1]].preferences[order[0]]
	n += diners[order[0]].preferences[order[len(order)-1]]

	for i := 0; i < len(order)-1; i++ {
		n += diners[order[i]].preferences[order[i+1]]
		n += diners[order[i+1]].preferences[order[i]]
	}

	return n
}

func permute(names []string) chan ([]string) {
	ch := make(chan ([]string), 1024)
	go func() {
		perm(len(names), names, ch)
		close(ch)
	}()

	return ch
}

func perm(k int, names []string, ch chan ([]string)) {
	if k == 1 {
		namesCopy := make([]string, len(names))
		copy(namesCopy, names)
		ch <- namesCopy
		return
	}

	perm(k-1, names, ch)

	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			names[i], names[k-1] = names[k-1], names[i]
		} else {
			names[0], names[k-1] = names[k-1], names[0]
		}
		perm(k-1, names, ch)
	}
}
