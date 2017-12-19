package circularlist

import "strings"

type List struct {
	curIdx int
	items  []*ListItem
}

type ListItem struct {
	Val int
}

func (l *List) Start() *ListItem {
	return l.items[0]
}

func (l *List) CurPos() *ListItem {
	return l.items[l.curIdx]
}

func (l *List) SetPos(p int) {
	l.curIdx = p
}

func New(val int) *List {
	list := List{}
	li := ListItem{Val: val}
	list.items = make([]*ListItem, 0, 5e7)
	list.items = append(list.items, &li)

	return &list
}

func (l *List) Advance(places int) {
	l.curIdx = (l.curIdx + places) % len(l.items)
}

func (l *List) Insert(val int) {
	new := ListItem{Val: val}
	l.items = append(l.items, &new)
	copy(l.items[l.curIdx+2:], l.items[l.curIdx+1:])
	l.items[l.curIdx+1] = &new
	l.curIdx++
}

func (l *List) SkipAndInsert(skipSize, iterations int) {
	for i := 0; i < iterations; i++ {
		l.Advance(skipSize)
		l.Insert(i + 1)
	}
}

func (l *List) String() string {
	nums := []string{}
	for _, item := range l.items {
		nums = append(nums, string(item.Val))
	}
	return strings.Join(nums, ",")
}
