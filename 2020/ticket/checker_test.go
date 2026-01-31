package ticket_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/ticket"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Checker", func() {
	var (
		data    io.Reader
		checker ticket.Checker
	)

	JustBeforeEach(func() {
		checker = ticket.NewChecker()
		checker.Load(data)
	})

	Describe("part A", func() {
		BeforeEach(func() {
			data = strings.NewReader(`
class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
`)
		})

		It("creates 3 rules", func() {
			Expect(checker.RuleCount()).To(Equal(3))
		})

		It("loads the nearby tickets", func() {
			Expect(checker.NearbyCount()).To(Equal(4))
		})

		It("gets the error rate", func() {
			Expect(checker.ErrorRate()).To(Equal(71))
		})
	})

	Describe("part B", func() {
		BeforeEach(func() {
			data = strings.NewReader(`
class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9
`)
		})

		It("can match fields to indices", func() {
			Expect(checker.FieldPositions()).To(Equal(map[string]int{
				"class": 1,
				"row":   0,
				"seat":  2,
			}))
		})
	})
})
