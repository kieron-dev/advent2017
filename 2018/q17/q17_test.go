package q17_test

import (
	"log"
	"os"
	"strings"

	"github.com/kieron-pivotal/advent2017/2018/q17"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q17", func() {

	It("draws correct pic", func() {
		in := strings.NewReader(`x=495, y=2..7
y=7, x=495..501
x=501, y=3..7
x=498, y=2..4
x=506, y=1..2
x=498, y=10..13
x=504, y=10..13
y=13, x=498..504
`)
		s := q17.NewSlice(in)
		s.Print()
		Expect(true).To(BeTrue())
	})

	FIt("can load the real input", func() {
		f, err := os.Open("input")
		if err != nil {
			log.Fatal(err)
		}

		s := q17.NewSlice(f)
		s.Print()
	})

})
