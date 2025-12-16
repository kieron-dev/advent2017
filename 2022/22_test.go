package two022_test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type maze22 struct {
	layout       [][]byte
	instructions string
}

func newMaze(in io.Reader) maze22 {
	var maze maze22

	inMaze := true
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			inMaze = false
			continue
		}

		if inMaze {
			maze.layout = append(maze.layout, []byte(line))
			continue
		}

		maze.instructions = line
	}

	return maze
}

func (m maze22) incr(r, c, d int) (int, int) {
	for {
		switch d {
		case 0:
			c = (c + 1) % len(m.layout[r])
		case 1:
			r = (r + 1) % len(m.layout)
		case 2:
			c = (c + len(m.layout[r]) - 1) % len(m.layout[r])
		case 3:
			r = (r + len(m.layout) - 1) % len(m.layout)
		}

		if c >= len(m.layout[r]) {
			continue
		}

		if m.layout[r][c] != ' ' {
			break
		}
	}

	return r, c
}

func (m maze22) move(r, c, d, n int) (int, int) {
	for n > 0 {
		nr, nc := m.incr(r, c, d)
		s := m.layout[nr][nc]
		if s == '#' {
			return r, c
		}
		if s == '.' {
			n--
		}
		r, c = nr, nc

	}
	return r, c
}

func (m maze22) followInstructions() (int, int, int) {
	var r, c, d int

	for m.layout[r][c] != '.' {
		c++
	}

	num := 0
	for _, i := range m.instructions {
		if i == 'R' {
			r, c = m.move(r, c, d, num)
			d = (d + 1) % 4
			num = 0
			continue
		}
		if i == 'L' {
			r, c = m.move(r, c, d, num)
			d = (d + 3) % 4
			num = 0
			continue
		}
		n, err := strconv.Atoi(string(i))
		Expect(err).NotTo(HaveOccurred())
		num = num*10 + n
	}

	r, c = m.move(r, c, d, num)

	return r + 1, c + 1, d
}

func (m maze22) getPassword() int {
	r, c, d := m.followInstructions()
	return 1000*r + 4*c + d
}

var _ = Describe("22", func() {
	example := `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5
`

	It("can load the example", func() {
		m := newMaze(strings.NewReader(example))

		Expect(m.layout).To(HaveLen(12))
		Expect(m.layout[11][14]).To(Equal(byte('#')))
		Expect(m.instructions).To(Equal("10R5L5R10L4R5L5"))
	})

	It("can follow instructions", func() {
		m := newMaze(strings.NewReader(example))
		r, c, d := m.followInstructions()

		Expect(r).To(Equal(6))
		Expect(c).To(Equal(8))
		Expect(d).To(Equal(0))
	})

	It("can go through the top", func() {
		m := newMaze(strings.NewReader(example))
		m.instructions = "L3R8L2"
		r, c, d := m.followInstructions()
		fmt.Printf("r, c, d = %d, %d, %d\n", r, c, d)

		Expect(r).To(Equal(12))
		Expect(c).To(Equal(13))
		Expect(d).To(Equal(3))
	})

	It("can do part A", func() {
		f, err := os.Open("input22")
		Expect(err).NotTo(HaveOccurred())
		m := newMaze(f)
		f.Close()

		Expect(m.getPassword()).To(Equal(60362))
	})

	It("does part B", func() {
		/* shape
		.AB
		.C
		.D
		EF
		*/
	})
})
