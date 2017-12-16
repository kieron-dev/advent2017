package defrag_test

import (
	"github.com/kieron-pivotal/advent2017/14/defrag"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Defrag", func() {
	It("should hash correctly", func() {
		Expect(defrag.Hash("flqrgnkx", 0)).To(HavePrefix("d4"))
	})

	It("counts bits in hex string", func() {
		Expect(defrag.BitCount("0f")).To(Equal(4))
		Expect(defrag.BitCount("ff")).To(Equal(8))
		Expect(defrag.BitCount("0fff")).To(Equal(12))
	})

	It("gets total right", func() {
		Expect(defrag.CountUsed("flqrgnkx")).To(Equal(8108))
	})
})
