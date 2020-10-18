package days_test

import (
	"github.com/kieron-dev/advent2017/advent2019"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q04", func() {
	It("does part A", func() {
		gen := advent2019.NewPWGen(347312, 805915, false)
		var err error
		i := 0
		for err == nil {
			_, err = gen.Next()
			i++
		}
		Expect(i).To(Equal(594))
	})

	It("does part B", func() {
		gen := advent2019.NewPWGen(347312, 805915, true)
		i := 1
		for {
			_, err := gen.Next()
			if err != nil {
				break
			}
			i++
		}
		Expect(i).To(Equal(364))
	})
})
