package days_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("20", func() {
	lim := 36_000_000
	It("does part A", func() {
		presents := map[int]int{}

		i := 1
		min := lim / 10
		for i < min {
			for j := 1; i*j <= min; j++ {
				presents[i*j] += 10 * i
				if presents[i*j] >= lim {
					n := i * j
					if n < min {
						min = n
					}
				}
			}
			i++
		}

		Expect(min).To(Equal(831600))
	})

	It("does part B", func() {
		presents := map[int]int{}

		i := 1
		min := 100000000
		for i < min {
			for j := 1; j <= 50 && i*j < min; j++ {
				presents[i*j] += 11 * i
				if presents[i*j] >= lim {
					n := i * j
					if n < min {
						min = n
					}
				}
			}
			i++
		}

		Expect(min).To(Equal(884520))
	})
})
