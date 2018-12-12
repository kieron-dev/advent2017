package q12

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

type Plants struct {
	state     map[int]rune
	rules     map[string]string
	min       int
	max       int
	iteration int
}

func (p *Plants) Step() {
	p.iteration++
	newState := map[int]rune{}
	newMin := p.max
	newMax := p.min
	for i := p.min - 5; i <= p.max+1; i++ {
		var buf bytes.Buffer
		for j := 0; j < 5; j++ {
			r, ok := p.state[i+j]
			if !ok {
				r = rune('.')
			}
			buf.WriteString(string(r))
		}
		match := buf.String()
		if p.rules[match] == "#" {
			newState[i+2] = rune('#')
			if i+2 < newMin {
				newMin = i + 2
			}
			newMax = i + 2
		} else {
			newState[i+2] = rune('.')
		}
	}
	p.state = newState
	p.min = newMin
	p.max = newMax
}

func NewPlants(r io.Reader) *Plants {
	p := Plants{}
	p.state = map[int]rune{}
	p.rules = map[string]string{}

	br := bufio.NewReader(r)
	line, err := br.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")

	p.LoadState(parts[2])

	for {
		line, err = br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		p.LoadRule(strings.TrimSpace(line))
	}

	return &p
}

func (p *Plants) HashPosSum() int {
	sum := 0
	for i, r := range p.state {
		if r == rune('#') {
			sum += i
		}
	}
	return sum
}

func (p *Plants) LoadRule(rule string) {
	if !strings.Contains(rule, "=>") {
		return
	}
	var matchStr, resultStr string
	fmt.Sscanf(rule, "%s => %s", &matchStr, &resultStr)
	p.rules[matchStr] = resultStr
}

func (p *Plants) LoadState(state string) {
	for i, r := range state {
		p.state[i] = r
		p.max = i
	}
}

func (p *Plants) State() string {
	var buf bytes.Buffer
	for i := p.min; i < p.max+1; i++ {
		buf.WriteString(string(p.state[i]))
	}
	return buf.String()
}

func (p *Plants) Rules() map[string]string {
	return p.rules
}
