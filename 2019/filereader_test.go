package advent2019_test

import (
	"io"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"advent2019"
)

var _ = Describe("Filereader", func() {
	var (
		file       io.Reader
		fileReader advent2019.FileReader
		fn         func(string)
	)

	Describe("counting lines", func() {
		var count int

		BeforeEach(func() {
			file = strings.NewReader("a\nb\nc")
			fileReader = advent2019.FileReader{}
			fn = func(line string) {
				count++
			}
		})

		It("can apply a function to each row of the file", func() {
			fileReader.Each(file, fn)
			Expect(count).To(Equal(3))
		})
	})
})
