package repairbot_test

import (
	"io/ioutil"
	"strings"

	"github.com/kieron-dev/advent2017/advent2019/grid"
	"github.com/kieron-dev/advent2017/advent2019/repairbot"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repairbot", func() {

	var (
		bot        *repairbot.Bot
		bytes      []byte
		oxygenDist int
	)

	BeforeEach(func() {
		bot = repairbot.New()
		var err error
		bytes, err = ioutil.ReadFile("../days/input15")
		if err != nil {
			panic(err)
		}
	})

	JustBeforeEach(func() {
		oxygenDist = bot.RunProg(strings.TrimSpace(string(bytes)))
	})

	It("gets a response to input", func() {
		status := bot.Move(grid.North)
		Expect(status).To(BeNumerically(">=", 0))
		Expect(status).To(BeNumerically("<=", 2))
	})

	Context("running graphically", func() {
		It("can run graphically", func() {
			Expect(oxygenDist).To(Equal(298))
		})
	})

})
