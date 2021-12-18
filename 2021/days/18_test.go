package days_test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("18", func() {
	It("can parse a term", func() {
		st := parseTerm(nil, "[[[[[9,8],1],2],3],4]")
		Expect(st.String()).To(Equal("[[[[[9,8],1],2],3],4]"))
	})

	It("can find a literal pair at depth >= 4", func() {
		st := parseTerm(nil, "[[[[[9,8],1],2],3],4]")
		Expect(st.firstDeepLiteral(4).String()).To(Equal("[9,8]"))
	})

	It("can find a literal pair at depth >= 4 on the right", func() {
		st := parseTerm(nil, "[7,[6,[5,[4,[3,2]]]]]")
		Expect(st.firstDeepLiteral(4).String()).To(Equal("[3,2]"))
	})

	It("can find the value before [3,2]", func() {
		st := parseTerm(nil, "[7,[6,[5,[4,[3,2]]]]]")
		t32 := st.firstDeepLiteral(4)
		Expect(t32.String()).To(Equal("[3,2]"))

		valBefore := t32.valueBefore()
		Expect(valBefore).ToNot(BeNil())
		Expect(valBefore.String()).To(Equal("4"))
	})

	It("can find the value after [9,8]", func() {
		st := parseTerm(nil, "[[[[[9,8],1],2],3],4]")
		t98 := st.firstDeepLiteral(4)
		Expect(t98.String()).To(Equal("[9,8]"))

		valAfter := t98.valueAfter()
		Expect(valAfter).ToNot(BeNil())
		Expect(valAfter.String()).To(Equal("1"))
	})

	It("can reduce", func() {
		st := parseTerm(nil, "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]")
		st.reduce()
		Expect(st.String()).To(Equal("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"))
	})

	DescribeTable("can explode a term", func(in, out string) {
		inTerm := parseTerm(nil, in)
		outTerm := parseTerm(nil, out)
		inTerm.explode()
		Expect(inTerm.String()).To(Equal(outTerm.String()))
	},
		Entry("1", "[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"),
		Entry("2", "[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"),
		Entry("3", "[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"),
		Entry("4", "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"),
	)

	It("can add", func() {
		t1 := parseTerm(nil, "[[[[4,3],4],4],[7,[[8,4],9]]]")
		t2 := parseTerm(nil, "[1,1]")
		t := add(t1, t2)
		Expect(t.String()).To(Equal("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"))
	})

	DescribeTable("magnitude", func(in string, mag int) {
		t := parseTerm(nil, in)
		Expect(t.magnitude()).To(Equal(mag))
	},
		Entry("1", "[1,9]", 21),
		Entry("2", "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488),
	)

	It("can do part A", func() {
		input, err := os.Open("input17")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		var last *snailTerm
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			term := parseTerm(nil, line)
			if last == nil {
				last = term
				continue
			}
			last = add(last, term)
		}

		Expect(last.magnitude()).To(Equal(2501))
	})

	It("can do part B", func() {
		input, err := os.Open("input17")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		terms := []string{}
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			terms = append(terms, line)
		}

		max := 0
		for i := 0; i < len(terms); i++ {
			for j := i + 1; j < len(terms); j++ {
				a1 := parseTerm(nil, terms[i])
				a2 := parseTerm(nil, terms[i])
				b1 := parseTerm(nil, terms[j])
				b2 := parseTerm(nil, terms[j])
				m1 := add(a1, b1).magnitude()
				m2 := add(a2, b2).magnitude()
				if m1 > max {
					max = m1
				}
				if m2 > max {
					max = m2
				}
			}
		}

		Expect(max).To(Equal(4935))
	})
})

type snailTerm struct {
	parent *snailTerm
	// either both leftTerm and rightTerm not nil, or val contains the number
	leftTerm  *snailTerm
	rightTerm *snailTerm
	val       int
}

func parseTerm(parent *snailTerm, t string) *snailTerm {
	if t[0] != '[' {
		return &snailTerm{parent: parent, val: AToI(t)}
	}

	level := 0
	mid := -1
	for i := 0; i < len(t); i++ {
		if t[i] == '[' {
			level++
		} else if t[i] == ']' {
			level--
		} else if t[i] == ',' && level == 1 {
			mid = i
			break
		}
	}

	p := &snailTerm{parent: parent}
	p.leftTerm = parseTerm(p, t[1:mid])
	p.rightTerm = parseTerm(p, t[mid+1:len(t)-1])

	return p
}

func (t snailTerm) String() string {
	if t.isValueTerm() {
		return strconv.Itoa(t.val)
	}

	return fmt.Sprintf("[%s,%s]", t.leftTerm.String(), t.rightTerm.String())
}

func (t snailTerm) isValueTerm() bool {
	return t.leftTerm == nil
}

func (t snailTerm) isLiteralPair() bool {
	return !t.isValueTerm() && t.leftTerm.isValueTerm() && t.rightTerm.isValueTerm()
}

func (t *snailTerm) firstDeepLiteral(depth int) *snailTerm {
	if depth <= 0 && t.isLiteralPair() {
		return t
	}

	if t.isValueTerm() {
		return nil
	}

	l := t.leftTerm.firstDeepLiteral(depth - 1)
	if l != nil {
		return l
	}

	return t.rightTerm.firstDeepLiteral(depth - 1)
}

func (t *snailTerm) valueBefore() *snailTerm {
	if t.isValueTerm() {
		return t
	}

	p := t.parent
	if p == nil {
		return nil
	}

	if p.leftTerm == t {
		return p.valueBefore()
	}

	vals := p.leftTerm.values()
	if len(vals) == 0 {
		return nil
	}
	return vals[len(vals)-1]
}

func (t *snailTerm) valueAfter() *snailTerm {
	if t.isValueTerm() {
		return t
	}

	p := t.parent
	if p == nil {
		return nil
	}

	if p.rightTerm == t {
		return p.valueAfter()
	}

	vals := p.rightTerm.values()
	if len(vals) == 0 {
		return nil
	}
	return vals[0]
}

func (t *snailTerm) values() []*snailTerm {
	if t.isValueTerm() {
		return []*snailTerm{t}
	}

	res := []*snailTerm{}
	res = append(res, t.leftTerm.values()...)
	res = append(res, t.rightTerm.values()...)

	return res
}

func (t *snailTerm) explode() bool {
	lit := t.firstDeepLiteral(4)
	if lit == nil {
		return false
	}

	before := lit.valueBefore()
	if before != nil {
		before.val += lit.leftTerm.val
	}

	after := lit.valueAfter()
	if after != nil {
		after.val += lit.rightTerm.val
	}

	lit.leftTerm = nil
	lit.rightTerm = nil
	lit.val = 0

	return true
}

func (t *snailTerm) split() bool {
	any := false
	for _, v := range t.values() {
		if v.val > 9 {
			v.leftTerm = &snailTerm{
				parent: v,
				val:    v.val / 2,
			}
			v.rightTerm = &snailTerm{
				parent: v,
				val:    v.val / 2,
			}
			if v.val%2 == 1 {
				v.rightTerm.val++
			}
			v.val = 0
			any = true
			break
		}
	}
	return any
}

func (t *snailTerm) reduce() {
	for {
		if t.explode() {
			continue
		}
		if t.split() {
			continue
		}
		break
	}
}

func add(t1, t2 *snailTerm) *snailTerm {
	t := &snailTerm{
		leftTerm:  t1,
		rightTerm: t2,
	}
	t1.parent = t
	t2.parent = t

	t.reduce()

	return t
}

func (t snailTerm) magnitude() int {
	if t.isValueTerm() {
		return t.val
	}

	return 3*t.leftTerm.magnitude() + 2*t.rightTerm.magnitude()
}
