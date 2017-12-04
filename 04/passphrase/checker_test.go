package passphrase_test

import (
	"github.com/kieron-pivotal/advent2017/04/passphrase"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Checker", func() {
	It("counts ok lines", func() {
		Expect(passphrase.Check([]string{"aa bb", "aa aa"})).To(Equal(1))
	})
})
