package linkedlist_test

import (
	"github.com/kieron-pivotal/advent2017/17/linkedlist"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Linkedlist", func() {
	It("can create a new list", func() {
		list := linkedlist.NewCiruclarList(0)
		Expect(list.CurPos.Next).To(Equal(list.CurPos))
	})

	It("can insert after an item", func() {
		list := linkedlist.NewCiruclarList(0)
		list.Insert(1)
		Expect(list.Start.Next).To(Equal(list.CurPos))
		Expect(list.CurPos.Next).To(Equal(list.Start))
	})

	It("can skip N items", func() {
		list := linkedlist.NewCiruclarList(0)
		list.Insert(1)
		list.Insert(2)
		list.Advance(1)
		Expect(list.CurPos.Val).To(Equal(0))
		list.Advance(5)
		Expect(list.CurPos.Val).To(Equal(2))
	})

	It("can skip and insert", func() {
		l := linkedlist.NewCiruclarList(0)
		l.SkipAndInsert(3, 2017)
		Expect(l.CurPos.Next.Val).To(Equal(638))
	})

	It("gives the answer to part 1", func() {
		l := linkedlist.NewCiruclarList(0)
		l.SkipAndInsert(328, 2017)
		Expect(l.CurPos.Next.Val).To(Equal(1670))
	})

	It("gives the answer to part 2", func() {
		l := linkedlist.NewCiruclarList(0)
		l.SkipAndInsert(328, 5e7)
		Expect(l.Start.Next.Val).To(Equal(-1))
	})
})
