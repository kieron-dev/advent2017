package twentytwentyfive_test

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sumInvalidIDsA(in []byte) int {
	inStr := string(in)
	inStr = strings.TrimSpace(inStr)
	ranges := strings.Split(inStr, ",")
	sum := 0
	for _, r := range ranges {
		vals := strings.Split(r, "-")
		from, err := strconv.Atoi(vals[0])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}

	outer:
		for i := from; i <= to; i++ {
			s := strconv.Itoa(i)
			l := len(s)
			if l%2 != 0 {
				continue
			}
			for j := 0; j < l/2; j++ {
				if s[j] != s[j+l/2] {
					continue outer
				}
			}
			sum += i
		}
	}
	return sum
}

func sumInvalidIDsB(in []byte) int {
	inStr := string(in)
	inStr = strings.TrimSpace(inStr)
	ranges := strings.Split(inStr, ",")
	sum := 0
	for _, r := range ranges {
		vals := strings.Split(r, "-")
		from, err := strconv.Atoi(vals[0])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}

	outer:
		for id := from; id <= to; id++ {
			idStr := strconv.Itoa(id)
			length := len(idStr)
			// e.g. 121212
			// try lengths 1, 2, 3
		blockLoop:
			for blockLen := 1; blockLen <= length/2; blockLen++ {
				// if n doesn't divide l pick next n
				if length%blockLen != 0 {
					continue
				}
				numBlocks := length / blockLen
				// iterate through the first block
				for i := 0; i < blockLen; i++ {
					// iterate through the blocks
					for block := 1; block < numBlocks; block++ {
						if idStr[i] != idStr[i+block*blockLen] {
							// try another block length
							continue blockLoop
						}
					}
				}
				sum += id
				// we're done, don't count any more block lengths for this number
				continue outer
			}
		}
	}
	return sum
}

func Test02a(t *testing.T) {
	in, err := os.ReadFile("input02")
	if err != nil {
		panic(err)
	}

	type testcase struct {
		in       []byte
		expected int
	}

	testcases := map[string]testcase{
		"ex01": {
			in:       []byte("11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"),
			expected: 1227775554,
		},
		"real": {
			in:       in,
			expected: 18952700150,
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, sumInvalidIDsA(tc.in))
		})
	}
}

func Test02b(t *testing.T) {
	in, err := os.ReadFile("input02")
	if err != nil {
		panic(err)
	}

	type testcase struct {
		in       []byte
		expected int
	}

	testcases := map[string]testcase{
		"ex01": {
			in:       []byte("11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"),
			expected: 4174379265,
		},
		"real": {
			in:       in,
			expected: 28858486244,
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, sumInvalidIDsB(tc.in))
		})
	}
}
