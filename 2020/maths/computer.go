// Package maths has this comment
package maths

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type Computer struct {
	expressions    []string
	plusPrecedence bool
}

func NewComputer(plusPrecedence bool) Computer {
	return Computer{
		plusPrecedence: plusPrecedence,
	}
}

func (c *Computer) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		line = strings.ReplaceAll(line, "(", "( ")
		line = strings.ReplaceAll(line, ")", " )")

		c.expressions = append(c.expressions, line)
	}
}

type operator int

const (
	noop operator = iota
	plus
	mult
)

type Interpreter struct {
	pos    int
	tokens []string
}

func (i *Interpreter) Result() int {
	// grammar
	//
	// expr: term (MUL term)*
	// term: factor (PLUS factor)*
	// factor: (INTEGER | (LPAREN expr RPAREN))

	return i.Expr()
}

func (i *Interpreter) Expr() int {
	val := i.Term()

	for i.tokens[i.pos] == "*" {
		i.pos++
		val *= i.Term()
	}

	return val
}

func (i *Interpreter) Term() int {
	val := i.Factor()

	for i.tokens[i.pos] == "+" {
		i.pos++
		val += i.Factor()
	}

	return val
}

func (i *Interpreter) Factor() int {
	n, err := strconv.Atoi(i.tokens[i.pos])
	if err == nil {
		i.pos++
		return n
	}

	if i.tokens[i.pos] == "(" {
		i.pos++
		val := i.Expr()

		if i.tokens[i.pos] != ")" {
			log.Fatalf("expected ')'")
		}

		i.pos++

		return val
	}

	log.Fatalf("expected an int, or a '('")

	return 0
}

func NewInterpreter(line string) Interpreter {
	return Interpreter{
		tokens: append(strings.Split(line, " "), "EOF"),
	}
}

func (c Computer) Result(idx int) int {
	if c.plusPrecedence {
		interpreter := NewInterpreter(c.expressions[idx])
		return interpreter.Result()
	}

	items := strings.Split(c.expressions[idx], " ")

	opStack := []operator{noop}
	valStack := []int{0}

	for _, item := range items {
		switch item {
		case "+":
			opStack[len(opStack)-1] = plus
		case "*":
			opStack[len(opStack)-1] = mult
		case "(":
			opStack = append(opStack, noop)
			valStack = append(valStack, 0)
		case ")":
			opStack = opStack[:len(opStack)-1]
			val := valStack[len(valStack)-1]
			valStack = valStack[:len(valStack)-1]
			switch opStack[len(opStack)-1] {
			case plus:
				valStack[len(valStack)-1] += val
			case mult:
				valStack[len(valStack)-1] *= val
			default:
				valStack[len(valStack)-1] = val
			}
		default:
			val, err := strconv.Atoi(item)
			if err != nil {
				log.Fatalf("atoi failed: %v", err)
			}

			switch opStack[len(opStack)-1] {
			case plus:
				valStack[len(valStack)-1] += val
			case mult:
				valStack[len(valStack)-1] *= val
			default:
				valStack[len(valStack)-1] = val
			}
		}
	}

	if len(valStack) != 1 {
		log.Fatalf("valStack wrong at exit: %v", valStack)
	}

	if len(opStack) != 1 {
		log.Fatalf("opStack wrong at exit: %v", opStack)
	}

	return valStack[0]
}

func (c Computer) SumResults() int {
	sum := 0

	for i := range c.expressions {
		sum += c.Result(i)
	}

	return sum
}
