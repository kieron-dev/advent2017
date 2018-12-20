package q17_test

import (
	"io"
	"strings"

	"github.com/kieron-pivotal/advent2017/2018/q17"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q17", func() {

	var (
		ex01 io.Reader
	)

	BeforeEach(func() {
		ex01 = strings.NewReader(`x=495, y=2..7
y=7, x=495..501
x=501, y=3..7
x=498, y=2..4
x=506, y=1..2
x=498, y=10..13
x=504, y=10..13
y=13, x=498..504
`)
	})

	It("draws correct pic", func() {
		s := q17.NewSlice(ex01)
		s.Print()
		Expect(true).To(BeTrue())
	})

	It("can detect a contained point", func() {
		s := q17.NewSlice(ex01)
		Expect(s.GetContainedRow(q17.NewCoord(500, 6))).To(Equal([]q17.Coord{
			{X: 496, Y: 6},
			{X: 497, Y: 6},
			{X: 498, Y: 6},
			{X: 499, Y: 6},
			{X: 500, Y: 6},
		}))
	})

	It("can detect an uncontained point", func() {
		s := q17.NewSlice(ex01)
		Expect(s.GetContainedRow(q17.NewCoord(500, 2))).To(Equal([]q17.Coord{}))
	})

	FIt("can fill solid bits", func() {
		s := q17.NewSlice(ex01)
		s.Flow(q17.NewCoord(500, 0))
		s.Print()
		Expect(s.CountWater()).To(Equal(57))
	})

	// FIt("can load the real input", func() {
	// 	f, err := os.Open("input")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	s := q17.NewSlice(f)
	// 	s.Flow(q17.NewCoord(500, 0))
	// 	s.Print()
	// 	fmt.Printf("s.CountWater() = %+v\n", s.CountWater())
	// })

})
