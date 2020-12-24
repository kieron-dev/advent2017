package cards

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type Stack struct {
	slice []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Pop() int {
	if len(s.slice) == 0 {
		log.Fatalf("cannot pop off an empty stack")
	}

	top := s.slice[0]
	s.slice = s.slice[1:]

	return top
}

func (s *Stack) Push(e int) {
	s.slice = append(s.slice, e)
}

func (s Stack) Elements() []int {
	return s.slice[:]
}

func (s Stack) Empty() bool {
	return len(s.slice) == 0
}

func (s Stack) Len() int {
	return len(s.slice)
}

func (s *Stack) Set(elems []int) {
	s.slice = make([]int, len(elems))
	copy(s.slice, elems)
}

func (s Stack) String() string {
	r := ""
	for _, n := range s.slice {
		r += strconv.Itoa(n) + ","
	}

	return r
}

type Combat struct {
	stack1       *Stack
	stack2       *Stack
	recursive    bool
	visited      map[string]bool
	infiniteLoop bool
}

func NewCombat() Combat {
	return Combat{
		stack1:  NewStack(),
		stack2:  NewStack(),
		visited: map[string]bool{},
	}
}

func (c *Combat) SetRecursive(r bool) {
	c.recursive = r
}

func (c *Combat) SetStacks(s1, s2 []int) {
	c.stack1.Set(s1)
	c.stack2.Set(s2)
}

func (c *Combat) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	currentPlayer := 1

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if line == "Player 1:" {
			continue
		}

		if line == "Player 2:" {
			currentPlayer = 2
			continue
		}

		card, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("atoi failed: %v", err)
		}

		if currentPlayer == 1 {
			c.stack1.Push(card)
		} else {
			c.stack2.Push(card)
		}
	}
}

func (c Combat) Stack(i int) []int {
	if i == 1 {
		return c.stack1.Elements()
	}

	return c.stack2.Elements()
}

func (c *Combat) Play() {
	key := fmt.Sprintf("%s:%s", c.stack1.String(), c.stack2.String())
	if c.visited[key] {
		c.infiniteLoop = true
		return
	}
	c.visited[key] = true

	t1 := c.stack1.Pop()
	t2 := c.stack2.Pop()

	if c.recursive && t1 <= c.stack1.Len() && t2 <= c.stack2.Len() {
		g := NewCombat()
		g.SetStacks(c.stack1.Elements()[:t1], c.stack2.Elements()[:t2])
		for g.Winner() == 0 {
			g.Play()
		}
		winner := g.Winner()

		if winner == 1 {
			c.stack1.Push(t1)
			c.stack1.Push(t2)
		} else {
			c.stack2.Push(t2)
			c.stack2.Push(t1)
		}

		return
	}

	if t1 > t2 {
		c.stack1.Push(t1)
		c.stack1.Push(t2)
	} else {
		c.stack2.Push(t2)
		c.stack2.Push(t1)
	}
}

func (c Combat) Winner() int {
	if c.infiniteLoop {
		return 1
	}

	if c.stack1.Empty() {
		return 2
	}

	if c.stack2.Empty() {
		return 1
	}

	return 0
}

func (c Combat) Score(i int) int {
	var elems []int

	if i == 1 {
		elems = c.stack1.Elements()
	} else {
		elems = c.stack2.Elements()
	}

	score := 0
	mult := len(elems)

	for _, n := range elems {
		score += mult * n
		mult--
	}

	return score
}
