// Package password verifies password policy
package password

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Password struct {
	from, to int
	char     string
	password string
}

type Checker struct {
	passwords []Password
}

func NewChecker() Checker {
	return Checker{}
}

func (c *Checker) Load(data io.Reader) {
	re := regexp.MustCompile(`(\d+)-(\d+) (.): (.*)`)

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		matches := re.FindStringSubmatch(line)
		if matches == nil {
			log.Fatalf("regex didn't match line: %s", line)
		}

		from, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Fatalf("strconv.Atoi problem with %q", matches[1])
		}

		to, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatalf("strconv.Atoi problem with %q", matches[2])
		}

		c.passwords = append(c.passwords, Password{
			from:     from,
			to:       to,
			char:     matches[3],
			password: matches[4],
		})

	}
}

func (c Checker) CorrectCount() int {
	n := 0

	for _, pw := range c.passwords {
		if pw.IsValid() {
			n++
		}
	}

	return n
}

func (c Checker) CorrectCountNew() int {
	n := 0

	for _, pw := range c.passwords {
		if pw.IsValidNew() {
			n++
		}
	}

	return n
}

func (p Password) IsValid() bool {
	n := 0

	for _, c := range p.password {
		if byte(c) == p.char[0] {
			n++
		}
	}

	return p.from <= n && n <= p.to
}

func (p Password) IsValidNew() bool {
	match1 := p.password[p.from-1] == p.char[0]
	match2 := p.password[p.to-1] == p.char[0]

	return (match1 || match2) && !(match1 && match2)
}
