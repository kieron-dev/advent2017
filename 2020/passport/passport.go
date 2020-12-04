// Package passport handles passport machines and validation
package passport

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func NewPassport(s string) Passport {
	var p Passport

	for _, pair := range strings.Split(s, " ") {
		if pair == "" {
			continue
		}

		items := strings.Split(pair, ":")
		if len(items) != 2 {
			log.Fatalf("parsing failed for %q", pair)
		}

		k, v := items[0], items[1]

		switch k {
		case "byr":
			p.BirthYear = v
		case "iyr":
			p.IssueYear = v
		case "eyr":
			p.ExpirationYear = v
		case "hgt":
			p.Height = v
		case "hcl":
			p.HairColor = v
		case "ecl":
			p.EyeColor = v
		case "pid":
			p.PassportID = v
		case "cid":
			p.CountryID = v
		}
	}

	return p
}

func (p Passport) IsValid() bool {
	return p.BirthYear != "" &&
		p.IssueYear != "" &&
		p.ExpirationYear != "" &&
		p.Height != "" &&
		p.EyeColor != "" &&
		p.HairColor != "" &&
		p.PassportID != ""
}

func (p Passport) IsStrictlyValid() bool {
	return validInt(p.BirthYear, 4, 1920, 2002) &&
		validInt(p.IssueYear, 4, 2010, 2020) &&
		validInt(p.ExpirationYear, 4, 2020, 2030) &&
		(validLen(p.Height, "cm", 150, 193) || validLen(p.Height, "in", 59, 76)) &&
		validRGB(p.HairColor) &&
		validColor(p.EyeColor) &&
		validNumber(p.PassportID, 9)
}

func validInt(s string, l, min, max int) bool {
	if len(s) != l {
		return false
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return n >= min && n <= max
}

func validLen(s, unit string, min, max int) bool {
	if !strings.HasSuffix(s, unit) {
		return false
	}

	s = s[:len(s)-len(unit)]

	n, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return n >= min && n <= max
}

type Manager struct {
	passports []Passport
}

var reRGB = regexp.MustCompile(`^#[0-9a-f]{6}$`)

func validRGB(s string) bool {
	return reRGB.MatchString(s)
}

var validColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func validColor(s string) bool {
	for _, c := range validColors {
		if s == c {
			return true
		}
	}

	return false
}

func validNumber(s string, l int) bool {
	reNumLen := regexp.MustCompile(fmt.Sprintf(`^\d{%d}$`, l))
	return reNumLen.MatchString(s)
}

func NewManager() Manager {
	return Manager{}
}

func (m *Manager) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	buf := ""
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" && buf != "" {
			m.loadPassport(buf)
			buf = ""
			continue
		}
		buf += " " + line
	}

	m.loadPassport(buf)
}

func (m *Manager) loadPassport(passportStr string) {
	p := NewPassport(passportStr)
	m.passports = append(m.passports, p)
}

func (m Manager) Count() int {
	return len(m.passports)
}

func (m Manager) ValidCount() int {
	n := 0

	for _, p := range m.passports {
		if p.IsValid() {
			n++
		}
	}

	return n
}

func (m Manager) StrictValidCount() int {
	n := 0

	for _, p := range m.passports {
		if p.IsStrictlyValid() {
			n++
		}
	}

	return n
}
