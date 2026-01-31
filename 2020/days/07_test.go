package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/tree"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("07", func() {
	var (
		data *os.File
		bags tree.Bags
	)

	BeforeEach(func() {
		var err error
		data, err := os.Open("./input07")
		Expect(err).NotTo(HaveOccurred())

		bags = tree.NewBags()
		bags.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("part A", func() {
		Expect(bags.NumOuterContaining("shiny gold")).To(Equal(208))
	})

	It("part B", func() {
		Expect(bags.BagsInside("shiny gold")).To(Equal(1664))
	})
})
