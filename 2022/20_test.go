package two022_test

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var example20 = `1
2
-3
3
-2
0
4`

func (nl *numList) initPtrs(nums []int) {
	ptrs := make([]*num, len(nums))
	for i, n := range nums {
		ptrs[i] = &num{val: n}
		if n == 0 {
			nl.zero = ptrs[i]
		}
	}

	for i, p := range ptrs {
		if i == 0 {
			p.l = ptrs[len(ptrs)-1]
		} else {
			p.l = ptrs[i-1]
		}
		if i == len(ptrs)-1 {
			p.r = ptrs[0]
		} else {
			p.r = ptrs[i+1]
		}
	}

	nl.nums = ptrs
}

func newNumList(in io.Reader) *numList {
	bs, err := io.ReadAll(in)
	Expect(err).NotTo(HaveOccurred())

	numStrs := bytes.Fields(bs)
	nums := make([]int, len(numStrs))
	for i, s := range numStrs {
		n, err := strconv.Atoi(string(s))
		Expect(err).NotTo(HaveOccurred())
		nums[i] = n
	}

	nl := &numList{
		len: len(nums),
	}
	nl.initPtrs(nums)

	return nl
}

type numList struct {
	nums []*num
	zero *num
	len  int
}

type num struct {
	l   *num
	r   *num
	val int
}

func (nl *numList) move(p *num) {
	amount := p.val % (nl.len - 1)
	if amount == 0 {
		return
	}

	if p.val < 0 {
		amount += nl.len - 1
	}

	nl.moveRight(p, amount)
}

func (nl *numList) moveRight(p *num, n int) {
	p.l.r = p.r
	p.r.l = p.l

	newLeft := p
	for i := 0; i < n; i++ {
		newLeft = newLeft.r
	}

	p.l = newLeft
	p.r = newLeft.r
	newLeft.r = p
	p.r.l = p
}

func (nl *numList) rotate() {
	for i := 0; i < nl.len; i++ {
		nl.move(nl.nums[i])
	}
}

func (nl numList) coords() [3]int {
	z := nl.zero

	ret := [3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			z = z.r
		}
		ret[i] = z.val
	}

	return ret
}

func (nl numList) coordsSum() int {
	s := 0
	for _, n := range nl.coords() {
		s += n
	}

	return s
}

func (nl numList) scale(n int) {
	for _, p := range nl.nums {
		p.val *= n
	}
}

func (nl numList) vals() []int {
	p := nl.zero
	ret := []int{}
	for {
		ret = append(ret, p.val)
		p = p.r
		if p == nl.zero {
			break
		}
	}

	return ret
}

var _ = Describe("20", func() {
	It("do the example", func() {
		nl := newNumList(strings.NewReader(example20))
		nl.rotate()
		Expect(nl.coords()).To(Equal([3]int{4, -3, 2}))
	})

	It("does part A", func() {
		f, err := os.Open("input20")
		Expect(err).NotTo(HaveOccurred())

		nl := newNumList(f)
		f.Close()

		nl.rotate()
		Expect(nl.coordsSum()).To(Equal(4914))
	})

	It("does example part B style", func() {
		nl := newNumList(strings.NewReader(example20))
		nl.scale(811589153)
		for i := 0; i < 10; i++ {
			nl.rotate()
		}
		Expect(nl.coordsSum()).To(Equal(1623178306))
	})

	It("does part B", func() {
		f, err := os.Open("input20")
		Expect(err).NotTo(HaveOccurred())

		nl := newNumList(f)
		f.Close()

		nl.scale(811589153)

		for i := 0; i < 10; i++ {
			nl.rotate()
		}

		Expect(nl.coordsSum()).To(Equal(7973051839072))
	})

	It("does reddit example", func() {
		nl := newNumList(strings.NewReader("1 2 -3 3 -2 0 8"))
		nl.rotate()
		p := nl.zero
		for {
			p = p.r
			if p == nl.zero {
				break
			}
		}
		Expect(nl.coordsSum()).To(Equal(7))
	})

	DescribeTable("moving", func(in string, res []int) {
		nl := newNumList(strings.NewReader(in))
		nl.move(nl.nums[0])
		Expect(nl.vals()).To(Equal(res))
	},
		Entry("1", "1 2 3 0", []int{0, 2, 1, 3}),
		Entry("2", "3 2 1 0", []int{0, 3, 2, 1}),
	)
})
