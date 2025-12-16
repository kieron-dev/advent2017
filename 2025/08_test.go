package twentytwentyfive_test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ex08 = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689
`

func Test08a(t *testing.T) {
	real, err := os.Open("input08")
	Check(err)
	type tc struct {
		in         io.Reader
		iterations int
		part       int
		expected   int
	}

	tcs := map[string]tc{
		"ex01":  {in: strings.NewReader(ex08), iterations: 10, part: 1, expected: 40},
		"part1": {in: real, iterations: 1000, part: 1, expected: 62186},
		"ex02":  {in: strings.NewReader(ex08), iterations: 10, part: 2, expected: 25272},
		"part2": {in: real, iterations: 1000, part: 2, expected: 8420405530},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, connectN(tc.in, tc.iterations, tc.part))
		})
	}
}

type Coord3d [3]int

// don't bother taking the square root, as it won't affect anything
func (c Coord3d) Dist2(d Coord3d) int {
	sum := 0
	for i := range 3 {
		sum += (c[i] - d[i]) * (c[i] - d[i])
	}

	return sum
}

func (c Coord3d) String() string {
	return fmt.Sprintf("(%d, %d, %d)", c[0], c[1], c[2])
}

type Distance struct {
	to       Coord3d
	from     Coord3d
	distance int
}

func connectN(in io.Reader, iterations int, part int) int {
	distances := []Distance{}
	coords := []Coord3d{}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		bits := strings.Split(line, ",")
		var bitNs [3]int
		for i := range 3 {
			n, err := strconv.Atoi(bits[i])
			Check(err)
			bitNs[i] = n
		}
		coord := Coord3d(bitNs)
		for _, other := range coords {
			dist := coord.Dist2(other)
			distances = append(distances, Distance{to: other, from: coord, distance: dist})
		}
		coords = append(coords, coord)
	}

	sort.Slice(distances, func(a, b int) bool {
		return distances[a].distance < distances[b].distance
	})

	groupMap := map[Coord3d]int{}
	for i, c := range coords {
		groupMap[c] = i + 1
	}
	gCount := len(groupMap)

	i := 0
	for _, d := range distances {
		if part == 1 && i >= iterations {
			break
		}
		i++

		if groupMap[d.from] == groupMap[d.to] {
			continue
		}

		g1 := groupMap[d.to]
		g2 := groupMap[d.from]

		for x, grp := range groupMap {
			if grp == g2 {
				groupMap[x] = g1
			}
		}
		gCount--

		if gCount == 1 && part == 2 {
			return d.to[0] * d.from[0]
		}
	}

	// fmt.Println(groupMap)
	freqs := map[int]int{}
	for _, g := range groupMap {
		freqs[g]++
	}

	// fmt.Println(freqs)

	nums := []int{}
	for _, f := range freqs {
		nums = append(nums, f)
	}

	sort.Ints(nums)
	l := len(nums)

	return nums[l-1] * nums[l-2] * nums[l-3]
}
