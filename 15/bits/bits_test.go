package bits_test

import (
	"github.com/kieron-pivotal/advent2017/15/bits"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bits", func() {
	It("Gets Next A", func() {
		n := int64(65)
		for i := 0; i < 5; i++ {
			n = bits.GetNextA(n)
		}
		Expect(n).To(Equal(int64(1352636452)))
	})

	It("gets next B", func() {
		Expect(bits.GetNextB(int64(8921))).To(Equal(int64(430625591)))
		Expect(bits.GetNextB(int64(430625591))).To(Equal(int64(1233683848)))
		Expect(bits.GetNextB(int64(1233683848))).To(Equal(int64(1431495498)))
		Expect(bits.GetNextB(int64(1431495498))).To(Equal(int64(137874439)))
		Expect(bits.GetNextB(int64(137874439))).To(Equal(int64(285222916)))
	})

	It("Gets Next 5 Bs", func() {
		n := int64(8921)
		for i := 0; i < 5; i++ {
			n = bits.GetNextB(n)
		}
		Expect(n).To(Equal(int64(285222916)))
	})

	It("counts lower 16 matches", func() {
		Expect(bits.CountLowerMatches(65, 8921, bits.GetNextA, bits.GetNextB, 4e7)).To(Equal(588))
	})

	It("solves part 1", func() {
		Expect(bits.CountLowerMatches(722, 354, bits.GetNextA, bits.GetNextB, 4e7)).To(Equal(0))
	})

	It("solves part 2", func() {
		Expect(bits.CountLowerMatches(722, 354, bits.GetNextA4, bits.GetNextB8, 5e6)).To(Equal(0))
	})
})
