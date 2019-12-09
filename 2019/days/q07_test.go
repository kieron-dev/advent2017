package days_test

import (
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019"
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
		max := big.NewInt(0)

		for _, perm := range permHelper.All([]int64{0, 1, 2, 3, 4}) {
			arr := advent2019.NewArray(5)
			arr.SetProgram(strings.TrimSpace(string(contents)))
			arr.SetPhase(perm)
			arr.WriteInitialInput(0)
			arr.Run()
			val := arr.GetResult()
			if val.Cmp(max) > 0 {
				max.Set(val)
			}
		}
		Expect(max.String()).To(Equal("262086"))
	})

	It("does part B", func() {
		max := big.NewInt(0)

		for _, perm := range permHelper.All([]int64{5, 6, 7, 8, 9}) {
			arr := advent2019.NewFeedbackArray(5)
			arr.SetProgram(strings.TrimSpace(string(contents)))
			arr.SetPhase(perm)
			arr.WriteInitialInput(0)
			arr.Run()
			val := arr.GetResult()
			if val.Cmp(max) > 0 {
				max.Set(val)
			}
		}
		Expect(max.String()).To(Equal("5371621"))
	})
})
