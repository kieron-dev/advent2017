package two023_test

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type game struct {
	id    int
	draws []draw
}

func newGame(s string) game {
	var g game
	parts := strings.Split(s, ":")
	Expect(parts).To(HaveLen(2))
	id, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
	Expect(err).NotTo(HaveOccurred())
	g.id = id

	for _, bit := range strings.Split(parts[1], ";") {
		g.draws = append(g.draws, newDraw(strings.TrimSpace(bit)))
	}

	return g
}

func (g game) possible(red, green, blue int) bool {
	for _, d := range g.draws {
		if d.red > red {
			return false
		}
		if d.green > green {
			return false
		}
		if d.blue > blue {
			return false
		}
	}
	return true
}

func (g game) power() int {
	var red, green, blue int
	for _, d := range g.draws {
		if d.red > red {
			red = d.red
		}
		if d.green > green {
			green = d.green
		}
		if d.blue > blue {
			blue = d.blue
		}
	}

	return red * green * blue
}

type draw struct {
	red   int
	green int
	blue  int
}

func newDraw(s string) draw {
	parts := strings.Split(s, ",")
	d := draw{}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		bits := strings.Split(part, " ")
		Expect(bits).To(HaveLen(2))
		num, err := strconv.Atoi(bits[0])
		Expect(err).NotTo(HaveOccurred())
		switch bits[1] {
		case "red":
			d.red = num
		case "green":
			d.green = num
		case "blue":
			d.blue = num
		}
	}
	return d
}

var _ = Describe("02", func() {
	It("does part A", func() {
		f, err := os.Open("input02")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		sum := 0
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}
			g := newGame(line)
			if g.possible(12, 13, 14) {
				sum += g.id
			}
		}

		Expect(sum).To(Equal(2720))
	})

	It("does part B", func() {
		f, err := os.Open("input02")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		sum := 0
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}
			g := newGame(line)
			sum += g.power()
		}

		Expect(sum).To(Equal(71535))
	})
})
