package patterns

import (
	"strings"
)

type Art struct {
	pattern []string
	rules   map[string]string
}

func New() *Art {
	art := Art{}
	art.pattern = []string{
		".#.",
		"..#",
		"###",
	}
	art.rules = map[string]string{}
	return &art
}

func (a *Art) Size() int {
	return len(a.pattern)
}

func (a *Art) Pattern() []string {
	return a.pattern
}

func (a *Art) GetKeys(x, y int) []string {
	square := a.GetSquare(x, y)
	out := []string{}
	for i := 0; i < 2; i++ {
		if i == 1 {
			FlipSquare(square)
		}
		for j := 0; j < 4; j++ {
			out = append(out, SquareToString(square))
			RotateSquare(square)
		}
	}
	return out
}

func (a *Art) getComponentSize() int {
	l := 3
	if len(a.pattern)%2 == 0 {
		l = 2
	}
	return l
}

func (a *Art) GetSquare(x, y int) [][]byte {
	l := a.getComponentSize()
	sq := make([][]byte, l)
	for i := 0; i < l; i++ {
		sq[i] = make([]byte, l)
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			sq[i][j] = a.pattern[x*l+i][y*l+j]
		}
	}
	return sq
}

func (a *Art) AddRule(match, newPattern string) {
	a.rules[match] = newPattern
}

func (a *Art) GetNewPattern(x, y int) []string {
	for _, key := range a.GetKeys(x, y) {
		if pattern, ok := a.rules[key]; ok {
			return strings.Split(pattern, "/")
		}
	}
	return []string{}
}

func (a *Art) Advance() {
	newPattern := []string{}
	l := a.getComponentSize()
	s := a.Size()
	for i := 0; i*l < s; i++ {
		for j := 0; j*l < s; j++ {
			p := a.GetNewPattern(i, j)
			if len(newPattern) < (i+1)*(l+1) {
				newPattern = append(newPattern, p...)
			} else {
				for k, line := range p {
					newPattern[(l+1)*i+k] += line
				}
			}
		}
	}
	a.pattern = newPattern
}

func (a *Art) OnCount() int {
	c := 0
	for _, s := range a.pattern {
		c += strings.Count(s, "#")
	}
	return c
}

func RotateSquare(sq [][]byte) {
	l := len(sq) - 1
	for i := 0; i < l; i++ {
		saved := sq[0][i]
		sq[0][i] = sq[l-i][0]
		sq[l-i][0] = sq[l][l-i]
		sq[l][l-i] = sq[i][l]
		sq[i][l] = saved
	}
}

func FlipSquare(sq [][]byte) {
	l := len(sq)
	for i := 0; i < l; i++ {
		save := sq[i][0]
		sq[i][0] = sq[i][l-1]
		sq[i][l-1] = save
	}
}

func SquareToString(sq [][]byte) string {
	l := len(sq)
	s := ""
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			s += string(sq[i][j])
		}
		if i < l-1 {
			s += "/"
		}
	}
	return s
}
