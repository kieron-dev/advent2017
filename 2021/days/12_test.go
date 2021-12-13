package days_test

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type ANode struct {
	name  string
	conns []*ANode
}

func newANode(name string) *ANode {
	return &ANode{name: name}
}

var _ = Describe("12", func() {
	It("does part A", func() {
		nodes := getInput12()
		visited := map[*ANode]bool{}

		start := nodes["start"]
		c := start.PathsToEnd(visited)
		Expect(c).To(Equal(5756))
	})

	It("does part B", func() {
		nodes := getInput12()
		visited := map[*ANode]int{}

		start := nodes["start"]
		c := start.PathsToEnd2(visited)
		Expect(c).To(Equal(144603))
	})
})

func allLCLessMax1(visited map[*ANode]int) bool {
	for n, c := range visited {
		if !isLower(n.name) {
			continue
		}
		if c > 1 {
			return false
		}
	}

	return true
}

func isLower(s string) bool {
	return strings.ToLower(s) == s
}

func (a *ANode) PathsToEnd2(visited map[*ANode]int) int {
	if a.name == "end" {
		return 1
	}
	if a.name == "start" && visited[a] > 0 {
		return 0
	}
	if isLower(a.name) {
		if !allLCLessMax1(visited) {
			if visited[a] > 0 {
				return 0
			}
		}
		if visited[a] > 1 {
			return 0
		}
	}

	newMap := map[*ANode]int{}
	for k, v := range visited {
		newMap[k] = v
	}
	newMap[a]++

	count := 0

	for _, c := range a.conns {
		count += c.PathsToEnd2(newMap)
	}

	return count
}

func (a *ANode) PathsToEnd(visited map[*ANode]bool) int {
	if a.name == "end" {
		return 1
	}

	newMap := map[*ANode]bool{}
	for k, v := range visited {
		newMap[k] = v
	}
	newMap[a] = true

	count := 0

	for _, c := range a.conns {
		if isLower(c.name) && visited[c] {
			continue
		}
		count += c.PathsToEnd(newMap)
	}

	return count
}

func getInput12() map[string]*ANode {
	input, err := os.Open("input12")
	Expect(err).NotTo(HaveOccurred())
	defer input.Close()

	nodes := map[string]*ANode{}

	re := regexp.MustCompile(`^(.*)-(.*)$`)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		Expect(matches).To(HaveLen(3))

		from, ok := nodes[matches[1]]
		if !ok {
			from = newANode(matches[1])
			nodes[matches[1]] = from
		}
		to, ok := nodes[matches[2]]
		if !ok {
			to = newANode(matches[2])
			nodes[matches[2]] = to
		}
		from.conns = append(from.conns, to)
		to.conns = append(to.conns, from)
	}
	return nodes
}
