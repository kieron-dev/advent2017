package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/navigation"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("12", func() {
	var (
		data *os.File
		ship navigation.Ship
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input12")
		Expect(err).NotTo(HaveOccurred())

		ship = navigation.NewShip()
		ship.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("part A", func() {
		ship.Move()
		Expect(ship.ManhattanDistance()).To(Equal(2847))
	})

	It("part B", func() {
		ship.MoveNew()
		Expect(ship.ManhattanDistance()).To(Equal(29839))
	})
})
