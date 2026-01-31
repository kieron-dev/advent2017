package seating_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/seating"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Seating", func() {
	var (
		data io.Reader
		plan seating.Plan
	)

	BeforeEach(func() {
		data = strings.NewReader(`
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL
`)
		plan = seating.NewPlan()
		plan.Load(data)
	})

	It("gets the highest id", func() {
		Expect(plan.MaxSeatID()).To(Equal(820))
	})
})
