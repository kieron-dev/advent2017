package q22_test

import (
	"github.com/kieron-pivotal/advent2017/2018/q22"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q22", func() {

	It("can calculate the geologic indices", func() {
		m := q22.NewMap(q22.C(10, 10), 510)

		Expect(m.GeologicIndex(q22.C(0, 0))).To(Equal(0))
		Expect(m.GeologicIndex(q22.C(1, 0))).To(Equal(16807))
		Expect(m.GeologicIndex(q22.C(0, 1))).To(Equal(48271))
		Expect(m.GeologicIndex(q22.C(1, 1))).To(Equal(145722555))
	})

	It("can calculate the erosion levels", func() {
		m := q22.NewMap(q22.C(10, 10), 510)

		Expect(m.ErosionLevel(q22.C(0, 0))).To(Equal(510))
		Expect(m.ErosionLevel(q22.C(1, 0))).To(Equal(17317))
		Expect(m.ErosionLevel(q22.C(0, 1))).To(Equal(8415))
		Expect(m.ErosionLevel(q22.C(1, 1))).To(Equal(1805))
		Expect(m.ErosionLevel(q22.C(10, 10))).To(Equal(510))
	})

	It("can calculate the cave types", func() {
		m := q22.NewMap(q22.C(10, 10), 510)

		Expect(m.Type(q22.C(0, 0))).To(Equal(q22.Rocky))
		Expect(m.Type(q22.C(1, 0))).To(Equal(q22.Wet))
		Expect(m.Type(q22.C(0, 1))).To(Equal(q22.Rocky))
		Expect(m.Type(q22.C(1, 1))).To(Equal(q22.Narrow))
		Expect(m.Type(q22.C(10, 10))).To(Equal(q22.Rocky))
	})

	It("can calculate the risk level from 0,0 to target", func() {
		m := q22.NewMap(q22.C(10, 10), 510)

		Expect(m.RiskLevel()).To(Equal(114))
	})

	It("can get the type of an out-of-bounds cell", func() {
		m := q22.NewMap(q22.C(10, 10), 510)

		Expect(m.Type(q22.C(12, 13))).To(Equal(q22.Wet))
	})

	It("can get the min time to target", func() {
		m := q22.NewMap(q22.C(10, 10), 510)

		Expect(m.ShortestToTarget()).To(Equal(45))
	})

})
