package q18_test

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kieron-pivotal/advent2017/2018/q18"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q18", func() {

	var (
		ex01 io.Reader
	)

	BeforeEach(func() {
		ex01 = strings.NewReader(`.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|.`)
	})

	It("can do step 1", func() {
		a := q18.NewArea(ex01)
		a.Print()
		fmt.Println("")
		a.Step()
		a.Print()
		Expect(true).To(BeTrue())
	})

	It("gets the calc right", func() {
		a := q18.NewArea(ex01)
		for i := 0; i < 10; i++ {
			a.Step()
		}
		a.Print()
		fmt.Printf("a.Score() = %+v\n", a.Score())
	})

	It("can work out periods", func() {
		f, err := os.Open("input")
		Expect(err).NotTo(HaveOccurred())
		a := q18.NewArea(f)
		start, period := a.GetPeriod()
		fmt.Printf("start = %+v\n", start)
		fmt.Printf("period = %+v\n", period)

		for i := 0; i < start; i++ {
			a.Step()
		}
		score := a.Score()
		for i := 0; i < period; i++ {
			a.Step()
		}
		Expect(a.Score()).To(Equal(score))
	})

	It("can go far in the future", func() {
		f, err := os.Open("input")
		Expect(err).NotTo(HaveOccurred())
		a := q18.NewArea(f)
		score := a.GetBigFutureScore(1000)

		a.Reset()

		for i := 0; i < 1000; i++ {
			a.Step()
		}
		Expect(a.Score()).To(Equal(score))
	})

})
