package q23_test

import (
	"io"
	"os"
	"strings"

	"github.com/kieron-pivotal/advent2017/2018/q23"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q23", func() {
	var (
		ex01 io.Reader
		ex02 io.Reader
	)

	BeforeEach(func() {
		ex01 = strings.NewReader(`pos=<0,0,0>, r=4
pos=<1,0,0>, r=1
pos=<4,0,0>, r=3
pos=<0,2,0>, r=1
pos=<0,5,0>, r=3
pos=<0,0,3>, r=1
pos=<1,1,1>, r=1
pos=<1,1,2>, r=1
pos=<1,3,1>, r=1`)
		ex02 = strings.NewReader(`pos=<10,12,12>, r=2
pos=<12,14,12>, r=2
pos=<16,12,12>, r=4
pos=<14,14,14>, r=6
pos=<50,50,50>, r=200
pos=<10,10,10>, r=5`)
	})

	It("can find strongest nanobot", func() {
		t := q23.NewTeleport(ex01)
		strongest := t.Strongest()
		Expect(strongest.SignalRadius).To(Equal(4))
		Expect(strongest.Coord).To(Equal(q23.Coord{}))
	})

	It("can get the in range count", func() {
		t := q23.NewTeleport(ex01)
		strongest := t.Strongest()
		Expect(t.InRange(strongest)).To(Equal(7))
	})

	It("can get limits", func() {
		t := q23.NewTeleport(ex01)
		min, max := t.GetLimits()
		Expect(min.X).To(Equal(-4))
		Expect(min.Y).To(Equal(-4))
		Expect(min.Z).To(Equal(-4))
		Expect(max.X).To(Equal(7))
		Expect(max.Y).To(Equal(8))
		Expect(max.Z).To(Equal(4))
	})

	It("can sample the real input", func() {
		f, err := os.Open("input")
		Expect(err).NotTo(HaveOccurred())
		t := q23.NewTeleport(f)
		min, max := t.GetLimits()
		step := q23.GetStepForSample(min, max, 32)
		inRange := t.Sample(min, max, step)
		Expect(inRange).To(BeNumerically(">", 100))
	})

	It("can find the best coord", func() {
		t := q23.NewTeleport(ex02)
		coord := t.FindBestCoord(24)
		// cube := t.GetBestSubCube(8)
		// coord := t.GetBestInCube(cube[0], cube[1])
		Expect(coord).To(Equal(q23.C(12, 12, 12)))
	})

})
