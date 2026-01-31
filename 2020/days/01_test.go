package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/money"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("01", func() {
	var (
		moneyWrangler money.Wrangler
		data          *os.File
	)

	BeforeEach(func() {
		moneyWrangler = money.NewWrangler()

		var err error
		data, err = os.Open("./input01")
		Expect(err).NotTo(HaveOccurred())

		moneyWrangler.LoadExpenses(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("part A", func() {
		Expect(moneyWrangler.ProductFor(2020)).To(Equal(605364))
	})

	It("part B", func() {
		Expect(moneyWrangler.Product3For(2020)).To(Equal(128397680))
	})
})
