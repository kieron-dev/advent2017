package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/image"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("20", func() {
	var (
		data   *os.File
		sorter image.Sorter
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input20")
		Expect(err).NotTo(HaveOccurred())

		sorter = image.NewSorter()
		sorter.Load(data)
	})

	It("does part A", func() {
		Expect(sorter.CornerProduct()).To(Equal(16192267830719))
	})

	It("does part B", func() {
		Expect(sorter.Solve()).To(Equal(1909))
	})
})
