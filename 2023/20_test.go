package two023_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type moduleKind string

const (
	flipFlop    moduleKind = "%"
	conjunction moduleKind = "&"
	broadcaster moduleKind = "broadcaster"
	output      moduleKind = "output"
)

type pulse int

const (
	_ pulse = iota
	lowPulse
	highPulse
)

func (p pulse) String() string {
	switch p {
	case lowPulse:
		return "low"
	case highPulse:
		return "high"
	}
	return ""
}

type module interface {
	receive(from string, p pulse) (pulse, bool)
	setInput(from string)
	name() string
}

type flipFlopModule struct {
	on     bool
	myname string
}

func (m flipFlopModule) name() string {
	return m.myname
}

func (m *flipFlopModule) receive(from string, p pulse) (pulse, bool) {
	if p == highPulse {
		return 0, false
	}

	m.on = !m.on
	if m.on {
		return highPulse, true
	}

	return lowPulse, true
}

func (m flipFlopModule) setInput(string) {}

type conjunctionModule struct {
	inputs map[string]pulse
	myname string
}

func (m conjunctionModule) name() string {
	return m.myname
}

func (m *conjunctionModule) setInput(input string) {
	if m.inputs == nil {
		m.inputs = map[string]pulse{}
	}
	m.inputs[input] = lowPulse
}

func (m *conjunctionModule) receive(from string, p pulse) (pulse, bool) {
	m.inputs[from] = p
	for _, p := range m.inputs {
		if p == lowPulse {
			return highPulse, true
		}
	}
	return lowPulse, true
}

type broadcastModule struct{}

func (m broadcastModule) name() string {
	return "broadcaster"
}

func (m broadcastModule) receive(_ string, p pulse) (pulse, bool) {
	return p, true
}

func (m broadcastModule) setInput(string) {}

type outputModule struct{}

func (m outputModule) name() string {
	return "output"
}

func (m outputModule) receive(_ string, p pulse) (pulse, bool) {
	return 0, false
}

func (m outputModule) setInput(string) {}

type network struct {
	parts   map[string]module
	routing map[string][]string
	rxHit   bool
}

func newNetwork() network {
	n := network{}
	n.parts = map[string]module{"output": &outputModule{}, "rx": &outputModule{}}
	n.routing = map[string][]string{"output": {}}

	return n
}

func newModule(line string) module {
	if line == string(broadcaster) {
		return &broadcastModule{}
	}
	if line == string(output) {
		return &outputModule{}
	}

	if line[0:1] == string(flipFlop) {
		return &flipFlopModule{myname: line[1:]}
	}

	if line[0:1] == string(conjunction) {
		return &conjunctionModule{myname: line[1:]}
	}

	return nil
}

func (n network) completeConjunctions() {
	cs := map[string]bool{}
	for k, m := range n.parts {
		switch m.(type) {
		case *conjunctionModule:
			cs[k] = true
		}
	}

	for k, v := range n.routing {
		for _, t := range v {
			if cs[t] {
				n.parts[t].setInput(k)
			}
		}
	}
}

func loadNetwork(filename string) network {
	f, err := os.Open(filename)
	Expect(err).NotTo(HaveOccurred())
	defer f.Close()

	n := newNetwork()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, " -> ")
		Expect(parts).To(HaveLen(2))
		m := newModule(parts[0])
		n.parts[m.name()] = m
		targets := strings.Split(parts[1], ", ")
		n.routing[m.name()] = targets
	}

	n.completeConjunctions()

	return n
}

type signal struct {
	p        pulse
	from, to module
}

func (n *network) broadcast() (int, int) {
	m := n.parts["broadcaster"]
	q := []signal{{p: lowPulse, from: nil, to: m}}
	lows := 0
	highs := 0

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		from := ""
		if cur.from != nil {
			from = cur.from.name()
		}
		var newPulse pulse
		var ok bool
		if cur.to != nil {
			newPulse, ok = cur.to.receive(from, cur.p)
		}
		if cur.p == highPulse {
			highs++
		} else {
			if cur.to.name() == "rx" {
				n.rxHit = true
			}
			lows++
		}
		fmt.Printf("%s -%s-> %s\n", from, cur.p, cur.to.name())
		if !ok {
			continue
		}
		for _, t := range n.routing[cur.to.name()] {
			q = append(q, signal{p: newPulse, from: cur.to, to: n.parts[t]})
		}
	}

	return lows, highs
}

var _ = Describe("20", func() {
	It("does part A", func() {
		n := loadNetwork("input20")
		var lows, highs int
		for i := 0; i < 1000; i++ {
			l, h := n.broadcast()
			lows += l
			highs += h
		}
		Expect(lows * highs).To(Equal(899848294))
	})

	FIt("does part B", func() {
		n := loadNetwork("input20")
		i := 0
		for {
			n.broadcast()
			if n.rxHit {
				break
			}
			i++
			// if i%100000 == 0 {
			// 	fmt.Printf("%d\n", i)
			// }
			fmt.Printf("n.parts[rx] = %+v\n", n.parts["dn"])
			fmt.Println("")
		}

		Expect(i).To(Equal(0))
	})
})
