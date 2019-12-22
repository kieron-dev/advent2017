package fft_test

import (
	"fmt"

	"github.com/kieron-pivotal/advent2017/advent2019/fft"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("FFT", func() {

	var (
		t *fft.Transform
	)

	BeforeEach(func() {
		t = fft.NewTransform()
	})

	DescribeTable("easy transforms", func(input []int, phases int, expectedOutput []int) {
		for i := 0; i < phases; i++ {
			input = t.Process(input)
		}
		Expect(input).To(Equal(expectedOutput))
	},
		Entry("1", []int{1, 2, 3, 4, 5, 6, 7, 8}, 1, []int{4, 8, 2, 2, 6, 1, 5, 8}),
		Entry("2", []int{1, 2, 3, 4, 5, 6, 7, 8}, 2, []int{3, 4, 0, 4, 0, 4, 3, 8}),
		Entry("3", []int{1, 2, 3, 4, 5, 6, 7, 8}, 3, []int{0, 3, 4, 1, 5, 5, 1, 8}),
		Entry("4", []int{1, 2, 3, 4, 5, 6, 7, 8}, 4, []int{0, 1, 0, 2, 9, 4, 9, 8}),
	)

	DescribeTable("longer transforms", func(input string, phases int, expectedOutput8 []int) {
		in := fft.StringToSlice(input)
		for i := 0; i < phases; i++ {
			in = t.Process(in)
		}
		Expect(in[:8]).To(Equal(expectedOutput8))
	},
		Entry("1", "80871224585914546619083218645595", 100, []int{2, 4, 1, 7, 6, 1, 7, 6}),
		Entry("2", "19617804207202209144916044189917", 100, []int{7, 3, 7, 4, 5, 4, 1, 8}),
		Entry("3", "69317163492948606335995924319873", 100, []int{5, 2, 4, 3, 2, 1, 3, 3}),
	)

	DescribeTable("partial transforms", func(input string, phases int, expectedOutput8 []int) {
		totalLen := 10000 * len(input)
		offset := fft.GetOffset(input)
		Expect(fmt.Sprintf("%07d", offset)).To(Equal(input[:7]))

		ints := fft.StringToSlice(input)

		requiredLen := totalLen - offset
		inputInts := make([]int, requiredLen)
		for i := 0; i < requiredLen; i++ {
			inputInts[i] = ints[(offset+i)%len(input)]
		}
		for i := 0; i < phases; i++ {
			inputInts = t.ProcessForOffset(inputInts)
		}
		Expect(inputInts[:8]).To(Equal(expectedOutput8))
	},
		Entry("1", "03036732577212944063491565474664", 100, []int{8, 4, 4, 6, 2, 0, 2, 6}),
		Entry("2", "02935109699940807407585447034323", 100, []int{7, 8, 7, 2, 5, 2, 7, 0}),
		Entry("3", "03081770884921959731165446850517", 100, []int{5, 3, 5, 5, 3, 7, 3, 1}),
	)
})
