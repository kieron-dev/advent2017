package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/rule"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("19", func() {
	var (
		data      *os.File
		validator rule.Validator
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input19")
		Expect(err).NotTo(HaveOccurred())

		validator = rule.NewValidator()
		validator.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("does part A", func() {
		Expect(validator.ValidCount()).To(Equal(265))
	})

	It("does part B", func() {
		validator.SetNewRules()
		Expect(validator.ValidCount()).To(Equal(394))
	})
})
