package twentysixteen_test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ex10 = `value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2`

func Test10a(t *testing.T) {
	real, err := os.Open("input10")
	Check(err)

	type tc struct {
		in       io.Reader
		expected int
		low      int
		high     int
	}

	tcs := map[string]tc{
		"ex01": {
			in:       strings.NewReader(ex10),
			expected: 2,
			low:      2,
			high:     5,
		},
		"real": {
			in:       real,
			expected: 98,
			low:      17,
			high:     61,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, botForComparison(tc.in, tc.low, tc.high))
		})
	}
}

type factory struct {
	bots    map[int]*bot
	outputs map[int]*output
}

func NewFactory() factory {
	return factory{
		bots:    map[int]*bot{},
		outputs: map[int]*output{},
	}
}

func (f factory) GetBot(idx int) *bot {
	b, ok := f.bots[idx]
	if !ok {
		b = &bot{id: idx}
		f.bots[idx] = b
	}

	return b
}

func (f factory) GetOutput(idx int) *output {
	o, ok := f.outputs[idx]
	if !ok {
		o = &output{id: idx}
		f.outputs[idx] = o
	}

	return o
}

func (f factory) FollowRules(l, h int) int {
	ret := -1
	for {
		any := false
		for idx, bot := range f.bots {
			if len(bot.chips) == 2 && bot.chips[0] == l && bot.chips[1] == h {
				ret = idx
			}
			if bot.FollowRules(f) {
				any = true
			}
		}
		if !any {
			break
		}
	}
	return ret
}

type bot struct {
	id         int
	chips      []int
	lowToType  string
	highToType string
	lowToIdx   int
	highToIdx  int
}

func (b *bot) Add(val int) {
	b.chips = append(b.chips, val)
	sort.Ints(b.chips)
}

func (b *bot) FollowRules(f factory) bool {
	if len(b.chips) != 2 {
		return false
	}

	if b.lowToType == "bot" {
		to := f.GetBot(b.lowToIdx)
		to.Add(b.chips[0])
	} else {
		to := f.GetOutput(b.lowToIdx)
		to.Add(b.chips[0])
	}

	if b.highToType == "bot" {
		to := f.GetBot(b.highToIdx)
		to.Add(b.chips[1])
	} else {
		to := f.GetOutput(b.highToIdx)
		to.Add(b.chips[1])
	}

	b.chips = []int{}

	return true
}

type output struct {
	id    int
	chips []int
}

func (o *output) Add(idx int) {
	o.chips = append(o.chips, idx)
	sort.Ints(o.chips)
}

func botForComparison(in io.Reader, valueLow, valueHigh int) int {
	factory := NewFactory()

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		switch {
		case strings.Contains(line, "value"):
			bits := strings.Split(line[6:], " goes to bot ")
			val, err := strconv.Atoi(bits[0])
			Check(err)
			toBot, err := strconv.Atoi(bits[1])
			Check(err)
			bot := factory.GetBot(toBot)
			bot.Add(val)
		case strings.Contains(line, "gives"):
			bits1 := strings.Split(line[4:], " gives low to ")
			fromBot, err := strconv.Atoi(bits1[0])
			Check(err)
			bits2 := strings.Split(bits1[1], " and high to ")
			bot := factory.GetBot(fromBot)
			if strings.Contains(bits2[0], "bot") {
				toBot, err := strconv.Atoi(bits2[0][4:])
				Check(err)
				bot.lowToIdx = toBot
				bot.lowToType = "bot"
			} else {
				toBot, err := strconv.Atoi(bits2[0][7:])
				Check(err)
				bot.lowToIdx = toBot
				bot.lowToType = "output"
			}
			if strings.Contains(bits2[1], "bot") {
				toBot, err := strconv.Atoi(bits2[1][4:])
				Check(err)
				bot.highToIdx = toBot
				bot.highToType = "bot"
			} else {
				toBot, err := strconv.Atoi(bits2[1][7:])
				Check(err)
				bot.highToIdx = toBot
				bot.highToType = "output"
			}
		}
	}

	b := factory.FollowRules(valueLow, valueHigh)

	for i, o := range factory.outputs {
		fmt.Println(i, o)
	}

	return b
}
