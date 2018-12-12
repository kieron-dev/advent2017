package q12

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"strings"
)

type Plants struct {
	state map[int]bool
	rules [][5]bool
	min   int
	max   int
}

func NewPlants(r io.Reader) *Plants {
	p := Plants{}
	p.state = map[int]bool{}
	p.rules = [][5]bool{}

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

func (p *Plants) LoadRule(rule string) {
	if !strings.Contains(rule, "=>") {
		return
	}

}

func (p *Plants) LoadState(state string) {
	for i, r := range state {
		if r == rune('#') {
			p.state[i] = true
		}
		p.max = i
	}
}

func (p *Plants) State() string {
	var buf bytes.Buffer
	for i := p.min; i < p.max+1; i++ {
		if p.state[i] {
			buf.WriteString("#")
		} else {
			buf.WriteString(".")
		}
	}
	return buf.String()
}

func (p *Plants) Rules() []string {
	return []string{}
}
