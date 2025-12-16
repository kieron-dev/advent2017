package twentysixteen_test

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Day04a(in io.Reader) int {
	return 0
}

type Room struct {
	letterGroups []string
	number       int
	checksum     string
}

// aaaaa-bbb-z-y-x-123[abxyz]
func NewRoom(code string) Room {
	p := strings.Index(code, "[")
	checksum := code[p+1 : len(code)-1]
	code = code[:p]

	p = strings.LastIndex(code, "-")
	numStr := code[p+1:]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	code = code[:p]

	letterGroups := strings.FieldsFunc(code, func(r rune) bool {
		return r == '-'
	})

	return Room{
		letterGroups: letterGroups,
		number:       num,
		checksum:     checksum,
	}
}

func (r Room) IsValid() bool {
	return r.Checksum() == r.checksum
}

func (r Room) Checksum() string {
	freqs := map[rune]int{}

	for _, g := range r.letterGroups {
		for _, r := range g {
			freqs[r]++
		}
	}

	type runefreq struct {
		r rune
		f int
	}

	freqObjs := make([]runefreq, 0, len(freqs))
	for r, f := range freqs {
		freqObjs = append(freqObjs, runefreq{r, f})
	}

	sort.Slice(freqObjs, func(a, b int) bool {
		if freqObjs[a].f == freqObjs[b].f {
			return freqObjs[a].r < freqObjs[b].r
		}
		return freqObjs[a].f > freqObjs[b].f
	})

	var ret strings.Builder
	for i := range 5 {
		ret.WriteRune(freqObjs[i].r)
	}

	return ret.String()
}

func (r Room) Decrypt() string {
	var out strings.Builder

	for i, g := range r.letterGroups {
		if i > 0 {
			out.WriteRune(' ')
		}
		for _, c := range g {
			c += rune(r.number % 26)
			if c > 'z' {
				c -= 26
			}
			out.WriteRune(c)
		}
	}

	return out.String()
}

func TestIsValid(t *testing.T) {
	type testcase struct {
		code     string
		expected bool
	}
	testcases := map[string]testcase{
		"ex01": {code: "aaaaa-bbb-z-y-x-123[abxyz]", expected: true},
		"ex02": {code: "a-b-c-d-e-f-g-h-987[abcde]", expected: true},
		"ex03": {code: "not-a-real-room-404[oarel]", expected: true},
		"ex04": {code: "totally-real-room-200[decoy]", expected: false},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			r := NewRoom(tc.code)
			assert.Equal(t, tc.expected, r.IsValid())
		})
	}
}

func TestDecrypt(t *testing.T) {
	type testcase struct {
		code     string
		expected string
	}
	testcases := map[string]testcase{
		"ex01": {code: "qzmt-zixmtkozy-ivhz-343[abcde]", expected: "very encrypted name"},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			r := NewRoom(tc.code)
			assert.Equal(t, tc.expected, r.Decrypt())
		})
	}
}

func Test04a(t *testing.T) {
	in, err := os.Open("input04")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scanner := bufio.NewScanner(in)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		r := NewRoom(line)
		if r.IsValid() {
			sum += r.number
		}
	}

	assert.Equal(t, 173787, sum)
}

func Test04b(t *testing.T) {
	in, err := os.Open("input04")
	if err != nil {
		panic(err)
	}
	defer func() { _ = in.Close() }()

	scanner := bufio.NewScanner(in)
	var sectorID int
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		r := NewRoom(line)
		if strings.Contains(r.Decrypt(), "northpole") {
			sectorID = r.number
		}
	}

	assert.Equal(t, 548, sectorID)
}
