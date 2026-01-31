package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/ingredient"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("21", func() {
	var (
		data    *os.File
		checker ingredient.Checker
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input21")
		Expect(err).NotTo(HaveOccurred())

		checker = ingredient.NewChecker()
		checker.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("does part A", func() {
		Expect(checker.NonAllergenCount()).To(Equal(2317))
	})

	It("does part B", func() {
		Expect(checker.AllergenIngredients()).To(Equal("kbdgs,sqvv,slkfgq,vgnj,brdd,tpd,csfmb,lrnz"))
	})
})
