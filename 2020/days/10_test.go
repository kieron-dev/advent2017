package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/jolt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("10", func() {
	var (
		data      *os.File
		organiser jolt.Organiser
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input10")
		Expect(err).NotTo(HaveOccurred())

		organiser = jolt.NewOrganiser()
		organiser.Load(data)
	})

	It("part A", func() {
		diffs := organiser.GetDiffs()
		Expect(diffs[1] * diffs[3]).To(Equal(2059))
	})

	It("part B", func() {
		Expect(organiser.Combinations()).To(Equal(86812553324672))
	})
})
