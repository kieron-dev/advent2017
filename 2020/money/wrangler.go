package money

import (
	"bufio"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Wrangler struct {
	items []int
}

func NewWrangler() Wrangler {
	return Wrangler{}
}

func (w *Wrangler) LoadExpenses(data io.Reader) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		n, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			log.Fatalf("conv-int: %v", err)
		}
		w.items = append(w.items, n)
	}

	sort.Ints(w.items)
}

func (w Wrangler) GetSummingTo(sum int) []int {
	for i := 0; i < len(w.items)-1; i++ {
		for j := i + 1; j < len(w.items); j++ {
			s := w.items[i] + w.items[j]

			if s > sum {
				break
			}

			if s == sum {
				return []int{w.items[i], w.items[j]}
			}
		}
	}

	return []int{}
}

func (w Wrangler) Get3SummingTo(sum int) []int {
	for i := 0; i < len(w.items)-2; i++ {
		if w.items[i] > sum {
			break
		}
		for j := i + 1; j < len(w.items)-1; j++ {
			if w.items[i]+w.items[j] > sum {
				break
			}
			for k := j + 1; k < len(w.items); k++ {
				s := w.items[i] + w.items[j] + w.items[k]

				if s > sum {
					break
				}

				if s == sum {
					return []int{w.items[i], w.items[j], w.items[k]}
				}
			}
		}
	}
	return []int{}
}

func (w Wrangler) ProductFor(sum int) int {
	sumTo := w.GetSummingTo(sum)

	return sumTo[0] * sumTo[1]
}

func (w Wrangler) Product3For(sum int) int {
	sumTo := w.Get3SummingTo(sum)

	return sumTo[0] * sumTo[1] * sumTo[2]
}
