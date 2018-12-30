package q23_test

import (
	"io"
	"strings"

	"github.com/kieron-pivotal/advent2017/2018/q23"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q23", func() {
	var (
		ex01 io.Reader
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

})
