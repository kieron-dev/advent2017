package two023_test

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type direction coord

// const really
var (
	left  = direction{0, -1}
	right = direction{0, 1}
	up    = direction{-1, 0}
	down  = direction{1, 0}
)

type beam struct {
	position coord
	dir      direction
}

func (b *beam) advance() {
	b.position = b.position.add(coord(b.dir))
}

func (b beam) key() string {
	return fmt.Sprintf("%v", b)
}

type mirrorMap struct {
	plan      []string
	beams     []beam
	energized map[coord]bool
}

func (m *mirrorMap) reset() {
	m.beams = []beam{}
	m.energized = map[coord]bool{}
}

func (m mirrorMap) traceBeams() {
	visited := map[string]bool{}

	for len(m.beams) > 0 {
		cur := m.beams[0]
		m.beams = m.beams[1:]

		for {
			cur.advance()
			if cur.position[0] < 0 || cur.position[0] >= len(m.plan) ||
				cur.position[1] < 0 || cur.position[1] >= len(m.plan[0]) {
				break
			}
			if visited[cur.key()] {
				break
			}
			visited[cur.key()] = true

			switch m.plan[cur.position[0]][cur.position[1]] {
			case '.':
				// nothing
			case '/':
				switch cur.dir {
				case left:
					cur.dir = down
				case right:
					cur.dir = up
				case up:
					cur.dir = right
				case down:
					cur.dir = left
				}
			case '\\':
				switch cur.dir {
				case left:
					cur.dir = up
				case right:
					cur.dir = down
				case up:
					cur.dir = left
				case down:
					cur.dir = right
				}
			case '|':
				if cur.dir == left || cur.dir == right {
					cur.dir = up
					m.beams = append(m.beams, beam{position: cur.position, dir: down})
				}
			case '-':
				if cur.dir == up || cur.dir == down {
					cur.dir = left
					m.beams = append(m.beams, beam{position: cur.position, dir: right})
				}
			}

			m.energized[cur.position] = true
		}
	}
}

func newMirrorMap(filename string) mirrorMap {
	f, err := os.Open(filename)
	Expect(err).NotTo(HaveOccurred())
	defer f.Close()

	var m mirrorMap
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		m.plan = append(m.plan, line)
	}
	m.beams = []beam{{position: coord{0, -1}, dir: right}}
	m.energized = map[coord]bool{}

	return m
}

var _ = Describe("16", func() {
	It("does part A", func() {
		m := newMirrorMap("input16")
		m.traceBeams()
		Expect(len(m.energized)).To(Equal(7472))
	})

	It("does part B", func() {
		max := 0
		m := newMirrorMap("input16")
		for r := 0; r < len(m.plan); r++ {
			m.reset()
			m.beams = append(m.beams, beam{
				position: coord{r, -1},
				dir:      right,
			})
			m.traceBeams()
			e := len(m.energized)
			if e > max {
				max = e
			}
		}
		for r := 0; r < len(m.plan); r++ {
			m.reset()
			m.beams = append(m.beams, beam{
				position: coord{r, len(m.plan[0])},
				dir:      left,
			})
			m.traceBeams()
			e := len(m.energized)
			if e > max {
				max = e
			}
		}
		for c := 0; c < len(m.plan[0]); c++ {
			m.reset()
			m.beams = append(m.beams, beam{
				position: coord{-1, c},
				dir:      down,
			})
			m.traceBeams()
			e := len(m.energized)
			if e > max {
				max = e
			}
		}
		for c := 0; c < len(m.plan[0]); c++ {
			m.reset()
			m.beams = append(m.beams, beam{
				position: coord{len(m.plan), c},
				dir:      up,
			})
			m.traceBeams()
			e := len(m.energized)
			if e > max {
				max = e
			}
		}
		Expect(max).To(Equal(7716))
	})
})
