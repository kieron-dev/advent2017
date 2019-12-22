package vacuumbot_test

import (
	"io/ioutil"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019/vacuumbot"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vacuumbot", func() {
	var (
		s    *vacuumbot.System
		prog string
	)

	BeforeEach(func() {
		s = vacuumbot.NewSystem()

		contents, err := ioutil.ReadFile("../days/input17")
		if err != nil {
			panic(err)
		}
		prog = strings.TrimSpace(string(contents))
	})

	It("can get some output", func() {
		s.SetProg(prog)
		s.Run()

		s.AcquireGrid()
		xos := s.GetCrossovers()
		Expect(xos).ToNot(BeEmpty())
	})
})
