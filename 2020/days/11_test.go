package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/gameoflife"
	. "github.com/onsi/ginkgo"
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
		plan.Stabilise()
		Expect(plan.OccupiedSeats()).To(Equal(-1))
	})
})
