package circularlist_test

import (
	"github.com/kieron-pivotal/advent2017/17/circularlist"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("circularlist", func() {
	It("can create a new list", func() {
		list := circularlist.New(0)
		list.Advance(1)
		Expect(list.CurPos()).To(Equal(list.Start()))
	})

	It("can insert after an item", func() {
		list := circularlist.New(0)
		list.Insert(1)
		newItem := list.CurPos()

		list.Advance(1)
		Expect(list.CurPos()).To(Equal(list.Start()))
		list.Advance(1)
		Expect(list.CurPos()).To(Equal(newItem))
	})

	It("can skip N items", func() {
		list := circularlist.New(0)
		list.Insert(1)
		list.Insert(2)
		list.Advance(1)
		Expect(list.CurPos().Val).To(Equal(0))
		list.Advance(5)
		Expect(list.CurPos().Val).To(Equal(2))
	})

	It("can skip and insert", func() {
		l := circularlist.New(0)
		l.SkipAndInsert(3, 2017)
		l.Advance(1)
		Expect(l.CurPos().Val).To(Equal(638))
	})

	It("gives the answer to part 1", func() {
		l := circularlist.New(0)
		l.SkipAndInsert(328, 2017)
		l.Advance(1)
		Expect(l.CurPos().Val).To(Equal(1670))
	})

	It("gives the answer to part 2", func() {
		l := circularlist.New(0)
		l.SkipAndInsert(328, 5e7)
		l.SetPos(1)
		Expect(l.CurPos().Val).To(Equal(-1))
	})
})
