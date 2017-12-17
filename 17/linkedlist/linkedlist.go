package linkedlist

type ListItem struct {
	Val  int
	Next *ListItem
}

func NewCiruclarList(val int) *ListItem {
	li := ListItem{Val: val}
	li.Next = &li
	return &li
}

func (li *ListItem) Insert(val int) *ListItem {
	new := ListItem{Val: val}
	new.Next = li.Next
	li.Next = &new
	return &new
}

func (li *ListItem) Advance(places int) *ListItem {
	ret := li
	for i := 0; i < places; i++ {
		ret = ret.Next
	}
	return ret
}

func (li *ListItem) SkipAndInsert(skipSize, iterations int) *ListItem {
	cur := li
	for i := 0; i < iterations; i++ {
		cur = cur.Advance(skipSize)
		cur = cur.Insert(i + 1)
	}
	return cur.Next
}
