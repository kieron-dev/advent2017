package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/ticket"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("16", func() {
	var (
		data    *os.File
		checker ticket.Checker
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input16")
		Expect(err).NotTo(HaveOccurred())

		checker = ticket.NewChecker()
		checker.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("does part A", func() {
		Expect(checker.ErrorRate()).To(Equal(23115))
	})

	It("does part B", func() {
		Expect(checker.DepartureProduct()).To(Equal(239727793813))
	})
})
