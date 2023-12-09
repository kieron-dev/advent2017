package two023_test

import (
	"bufio"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type directions struct {
	turns string
	nodes map[string][2]string
}

func newDirections(fileName string) directions {
	f, err := os.Open(fileName)
	Expect(err).NotTo(HaveOccurred())

	d := directions{}
	d.nodes = map[string][2]string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if len(d.turns) == 0 {
			d.turns = line
			continue
		}

		parts := strings.Split(line, "=")
		key := strings.TrimSpace(parts[0])

		right := strings.ReplaceAll(parts[1], "(", "")
		right = strings.ReplaceAll(right, ")", "")
		parts2 := strings.Split(right, ",")
		d.nodes[key] = [2]string{strings.TrimSpace(parts2[0]), strings.TrimSpace(parts2[1])}
	}

	return d
}

func (d directions) steps(from, to string) int {
	i := 0
	cur := from
	for cur != to {
		next := 0
		if d.turns[i%len(d.turns)] == 'R' {
			next = 1
		}
		cur = d.nodes[cur][next]
		i++
	}
	return i
}

func (d directions) allStepsTillZ() map[string][]int {
	res := map[string][]int{}

	for k := range d.nodes {
		if strings.HasSuffix(k, "A") {
			res[k] = d.stepsTillZ(k)
		}
	}

	return res
}

func (d directions) stepsTillZ(from string) []int {
	res := []int{}
	visited := map[string]int{}
	i := 0
	cur := from
	for visited[cur] < 5 {
		next := 0
		if d.turns[i%len(d.turns)] == 'R' {
			next = 1
		}
		i++
		cur = d.nodes[cur][next]
		if strings.HasSuffix(cur, "Z") {
			res = append(res, i)
			visited[cur]++
		}
	}

	return res
}

var _ = Describe("08", func() {
	It("does part A", func() {
		dirs := newDirections("input08")
		Expect(dirs.steps("AAA", "ZZZ")).To(Equal(19637))
	})

	It("does part B", func() {
		dirs := newDirections("input08")
		res := dirs.allStepsTillZ()
		nums := []int{}
		for _, v := range res {
			nums = append(nums, v[0])
		}

		val := lcm(nums)
		Expect(val).To(Equal(8811050362409))
	})
})

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(nums []int) int {
	if len(nums) == 2 {
		return nums[0] * nums[1] / gcd(nums[0], nums[1])
	}

	l := len(nums)
	a := nums[l-2]
	b := nums[l-1]
	nums[l-2] = a * b / gcd(a, b)

	return lcm(nums[:l-1])
}
