package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/password"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("02", func() {
	var (
		input   *os.File
		checker password.Checker
	)

	BeforeEach(func() {
		var err error
		input, err = os.Open("./input02")
		Expect(err).NotTo(HaveOccurred())

		checker = password.NewChecker()

		checker.Load(input)
	})

	It("part A", func() {
		Expect(checker.CorrectCount()).To(Equal(538))
	})

	It("part B", func() {
		Expect(checker.CorrectCountNew()).To(Equal(489))
	})
})
