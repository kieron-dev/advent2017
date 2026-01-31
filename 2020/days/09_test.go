package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/cipher"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("09", func() {
	var (
		data *os.File
		code cipher.Xmas
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input09")
		Expect(err).NotTo(HaveOccurred())

		code = cipher.NewXmas()
		code.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("part A", func() {
		Expect(code.FirstError(25)).To(Equal(22477624))
	})

	It("part B", func() {
		Expect(code.EncryptionWeakness(25)).To(Equal(2980044))
	})
})
