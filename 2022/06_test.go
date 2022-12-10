package two022_test

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("06", func() {
	It("does part A", func() {
		bs, err := ioutil.ReadFile("input06")
		Expect(err).NotTo(HaveOccurred())

		res := -1

		set := map[byte]int{}
		for i := 0; i < 3; i++ {
			set[bs[i]]++
		}

		for i := 0; i < len(bs)-3; i++ {
			set[bs[i+3]]++

			if len(set) == 4 {
				res = i + 4
				break
			}

			set[bs[i]]--
			if set[bs[i]] == 0 {
				delete(set, bs[i])
			}
		}

		Expect(res).To(Equal(-2))
	})

	It("does part B", func() {
		bs, err := ioutil.ReadFile("input06")
		Expect(err).NotTo(HaveOccurred())

		res := -1

		set := map[byte]int{}
		for i := 0; i < 13; i++ {
			set[bs[i]]++
		}

		for i := 0; i < len(bs)-13; i++ {
			set[bs[i+13]]++

			if len(set) == 14 {
				res = i + 14
				break
			}

			set[bs[i]]--
			if set[bs[i]] == 0 {
				delete(set, bs[i])
			}
		}

		Expect(res).To(Equal(-2))
	})
})
