package advent2019_test

import (
	"github.com/kieron-pivotal/advent2017/advent2019"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("PWGen", func() {
	var (
		generator *advent2019.PWGen
		min       int
		max       int
		exactly2  bool
	)
	BeforeEach(func() {
		min = 200000
		max = 500000
		exactly2 = false
	})

	JustBeforeEach(func() {
		generator = advent2019.NewPWGen(min, max, exactly2)
	})

	It("iterates through 6 digit numbers", func() {
		Expect(generator.Current()).To(Equal(222222))
		Expect(generator.Next()).To(Equal(222223))
	})

	It("only includes numbers with two identical consecutive digits", func() {
		for {
			cur := generator.Current()
			Expect(cur).ToNot(Equal(234567))
			_, err := generator.Next()
			if err != nil {
				break
			}
		}
	})

	It("goes up to the max", func() {
		var err error
		var last int
		for err == nil {
			var n int
			n, err = generator.Next()
			if err == nil {
				last = n
			}
		}
		Expect(last).To(Equal(499999))
	})
})
