package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/customs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("06", func() {
	var (
		data  *os.File
		forms customs.Forms
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input06")
		Expect(err).NotTo(HaveOccurred())

		forms = customs.NewForms()
		forms.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("part A", func() {
		Expect(forms.GroupSum()).To(Equal(6903))
	})

	It("part B", func() {
		Expect(forms.WholeGroupSum()).To(Equal(3493))
	})
})
