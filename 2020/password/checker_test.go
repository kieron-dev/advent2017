package password_test

import (
	"io"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kieron-dev/adventofcode/2020/password"
)

var _ = Describe("Checker", func() {
	var (
		checker password.Checker
		input   io.Reader
	)

	BeforeEach(func() {
		input = strings.NewReader(`
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`)
		checker = password.NewChecker()
		checker.Load(input)
	})

	It("can verify passwords", func() {
		Expect(checker.CorrectCount()).To(Equal(2))
	})
})
