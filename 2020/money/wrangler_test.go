package money_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/money"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wrangler", func() {
	var (
		wrangler money.Wrangler
		data     io.Reader
	)

	BeforeEach(func() {
		wrangler = money.NewWrangler()

		data = strings.NewReader(`
1
2
8
9
3
`)
		wrangler.LoadExpenses(data)
	})

	It("can find numbers summing to a value", func() {
		Expect(wrangler.GetSummingTo(12)).To(ConsistOf(3, 9))
	})

	It("can multiply the sum numbers", func() {
		Expect(wrangler.ProductFor(12)).To(Equal(27))
	})

	It("can find three numbers summing to a value", func() {
		Expect(wrangler.Get3SummingTo(20)).To(ConsistOf(3, 8, 9))
	})
})
