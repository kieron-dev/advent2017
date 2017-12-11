package hash_test

import (
	"github.com/kieron-pivotal/advent2017/10/hash"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hash", func() {
	var list []int
	BeforeEach(func() {
		list = []int{0, 1, 2, 3, 4, 5, 6, 7}
	})

	It("does circular gets", func() {
		Expect(hash.Get(list, 0)).To(Equal(0))
		Expect(hash.Get(list, 1)).To(Equal(1))
		Expect(hash.Get(list, 7)).To(Equal(7))
		Expect(hash.Get(list, 8)).To(Equal(0))
		Expect(hash.Get(list, 9)).To(Equal(1))
	})

	It("does circular sets", func() {
		hash.Set(list, 0, 10)
		Expect(hash.Get(list, 0)).To(Equal(10))
		hash.Set(list, 8, 20)
		Expect(hash.Get(list, 0)).To(Equal(20))
		hash.Set(list, 17, 30)
		Expect(hash.Get(list, 1)).To(Equal(30))
	})

	It("does a simple circular reverse", func() {
		hash.Reverse(list, 0, 1)
		Expect(hash.Get(list, 0)).To(Equal(1))
		Expect(hash.Get(list, 1)).To(Equal(0))
	})

	It("does a longer simple reverse", func() {
		hash.Reverse(list, 1, 3)
		Expect(list).To(Equal([]int{0, 3, 2, 1, 4, 5, 6, 7}))
	})

	It("does a circular reverse", func() {
		hash.Reverse(list, 7, 9)
		Expect(list).To(Equal([]int{0, 7, 2, 3, 4, 5, 6, 1}))
	})

	It("computes hash", func() {
		Expect(hash.Compute([]int{3, 4, 1, 5}, 5)).To(Equal(12))
	})
})
