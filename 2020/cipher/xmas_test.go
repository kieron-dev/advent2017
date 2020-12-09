package cipher_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/cipher"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Xmas", func() {
	var (
		data io.Reader
		code cipher.Xmas
	)

	BeforeEach(func() {
		data = strings.NewReader(`
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`)
		code = cipher.NewXmas()
		code.Load(data)
	})

	It("finds the error in a size 5 xmas cipher", func() {
		Expect(code.FirstError(5)).To(Equal(127))
	})

	It("can find the list of contiguous numbers summing to error", func() {
		Expect(code.WeaknessNums(5)).To(Equal([]int{15, 25, 47, 40}))
	})

	It("can return the encryption weakness", func() {
		Expect(code.EncryptionWeakness(5)).To(Equal(62))
	})
})
