package linkedlist_test

import (
	"github.com/kieron-pivotal/advent2017/17/linkedlist"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Linkedlist", func() {
	It("can create a new list", func() {
		li := linkedlist.NewCiruclarList(0)
		Expect(li.Next).To(Equal(li))
	})

	It("can insert after an item", func() {
		li0 := linkedlist.NewCiruclarList(0)
		li1 := li0.Insert(1)
		Expect(li0.Next).To(Equal(li1))
		Expect(li1.Next).To(Equal(li0))
	})

	It("can skip N items", func() {
		i0 := linkedlist.NewCiruclarList(0)
		i1 := i0.Insert(1)
		i2 := i1.Insert(2)
		i := i2.Advance(1)
		Expect(i).To(Equal(i0))
		i = i2.Advance(5)
		Expect(i).To(Equal(i1))
	})

	It("can skip and insert", func() {
		li := linkedlist.NewCiruclarList(0)
		next := li.SkipAndInsert(3, 2017)
		Expect(next.Val).To(Equal(638))
	})

	It("gives the answer to part 1", func() {
		li := linkedlist.NewCiruclarList(0)
		next := li.SkipAndInsert(328, 2017)
		Expect(next.Val).To(Equal(-1))
	})
})
