package days_test

import (
	"io"
	"os"

	"github.com/kieron-dev/advent2017/advent2019/manyworlds"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q18", func() {
	var (
		w         *manyworlds.World
		input     io.Reader
		inputName string
	)

	JustBeforeEach(func() {
		var err error
		input, err = os.Open(inputName)
		if err != nil {
			panic(err)
		}
		w = manyworlds.NewWorld()
		w.LoadMap(input)
	})

	Context("part A", func() {
		BeforeEach(func() {
			inputName = "input18"
		})

		It("does part A", func() {
			Expect(w.MinStepsToCollectKeys()).To(Equal(4668))
		})
	})

	Context("part B", func() {
		BeforeEach(func() {
			if os.Getenv("INCLUDE_SLOW") != "true" {
				Skip(`"$INCLUDE_SLOW" != "true"`)
			}
			inputName = "input18b"
		})

		It("does part B", func() {
			Expect(w.MinStepsToCollectKeys()).To(Equal(1910))
		})
	})
})
