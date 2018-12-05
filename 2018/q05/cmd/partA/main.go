package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	chain := string(bytes)
	l := reduce(chain)
	fmt.Println(l)

	minL := len(chain)
	for r := rune('a'); r <= rune('z'); r++ {
		cr := unicode.ToUpper(r)
		test := strings.Replace(chain, string(r), "", -1)
		test = strings.Replace(test, string(cr), "", -1)
		l = reduce(test)
		if l < minL {
			minL = l
		}
	}
	fmt.Printf("minL = %+v\n", minL)
}

func reduce(chain string) int {
	stack := NewStack()
	for _, r := range chain {
		if r == rune('\n') {
			continue
		}
		prev := stack.Peek()
		if r != prev && strings.ToUpper(string(r)) == strings.ToUpper(string(prev)) {
			stack.Pop()
		} else {
			stack.Push(r)
		}
	}
	return stack.Len()
}

type Stack struct {
	items  []rune
	length int
}

func NewStack() *Stack {
	s := Stack{}
	s.items = []rune{}
	return &s
}

func (s *Stack) Peek() rune {
	if s.length == 0 {
		return rune(0)
	}
	return s.items[s.length-1]
}

func (s *Stack) Push(r rune) {
	s.items = append(s.items, r)
	s.length++
}

func (s *Stack) Pop() rune {
	if s.length == 0 {
		log.Fatal("empty stack")
	}
	r := s.Peek()
	s.length--
	s.items = s.items[0:s.length]
	return r
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Print() {
	i := 0
	for _, r := range s.items {
		i++
		fmt.Printf("%c", r)
	}
	fmt.Printf("i = %+v\n", i)
	fmt.Println()
}
