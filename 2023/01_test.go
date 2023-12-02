package two023_test

import (
	"bytes"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("01", func() {
	It("does part A", func() {
		bs, err := ioutil.ReadFile("input01")
		Expect(err).NotTo(HaveOccurred())

		sum := 0
		first := regexp.MustCompile(`^\D*(\d)`)
		last := regexp.MustCompile(`(\d)\D*$`)
		for _, bw := range bytes.Fields(bs) {
			fmatch := first.FindSubmatch(bw)
			lmatch := last.FindSubmatch(bw)

			sum += 10 * int(fmatch[1][0]-byte('0'))
			sum += int(lmatch[1][0] - byte('0'))
		}

		Expect(sum).To(Equal(55002))
	})

	It("does part B", func() {
		bs, err := ioutil.ReadFile("input01")
		Expect(err).NotTo(HaveOccurred())

		sum := 0
		for _, bw := range bytes.Fields(bs) {
			sum += 10 * firstDigit(string(bw))
			sum += lastDigit(string(bw))
		}

		Expect(sum).To(Equal(55093))
	})
})

var nums = []string{
	"one", "two", "three", "four", "five",
	"six", "seven", "eight", "nine",
}

func firstDigit(s string) int {
	min := len(s)
	minN := 0
	for i, ns := range nums {
		spos := strings.Index(s, ns)
		if spos > -1 && spos < min {
			min = spos
			minN = i + 1
		}
		ds := strconv.Itoa(i + 1)
		dpos := strings.Index(s, ds)
		if dpos > -1 && dpos < min {
			min = dpos
			minN = i + 1
		}
	}

	return minN
}

func lastDigit(s string) int {
	max := -1
	maxN := 0
	for i, ns := range nums {
		spos := strings.LastIndex(s, ns)
		if spos > -1 && spos > max {
			max = spos
			maxN = i + 1
		}
		ds := strconv.Itoa(i + 1)
		dpos := strings.LastIndex(s, ds)
		if dpos > -1 && dpos > max {
			max = dpos
			maxN = i + 1
		}
	}

	return maxN
}
