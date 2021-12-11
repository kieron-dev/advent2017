package days_test

import (
	"crypto/md5"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("04", func() {
	prefix := "iwrupvqb"

	It("does part A", func() {
		i := 0
		for {
			str := prefix + strconv.Itoa(i)
			hash := md5.Sum([]byte(str))

			if hash[0] == 0 && hash[1] == 0 && hash[2] < 16 {
				break
			}

			i++
		}

		Expect(i).To(Equal(346386))
	})

	// Too slow!
	// It("does part B", func() {
	// 	i := 0
	// 	for {
	// 		str := prefix + strconv.Itoa(i)
	// 		hash := md5.Sum([]byte(str))

	// 		if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
	// 			break
	// 		}

	// 		i++
	// 	}

	// 	Expect(i).To(Equal(9958218))
	// })
})
