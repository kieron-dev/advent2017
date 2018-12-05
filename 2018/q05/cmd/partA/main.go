package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	chain := string(bytes)
	stack := NewStack()
	for _, r := range chain {
		prev := stack.Peek()
		if r != prev && strings.Upper(r) == strings.Upper(prev) {
			stack.Pop()
		} else {
			stack.Push(r)
		}
	}
	fmt.Println(stack.Len())
}

type Stack struct {
	items  []rune
	length int
}

func NewStack() Stack {
	s := Stack{}
	s.items = []rune{}
	return s
}

func (s Stack) Peek() rune {
	if s.length == 0 {
		return rune('\0')
	}
	return s.items[s.length-1]
}

func (s Stack) Push(r rune) {
	s.items = append(s.items, r)
	s.length++
}

func (s Stack) Pop() rune {
	if s.length == 0 {
		log.Fatal("empty stack")
	}
	r := s.Peek()
	s.length--
	s.items = s.items[0:s.length]
	return r
}

func (s Stack) Len() int {
	return s.length
}
