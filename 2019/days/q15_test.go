package days_test

import (
	"io/ioutil"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019/repairbot"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q15", func() {
	var (
		bot   *repairbot.Bot
		bytes []byte
	)

	BeforeEach(func() {
		bot = repairbot.New()
		var err error
		bytes, err = ioutil.ReadFile("input15")
		if err != nil {
			panic(err)
		}
	})

	It("does part A", func() {
		oxygenDist := bot.RunProg(strings.TrimSpace(string(bytes)))
		Expect(oxygenDist).To(Equal(298))
	})

	It("does part B", func() {
		bot.RunProg(strings.TrimSpace(string(bytes)))
		Expect(bot.TimeToFillArea()).To(Equal(346))
	})
})
