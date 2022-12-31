package two022_test

import (
	"bytes"
	"os"
	"regexp"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type valve struct {
	name     string
	rate     int
	linkedTo []*valve
}

var _ = Describe("16", func() {
	var (
		valves               map[string]*valve
		interestingDistances map[string]map[string]int
		visited              map[string]int
		maxTime              int
	)

	loadValves := func(f []byte) {
		re := regexp.MustCompile(`Valve (.*) has flow rate=(.*); tunnels? leads? to valves? (.*)$`)
		for _, line := range bytes.Split(f, []byte("\n")) {
			if len(line) == 0 {
				continue
			}
			matches := re.FindSubmatch(line)
			valveName := string(matches[1])
			rate := stoi(string(matches[2]))
			tunnels := strings.Split(string(matches[3]), ", ")

			v := &valve{
				name: valveName,
				rate: rate,
			}

			for _, link := range tunnels {
				if _, ok := valves[link]; !ok {
					pv := &valve{name: link}
					valves[link] = pv
				}
				v.linkedTo = append(v.linkedTo, valves[link])
			}

			valves[valveName] = v
		}
	}

	distanceTo := func(from, to string) int {
		visited := map[string]string{}
		distances := map[string]int{"AA": 0}
		queue := []string{from}
		for len(queue) != 0 {
			cur := queue[0]
			queue = queue[1:]
			if visited[cur] == "done" {
				continue
			}
			if cur == to {
				return distances[cur]
			}

			visited[cur] = "doing"

			v := valves[cur]
			for _, link := range v.linkedTo {
				if visited[link.name] != "" {
					continue
				}
				queue = append(queue, link.name)
				distances[link.name] = distances[cur] + 1
			}

			visited[cur] = "done"
		}
		return -1
	}

	loadInterestingDistances := func() []string {
		interestingValveMap := map[string]bool{"AA": true}
		interestingValves := []string{"AA"}
		for name, v := range valves {
			if v.rate > 0 {
				interestingValveMap[name] = true
				interestingValves = append(interestingValves, name)
			}
		}

		for i, vname := range interestingValves {
			if interestingDistances[vname] == nil {
				interestingDistances[vname] = map[string]int{}
			}
			for j := i + 1; j < len(interestingValves); j++ {
				dist := distanceTo(vname, interestingValves[j])
				if dist > -1 {
					interestingDistances[vname][interestingValves[j]] = dist
					if interestingDistances[interestingValves[j]] == nil {
						interestingDistances[interestingValves[j]] = map[string]int{}
					}
					interestingDistances[interestingValves[j]][vname] = dist
				}
			}
		}

		return interestingValves
	}

	var backtrackPartA func(string, int, int) int
	backtrackPartA = func(from string, time, score int) int {
		if time > maxTime {
			return score
		}
		newscore := score + (maxTime-time)*valves[from].rate
		bestrate := newscore
		visited[from] = time
		for to, dist := range interestingDistances[from] {
			if _, ok := visited[to]; ok {
				continue
			}
			r := backtrackPartA(to, time+dist+1, newscore)
			if r > bestrate {
				bestrate = r
			}
		}
		delete(visited, from)
		return bestrate
	}

	var backtrackPartB func(string, int, int, map[string]bool) int
	backtrackPartB = func(from string, time, score int, allowed map[string]bool) int {
		if time > maxTime {
			return score
		}
		newscore := score + (maxTime-time)*valves[from].rate
		bestrate := newscore
		visited[from] = time
		for to, dist := range interestingDistances[from] {
			if !allowed[to] {
				continue
			}
			if _, ok := visited[to]; ok {
				continue
			}
			r := backtrackPartB(to, time+dist+1, newscore, allowed)
			if r > bestrate {
				bestrate = r
			}
		}
		delete(visited, from)
		return bestrate
	}

	var subsetsWithLenRec func(u []string, n, k, i int, currentSubset *[]string, res *[][]string)
	subsetsWithLenRec = func(u []string, n, k, i int, currentSubset *[]string, res *[][]string) {
		if n < k {
			return
		}

		if k == 0 {
			soln := make([]string, len(*currentSubset))
			copy(soln, *currentSubset)
			*res = append(*res, soln)
			return
		}

		subsetsWithLenRec(u, n-1, k, i+1, currentSubset, res)
		*currentSubset = append(*currentSubset, u[i])
		subsetsWithLenRec(u, n-1, k-1, i+1, currentSubset, res)
		*currentSubset = (*currentSubset)[:len(*currentSubset)-1]
	}

	subsetsWithLen := func(u []string, l int) [][]string {
		var res [][]string
		var currentSubset []string
		subsetsWithLenRec(u, len(u), l, 0, &currentSubset, &res)
		return res
	}

	BeforeEach(func() {
		valves = map[string]*valve{}
		interestingDistances = map[string]map[string]int{}
		visited = map[string]int{}
		maxTime = -1
	})

	It("does part A", func() {
		f, err := os.ReadFile("input16")
		Expect(err).NotTo(HaveOccurred())

		loadValves(f)
		loadInterestingDistances()

		maxTime = 30
		ans := backtrackPartA("AA", 0, 0)
		Expect(ans).To(Equal(1940))
	})

	complements := func(u, s []string) (orig, compl map[string]bool) {
		m := map[string]bool{}
		for _, a := range s {
			m[a] = true
		}

		res := map[string]bool{}
		for _, a := range u {
			if !m[a] {
				res[a] = true
			}
		}

		return m, res
	}

	It("does part B", func() {
		f, err := os.ReadFile("input16")
		Expect(err).NotTo(HaveOccurred())

		loadValves(f)
		interestingValves := loadInterestingDistances()
		interestingValves = interestingValves[1:]
		maxTime = 26
		var maxScore int
		for i := 1; i <= len(interestingValves)/2; i++ {
			for _, s := range subsetsWithLen(interestingValves, i) {
				a, b := complements(interestingValves, s)
				scoreA := backtrackPartB("AA", 0, 0, a)
				scoreB := backtrackPartB("AA", 0, 0, b)
				if scoreA+scoreB > maxScore {
					maxScore = scoreA + scoreB
				}
			}
		}
		Expect(maxScore).To(Equal(2469))
	})
})
