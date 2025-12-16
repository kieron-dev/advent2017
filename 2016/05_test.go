package twentysixteen_test

import (
	"crypto/md5"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test05a(t *testing.T) {
	type testcase struct {
		doorID   string
		expected string
	}

	testcases := map[string]testcase{
		"ex01": {
			doorID:   "abc",
			expected: "18f47a30",
		},
		"real": {
			doorID:   "ugkcyxxp",
			expected: "d4cd2ee1",
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, DoorPasswordA(tc.doorID))
		})
	}
}

func Test05b(t *testing.T) {
	type testcase struct {
		doorID   string
		expected string
	}

	testcases := map[string]testcase{
		"ex01": {
			doorID:   "abc",
			expected: "05ace8e3",
		},
		"real": {
			doorID:   "ugkcyxxp",
			expected: "f2c730e5",
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, DoorPasswordB(tc.doorID))
		})
	}
}

func DoorPasswordA(id string) string {
	i := 0
	pw := make([]byte, 0, 8)
	for {
		in := fmt.Sprintf("%s%d", id, i)
		hash := md5.Sum([]byte(in))
		out := fmt.Sprintf("%x", hash)
		if strings.HasPrefix(out, "00000") {
			pw = append(pw, out[5])
		}
		if len(pw) == 8 {
			break
		}
		i++
	}

	return string(pw)
}

func DoorPasswordB(id string) string {
	i := 0
	set := 0
	pw := make([]byte, 8)
	for {
		in := fmt.Sprintf("%s%d", id, i)
		hash := md5.Sum([]byte(in))
		out := fmt.Sprintf("%x", hash)
		if strings.HasPrefix(out, "00000") {
			pos := int(out[5] - '0')
			if pos < 8 && pw[pos] == 0 {
				pw[pos] = out[6]
				set++
			}
		}
		if set == 8 {
			break
		}
		i++
	}

	return string(pw)
}
