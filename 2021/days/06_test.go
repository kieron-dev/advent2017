package days_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("06", func() {
	It("does the example", func() {
		in := strings.NewReader("3,4,3,1,2")

		numBytes, err := ioutil.ReadAll(in)
		Expect(err).NotTo(HaveOccurred())
		nums := parseNumList(string(numBytes), ",")

		states := make([]int, 9)
		for _, n := range nums {
			states[n]++
		}

		for i := 0; i < 18; i++ {
			states = iterateLanternFish(states)
		}

		Expect(countLanternFish(states)).To(Equal(26))

		for i := 18; i < 80; i++ {
			states = iterateLanternFish(states)
		}

		Expect(countLanternFish(states)).To(Equal(5934))
	})

	It("does part A", func() {
		in, err := os.Open("input06")
		Expect(err).NotTo(HaveOccurred())
		defer in.Close()

		numBytes, err := ioutil.ReadAll(in)
		Expect(err).NotTo(HaveOccurred())
		numBytes = bytes.TrimSpace(numBytes)
		nums := parseNumList(string(numBytes), ",")

		states := make([]int, 9)
		for _, n := range nums {
			states[n]++
		}

		for i := 0; i < 80; i++ {
			states = iterateLanternFish(states)
		}

		Expect(countLanternFish(states)).To(Equal(365862))
	})

	It("does part B", func() {
		in, err := os.Open("input06")
		Expect(err).NotTo(HaveOccurred())
		defer in.Close()

		numBytes, err := ioutil.ReadAll(in)
		Expect(err).NotTo(HaveOccurred())
		numBytes = bytes.TrimSpace(numBytes)
		nums := parseNumList(string(numBytes), ",")

		states := make([]int, 9)
		for _, n := range nums {
			states[n]++
		}

		for i := 0; i < 256; i++ {
			states = iterateLanternFish(states)
		}

		Expect(countLanternFish(states)).To(Equal(1653250886439))
	})
})

func countLanternFish(states []int) int {
	count := 0
	for _, n := range states {
		count += n
	}

	return count
}

func iterateLanternFish(states []int) []int {
	newStates := make([]int, 9)

	for i := 0; i < 9; i++ {
		newStates[i] = states[(i+1)%9]
	}
	newStates[6] += states[0]

	return newStates
}
