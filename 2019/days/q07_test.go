package days_test

import (
	"io/ioutil"
	"strings"

	"github.com/kieron-dev/advent2017/advent2019"
	"github.com/kieron-dev/advent2017/advent2019/intcode"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q07", func() {
	var (
		contents   []byte
		permHelper advent2019.Perms
	)

	BeforeEach(func() {
		var err error
		contents, err = ioutil.ReadFile("./input07")
		if err != nil {
			panic(err)
		}
		permHelper = advent2019.Perms{}
	})

	It("does part A", func() {
		max := 0

		for _, perm := range permHelper.All([]int{0, 1, 2, 3, 4}) {
			arr := intcode.NewArray(5)
			arr.SetProgram(strings.TrimSpace(string(contents)))
			arr.SetPhase(perm)
			arr.WriteInitialInput(0)
			arr.Run()
			val := arr.GetResult()
			if val > max {
				max = val
			}
		}
		Expect(max).To(Equal(262086))
	})

	It("does part B", func() {
		max := 0

		for _, perm := range permHelper.All([]int{5, 6, 7, 8, 9}) {
			arr := intcode.NewFeedbackArray(5)
			arr.SetProgram(strings.TrimSpace(string(contents)))
			arr.SetPhase(perm)
			arr.WriteInitialInput(0)
			arr.Run()
			val := arr.GetResult()
			if val > max {
				max = val
			}
		}
		Expect(max).To(Equal(5371621))
	})
})
