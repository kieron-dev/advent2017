package two023_test

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type quality string

const (
	X quality = "x"
	M quality = "m"
	A quality = "a"
	S quality = "s"
)

type tool struct {
	qualities map[quality]int
}

func newTool(line string) tool {
	line = line[1 : len(line)-1]
	parts := strings.Split(line, ",")
	Expect(parts).To(HaveLen(4))
	t := tool{}
	t.qualities = map[quality]int{}
	for i, q := range []quality{X, M, A, S} {
		part := parts[i]
		bits := strings.Split(part, "=")
		n, err := strconv.Atoi(bits[1])
		Expect(err).NotTo(HaveOccurred())
		t.qualities[q] = n
	}
	return t
}

func (t tool) sum() int {
	s := 0
	for _, v := range t.qualities {
		s += v
	}
	return s
}

type sign int

const (
	_ sign = iota
	GT
	LT
)

type rule struct {
	quality quality
	sign    sign
	val     int
	target  string
}

func newRule(s string) rule {
	r := rule{}

	if strings.Contains(s, ":") {
		r.quality = quality(s[0:1])
		if s[1] == '<' {
			r.sign = LT
		} else if s[1] == '>' {
			r.sign = GT
		}
		bits := strings.Split(s[2:], ":")
		val, err := strconv.Atoi(bits[0])
		Expect(err).NotTo(HaveOccurred())
		r.val = val
		r.target = bits[1]

		return r
	}

	return rule{target: s}
}

func (r rule) true(t tool) bool {
	switch r.sign {
	case GT:
		return t.qualities[r.quality] > r.val
	case LT:
		return t.qualities[r.quality] < r.val
	default:
		return true
	}
}

type process struct {
	name  string
	rules []rule
}

func newProcess(line string) process {
	p := process{}
	parts1 := strings.Split(line, "{")
	p.name = parts1[0]

	rest := parts1[1]
	rest = rest[:len(rest)-1]

	parts2 := strings.Split(rest, ",")
	for _, r := range parts2 {
		p.rules = append(p.rules, newRule(r))
	}
	return p
}

type qa struct {
	processes map[string]process
	tools     []tool
}

func newQA(filename string) qa {
	q := qa{}
	q.processes = map[string]process{}

	f, err := os.Open(filename)
	Expect(err).NotTo(HaveOccurred())
	defer f.Close()

	inTools := false
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			inTools = true
			continue
		}

		if !inTools {
			p := newProcess(line)
			q.processes[p.name] = p
			continue
		}

		q.tools = append(q.tools, newTool(line))

	}

	return q
}

func (q qa) isApproved(t tool) bool {
	p := "in"
	for p != "R" && p != "A" {
		p = q.process(p, t)
	}

	return p == "A"
}

func (q qa) process(p string, t tool) string {
	proc := q.processes[p]
	for _, r := range proc.rules {
		if r.true(t) {
			return r.target
		}
	}

	Fail("oops")
	return ""
}

func (q qa) sumApprovedTools() int {
	s := 0
	for _, t := range q.tools {
		if q.isApproved(t) {
			s += t.sum()
		}
	}
	return s
}

type toolRange struct {
	mins  [4]int
	maxes [4]int
}

func (q qa) waysToApproved() int {
	t := toolRange{
		mins:  [4]int{1, 1, 1, 1},
		maxes: [4]int{4000, 4000, 4000, 4000},
	}

	return q.dfs("in", 0, t)
}

var qindexes = map[quality]int{
	X: 0,
	M: 1,
	A: 2,
	S: 3,
}

func (q qa) dfs(cur string, idx int, t toolRange) int {
	if cur == "R" {
		return 0
	}
	if cur == "A" {
		s := 1
		for i := range t.mins {
			s *= t.maxes[i] - t.mins[i] + 1
		}
		return s
	}

	t1 := t
	t2 := t
	r := q.processes[cur].rules[idx]
	qidx := qindexes[r.quality]
	if r.sign == GT {
		t1.mins[qidx] = max(t1.mins[qidx], r.val+1)
		t2.maxes[qidx] = min(t2.maxes[qidx], r.val)
		return q.dfs(r.target, 0, t1) +
			q.dfs(cur, idx+1, t2)
	}
	if r.sign == LT {
		t1.maxes[qidx] = min(t1.maxes[qidx], r.val-1)
		t2.mins[qidx] = max(t2.mins[qidx], r.val)
		return q.dfs(r.target, 0, t1) +
			q.dfs(cur, idx+1, t2)
	}
	return q.dfs(r.target, 0, t)
}

var _ = Describe("19", func() {
	It("can do part A", func() {
		q := newQA("input19")

		s := q.sumApprovedTools()
		Expect(s).To(Equal(495298))
	})

	It("can do part B", func() {
		q := newQA("input19")

		s := q.waysToApproved()
		Expect(s).To(Equal(132186256794011))
	})
})
