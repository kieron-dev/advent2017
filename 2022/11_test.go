package two022_test

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Monkey struct {
	items   []int
	op      func(int) int
	test    func(int) bool
	ifTrue  int
	ifFalse int
	counter int
}

func (m *Monkey) Process(div3 bool) map[int][]int {
	res := map[int][]int{}

	for _, item := range m.items {
		w := m.op(item)
		if div3 {
			w /= 3
		}
		w %= 11 * 17 * 5 * 13 * 19 * 2 * 3 * 7
		if m.test(w) {
			res[m.ifTrue] = append(res[m.ifTrue], w)
		} else {
			res[m.ifFalse] = append(res[m.ifFalse], w)
		}
		m.counter++
	}
	m.items = nil

	return res
}

var _ = Describe("11", func() {
	It("does part A", func() {
		monkeys := loadMonkeys()

		for i := 0; i < 20; i++ {
			for _, monkey := range monkeys {
				moves := monkey.Process(true)
				for n, items := range moves {
					monkeys[n].items = append(monkeys[n].items, items...)
				}
			}
		}

		var counts []int
		for _, monkey := range monkeys {
			counts = append(counts, monkey.counter)
		}
		sort.Ints(counts)
		res := counts[len(counts)-2] * counts[len(counts)-1]
		Expect(res).To(Equal(54036))
	})

	It("does part B", func() {
		monkeys := loadMonkeys()

		for i := 0; i < 10000; i++ {
			for _, monkey := range monkeys {
				moves := monkey.Process(false)
				for n, items := range moves {
					monkeys[n].items = append(monkeys[n].items, items...)
				}
			}
		}

		var counts []int
		for _, monkey := range monkeys {
			counts = append(counts, monkey.counter)
		}
		sort.Ints(counts)
		res := counts[len(counts)-2] * counts[len(counts)-1]
		Expect(res).To(Equal(13237873355))
	})
})

func loadMonkeys() []*Monkey {
	f, err := os.Open("input11")
	Expect(err).NotTo(HaveOccurred())
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var monkeys []*Monkey
	var curMonkey *Monkey
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "Monkey"):
			curMonkey = &Monkey{}
			monkeys = append(monkeys, curMonkey)
		case strings.Contains(line, "Starting items:"):
			re := regexp.MustCompile(`(?:(\d+), )*?(\d+)`)
			for _, match := range re.FindAllStringSubmatch(line, -1) {
				curMonkey.items = append(curMonkey.items, stoi(match[0]))
			}
		case strings.Contains(line, "Operation:"):
			sop := "+"
			if strings.Contains(line, "*") {
				sop = "*"
			}
			bits := strings.Split(line, " ")
			if bits[len(bits)-1] == "old" {
				curMonkey.op = func(n int) int {
					return n * n
				}
				continue
			}

			arg := lastNumInLine(line)
			curMonkey.op = func(n int) int {
				if sop == "*" {
					return n * arg
				}
				return n + arg
			}
		case strings.Contains(line, "Test:"):
			arg := lastNumInLine(line)
			curMonkey.test = func(n int) bool {
				return n%arg == 0
			}
		case strings.Contains(line, "If true"):
			arg := lastNumInLine(line)
			curMonkey.ifTrue = arg
		case strings.Contains(line, "If false"):
			arg := lastNumInLine(line)
			curMonkey.ifFalse = arg
		}
	}

	return monkeys
}

func stoi(s string) int {
	n, err := strconv.Atoi(s)
	Expect(err).NotTo(HaveOccurred())
	return n
}

func lastNumInLine(line string) int {
	bits := strings.Split(line, " ")
	return stoi(bits[len(bits)-1])
}
