package jolt_test

import (
	"io"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kieron-dev/adventofcode/2020/jolt"
)

var _ = Describe("Organiser", func() {
	var (
		data      io.Reader
		organiser jolt.Organiser
	)

	BeforeEach(func() {
		data = strings.NewReader(`
16
10
15
5
1
11
7
19
6
12
4
`)
		organiser = jolt.NewOrganiser()
		organiser.Load(data)
	})

	It("counts the jolt diffs", func() {
		diffs := organiser.GetDiffs()
		Expect(diffs[1]).To(Equal(7))
		Expect(diffs[3]).To(Equal(5))
	})

	It("calcs the number of combinations possible", func() {
		Expect(organiser.Combinations()).To(Equal(8))
	})
})
