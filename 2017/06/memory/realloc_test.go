package memory_test

import (
	"github.com/kieron-pivotal/advent2017/06/memory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("realloc", func() {
	It("calcs cycle length", func() {
		Expect(memory.ReallocFirstCyclePos([]int{0, 2, 7, 0})).To(Equal(5))
	})
})
