package days_test

import (
	"io"
	"os"

	"github.com/kieron-pivotal/advent2017/advent2019/bodies"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q12", func() {
	var (
		input io.Reader
		s     *bodies.System
	)

	BeforeEach(func() {
		var err error
		input, err = os.Open("./input12")
		if err != nil {
			panic(err)
		}
		s = bodies.NewSystem()
		s.Load(input)
	})

	It("does part A", func() {
		for i := 0; i < 1000; i++ {
			s.Tick()
		}
		Expect(s.TotalEnergy()).To(Equal(10189))
	})

	It("does part B", func() {
		Expect(s.FirstRepeat()).To(Equal(int64(0)))
	})
})
