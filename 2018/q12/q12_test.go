package q12_test

import (
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q12"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q12", func() {

	var (
		r *os.File
	)

	BeforeEach(func() {
		var err error
		r, err = os.Open("example")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		r.Close()
	})

	Context("loading config", func() {
		It("can get initial state", func() {
			plants := q12.NewPlants(r)
			Expect(plants.State()).To(Equal("#..#.#..##......###...###"))
		})

		It("can load rules", func() {
			plants := q12.NewPlants(r)
			rules := plants.Rules()
			Expect(rules).To(HaveLen(14))
			Expect(rules["..#.."]).To(Equal("#"))
		})
	})

	Context("evolving", func() {
		It("can evolve some steps", func() {
			plants := q12.NewPlants(r)
			By("step 1", func() {
				plants.Step()
				Expect(plants.State()).To(Equal("#...#....#.....#..#..#..#"))
			})
			By("step 2", func() {
				plants.Step()
				Expect(plants.State()).To(Equal("##..##...##....#..#..#..##"))
			})
			By("step 3", func() {
				plants.Step()
				Expect(plants.State()).To(Equal("#.#...#..#.#....#..#..#...#"))
			})
		})
	})

	It("has correct hash sum after 20 steps", func() {
		plants := q12.NewPlants(r)
		for i := 0; i < 20; i++ {
			plants.Step()
		}
		Expect(plants.HashPosSum()).To(Equal(325))
	})
})
