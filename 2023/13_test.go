package two023_test

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type volco struct {
	terrain [][]byte
}

func (v volco) print() {
	fmt.Println()
	for _, l := range v.terrain {
		fmt.Println(string(l))
	}
	fmt.Println()
}

func (v volco) horizMirror(ignore int) (int, bool) {
	for r := 1; r < len(v.terrain); r++ {
		if r == ignore {
			continue
		}
		if v.isHorizSymm(r) {
			return r, true
		}
	}
	return 0, false
}

func (v volco) isHorizSymm(r int) bool {
	for t, b := r-1, r; t >= 0 && b < len(v.terrain); t, b = t-1, b+1 {
		for c := 0; c < len(v.terrain[0]); c++ {
			if v.terrain[t][c] != v.terrain[b][c] {
				return false
			}
		}
	}
	return true
}

func (v volco) vertMirror(ignore int) (int, bool) {
	for c := 1; c < len(v.terrain[0]); c++ {
		if c == ignore {
			continue
		}
		if v.isVertSymm(c) {
			return c, true
		}
	}
	return 0, false
}

func (v volco) isVertSymm(c int) bool {
	for l, r := c-1, c; l >= 0 && r < len(v.terrain[0]); l, r = l-1, r+1 {
		for y := 0; y < len(v.terrain); y++ {
			if v.terrain[y][l] != v.terrain[y][r] {
				return false
			}
		}
	}

	return true
}

func (v volco) newScore() int {
	oldScore, ok := v.score(-1)
	Expect(ok).To(BeTrue())

	for r := 0; r < len(v.terrain); r++ {
		for c := 0; c < len(v.terrain[0]); c++ {
			oldVal := v.terrain[r][c]
			var newVal byte = '#'
			if oldVal == '#' {
				newVal = '.'
			}
			v.terrain[r][c] = newVal
			if s, ok := v.score(oldScore); ok {
				return s
			}
			v.terrain[r][c] = oldVal
		}
	}

	v.print()
	Fail("oops")
	return 0
}

func (v volco) score(ignore int) (int, bool) {
	if n, ok := v.vertMirror(ignore); ok {
		return n, true
	}
	if n, ok := v.horizMirror(ignore / 100); ok {
		return 100 * n, true
	}
	return 0, false
}

func loadVolcos(filename string) []volco {
	f, err := os.Open(filename)
	Expect(err).NotTo(HaveOccurred())
	defer f.Close()

	scanner := bufio.NewScanner(f)
	volcos := []volco{}
	rows := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			volcos = append(volcos, volco{terrain: rows})
			rows = [][]byte{}
			continue
		}
		rows = append(rows, []byte(line))
	}
	if len(rows) > 0 {
		volcos = append(volcos, volco{terrain: rows})
	}

	return volcos
}

var _ = Describe("13", func() {
	It("can do part A", func() {
		sum := 0
		for _, volco := range loadVolcos("input13") {
			s, ok := volco.score(-1)
			Expect(ok).To(BeTrue())
			sum += s
		}

		Expect(sum).To(Equal(35210))
	})

	It("can do part B", func() {
		sum := 0
		for _, volco := range loadVolcos("input13") {
			sum += volco.newScore()
		}

		Expect(sum).To(Equal(31974))
	})
})
