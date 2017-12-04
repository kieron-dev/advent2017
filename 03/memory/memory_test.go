package memory_test

import (
	"github.com/kieron-pivotal/advent2017/03/memory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Memory", func() {
	It("calcs distance", func() {
		Expect(memory.Distance(1)).To(Equal(0))

		Expect(memory.Distance(2)).To(Equal(1))
		Expect(memory.Distance(3)).To(Equal(2))
		Expect(memory.Distance(4)).To(Equal(1))
		Expect(memory.Distance(5)).To(Equal(2))
		Expect(memory.Distance(6)).To(Equal(1))
		Expect(memory.Distance(7)).To(Equal(2))
		Expect(memory.Distance(8)).To(Equal(1))
		Expect(memory.Distance(9)).To(Equal(2))

		Expect(memory.Distance(10)).To(Equal(3))
		Expect(memory.Distance(11)).To(Equal(2))
		Expect(memory.Distance(12)).To(Equal(3))
		Expect(memory.Distance(13)).To(Equal(4))
		Expect(memory.Distance(14)).To(Equal(3))
		Expect(memory.Distance(15)).To(Equal(2))
		Expect(memory.Distance(16)).To(Equal(3))
		Expect(memory.Distance(17)).To(Equal(4))

		Expect(memory.Distance(26)).To(Equal(5))
		Expect(memory.Distance(27)).To(Equal(4))
		Expect(memory.Distance(28)).To(Equal(3))
	})

	It("calcs weirdSum", func() {
		Expect(memory.WeirdSum(1)).To(Equal(2))
		Expect(memory.WeirdSum(2)).To(Equal(4))
		Expect(memory.WeirdSum(4)).To(Equal(5))
		Expect(memory.WeirdSum(5)).To(Equal(10))
		Expect(memory.WeirdSum(747)).To(Equal(806))

	})
})
