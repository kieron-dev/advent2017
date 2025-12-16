package twentytwentyfive_test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const size = 10

func TestMinButtonPresses(t *testing.T) {
	type tc struct {
		required         string
		buttonSchematics [][]int
		expected         int
	}

	tcs := map[string]tc{
		"ex01": {
			required:         ".##.",
			buttonSchematics: [][]int{{3}, {1, 3}, {2}, {2, 3}, {0, 2}, {0, 1}},
			expected:         2,
		},
		"ex02": {
			required:         "...#.",
			buttonSchematics: [][]int{{0, 2, 3, 4}, {2, 3}, {0, 4}, {0, 1, 2}, {1, 2, 3, 4}},
			expected:         3,
		},
		"ex03": {
			required:         ".###.#",
			buttonSchematics: [][]int{{0, 1, 2, 3, 4}, {0, 3, 4}, {0, 1, 2, 4, 5}, {1, 2}},
			expected:         2,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, minPresses(tc.required, tc.buttonSchematics))
		})
	}
}

func TestMinJoltagePresses(t *testing.T) {
	type tc struct {
		required         []int
		buttonSchematics [][]int
		expected         int
	}

	tcs := map[string]tc{
		"ex01": {
			required:         []int{3, 5, 4, 7},
			buttonSchematics: [][]int{{3}, {1, 3}, {2}, {2, 3}, {0, 2}, {0, 1}},
			expected:         10,
		},
		"ex02": {
			required:         []int{7, 5, 12, 7, 2},
			buttonSchematics: [][]int{{0, 2, 3, 4}, {2, 3}, {0, 4}, {0, 1, 2}, {1, 2, 3, 4}},
			expected:         12,
		},
		"ex03": {
			required:         []int{10, 11, 11, 5, 10, 5},
			buttonSchematics: [][]int{{0, 1, 2, 3, 4}, {0, 3, 4}, {0, 1, 2, 4, 5}, {1, 2}},
			expected:         11,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, minJoltagePresses(tc.required, tc.buttonSchematics))
		})
	}
}

func Test10a(t *testing.T) {
	real, err := os.Open("input10")
	Check(err)
	type tc struct {
		in       io.Reader
		expected int
	}

	tcs := map[string]tc{
		"ex01": {
			in: strings.NewReader(`[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
`),
			expected: 7,
		},
		"real": {
			in:       real,
			expected: 0,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, totalMinPresses(tc.in))
		})
	}
}

func Test10b(t *testing.T) {
	real, err := os.Open("input10")
	Check(err)
	type tc struct {
		in       io.Reader
		expected int
	}

	tcs := map[string]tc{
		"ex01": {
			in: strings.NewReader(`[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
`),
			expected: 33,
		},
		"real": {
			in:       real,
			expected: 0,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, totalMinJoltagePresses(tc.in))
		})
	}
}

func totalMinPresses(in io.Reader) int {
	sum := 0
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		lastSqBr := strings.Index(line, "]")
		required := line[1:lastSqBr]

		firstRoundBr := strings.Index(line, "(")
		lastRoundBr := strings.LastIndex(line, ")")
		fields := strings.FieldsFunc(line[firstRoundBr:lastRoundBr], func(r rune) bool {
			if r == '(' || r == ')' || r == ' ' {
				return true
			}
			return false
		})
		schematics := [][]int{}
		for _, field := range fields {
			bits := strings.Split(field, ",")
			bitInts := []int{}
			for _, b := range bits {
				n, err := strconv.Atoi(b)
				Check(err)
				bitInts = append(bitInts, n)
			}
			schematics = append(schematics, bitInts)
		}
		sum += minPresses(required, schematics)
	}

	return sum
}

const numWorkers = 8

var (
	workerChan = make(chan (string), 158)
	resultChan = make(chan (int), 158)
)

func totalMinJoltagePresses(in io.Reader) int {
	go startWorkers()

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" || line[0] == '#' {
			continue
		}

		workerChan <- line
	}
	close(workerChan)

	sum := 0
	for res := range resultChan {
		sum += res
	}

	return sum
}

func startWorkers() {
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := range numWorkers {
		go func(n int) {
			for line := range workerChan {
				resultChan <- processLine(line)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(resultChan)
}

func processLine(line string) int {
	firstCurlyBr := strings.Index(line, "{")
	lastCurlyBr := strings.LastIndex(line, "}")
	requireds := []int{}
	bits := strings.Split(line[firstCurlyBr+1:lastCurlyBr], ",")
	for _, b := range bits {
		n, err := strconv.Atoi(b)
		Check(err)
		requireds = append(requireds, n)
	}

	firstRoundBr := strings.Index(line, "(")
	lastRoundBr := strings.LastIndex(line, ")")
	fields := strings.FieldsFunc(line[firstRoundBr:lastRoundBr], func(r rune) bool {
		if r == '(' || r == ')' || r == ' ' {
			return true
		}
		return false
	})
	schematics := [][]int{}
	for _, field := range fields {
		bits := strings.Split(field, ",")
		bitInts := []int{}
		for _, b := range bits {
			n, err := strconv.Atoi(b)
			Check(err)
			bitInts = append(bitInts, n)
		}
		schematics = append(schematics, bitInts)
	}
	return minJoltagePresses(requireds, schematics)
}

func minPresses(required string, schematics [][]int) int {
	start := strings.Repeat(".", len(required))
	queue := []string{start}
	visited := map[string]bool{}
	distances := map[string]int{start: 0}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if visited[cur] {
			continue
		}
		if cur == required {
			return distances[cur]
		}
		visited[cur] = true
		for _, sc := range schematics {
			nxt := applySchematic(cur, sc)
			nxtDist, ok := distances[nxt]
			if !ok {
				nxtDist = 100000
			}
			distances[nxt] = min(nxtDist, distances[cur]+1)
			queue = append(queue, nxt)
		}
	}

	return -1
}

var cache map[[size]int]int

func minJoltagePresses2(from [size]int, to [size]int, schematics [][]int) int {
	if val, ok := cache[from]; ok {
		return val
	}

	if from == to {
		return 0
	}

	min := 10000000
	soln := false
outer:
	for _, sc := range schematics {
		nxt := applyJoltageSchematic(from, sc)
		for i, n := range nxt {
			if n > to[i] {
				continue outer
			}
		}

		val := minJoltagePresses2(nxt, to, schematics)
		if val >= 0 {
			soln = true
			if val+1 < min {
				min = val + 1
			}
		}
	}

	if soln {
		cache[from] = min
		return min
	}

	cache[from] = -1
	return -1
}

type SumConstraint struct {
	buttons []int
	sum     int
}

func (c SumConstraint) Check(presses []int) bool {
	sum := 0
	for _, b := range c.buttons {
		sum += presses[b]
	}
	return sum <= c.sum
}

func minJoltagePresses(required []int, schematics [][]int) int {
	var constraints []SumConstraint
	minButtonsCount := 100
	var minConstraint SumConstraint
	for pos, joltage := range required {
		c := SumConstraint{}
		for i, sch := range schematics {
			for _, n := range sch {
				if n == pos {
					c.buttons = append(c.buttons, i)
				}
			}
		}
		c.sum = joltage
		if len(c.buttons) < minButtonsCount {
			minButtonsCount = len(c.buttons)
			minConstraint = c
		}
		sort.Ints(c.buttons)
		constraints = append(constraints, c)
	}
	presses := make([]int, len(schematics))

	newOrder := make([]int, len(minConstraint.buttons))
	copy(newOrder, minConstraint.buttons)
	btnMap := map[int]bool{}
	for _, b := range minConstraint.buttons {
		btnMap[b] = true
	}
	for i := range len(presses) {
		if btnMap[i] {
			continue
		}
		newOrder = append(newOrder, i)
	}

	solns := solveMinPresses(required, schematics, constraints, presses, minConstraint, newOrder, 0)
	sort.Ints(solns)
	fmt.Println("RESULT:", required, solns[0])
	return solns[0]
}

func solveMinPresses(required []int, schematics [][]int, constraints []SumConstraint, presses []int, minConstraint SumConstraint, newOrder []int, pos int) []int {
	if pos >= len(presses) {
		return nil
	}

	actualPos := newOrder[pos]

	if pos == len(minConstraint.buttons)-1 {
		presses[actualPos] = minConstraint.sum
		for i := 0; i < len(minConstraint.buttons)-1; i++ {
			presses[actualPos] -= presses[minConstraint.buttons[i]]
		}
		soln := solveMinPresses(required, schematics, constraints, presses, minConstraint, newOrder, pos+1)
		presses[actualPos] = 0
		return soln
	}

	i := 0
	var solns []int
outer:
	for {

		presses[actualPos] = i
		for _, c := range constraints {
			if !c.Check(presses) {
				for j := pos; j < len(presses); j++ {
					presses[newOrder[j]] = 0
				}
				break outer
			}
		}
		s := make([]int, len(required))
		for pos, count := range presses {
			for _, b := range schematics[pos] {
				s[b] += count
			}
		}
		if slices.Equal(s, required) {
			sumPresses := 0
			for _, p := range presses {
				sumPresses += p
			}
			solns = append(solns, sumPresses)
			fmt.Println(required, "soln", sumPresses)
		}

		res := solveMinPresses(required, schematics, constraints, presses, minConstraint, newOrder, pos+1)
		solns = append(solns, res...)
		i++
	}

	return solns
}

func minJoltagePressesOld(requiredSl []int, schematics [][]int) int {
	from := [size]int{}
	// max length of things is 10
	var start, required [size]int
	copy(required[:], requiredSl)
	cache = map[[size]int]int{}
	return minJoltagePresses2(from, required, schematics)

	queue := [][size]int{start}
	visited := map[[size]int]bool{}
	distances := map[[size]int]int{start: 0}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if visited[cur] {
			continue
		}
		if cur == required {
			return distances[cur]
		}
		visited[cur] = true
	outer:
		for _, sc := range schematics {
			nxt := applyJoltageSchematic(cur, sc)
			for i, n := range nxt {
				if n > required[i] {
					continue outer
				}
			}
			nxtDist, ok := distances[nxt]
			if !ok {
				nxtDist = 10000000
			}
			distances[nxt] = min(nxtDist, distances[cur]+1)
			queue = append(queue, nxt)
		}
	}

	return -1
}

func applyJoltageSchematic(cur [size]int, sc []int) [size]int {
	var nxt [size]int
	copy(nxt[:], cur[:])
	for _, n := range sc {
		nxt[n]++
	}
	return nxt
}

func applySchematic(cur string, sc []int) string {
	bs := []byte(cur)
	for _, n := range sc {
		if bs[n] == '.' {
			bs[n] = '#'
		} else {
			bs[n] = '.'
		}
	}
	return string(bs)
}
