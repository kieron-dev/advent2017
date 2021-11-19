package days_test

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("01", func() {
	It("does part A", func() {
		contents, err := ioutil.ReadFile("input01")
		Expect(err).NotTo(HaveOccurred())

		pos := 0
		for _, b := range contents {
			switch b {
			case '(':
				pos++
			case ')':
				pos--
			}
		}

		Expect(pos).To(Equal(138))
	})

	It("does part B", func() {
		contents, err := ioutil.ReadFile("input01")
		Expect(err).NotTo(HaveOccurred())

		pos := 0
		step := 0

		for i, b := range contents {
			switch b {
			case '(':
				pos++
			case ')':
				pos--
			}

			if pos == -1 {
				step = i + 1
				break
			}
		}

		Expect(step).To(Equal(1771))
	})
})
