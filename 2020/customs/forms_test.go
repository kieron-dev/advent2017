package customs_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/customs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Forms", func() {
	var (
		forms customs.Forms
		data  io.Reader
	)

	BeforeEach(func() {
		data = strings.NewReader(`
abc

a
b
c

ab
ac

a
a
a
a

b
`)
		forms = customs.NewForms()
		forms.Load(data)
	})

	It("creates 5 groups", func() {
		Expect(forms.GroupCount()).To(Equal(5))
	})

	It("can calculate the sum across groups", func() {
		Expect(forms.GroupSum()).To(Equal(11))
	})

	It("can calculate the whole group sum", func() {
		Expect(forms.WholeGroupSum()).To(Equal(6))
	})
})
