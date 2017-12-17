package linkedlist

type List struct {
	Start  *ListItem
	Len    int
	CurPos *ListItem
}

type ListItem struct {
	Val  int
	Next *ListItem
}

func NewCiruclarList(val int) *List {
	list := List{Len: 1}
	li := ListItem{Val: val}
	li.Next = &li
	list.Start = &li
	list.CurPos = &li
	return &list
}

func (l *List) Insert(val int) {
	new := ListItem{Val: val}
	new.Next = l.CurPos.Next
	l.CurPos.Next = &new
	l.CurPos = &new
	l.Len++
}

func (l *List) Advance(places int) {
	for i := 0; i < places%l.Len; i++ {
		l.CurPos = l.CurPos.Next
	}
}

func (l *List) SkipAndInsert(skipSize, iterations int) {
	for i := 0; i < iterations; i++ {
		l.Advance(skipSize)
		l.Insert(i + 1)
	}
}
