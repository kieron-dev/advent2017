package network

import (
	"github.com/kieron-dev/advent2017/advent2019/intcode"
)

type NIC struct {
	computer *intcode.Computer
	in       chan int
	out      chan int
	stop     bool
}

func NewNIC() *NIC {
	nic := new(NIC)
	nic.in = make(chan int, 1)
	nic.out = make(chan int, 100)
	nic.computer = intcode.NewComputer(nic.in, nic.out)

	return nic
}

func (n *NIC) Start(prog string) {
	n.computer.SetInput(prog)

	go n.computer.Calculate()
}

func (n *NIC) Input(i int) {
	n.in <- i
}

func (n *NIC) Output() chan int {
	return n.out
}

type TokenRing struct {
	nics         []*NIC
	size         int
	queues       map[int][]int
	specialQueue []int
}

func NewNetwork(size int, prog string) *TokenRing {
	net := new(TokenRing)

	for i := 0; i < size; i++ {
		nic := NewNIC()
		nic.Start(prog)
		nic.Input(i)
		net.nics = append(net.nics, nic)
	}

	net.size = size
	net.queues = map[int][]int{}
	net.specialQueue = []int{}

	return net
}

func (t *TokenRing) Start() {
	changed := 2

	for changed > 0 {
		changed -= 1

		for i := 0; i < t.size; i++ {
			if len(t.queues[i]) == 0 {
				t.nics[i].Input(-1)
			} else {
				for _, v := range t.queues[i] {
					t.nics[i].Input(v)
				}
				t.queues[i] = []int{}
				changed = 2
			}

			if t.tryReadNIC(i) {
				changed = 2
			}
		}
	}
}

func (t *TokenRing) tryReadNIC(i int) bool {
	done := false
	changed := false

	for !done {
		select {
		case c := <-t.nics[i].out:
			x := <-t.nics[i].out
			y := <-t.nics[i].out
			t.Enqueue(i, c, x, y)
			changed = true
		default:
			done = true
		}
	}

	return changed
}

func (t *TokenRing) SpecialQueueHead() []int {
	if len(t.specialQueue) > 1 {
		return t.specialQueue[:2]
	}
	return nil
}

func (t *TokenRing) SpecialQueueTail() []int {
	if len(t.specialQueue) > 1 {
		return t.specialQueue[len(t.specialQueue)-2:]
	}
	return nil
}

func (t *TokenRing) Enqueue(i, c, x, y int) {
	if c == 255 {
		t.specialQueue = append(t.specialQueue, x, y)
		return
	}

	if c >= 0 && c < 50 {
		t.queues[c] = append(t.queues[c], x, y)
	}
}
