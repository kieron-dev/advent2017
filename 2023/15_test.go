package two023_test

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func hash(s string) int {
	v := 0
	for i := range s {
		b := s[i]
		v += int(b)
		v *= 17
		v %= 256
	}

	return v
}

type lens struct {
	prev, next *lens
	label      string
	power      int
}

func (l *lens) print() {
	cur := l
	for cur != nil {
		fmt.Printf("[%s %d] ", cur.label, cur.power)
		cur = cur.next
	}
}

func (l *lens) len() int {
	if l == nil {
		return 0
	}
	r := 0
	cur := l
	for cur != nil {
		r++
		cur = cur.next
	}
	return r
}

func (l *lens) remove(label string) *lens {
	if l == nil {
		return nil
	}

	if l.label == label {
		if l.next != nil {
			l.next.prev = nil
		}
		return l.next
	}

	cur := l
	for cur != nil {
		if cur.label == label {
			if cur.prev != nil {
				cur.prev.next = cur.next
				if cur.next != nil {
					cur.next.prev = cur.prev
				}
			}
			return l
		}
		cur = cur.next
	}

	return l
}

func (l *lens) add(label string, power int) *lens {
	if l == nil {
		return &lens{label: label, power: power}
	}

	cur := l
	var prev *lens
	for cur != nil {
		if cur.label == label {
			cur.power = power
			return l
		}
		prev = cur
		cur = cur.next
	}
	prev.next = &lens{prev: prev, label: label, power: power}
	return l
}

func (l *lens) lensPower() int {
	cur := l
	power := 0
	i := 1
	for cur != nil {
		power += i * cur.power
		i++
		cur = cur.next
	}
	return power
}

func loadStrings(filename string) []string {
	bs, err := os.ReadFile(filename)
	Expect(err).NotTo(HaveOccurred())
	return strings.Split(strings.TrimSpace(string(bs)), ",")
}

var _ = Describe("15", func() {
	It("does part A", func() {
		sum := 0
		for _, s := range loadStrings("input15") {
			sum += hash(s)
		}

		Expect(sum).To(Equal(509167))
	})

	It("does part B", func() {
		boxes := make([]*lens, 256)

		for _, s := range loadStrings("input15") {
			if strings.Contains(s, "-") {
				label := strings.Split(s, "-")[0]
				box := hash(label)
				boxes[box] = boxes[box].remove(label)
				continue
			}

			parts := strings.Split(s, "=")
			box := hash(parts[0])
			n, err := strconv.Atoi(parts[1])
			Expect(err).NotTo(HaveOccurred())
			boxes[box] = boxes[box].add(parts[0], n)
		}

		sum := 0
		for i, b := range boxes {
			if b != nil {
				sum += (i + 1) * b.lensPower()
			}
		}

		Expect(sum).To(Equal(259333))
	})
})

func printBoxes(boxes []*lens) {
	for i, b := range boxes {
		if b != nil {
			fmt.Printf("Box %3d: ", i)
			b.print()
			fmt.Println()
		}
	}
}
