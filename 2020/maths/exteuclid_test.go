package maths_test

import (
	"math/big"

	"github.com/kieron-dev/adventofcode/2020/maths"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ExtEuclid", func() {
	It("can calculate extended euclid", func() {
		d, s, t := maths.ExtEuclid(big.NewInt(3), big.NewInt(5))

		Expect(d.Int64()).To(Equal(int64(1)))
		Expect(s.Int64()).To(Equal(int64(2)))
		Expect(t.Int64()).To(Equal(int64(-1)))
	})
})
