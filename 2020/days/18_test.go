package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/maths"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("18", func() {
	var (
		data           *os.File
		computer       maths.Computer
		plusPrecedence bool
	)

	JustBeforeEach(func() {
		var err error
		data, err = os.Open("./input18")
		Expect(err).NotTo(HaveOccurred())

		computer = maths.NewComputer(plusPrecedence)
		computer.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	Context("equal precedence", func() {
		BeforeEach(func() {
			plusPrecedence = false
		})

		It("does part A", func() {
			Expect(computer.SumResults()).To(Equal(3159145843816))
		})
	})

	Context("plus precedence", func() {
		BeforeEach(func() {
			plusPrecedence = true
		})

		It("does part B", func() {
			Expect(computer.SumResults()).To(Equal(55699621957369))
		})
	})
})
