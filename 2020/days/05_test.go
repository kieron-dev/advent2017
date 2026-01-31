package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/seating"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("05", func() {
	var (
		data *os.File
		plan seating.Plan
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input05")
		Expect(err).NotTo(HaveOccurred())

		plan = seating.NewPlan()
		plan.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("part A", func() {
		Expect(plan.MaxSeatID()).To(Equal(801))
	})

	It("part B", func() {
		Expect(plan.MissingSeat()).To(Equal(597))
	})
})
