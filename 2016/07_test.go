package twentysixteen_test

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type IPAddrPart string

func (i IPAddrPart) HasABBA() bool {
	for j := 3; j < len(i); j++ {
		if i[j-3] == i[j] && i[j-2] == i[j-1] && i[j-1] != i[j] {
			return true
		}
	}
	return false
}

func (i IPAddrPart) GetABAs() []string {
	var ret []string
	for j := 2; j < len(i); j++ {
		if i[j-2] == i[j] && i[j-1] != i[j] {
			ret = append(ret, string(i[j-2:j+1]))
		}
	}
	return ret
}

func (i IPAddrPart) HasBAB(aba string) bool {
	bab := aba[1:] + aba[1:2]
	return strings.Contains(string(i), bab)
}

type IPAddr struct {
	main  []IPAddrPart
	hyper []IPAddrPart
}

func NewIPAddr(addr string) IPAddr {
	var ip IPAddr

	var cur strings.Builder
	for _, r := range addr {
		if r == '[' {
			ip.main = append(ip.main, IPAddrPart(cur.String()))
			cur = strings.Builder{}
			continue
		}
		if r == ']' {
			ip.hyper = append(ip.hyper, IPAddrPart(cur.String()))
			cur = strings.Builder{}
			continue
		}
		cur.WriteRune(r)
	}

	// we always end outside of a hyper (I think)
	ip.main = append(ip.main, IPAddrPart(cur.String()))

	return ip
}

func (i IPAddr) SupportsTLS() bool {
	for _, h := range i.hyper {
		if h.HasABBA() {
			return false
		}
	}

	for _, m := range i.main {
		if m.HasABBA() {
			return true
		}
	}

	return false
}

func (i IPAddr) SupportsSSL() bool {
	for _, m := range i.main {
		for _, aba := range m.GetABAs() {
			for _, h := range i.hyper {
				if h.HasBAB(aba) {
					return true
				}
			}
		}
	}

	return false
}

func TestSupportsTLS(t *testing.T) {
	type testcase struct {
		ipAddr   string
		expected bool
	}
	testcases := map[string]testcase{
		"ex01": {ipAddr: "abba[mnop]qrst", expected: true},
		"ex02": {ipAddr: "abcd[bddb]xyyx", expected: false},
		"ex03": {ipAddr: "aaaa[qwer]tyui", expected: false},
		"ex04": {ipAddr: "ioxxoj[asdfgh]zxcvbn", expected: true},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewIPAddr(tc.ipAddr).SupportsTLS())
		})
	}
}

func TestDay07a(t *testing.T) {
	in, err := os.Open("input07")
	if err != nil {
		panic(err)
	}
	defer func() { _ = in.Close() }()

	count := 0
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if NewIPAddr(line).SupportsTLS() {
			count++
		}
	}

	assert.Equal(t, 105, count)
}

func TestDay07b(t *testing.T) {
	in, err := os.Open("input07")
	if err != nil {
		panic(err)
	}
	defer func() { _ = in.Close() }()

	count := 0
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if NewIPAddr(line).SupportsSSL() {
			count++
		}
	}

	assert.Equal(t, 258, count)
}
