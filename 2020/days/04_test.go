package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/passport"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("04", func() {
	var (
		data    *os.File
		manager passport.Manager
	)

	BeforeEach(func() {
		var err error
		data, err := os.Open("./input04")
		Expect(err).NotTo(HaveOccurred())

		manager = passport.NewManager()
		manager.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("part A", func() {
		Expect(manager.ValidCount()).To(Equal(264))
	})

	It("part B", func() {
		Expect(manager.StrictValidCount()).To(Equal(224))
	})
})
