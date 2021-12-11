package days_test

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Coord struct {
	x, y int
}

func NewCoord(x, y int) Coord {
	return Coord{x: x, y: y}
}

var _ = Describe("03", func() {
	It("does part A", func() {
		input, err := ioutil.ReadFile("input03")
		Expect(err).NotTo(HaveOccurred())

		count := map[Coord]int{
			NewCoord(0, 0): 2,
		}

		x, y := 0, 0

		for _, c := range input {
			switch c {
			case '<':
				x--
			case '>':
				x++
			case 'v':
				y--
			case '^':
				y++
			}
			count[NewCoord(x, y)]++
		}

		Expect(len(count)).To(Equal(2081))
	})

	It("does part B", func() {
		input, err := ioutil.ReadFile("input03")
		Expect(err).NotTo(HaveOccurred())

		count := map[Coord]int{
			NewCoord(0, 0): 2,
		}

		x := []int{0, 0}
		y := []int{0, 0}

		for i, c := range input {
			switch c {
			case '<':
				x[i%2]--
			case '>':
				x[i%2]++
			case 'v':
				y[i%2]--
			case '^':
				y[i%2]++
			}
			count[NewCoord(x[i%2], y[i%2])]++
		}

		Expect(len(count)).To(Equal(2341))
	})
})
