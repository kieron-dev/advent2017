package two022_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"sort"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("13", func() {
	It("does part A", func() {
		f, err := ioutil.ReadFile("input13")
		Expect(err).NotTo(HaveOccurred())

		var lines [][]any
		for _, l := range bytes.Fields(f) {
			var line []any
			err := json.Unmarshal(l, &line)
			Expect(err).NotTo(HaveOccurred())
			lines = append(lines, line)
		}

		var sum int
		for i := 0; i < len(lines); i += 2 {
			left := lines[i]
			right := lines[i+1]

			if compare(left, right) < 0 {
				sum += 1 + i/2
			}
		}

		Expect(sum).To(Equal(4894))
	})

	It("does part B", func() {
		f, err := ioutil.ReadFile("input13")
		Expect(err).NotTo(HaveOccurred())

		var lines [][]any
		for _, l := range bytes.Fields(f) {
			var line []any
			err := json.Unmarshal(l, &line)
			Expect(err).NotTo(HaveOccurred())
			lines = append(lines, line)
		}

		lines = append(lines, []any{[]any{2.0}}, []any{[]any{6.0}})

		sort.Slice(lines, func(a, b int) bool {
			return compare(lines[a], lines[b]) < 0
		})

		var posA, posB int
		for i := range lines {
			bs, err := json.Marshal(lines[i])
			Expect(err).NotTo(HaveOccurred())
			str := string(bs)
			if str == "[[2]]" {
				posA = i + 1
			}
			if str == "[[6]]" {
				posB = i + 1
			}
		}
		Expect(posA * posB).To(Equal(24180))
	})
})

func compare(l, r []any) int {
	for i := range l {
		if i == len(r) {
			return 1
		}
		aList, aIsList := l[i].([]any)
		bList, bIsList := r[i].([]any)

		if aIsList && bIsList {
			c := compare(aList, bList)
			if c == 0 {
				continue
			}
			return c
		}

		if !aIsList && !bIsList {
			a := int(l[i].(float64))
			b := int(r[i].(float64))
			if a == b {
				continue
			}
			return a - b
		}

		var c int
		if !aIsList {
			c = compare([]any{l[i]}, bList)
		} else {
			c = compare(aList, []any{r[i]})
		}
		if c == 0 {
			continue
		}
		return c
	}

	if len(l) < len(r) {
		return -1
	}

	return 0
}
