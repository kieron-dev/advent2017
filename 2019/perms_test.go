package advent2019_test

import (
	"fmt"

	"github.com/kieron-pivotal/advent2017/advent2019"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Perms", func() {

	It("produces 6 distinct perms of 3 chars", func() {
		permHelper := advent2019.Perms{}
		all := permHelper.All([]int{1, 2, 3})

		m := map[string]bool{}
		for _, p := range all {
			key := fmt.Sprintf("%d%d%d", p[0], p[1], p[2])
			m[key] = true
		}

		Expect(m).To(HaveLen(6))
	})

})
