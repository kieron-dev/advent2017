package dance

import (
	"strconv"
	"strings"
)

type Dancers map[rune]int

func New(numDancers int) Dancers {
	d := Dancers{}
	label := 'a'
	for i := 0; i < numDancers; i++ {
		d[label] = i
		label++
	}
	return d
}

func (d Dancers) Spin(num int) {
	l := len(d)
	for k, v := range d {
		d[k] = (v + num) % l
	}
}

func (d Dancers) Exchange(a, b int) {
	var (
		achar rune
		bchar rune
	)
	for k, v := range d {
		if v == a {
			achar = k
		}
		if v == b {
			bchar = k
		}
	}
	d.Swap(achar, bchar)
}

func (d Dancers) Swap(a, b rune) {
	apos := d[a]
	d[a] = d[b]
	d[b] = apos
}

func (d Dancers) Print() string {
	list := make([]rune, len(d))
	for k, v := range d {
		list[v] = k
	}
	s := ""
	for _, r := range list {
		s += string(r)
	}
	return s
}

func (d Dancers) Move(move string) {
	switch move[0] {
	case 's':
		n, _ := strconv.Atoi(string(move[1:]))
		d.Spin(n)
	case 'x':
		nums := strings.Split(string(move[1:]), "/")
		from, _ := strconv.Atoi(nums[0])
		to, _ := strconv.Atoi(nums[1])
		d.Exchange(from, to)
	case 'p':
		chars := strings.Split(string(move[1:]), "/")
		d.Swap(rune(chars[0][0]), rune(chars[1][0]))
	}
}

func (d Dancers) ProcessMoves(moves []string) {
	for _, m := range moves {
		d.Move(m)
	}
}

func (d Dancers) IsOriginalOrder() bool {
	for i := 0; i < len(d); i++ {
		if d[rune('a'+i)] != i {
			return false
		}
	}
	return true
}
