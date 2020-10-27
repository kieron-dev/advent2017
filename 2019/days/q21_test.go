package days_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/kieron-dev/advent2017/advent2019/springdroid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Q21", func() {
	var (
		s     *springdroid.Droid
		out   *gbytes.Buffer
		m     sync.Mutex
		resCh chan int
	)

	BeforeEach(func() {
		s = springdroid.NewDroid()
		prog, err := ioutil.ReadFile("input21")
		if err != nil {
			panic(err)
		}
		s.LoadProgram(strings.TrimSpace(string(prog)))
		s.RunProgram()
		out = gbytes.NewBuffer()
		resCh = make(chan int)

		go func() {
			for {
				c := s.Output()
				if c < 256 {
					m.Lock()
					out.Write([]byte{byte(c)})
					fmt.Printf("%c", c)
					m.Unlock()
				} else {
					resCh <- c
					return
				}
			}
		}()

		Eventually(out).Should(gbytes.Say("Input instructions:"))
	})

	It("does part A", func() {
		s.Input(`
NOT B J
NOT C T
OR T J
AND D J
NOT A T
OR T J
WALK
`)
		var n int
		Eventually(resCh).Should(Receive(&n))
		Expect(n).To(Equal(19354437))
	})

	It("does part B", func() {
		s.Input(`
NOT B J
NOT C T
OR T J
AND D J
AND H J
NOT A T
OR T J
RUN
`)
		var n int
		Eventually(resCh).Should(Receive(&n))
		Expect(n).To(Equal(1145373084))
	})
})
