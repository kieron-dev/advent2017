package days_test

import (
	"io/ioutil"
	"strings"

	"github.com/kieron-dev/advent2017/advent2019/vacuumbot"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Q17", func() {
	var (
		s    *vacuumbot.System
		prog string
	)

	BeforeEach(func() {
		s = vacuumbot.NewSystem()

		contents, err := ioutil.ReadFile("input17")
		if err != nil {
			panic(err)
		}
		prog = strings.TrimSpace(string(contents))
	})

	It("can get some output", func() {
		s.SetProg(prog)
		s.Run()
		s.AcquireGrid()

		sum := 0
		for _, x := range s.GetCrossovers() {
			sum += x.X() * x.Y()
		}
		Expect(sum).To(Equal(10064))
	})

	It("can control the robot", func() {
		s.SetProg(prog)
		s.Poke(0, 2)
		s.Run()

		buf := gbytes.NewBuffer()
		ch := make(chan int)
		go func() {
			for n := range s.GetOutputChan() {
				if n < 256 {
					_, err := buf.Write([]byte{byte(n)})
					if err != nil {
						panic(err)
					}

				} else {
					ch <- n
				}
			}
		}()

		Eventually(buf).Should(gbytes.Say("Main:"))
		s.Input("A,A,B,C,B,C,B,C,B,A")

		Eventually(buf).Should(gbytes.Say("Function A:"))
		s.Input("L,10,L,8,R,8,L,8,R,6")

		Eventually(buf).Should(gbytes.Say("Function B:"))

		s.Input("R,6,R,8,R,8")

		Eventually(buf).Should(gbytes.Say("Function C:"))
		s.Input("R,6,R,6,L,8,L,10")

		Eventually(buf).Should(gbytes.Say("Continuous video feed?"))
		s.Input("n")

		answer := <-ch
		Expect(answer).To(Equal(1197725))
	})
})
