package q25_test

import (
	"io"
	"strings"

	"github.com/kieron-pivotal/advent2017/2018/q25"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q25", func() {

	DescribeTable("can partition into constellations", func(points io.Reader, expectedNum int) {
		s := q25.NewSpace(points)
		Expect(s.Partition()).To(Equal(expectedNum))
	},
		Entry("ex1", strings.NewReader(`0,0,0,0
3,0,0,0
0,3,0,0
0,0,3,0
0,0,0,3
0,0,0,6
9,0,0,0
12,0,0,0`), 2),

		Entry("ex2", strings.NewReader(`-1,2,2,0
0,0,2,-2
0,0,0,-2
-1,2,0,0
-2,-2,-2,2
3,0,2,-1
-1,3,2,2
-1,0,-1,0
0,2,1,-2
3,0,0,0`), 4),
	)

})
