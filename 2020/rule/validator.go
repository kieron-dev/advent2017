// Package rule validates stuff
package rule

import (
	"bufio"
	"io"
	"strings"
)

type Validator struct {
	rules   map[string]string
	entries []string
}

func NewValidator() Validator {
	return Validator{
		rules: map[string]string{},
	}
}

func (v Validator) IsValid(idx int) bool {
	str := v.entries[idx]

	remainders := v.CheckRule("0", []string{str})
	for _, r := range remainders {
		if r == "" {
			return true
		}
	}

	return false
}

func (v Validator) CheckRule(idx string, strs []string) (remainders []string) {
	okRemainders := []string{}
	rule := v.rules[idx]

	for _, str := range strs {
		if str == "" {
			continue
		}

		if rule[0] == '"' {
			if str[0] == rule[1] {
				okRemainders = append(okRemainders, str[1:])
			}

			continue
		}

		for _, subrule := range strings.Split(rule, "|") {
			workStrs := []string{str}

			for _, ruleNum := range strings.Split(strings.TrimSpace(subrule), " ") {
				workStrs = v.CheckRule(ruleNum, workStrs)

				if len(workStrs) == 0 {
					break
				}

			}

			okRemainders = append(okRemainders, workStrs...)
		}
	}

	return okRemainders
}

func (v *Validator) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			v.rules[parts[0]] = strings.TrimSpace(parts[1])
			continue
		}

		v.entries = append(v.entries, line)
	}
}

func (v Validator) ValidCount() int {
	n := 0

	for i := range v.entries {
		if v.IsValid(i) {
			n++
		}
	}

	return n
}

func (v *Validator) SetNewRules() {
	v.rules["8"] = "42 | 42 8"
	v.rules["11"] = "42 31 | 42 11 31"
}
