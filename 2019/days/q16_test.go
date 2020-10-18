package days_test

import (
	"io/ioutil"
	"strings"

	"github.com/kieron-dev/advent2017/advent2019/fft"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q16", func() {
	var (
		t     *fft.Transform
		inStr string
	)

	BeforeEach(func() {
		t = fft.NewTransform()
		contents, err := ioutil.ReadFile("input16")
		if err != nil {
			panic(err)
		}
		inStr = strings.TrimSpace(string(contents))
	})

	It("does part A", func() {
		input := fft.StringToSlice(inStr)
		for i := 0; i < 100; i++ {
			input = t.Process(input)
		}
		Expect(input[:8]).To(Equal([]int{7, 3, 1, 2, 7, 5, 2, 3}))
	})

	It("does part B", func() {
		totalLen := 10000 * len(inStr)
		offset := fft.GetOffset(inStr)

		ints := fft.StringToSlice(inStr)

		requiredLen := totalLen - offset
		inputInts := make([]int, requiredLen)
		for i := 0; i < requiredLen; i++ {
			inputInts[i] = ints[(offset+i)%len(inStr)]
		}
		for i := 0; i < 100; i++ {
			inputInts = t.ProcessForOffset(inputInts)
		}
		Expect(inputInts[:8]).To(Equal([]int{8, 0, 2, 8, 4, 4, 2, 0}))
	})
})
