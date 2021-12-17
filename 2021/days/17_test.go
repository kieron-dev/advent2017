package days_test

import (
	"regexp"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type probe struct {
	posX int
	posY int
	velX int
	velY int
	maxY int
}

func newProbe(velX, velY int) probe {
	return probe{
		velX: velX,
		velY: velY,
	}
}

func (p *probe) step() {
	p.posX += p.velX
	p.posY += p.velY
	if p.posY > p.maxY {
		p.maxY = p.posY
	}
	if p.velX > 0 {
		p.velX--
	} else if p.velX < 0 {
		p.velX++
	}
	p.velY--
}

func (p probe) in(x1, x2, y1, y2 int) bool {
	return p.posX >= x1 &&
		p.posX <= x2 &&
		p.posY >= y1 &&
		p.posY <= y2
}

var _ = Describe("17", func() {
	It("does the example", func() {
		input := "target area: x=20..30, y=-10..-5"
		x1, x2, y1, y2 := parseInput17(input)

		max := 0
		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				probe := newProbe(i, j)
				hit := false
				for probe.posY >= y1 {
					if probe.in(x1, x2, y1, y2) {
						hit = true
						break
					}
					probe.step()
				}
				if hit && probe.maxY > max {
					max = probe.maxY
				}
			}
		}

		Expect(max).To(Equal(45))
	})

	It("does part A", func() {
		input := "target area: x=277..318, y=-92..-53"
		x1, x2, y1, y2 := parseInput17(input)

		max := 0
		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				probe := newProbe(i, j)
				hit := false
				for probe.posY >= y1 {
					if probe.in(x1, x2, y1, y2) {
						hit = true
						break
					}
					probe.step()
				}
				if hit && probe.maxY > max {
					max = probe.maxY
				}
			}
		}

		Expect(max).To(Equal(4186))
	})

	It("does part B", func() {
		input := "target area: x=277..318, y=-92..-53"
		x1, x2, y1, y2 := parseInput17(input)

		count := 0
		for i := 0; i < 1000; i++ {
			for j := -100; j < 100; j++ {
				probe := newProbe(i, j)
				hit := false
				for probe.posY >= y1 {
					if probe.in(x1, x2, y1, y2) {
						hit = true
						break
					}
					probe.step()
				}
				if hit {
					count++
				}
			}
		}

		Expect(count).To(Equal(2709))
	})
})

func parseInput17(s string) (int, int, int, int) {
	re := regexp.MustCompile(`target area: x=(\d+)..(\d+), y=(-?\d+)..(-?\d+)`)
	matches := re.FindStringSubmatch(s)
	Expect(matches).To(HaveLen(5))

	return AToI(matches[1]), AToI(matches[2]), AToI(matches[3]), AToI(matches[4])
}
