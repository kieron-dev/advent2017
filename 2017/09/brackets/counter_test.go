package brackets_test

import (
	"github.com/kieron-pivotal/advent2017/09/brackets"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Counter", func() {

	It("counts unnested braces", func() {
		Expect(brackets.Count("")).To(Equal(0))
		Expect(brackets.Count("{}")).To(Equal(1))
		Expect(brackets.Count("{}{}")).To(Equal(2))
	})

	It("counts nested braces", func() {
		Expect(brackets.Count("{{}}")).To(Equal(3))
		Expect(brackets.Count("{{{}}}")).To(Equal(6))
		Expect(brackets.Count("{{}}{}")).To(Equal(4))
	})

	It("ignores plain garbage", func() {
		Expect(brackets.Count("{<{}>}")).To(Equal(1))
		Expect(brackets.Count("{<{<<}>}")).To(Equal(1))
	})

	It("uses ! to escape within garbage", func() {
		Expect(brackets.Count("{<{<!>{}>}")).To(Equal(1))
		Expect(brackets.Count("{<!>{}>}")).To(Equal(1))
	})
})
