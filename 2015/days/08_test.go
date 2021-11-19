package days_test

import (
	"bufio"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("08", func() {
	It("does part A", func() {
		file, err := os.Open("input08")
		Expect(err).NotTo(HaveOccurred())
		defer file.Close()

		scanner := bufio.NewScanner(file)
		diff := 0

		for scanner.Scan() {
			line := scanner.Text()

			l := len(line)
			chars := 0
			for i := 1; i < l-1; i++ {
				chars++
				if line[i] != '\\' {
					continue
				}

				if line[i+1] == '\\' || line[i+1] == '"' {
					i += 1
					continue
				}

				if line[i+1] == 'x' {
					i += 3
					continue
				}
			}

			diff += l - chars
		}

		Expect(diff).To(Equal(1342))
	})

	It("does part B", func() {
		file, err := os.Open("input08")
		Expect(err).NotTo(HaveOccurred())
		defer file.Close()

		scanner := bufio.NewScanner(file)
		diff := 0

		for scanner.Scan() {
			line := scanner.Text()

			l := len(line)
			chars := 2
			for i := 0; i < l; i++ {
				chars++
				switch line[i] {
				case '"':
					fallthrough
				case '\\':
					chars++
				}
			}

			diff += chars - l
		}

		Expect(diff).To(Equal(2074))
	})
})
