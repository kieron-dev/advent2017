package q21_test

import (
	"fmt"
	"os"

	"github.com/kieron-pivotal/advent2017/2018/q21"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q21", func() {

	It("runs part b a few times", func() {
		f, err := os.Open("input")
		Expect(err).NotTo(HaveOccurred())
		c := q21.NewComputer(f)
		c.SetRegisters(1634, 0, 0, 0, 0, 0)

		i := 0
		for { //i := 0; i < 4000; i++ {
			fmt.Printf("%d %v %q", c.IP, c.Registers, c.Instructions[c.IP])
			if !c.ExecuteNext() {
				break
			}
			i++
			if i%10000 == 0 {
				fmt.Printf("...i = %+v\n", i)
			}
			if i > 10000000 {
				break
			}
			fmt.Printf(" %v\n", c.Registers)
		}
		fmt.Printf("i = %+v\n", i)
	})
})
