package days_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("11", func() {
	input := "hxbxwxba"
	It("does part A", func() {
		Expect(nextPassword(input)).To(Equal("hxbxxyzz"))
	})

	It("does part B", func() {
		Expect(nextPassword(nextPassword(input))).To(Equal("hxcaabcc"))
	})
})

func hasDoubles(in []byte) bool {
	doubles := map[byte]int{}

	for i := 0; i < len(in)-1; i++ {
		if in[i] == in[i+1] {
			doubles[in[i]]++
		}
	}

	return len(doubles) > 1
}

func hasSeq(in []byte) bool {
	for i := 0; i < len(in)-2; i++ {
		if in[i+1] == in[i]+1 && in[i+2] == in[i]+2 {
			return true
		}
	}

	return false
}

func next(in []byte) []byte {
	l := len(in)
	carry := 0

	for i := l - 1; i >= 0; i-- {
		if carry > 0 {
			in[i], carry = inc(in[i])
			if carry == 0 {
				break
			}
			continue
		}
		in[i], carry = inc(in[i])

		if carry == 0 {
			break
		}
	}

	return in
}

func nextPassword(last string) string {
	input := []byte(last)
	for {
		input = next(input)
		if !hasSeq(input) || !hasDoubles(input) {
			continue
		}

		break
	}

	return string(input)
}

func inc(b byte) (byte, int) {
	if b == 'z' {
		return 'a', 1
	}

	b++
	for b == 'o' || b == 'i' || b == 'l' {
		b++
	}

	return b, 0
}
