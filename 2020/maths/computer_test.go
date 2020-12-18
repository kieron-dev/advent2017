package maths_test

import (
	"io"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kieron-dev/adventofcode/2020/maths"
)

var _ = Describe("Computer", func() {
	var (
		data           io.Reader
		computer       maths.Computer
		plusPrecedence bool
	)

	JustBeforeEach(func() {
		data = strings.NewReader(`
1 + 2 * 3 + 4 * 5 + 6
1 + (2 * 3) + (4 * (5 + 6))
2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2
`)

		computer = maths.NewComputer(plusPrecedence)
		computer.Load(data)
	})

	Context("equal precedence", func() {
		BeforeEach(func() {
			plusPrecedence = false
		})

		It("can get the results", func() {
			Expect(computer.Result(0)).To(Equal(71))
			Expect(computer.Result(1)).To(Equal(51))
			Expect(computer.Result(2)).To(Equal(26))
			Expect(computer.Result(3)).To(Equal(437))
			Expect(computer.Result(4)).To(Equal(12240))
			Expect(computer.Result(5)).To(Equal(13632))
		})
	})

	Context("plus precedence", func() {
		BeforeEach(func() {
			plusPrecedence = true
		})

		It("can get the results", func() {
			Expect(computer.Result(0)).To(Equal(231))
			Expect(computer.Result(1)).To(Equal(51))
			Expect(computer.Result(2)).To(Equal(46))
			Expect(computer.Result(3)).To(Equal(1445))
			Expect(computer.Result(4)).To(Equal(669060))
			Expect(computer.Result(5)).To(Equal(23340))
		})
	})
})
