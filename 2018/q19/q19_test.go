package q19_test

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kieron-pivotal/advent2017/2018/q19"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q19", func() {

	var (
		ex01 io.Reader
	)

	BeforeEach(func() {
		ex01 = strings.NewReader(`#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5`)
	})

	It("gets example 1 correct", func() {
		c := q19.NewComputer(ex01)
		a := c.Execute()
		Expect(a).To(Equal(6))
	})

	It("runs part A a few times", func() {
		f, err := os.Open("input")
		Expect(err).NotTo(HaveOccurred())
		c := q19.NewComputer(f)

		for { //i := 0; i < 4000; i++ {
			fmt.Printf("%d %q %v", c.IP, c.Instructions[c.IP], c.Registers)
			c.ExecuteNext()
			fmt.Printf(" %v\n", c.Registers)
		}
	})

	FIt("runs part b a few times", func() {
		f, err := os.Open("input")
		Expect(err).NotTo(HaveOccurred())
		c := q19.NewComputer(f)
		c.SetRegisters(1, 0, 0, 0, 0, 0)

		for i := 0; i < 400; i++ {
			fmt.Printf("%d %q %v", c.IP, c.Instructions[c.IP], c.Registers)
			c.ExecuteNext()
			fmt.Printf(" %v\n", c.Registers)
		}
	})
})
