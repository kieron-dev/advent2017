package days_test

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"

	"github.com/kieron-dev/advent2017/advent2019/springdroid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Q25", func() {
	var (
		d      *springdroid.Droid
		outBuf *gbytes.Buffer
	)

	BeforeEach(func() {
		d = springdroid.NewDroid()
		input, err := ioutil.ReadFile("input25")
		if err != nil {
			panic(err)
		}
		d.LoadProgram(strings.TrimSpace(string(input)))
		outBuf = gbytes.NewBuffer()
	})

	JustBeforeEach(func() {
		d.RunProgram()
	})

	It("can do part A", func() {
		go func() {
			for {
				b := d.Output()
				if b < 256 {
					fmt.Fprintf(outBuf, "%c", b)
				}
			}
		}()

		enterCommand := func(command string) {
			Eventually(outBuf).Should(gbytes.Say("Command?"))
			d.Input(command)
		}

		for _, cmd := range []string{
			"south", "take space law space brochure", "south", "take mouse", "south",
			"take astrolabe", "south", "take mug", "inv", "north", "north", "west", "north",
			"north", "take wreath", "south", "south", "east", "north", "west", "take sand",
			"north", "take manifold", "south", "west", "take monolith", "west",
		} {
			enterCommand(cmd)
		}

		items := []string{
			"monolith", "wreath", "mug", "astrolabe", "manifold", "sand",
			"mouse", "space law space brochure",
		}

		myItems := items[:]

	outer:
		for i := 0; i <= 256; i++ {
			for _, item := range myItems {
				enterCommand("drop " + item)
			}

			myItems = []string{}

			for j := 0; j < 8; j++ {
				if i&(1<<j) > 0 {
					enterCommand("take " + items[j])
					myItems = append(myItems, items[j])
				}
			}
			enterCommand("west")

			select {
			case <-outBuf.Detect("ejected"):
				continue outer
			case <-time.After(100 * time.Millisecond):
				break outer
			}

		}

		re := regexp.MustCompile(`typing (\d+) on`)
		matches := re.FindSubmatch(outBuf.Contents())
		Expect(matches).To(HaveLen(2))
		Expect(string(matches[1])).To(Equal("328960"))
	})
})
