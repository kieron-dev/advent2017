// Package customs deals with plane custom form submissions
package customs

import (
	"bufio"
	"io"
	"strings"
)

type Group struct {
	size    int
	answers map[rune]int
}

func NewGroup(size int) Group {
	return Group{
		size: size,
	}
}

func (g *Group) Load(answers string) {
	g.answers = map[rune]int{}

	for _, q := range answers {
		g.answers[q]++
	}
}

func (g Group) LenAnyAnswered() int {
	return len(g.answers)
}

func (g Group) LenAllAnswered() int {
	n := 0

	for _, c := range g.answers {
		if c == g.size {
			n++
		}
	}

	return n
}

type Forms struct {
	groups []Group
}

func NewForms() Forms {
	return Forms{}
}

func (f *Forms) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	buf := ""
	size := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" && size > 0 {
			f.AddGroup(buf, size)
			buf = ""
			size = 0
			continue
		}

		if line == "" {
			continue
		}

		buf += line
		size++
	}

	f.AddGroup(buf, size)
}

func (f *Forms) AddGroup(answers string, size int) {
	g := NewGroup(size)
	g.Load(answers)

	f.groups = append(f.groups, g)
}

func (f Forms) GroupCount() int {
	return len(f.groups)
}

func (f Forms) GroupSum() int {
	n := 0

	for _, g := range f.groups {
		n += g.LenAnyAnswered()
	}

	return n
}

func (f Forms) WholeGroupSum() int {
	n := 0

	for _, g := range f.groups {
		n += g.LenAllAnswered()
	}

	return n
}
