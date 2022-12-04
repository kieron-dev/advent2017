package two022_test

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("04", func() {
	It("does part A", func() {
		f, err := os.Open("input04")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		count := 0
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			var a1, a2, b1, b2 int
			_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &a1, &a2, &b1, &b2)
			Expect(err).NotTo(HaveOccurred())

			if (a1 >= b1 && a2 <= b2) || (b1 >= a1 && b2 <= a2) {
				count++
			}
		}

		Expect(count).To(Equal(464))
	})

	It("does part B", func() {
		f, err := os.Open("input04")
		Expect(err).NotTo(HaveOccurred())
		defer f.Close()

		count := 0
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			var a1, a2, b1, b2 int
			_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &a1, &a2, &b1, &b2)
			Expect(err).NotTo(HaveOccurred())

			if (a1 >= b1 && a1 <= b2) || (a2 >= b1 && a2 <= b2) ||
				(b1 >= a1 && b1 <= a2) || (b2 >= a1 && b2 <= a2) {
				count++
			}
		}

		Expect(count).To(Equal(770))
	})
})
