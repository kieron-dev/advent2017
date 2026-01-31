package bus_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/bus"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Finder", func() {
	var (
		data   io.Reader
		finder bus.Finder
	)

	BeforeEach(func() {
		data = strings.NewReader(`939
7,13,x,x,59,x,31,19
`)
		finder = bus.NewFinder()
		finder.Load(data)
	})

	It("calculates the right bus no. and wait", func() {
		busNo, wait := finder.Find()

		Expect(busNo).To(Equal(59))
		Expect(wait).To(Equal(5))
	})

	It("calcs the right magic number", func() {
		Expect(finder.SpecialTimestamp()).To(Equal(1068781))
	})
})
