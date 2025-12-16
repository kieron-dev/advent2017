package twentytwentyfive_test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test11a(t *testing.T) {
	real, err := os.Open("input11")
	Check(err)

	type tc struct {
		in       io.Reader
		expected int
	}

	tcs := map[string]tc{
		"ex01": {
			in: strings.NewReader(`
				aaa: you hhh
				you: bbb ccc
				bbb: ddd eee
				ccc: ddd eee fff
				ddd: ggg
				eee: out
				fff: out
				ggg: out
				hhh: ccc fff iii
				iii: out
				`),
			expected: 5,
		},
		"real": {
			in:       real,
			expected: 652,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, countPaths(tc.in))
		})
	}
}

func Test11b(t *testing.T) {
	real, err := os.Open("input11")
	Check(err)

	type tc struct {
		in       io.Reader
		expected int
	}

	tcs := map[string]tc{
		"ex01": {
			in: strings.NewReader(`
				svr: aaa bbb
				aaa: fft
				fft: ccc
				bbb: tty
				tty: ccc
				ccc: ddd eee
				ddd: hub
				hub: fff
				eee: dac
				dac: fff
				fff: ggg hhh
				ggg: out
				hhh: out
				`),
			expected: 2,
		},
		"real": {
			in:       real,
			expected: 362956369749210,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, countPathsWithDacAndFft(tc.in))
		})
	}
}

func countPaths(in io.Reader) int {
	connections := parseConnections(in)
	pathsToOut := countPathsDFS(connections, "you")

	return pathsToOut
}

func countPathsWithDacAndFft(in io.Reader) int {
	connections := parseConnections(in)
	pathCache := map[string]int{}
	pathsToOut := countPathsDFSWithDacAndFft(connections, "svr", false, false, pathCache)

	return pathsToOut
}

func countPathsDFSWithDacAndFft(cons map[string][]string, pos string, dac, fft bool, pathCache map[string]int) int {
	if pos == "out" {
		if dac && fft {
			return 1
		}
		return 0
	}

	key := fmt.Sprintf("%s:%v:%v", pos, dac, fft)
	if val, ok := pathCache[key]; ok {
		return val
	}

	if pos == "fft" {
		fft = true
	}
	if pos == "dac" {
		dac = true
	}

	paths := 0
	for _, next := range cons[pos] {
		paths += countPathsDFSWithDacAndFft(cons, next, dac, fft, pathCache)
	}

	pathCache[key] = paths

	return paths
}

func countPathsDFS(cons map[string][]string, pos string) int {
	if pos == "out" {
		return 1
	}

	paths := 0
	for _, next := range cons[pos] {
		paths += countPathsDFS(cons, next)
	}

	return paths
}

func parseConnections(in io.Reader) map[string][]string {
	connections := map[string][]string{}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		cPos := strings.Index(line, ":")
		from := line[:cPos]
		others := strings.Fields(line[cPos+2:])
		connections[from] = others
	}
	Check(scanner.Err())

	return connections
}
