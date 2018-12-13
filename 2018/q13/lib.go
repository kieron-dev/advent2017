package q13

import (
	"bufio"
	"io"
	"log"
	"sort"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Cart struct {
	Row           int
	Col           int
	Dir           Direction
	Intersections int
	Dead          bool
}

type Mine struct {
	Map   []string
	Carts []*Cart
}

func NewMine(r io.Reader) *Mine {
	m := Mine{}
	br := bufio.NewReader(r)
	row := 0
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		line = strings.Trim(line, "\n")

		for col := 0; col < len(line); col++ {
			if line[col] == '>' {
				m.AddCart(&Cart{Row: row, Col: col, Dir: Right})
			} else if line[col] == '<' {
				m.AddCart(&Cart{Row: row, Col: col, Dir: Left})
			} else if line[col] == '^' {
				m.AddCart(&Cart{Row: row, Col: col, Dir: Up})
			} else if line[col] == 'v' {
				m.AddCart(&Cart{Row: row, Col: col, Dir: Down})
			}
		}

		line = strings.Replace(line, ">", "-", -1)
		line = strings.Replace(line, "<", "-", -1)
		line = strings.Replace(line, "v", "|", -1)
		line = strings.Replace(line, "^", "|", -1)
		m.Map = append(m.Map, line)

		row++
	}

	return &m
}

func (m *Mine) AddCart(cart *Cart) {
	m.Carts = append(m.Carts, cart)
}

func (m *Mine) SortCarts() {
	sort.Slice(m.Carts, func(i, j int) bool {
		a := m.Carts[i]
		b := m.Carts[j]
		if a.Row == b.Row {
			return a.Col < b.Col
		}
		return a.Row < b.Row
	})
}

func (m *Mine) Move(c *Cart) {
	switch c.Dir {
	case Up:
		c.Row--
		switch m.Map[c.Row][c.Col] {
		case '/':
			c.Dir = Right
		case '\\':
			c.Dir = Left
		case '+':
			switch c.Intersections {
			case 0:
				c.Dir = Left
			case 2:
				c.Dir = Right
			}
			c.Intersections++
			c.Intersections %= 3
		}

	case Down:
		c.Row++
		switch m.Map[c.Row][c.Col] {
		case '/':
			c.Dir = Left
		case '\\':
			c.Dir = Right
		case '+':
			switch c.Intersections {
			case 0:
				c.Dir = Right
			case 2:
				c.Dir = Left
			}
			c.Intersections++
			c.Intersections %= 3
		}

	case Left:
		c.Col--
		switch m.Map[c.Row][c.Col] {
		case '/':
			c.Dir = Down
		case '\\':
			c.Dir = Up
		case '+':
			switch c.Intersections {
			case 0:
				c.Dir = Down
			case 2:
				c.Dir = Up
			}
			c.Intersections++
			c.Intersections %= 3
		}

	case Right:
		c.Col++
		switch m.Map[c.Row][c.Col] {
		case '/':
			c.Dir = Up
		case '\\':
			c.Dir = Down
		case '+':
			switch c.Intersections {
			case 0:
				c.Dir = Up
			case 2:
				c.Dir = Down
			}
			c.Intersections++
			c.Intersections %= 3
		}
	}
}

func (m *Mine) MoveCarts() *Cart {
	for _, cart := range m.Carts {
		m.Move(cart)
		if m.HasCrashed(cart) {
			return cart
		}
	}
	m.SortCarts()
	return nil
}

func (m *Mine) MoveCartsRemovingCrashes() {
	for _, cart := range m.Carts {
		if cart.Dead {
			continue
		}
		m.Move(cart)
		if m.HasCrashed(cart) {
			m.MarkCrashed(cart)
		}
	}
	m.SortCarts()
}

func (m *Mine) MarkCrashed(cart *Cart) {
	for _, c := range m.Carts {
		if c.Row == cart.Row && c.Col == cart.Col {
			c.Dead = true
		}
	}
}

func (m *Mine) HasCrashed(cart *Cart) bool {
	for _, c := range m.Carts {
		if c == cart || c.Dead {
			continue
		}
		if c.Row == cart.Row && c.Col == cart.Col {
			return true
		}
	}
	return false
}

func (m *Mine) RunTillOneLeft() (iteration int, cart *Cart) {
	for {
		iteration++
		m.MoveCartsRemovingCrashes()
		l := 0
		cart = nil
		for _, c := range m.Carts {
			if !c.Dead {
				l++
				cart = c
			}
		}

		if l <= 1 {
			return
		}
	}
}

func (m *Mine) RunTillCrash() (iteration int, crashedCart *Cart) {
	for {
		iteration++
		crashedCart = m.MoveCarts()
		if crashedCart != nil {
			return
		}
	}
}
