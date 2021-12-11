package days_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const dayInput = "3113322113"

var _ = Describe("10", func() {
	It("does part A", func() {
		num := dayInput

		for i := 0; i < 40; i++ {
			num = lookSay(num)
		}

		Expect(len(num)).To(Equal(329356))
	})

	It("does part B", func() {
		num := dayInput

		for i := 0; i < 50; i++ {
			num = lookSay(num)
		}

		Expect(len(num)).To(Equal(4666278))
	})
})

func lookSay(num string) string {
	out := strings.Builder{}

	i := 0
	for i < len(num) {
		cur := num[i]

		if i == len(num)-1 {
			out.WriteString("1" + string(cur))
			break
		}

		j := 1
		for num[i+j] == cur {
			j++
		}

		out.WriteString(fmt.Sprintf("%d%c", j, cur))

		i += j
	}

	return out.String()
}
