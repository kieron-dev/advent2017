package hexagons_test

import (
	"github.com/kieron-pivotal/advent2017/11/hexagons"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hexagons", func() {
	It("calcs correct distance", func() {
		Expect(hexagons.Distance([]string{"s", "s"})).To(Equal(2))
		Expect(hexagons.Distance([]string{"s", "n"})).To(Equal(0))
		Expect(hexagons.Distance([]string{"ne", "nw"})).To(Equal(1))
		Expect(hexagons.Distance([]string{"ne", "nw", "n", "sw", "se"})).To(Equal(1))
	})
})
