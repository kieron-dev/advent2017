package dance_test

import (
	"github.com/kieron-pivotal/advent2017/16/dance"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dance", func() {
	It("creates dancers with correct labels", func() {
		d := dance.New(2)
		a, ok := d['a']
		Expect(ok).To(BeTrue())
		Expect(a).To(Equal(0))

		b, ok := d['b']
		Expect(ok).To(BeTrue())
		Expect(b).To(Equal(1))
	})

	It("spins N dancers", func() {
		d := dance.New(8)
		d.Spin(3)
		Expect(d.String()).To(Equal("fghabcde"))
	})

	It("exchanges dancers", func() {
		d := dance.New(5)
		d.Exchange(3, 4)
		Expect(d.String()).To(Equal("abced"))
	})

	It("swaps dancers", func() {
		d := dance.New(5)
		d.Swap('a', 'e')
		Expect(d.String()).To(Equal("ebcda"))
	})

	It("does combo correctly", func() {
		d := dance.New(5)
		d.Spin(1)
		d.Exchange(3, 4)
		d.Swap('e', 'b')
		Expect(d.String()).To(Equal("baedc"))
	})

	It("recognises spin move", func() {
		d := dance.New(5)
		d.Move("s3")
		Expect(d.String()).To(Equal("cdeab"))
	})

	It("recognises exchange move", func() {
		d := dance.New(16)
		d.Move("x2/15")
		Expect(d.String()).To(Equal("abpdefghijklmnoc"))
	})

	It("recognises swap move", func() {
		d := dance.New(16)
		d.Move("pc/p")
		Expect(d.String()).To(Equal("abpdefghijklmnoc"))
	})

	It("correctly does move combo", func() {
		d := dance.New(5)
		d.Move("s1")
		d.Move("x3/4")
		d.Move("pe/b")
		Expect(d.String()).To(Equal("baedc"))
	})

	It("processes a slice of moves", func() {
		d := dance.New(5)
		d.ProcessMoves([]string{"s1", "x3/4", "pe/b"})
		Expect(d.String()).To(Equal("baedc"))
	})

	It("recognises new dancers as orig order", func() {
		Expect(dance.New(5).IsOriginalOrder()).To(BeTrue())
	})

	It("recognises a permutated list is not orig order", func() {
		d := dance.New(5)
		d.Move("s2")
		Expect(d.IsOriginalOrder()).To(BeFalse())
	})
})
