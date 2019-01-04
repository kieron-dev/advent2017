package q01_test

import (
	"io"
	"strings"

	"github.com/kieron-pivotal/advent2017/2018/q01"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q01", func() {

	var (
		ex01 io.Reader
	)

	BeforeEach(func() {
		ex01 = strings.NewReader(`+1
-2
+3
+1`)
	})

	It("can add some numbers", func() {
		c := q01.NewCalibrator(ex01)
		Expect(c.Add()).To(Equal(3))
	})

	It("can return first sum visited twice", func() {
		c := q01.NewCalibrator(ex01)
		Expect(c.FirstRepeat()).To(Equal(2))
	})

})
