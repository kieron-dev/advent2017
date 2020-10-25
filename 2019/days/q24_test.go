package days_test

import (
	"io/ioutil"

	"github.com/kieron-dev/advent2017/advent2019/life"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q24", func() {
	var (
		input []byte
		l     life.Life
	)

	BeforeEach(func() {
		var err error
		input, err = ioutil.ReadFile("./input24")
		Expect(err).NotTo(HaveOccurred())

		l = life.New(string(input))
	})

	It("does part A", func() {
		Expect(int(l.FirstRepeat())).To(Equal(18400817))
	})

	It("does part B", func() {
		line := life.NewLine(string(input))
		for i := 0; i < 200; i++ {
			line = line.Evolve()
		}
		Expect(line.CountBugs()).To(Equal(1944))
	})
})
