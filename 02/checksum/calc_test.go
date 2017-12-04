package checksum_test

import (
	"github.com/kieron-pivotal/advent2017/02/checksum"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calc", func() {
	It("calcs checksum", func() {
		Expect(checksum.Calc([][]int{{1, 2, 3}})).To(Equal(2))
	})
})

var _ = Describe("RowVal", func() {
	It("calcs row val", func() {
		Expect(checksum.RowVal([]int{1, 2, 3})).To(Equal(2))
	})
})
