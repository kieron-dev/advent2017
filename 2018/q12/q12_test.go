package q12_test

import (
	"io"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q12"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q12", func() {

	Context("loading config", func() {
		var (
			r io.Reader
		)

		BeforeEach(func() {
			var err error
			r, err = os.Open("example")
			Expect(err).NotTo(HaveOccurred())
		})

		It("can get initial state", func() {
			plants := q12.NewPlants(r)
			Expect(plants.State()).To(Equal("#..#.#..##......###...###"))
		})

		It("can load rules", func() {
			plants := q12.NewPlants(r)
			Expect(plants.Rules()).To(HaveLen(14))
		})
	})
})
