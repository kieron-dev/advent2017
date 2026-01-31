package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/gameoflife"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("11", func() {
	var (
		data *os.File
		plan gameoflife.SeatingPlan
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input11")
		Expect(err).NotTo(HaveOccurred())

		plan = gameoflife.NewSeatingPlan()
		plan.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("part A", func() {
		plan.Stabilise(false)
		Expect(plan.OccupiedSeats()).To(Equal(2310))
	})

	It("part B", func() {
		plan.Stabilise(true)
		Expect(plan.OccupiedSeats()).To(Equal(2074))
	})
})
