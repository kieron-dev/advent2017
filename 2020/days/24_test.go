package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/tiled"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("24", func() {
	var (
		data  *os.File
		floor tiled.Floor
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input24")
		Expect(err).NotTo(HaveOccurred())

		floor = tiled.NewFloor()
		floor.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("does part A", func() {
		Expect(floor.BlackCount()).To(Equal(277))
	})

	It("does part B", func() {
		for i := 0; i < 100; i++ {
			floor.Evolve()
		}

		Expect(floor.BlackCount()).To(Equal(3531))
	})
})
