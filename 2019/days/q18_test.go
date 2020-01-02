package days_test

import (
	"io"
	"os"

	"github.com/kieron-pivotal/advent2017/advent2019/manyworlds"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q18", func() {

	var (
		w     *manyworlds.World
		input io.Reader
	)

	BeforeEach(func() {
		var err error
		input, err = os.Open("input18")
		if err != nil {
			panic(err)
		}
		w = manyworlds.NewWorld()
		w.LoadMap(input)
	})

	It("does part A", func() {
		Expect(w.MinStepsToCollectKeys()).To(Equal(4668))
	})

	It("does part B", func() {

	})

})
