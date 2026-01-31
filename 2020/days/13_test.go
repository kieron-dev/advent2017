package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/bus"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("13", func() {
	var (
		data   *os.File
		finder bus.Finder
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input13")
		Expect(err).NotTo(HaveOccurred())

		finder = bus.NewFinder()
		finder.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("part A", func() {
		busNo, wait := finder.Find()
		Expect(busNo * wait).To(Equal(3882))
	})

	It("part B", func() {
		Expect(finder.SpecialTimestamp()).To(Equal(867295486378319))
	})
})
